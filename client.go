package chatbot

import (
	"context"
	"errors"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"google.golang.org/grpc"
)

// Client - ChatBotServiceClient
type Client struct {
	servAddr  string
	token     string
	sessionID string
	appType   chatbotpb.ChatAppType
	username  string
	conn      *grpc.ClientConn
	client    chatbotpb.ChatBotServiceClient
}

// NewClient - new ChatBotServiceClient
func NewClient(servAddr string, apptype chatbotpb.ChatAppType, token string, username string) *Client {
	return &Client{
		servAddr: servAddr,
		token:    token,
		appType:  apptype,
		username: username,
	}
}

// isValid - check myself
func (client *Client) isValid() error {
	if client.servAddr == "" {
		return chatbotbase.ErrNoServAddrInClient
	}

	if client.token == "" {
		return chatbotbase.ErrNoTokenInClient
	}

	return nil
}

// reset - reset
func (client *Client) reset() {
	if client.conn != nil {
		client.conn.Close()
	}

	client.conn = nil
	client.client = nil
}

// RegisterAppService - RegisterAppService
func (client *Client) RegisterAppService(ctx context.Context) error {

	err := client.isValid()
	if err != nil {
		return err
	}

	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}

		client.conn = conn
		client.client = chatbotpb.NewChatBotServiceClient(conn)
	}

	reply, err := client.client.RegisterAppService(ctx, &chatbotpb.RegisterAppService{
		AppServ: &chatbotpb.AppServInfo{
			AppType:  client.appType,
			Token:    client.token,
			Username: client.username,
		},
	})
	if err != nil {
		// if error, reset
		client.reset()

		return err
	}

	if reply.Error != "" {
		return errors.New(reply.Error)
	}

	client.sessionID = reply.SessionID

	return nil
}
