package chatbotcmdhelp

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// cmdStart - command start
type cmdStart struct {
}

// RunCommand - run command
func (cmd *cmdStart) RunCommand(ctx context.Context, serv *chatbot.Serv, params proto.Message,
	chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {

	return nil, nil
}

// ParseCommandLine - parse command line
func (cmd *cmdStart) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	proto.Message, error) {

	return nil, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand("start", &cmdStart{})
}
