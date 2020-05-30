package chatbot

import (
	"context"
	"io"
	"net"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	chatbotdb "github.com/zhs007/chatbot/db"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Serv - service
type Serv struct {
	Cfg         *Config
	lis         net.Listener
	grpcServ    *grpc.Server
	dbAppServ   *chatbotdb.AppServDB
	lstPlugins0 *PluginsList
	lstPlugins  *PluginsList
	lstPlugins2 *PluginsList
	MgrUser     UserMgr
	MgrText     *TextMgr
	Cmds        *CommondsList
	Core        ServiceCore
	MgrFile     *FileProcessorMgr
	mapChatMsgs *chatMsgMap
}

// NewChatBotServ -
func NewChatBotServ(cfg *Config, mgr UserMgr, core ServiceCore) (*Serv, error) {
	if cfg == nil {
		return nil, chatbotbase.ErrNoConfig
	}

	db, err := chatbotdb.NewAppServDB(cfg.DBPath, "", cfg.DBEngine, &appDataMgr{
		core: core,
	})
	if err != nil {
		return nil, err
	}

	mgrText, err := NewTextMgr(cfg)
	if err != nil {
		return nil, err
	}

	lis, err := net.Listen("tcp", cfg.BindAddr)
	if err != nil {
		chatbotbase.Error("NewChatBotServ", zap.Error(err))

		return nil, err
	}

	chatbotbase.Info("Listen", zap.String("addr", cfg.BindAddr))

	grpcServ := grpc.NewServer()

	serv := &Serv{
		Cfg:         cfg,
		lis:         lis,
		grpcServ:    grpcServ,
		lstPlugins0: NewPluginsList(),
		lstPlugins:  NewPluginsList(),
		lstPlugins2: NewPluginsList(),
		MgrUser:     mgr,
		MgrText:     mgrText,
		Cmds:        NewCommondsList(),
		MgrFile:     &FileProcessorMgr{},
		Core:        core,
		mapChatMsgs: newChatMsgMap(),
	}

	for _, v := range cfg.PluginsPreprocess {
		err = serv.lstPlugins0.AddPlugin(v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range cfg.Plugins {
		err = serv.lstPlugins.AddPlugin(v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range cfg.PluginsSecondLine {
		err = serv.lstPlugins2.AddPlugin(v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range cfg.Commands {
		err = serv.Cmds.AddCommand(v)
		if err != nil {
			return nil, err
		}
	}

	chatbotpb.RegisterChatBotServiceServer(grpcServ, serv)

	serv.dbAppServ = db

	chatbotbase.Info("NewChatBotServ OK.")

	return serv, nil
}

// NewSimpleChatBotServ - new a simple chatbot service with local user manager
func NewSimpleChatBotServ(cfg *Config, core ServiceCore) (*Serv, error) {
	mgr, err := NewLocalUserMgr(cfg.DBPath, "", cfg.DBEngine, core)
	if err != nil {
		chatbotbase.Error("NewSimpleChatBotServ:NewLocalUserMgr", zap.Error(err))

		return nil, err
	}

	return NewChatBotServ(cfg, mgr, core)
}

// Init - initial service
func (serv *Serv) Init(ctx context.Context) error {
	return serv.dbAppServ.Init(ctx, serv.Cfg.AppServ)
}

// Start - start a service
func (serv *Serv) Start(ctx context.Context) error {
	return serv.grpcServ.Serve(serv.lis)
}

// Stop - stop service
func (serv *Serv) Stop() {
	serv.lis.Close()

	return
}

// RegisterAppService - register app service
func (serv *Serv) RegisterAppService(ctx context.Context, ras *chatbotpb.RegisterAppService) (
	*chatbotpb.ReplyRegisterAppService, error) {

	asi, err := serv.dbAppServ.GetAppServ(ctx, ras.AppServ.Token)
	if err != nil {
		return &chatbotpb.ReplyRegisterAppService{
			Error: err.Error(),
		}, nil
	}

	if asi.AppType != ras.AppServ.AppType {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServType.Error(),
		}, nil
	}

	if asi.Username != ras.AppServ.Username {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServUserName.Error(),
		}, nil
	}

	asi, err = serv.dbAppServ.GenerateSessionID(ctx, asi)
	if err != nil {
		return &chatbotpb.ReplyRegisterAppService{
			Error: err.Error(),
		}, nil
	}

	return &chatbotpb.ReplyRegisterAppService{
		AppType:   ras.AppServ.AppType,
		SessionID: asi.Sessionid,
	}, nil
}

// SendChat - send chat
func (serv *Serv) SendChat(scs chatbotpb.ChatBotService_SendChatServer) error {
	var lst []*chatbotpb.ChatMsgStream

	for {
		in, err := scs.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			chatbotbase.Warn("SendChat:Recv",
				zap.Error(err))

			serv.replySendChatErr(scs, err)

			return err
		}

		isvalidtoken, err := serv.dbAppServ.CheckTokenSessionID(scs.Context(), in.Token, in.SessionID)
		if err != nil {
			chatbotbase.Warn("SendChat:CheckTokenSessionID",
				zap.Error(err))

			serv.replySendChatErr(scs, err)

			return err
		}

		if !isvalidtoken {
			chatbotbase.Warn("SendChat:isvalidtoken",
				zap.Error(chatbotbase.ErrServInvalidToken))

			serv.replySendChatErr(scs, chatbotbase.ErrServInvalidToken)

			return chatbotbase.ErrServInvalidToken
		}

		lst = append(lst, in)
	}

	cd, err := BuildChatMsg(lst)
	if err != nil {
		chatbotbase.Warn("SendChat:BuildChatMsg",
			zap.Error(err))

		serv.replySendChatErr(scs, err)

		return err
	}

	if cd == nil {
		chatbotbase.Warn("SendChat:ErrInvalidStream",
			zap.Error(chatbotbase.ErrInvalidStream))

		serv.replySendChatErr(scs, chatbotbase.ErrInvalidStream)

		return chatbotbase.ErrInvalidStream
	}

	ui, ud, err := serv.MgrUser.GetAppUserInfo(scs.Context(), cd.Token, cd.Uai)
	if err != nil || ui == nil {
		if err != nil {
			chatbotbase.Warn("SendChat:GetAppUserInfo",
				zap.Error(err))
		} else {
			chatbotbase.Warn("SendChat:GetAppUserInfo",
				zap.Error(chatbotbase.ErrServInvalidUserInfo))
		}

		serv.replySendChatErr(scs, chatbotbase.ErrServInvalidUserInfo)

		return chatbotbase.ErrServInvalidUserInfo
	}

	lstret0, err := serv.lstPlugins0.OnMessage(scs.Context(), serv, cd, ui, ud, scs)
	if err != nil {
		chatbotbase.Warn("SendChat:lstPlugins0.OnMessage",
			zap.Error(err))

		serv.replySendChatErr(scs, err)

		return err
	}

	lstret, err := serv.lstPlugins.OnMessage(scs.Context(), serv, cd, ui, ud, scs)
	if err != nil {
		chatbotbase.Warn("SendChat:lstPlugins.OnMessage",
			zap.Error(err))

		serv.replySendChatErr(scs, err)

		return err
	}

	lstret2, err := serv.lstPlugins2.OnMessageEx(scs.Context(), serv, cd, ui, ud, scs)
	if err != nil {
		chatbotbase.Warn("SendChat:lstPlugins2.OnMessageEx",
			zap.Error(err))

		serv.replySendChatErr(scs, err)

		return err
	}

	if lstret0 != nil {
		lstret = append(lstret0, lstret...)
	}

	if lstret2 != nil {
		lstret = append(lstret, lstret2...)
	}

	for _, v := range lstret {
		if cd.Gai != nil && !v.IsReplyPrivate {
			v.Gai = cd.Gai
		}

		lststream, err := BuildChatMsgStream(v)
		if err != nil {
			chatbotbase.Warn("SendChat:BuildChatMsgStream",
				zap.Error(err))

			serv.replySendChatErr(scs, err)

			return err
		}

		for _, sv := range lststream {
			err = scs.Send(sv)
			if err != nil {
				chatbotbase.Warn("SendChat:Send",
					zap.Error(err))

				serv.replySendChatErr(scs, err)

				return err
			}
		}
	}

	return nil
}

// replySendChatErr - reply a error for SendChat
func (serv *Serv) replySendChatErr(scs chatbotpb.ChatBotService_SendChatServer, err error) error {
	if err == nil {
		return serv.replySendChatErr(scs, chatbotbase.ErrServInvalidErr)
	}

	// chatbotbase.Warn("replySendChatErr", zap.Error(err))

	reply := &chatbotpb.ChatMsgStream{
		Error: err.Error(),
	}

	return scs.Send(reply)
}

// RequestChat - request chat
func (serv *Serv) RequestChat(req *chatbotpb.RequestChatData,
	ecs chatbotpb.ChatBotService_RequestChatServer) error {

	isvalidtoken, err := serv.dbAppServ.CheckTokenSessionID(ecs.Context(), req.Token, req.SessionID)
	if err != nil {
		chatbotbase.Warn("RequestChat:CheckTokenSessionID",
			zap.Error(err))

		serv.replyRequestChatErr(ecs, err)

		return err
	}

	if !isvalidtoken {
		chatbotbase.Warn("RequestChat:isvalidtoken",
			zap.Error(chatbotbase.ErrServInvalidToken))

		serv.replyRequestChatErr(ecs, chatbotbase.ErrServInvalidToken)

		return chatbotbase.ErrServInvalidToken
	}

	var lstret []*chatbotpb.ChatMsg

	serv.mapChatMsgs.getChatMsgs(req.Token, func(lst []*chatbotpb.ChatMsg) {
		if len(lst) > 0 {
			lstret = append(lstret, lst...)
		}
	})

	if len(lstret) > 0 {
		for _, v := range lstret {
			lststream, err := BuildChatMsgStream(v)
			if err != nil {
				chatbotbase.Warn("RequestChat:BuildChatMsgStream",
					zap.Error(err))

				serv.replyRequestChatErr(ecs, err)

				return err
			}

			for _, sv := range lststream {
				err = ecs.Send(sv)
				if err != nil {
					chatbotbase.Warn("RequestChat:Send",
						zap.Error(err))

					serv.replyRequestChatErr(ecs, err)

					return err
				}
			}
		}
	} else {
		reply := &chatbotpb.ChatMsgStream{
			IsNoMsg: true,
		}

		ecs.Send(reply)
	}

	return nil
}

// replyRequestChatErr - reply a error for RequestChat
func (serv *Serv) replyRequestChatErr(ecs chatbotpb.ChatBotService_RequestChatServer, err error) error {
	if err == nil {
		return serv.replyRequestChatErr(ecs, chatbotbase.ErrServInvalidErr)
	}

	// chatbotbase.Warn("replySendChatErr", zap.Error(err))

	reply := &chatbotpb.ChatMsgStream{
		Error: err.Error(),
	}

	return ecs.Send(reply)
}

// BuildBasicParamsMap - build basic params map
func (serv *Serv) BuildBasicParamsMap(chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, lang string) (
	map[string]interface{}, error) {

	locale, err := serv.MgrText.GetLocalizer(lang)
	if err != nil {
		return nil, err
	}

	chatbotname, err := locale.Localize(&i18n.LocalizeConfig{
		MessageID: serv.Cfg.ChatBotNameText,
	})
	if err != nil {
		return nil, err
	}

	fn := ""
	ft := ""
	fs := ""
	fh := ""

	if chat.File != nil {
		fn = chat.File.Filename
		ft = chat.File.FileType
		fs = chatbotbase.FormatFileSize(int64(len(chat.File.FileData)))
		fh = chatbotbase.MD5Buffer(chat.File.FileData)
	}

	return map[string]interface{}{
		"ChatBotName": chatbotname,
		"Name":        ui.Name,
		"UID":         ui.Uid,
		"TextChat":    chat.Msg,
		"FileName":    fn,
		"FileType":    ft,
		"FileSize":    fs,
		"FileHash":    fh,
	}, nil
}

// GetChatMsgLang - get chat message language
func (serv *Serv) GetChatMsgLang(chat *chatbotpb.ChatMsg) string {
	if chat.Uai.Lang != "" {
		return chat.Uai.Lang
	}

	return serv.Cfg.Language
}

// RequestCtrl - request control
func (serv *Serv) RequestCtrl(ctx context.Context, msg *chatbotpb.RequestCtrlData) (
	*chatbotpb.AppCtrlData, error) {

	return nil, nil
}

// SendCtrlResult - send control result
func (serv *Serv) SendCtrlResult(ctx context.Context, msg *chatbotpb.AppCtrlResult) (
	*chatbotpb.SCRResult, error) {

	return nil, nil
}

// PushChatMsgs - send ChatMsg List
func (serv *Serv) PushChatMsgs(token string, lst []*chatbotpb.ChatMsg) {
	serv.mapChatMsgs.addChatMsgs(token, lst)
}
