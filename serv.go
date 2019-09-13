package chatbot

import (
	"context"
	"net"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Serv - service
type Serv struct {
	cfg        *Config
	lis        net.Listener
	grpcServ   *grpc.Server
	mgrAppServ AppServMgr
}

// NewChatBotServ -
func NewChatBotServ(cfg *Config) (*Serv, error) {
	if cfg == nil {
		return nil, chatbotbase.ErrNoConfig
	}

	lis, err := net.Listen("tcp", cfg.BindAddr)
	if err != nil {
		chatbotbase.Error("NewChatBotServ", zap.Error(err))

		return nil, err
	}

	chatbotbase.Info("Listen", zap.String("addr", cfg.BindAddr))

	grpcServ := grpc.NewServer()

	serv := &Serv{
		cfg:      cfg,
		lis:      lis,
		grpcServ: grpcServ,
	}

	chatbotpb.RegisterChatBotServiceServer(grpcServ, serv)

	serv.mgrAppServ.Init(cfg)

	chatbotbase.Info("NewChatBotServ OK.")

	return serv, nil
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

	cfgAppServ := serv.mgrAppServ.getAppServ(ras.Token)
	if cfgAppServ == nil {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServToken.Error(),
		}, nil
	}

	if cfgAppServ.typeAppServ != ras.AppType {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServType.Error(),
		}, nil
	}

	if cfgAppServ.UserName != ras.Username {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServUserName.Error(),
		}, nil
	}

	return &chatbotpb.ReplyRegisterAppService{
		AppType: ras.AppType,
	}, nil
}

// sendChat - send chat
func (serv *Serv) SendChat(scs chatbotpb.ChatBotService_SendChatServer) error {
	return nil
}

// requestChat - request chat
func (serv *Serv) RequestChat(req *chatbotpb.RequestChatData,
	ecs chatbotpb.ChatBotService_RequestChatServer) error {

	return nil
}
