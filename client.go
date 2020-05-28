package chatbot

import (
	"context"
	"errors"
	"io"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"google.golang.org/grpc"
)

// Client - ChatBotServiceClient
type Client struct {
	servAddr  string
	token     string
	SessionID string
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

// onSendMsg - on send message to service
func (client *Client) onSendMsg() error {
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

	return nil
}

// RegisterAppService - RegisterAppService
func (client *Client) RegisterAppService(ctx context.Context) error {

	err := client.onSendMsg()
	if err != nil {
		// if error, reset
		client.reset()

		return err
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

	client.SessionID = reply.SessionID

	return nil
}

// SendChat - SendChat
func (client *Client) SendChat(ctx context.Context, chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {

	err := client.onSendMsg()
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	lst, err := BuildChatMsgStream(chat)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	stream, err := client.client.SendChat(ctx)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	var recverr error
	var lstrecv []*chatbotpb.ChatMsgStream
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)

				return
			}

			if err != nil {
				recverr = err

				close(waitc)

				return
			}

			lstrecv = append(lstrecv, in)
		}
	}()

	for _, cn := range lst {
		curerr := stream.Send(cn)
		if curerr != nil {
			// if error, close
			stream.CloseSend()

			// if error, reset
			client.reset()

			return nil, curerr
		}
	}

	stream.CloseSend()
	<-waitc

	if recverr != nil {
		// if error, reset
		client.reset()

		return nil, recverr
	}

	lstret, err := BuildChatMsgList(lstrecv)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	return lstret, nil
}

// RequestChat - RequestChat
func (client *Client) RequestChat(ctx context.Context) ([]*chatbotpb.ChatMsg, error) {

	err := client.onSendMsg()
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	in := &chatbotpb.RequestChatData{
		Token:     client.token,
		SessionID: client.SessionID,
	}

	stream, err := client.client.RequestChat(ctx, in)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	var recverr error
	var lstrecv []*chatbotpb.ChatMsgStream
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)

				return
			}

			if err != nil {
				recverr = err

				close(waitc)

				return
			}

			lstrecv = append(lstrecv, in)
		}
	}()

	<-waitc

	if recverr != nil {
		// if error, reset
		client.reset()

		return nil, recverr
	}

	lstret, err := BuildChatMsgList(lstrecv)
	if err != nil {
		// if error, reset
		client.reset()

		return nil, err
	}

	return lstret, nil
}
