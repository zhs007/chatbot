package chatbot

import (
	"context"
	"net"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotdb "github.com/zhs007/chatbot/db"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Serv - service
type Serv struct {
	cfg       *Config
	lis       net.Listener
	grpcServ  *grpc.Server
	dbAppServ *chatbotdb.AppServDB
}

// NewChatBotServ -
func NewChatBotServ(cfg *Config) (*Serv, error) {
	if cfg == nil {
		return nil, chatbotbase.ErrNoConfig
	}

	db, err := chatbotdb.NewAppServDB(cfg.DBPath, "", cfg.DBEngine)
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
		cfg:      cfg,
		lis:      lis,
		grpcServ: grpcServ,
	}

	chatbotpb.RegisterChatBotServiceServer(grpcServ, serv)

	serv.dbAppServ = db

	chatbotbase.Info("NewChatBotServ OK.")

	return serv, nil
}

// Init - initial service
func (serv *Serv) Init(ctx context.Context) error {
	return serv.dbAppServ.Init(ctx, serv.cfg.AppServ)
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
	return nil
}

// RequestChat - request chat
func (serv *Serv) RequestChat(req *chatbotpb.RequestChatData,
	ecs chatbotpb.ChatBotService_RequestChatServer) error {

	return nil
}
