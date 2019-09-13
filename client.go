package chatbot

import (
	"context"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"google.golang.org/grpc"
)

// Client - ChatBotServiceClient
type Client struct {
	servAddr string
	token    string
	conn     *grpc.ClientConn
	client   chatbotpb.ChatBotServiceClient
}

// NewClient - new ChatBotServiceClient
func NewClient(servAddr string, token string) *Client {
	return &Client{
		servAddr: servAddr,
		token:    token,
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
func (client *Client) RegisterAppService(ctx context.Context, apptype chatbotpb.ChatAppType,
	token string, username string) (*chatbotpb.ReplyRegisterAppService, error) {

	err := client.isValid()
	if err != nil {
		return nil, err
	}

	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}

		client.conn = conn
		client.client = chatbotpb.NewChatBotServiceClient(conn)
	}

	reply, err := client.client.RegisterAppService(ctx, &chatbotpb.RegisterAppService{
		AppType:  apptype,
		Token:    token,
		Username: username,
	})
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	return reply, nil
}
