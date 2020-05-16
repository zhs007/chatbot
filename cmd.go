package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/pb"
)

// Command - command
type Command interface {
	// RunCommand - run command
	RunCommand(ctx context.Context, serv *Serv, params proto.Message,
		chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
		scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error)
	// ParseCommandLine - parse command line
	ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (proto.Message, error)
}
