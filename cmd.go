package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// Command - command
type Command interface {
	// RunCommand - run command
	RunCommand(ctx context.Context, serv *Serv, params interface{},
		chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
		scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error)
	// ParseCommandLine - parse command line
	ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (interface{}, error)
	// OnMessage - get message
	OnMessage(ctx context.Context, serv *Serv, chat *chatbotpb.ChatMsg,
		ui *chatbotpb.UserInfo, ud proto.Message,
		scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error)
}
