package chatbotcmdhelp

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// CmdHelp - command help
type CmdHelp struct {
}

// RunCommand - run command
func (cmd *CmdHelp) RunCommand(ctx context.Context, params proto.Message,
	chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {

	return nil, nil
}

// ParseCommandLine - parse command line
func (cmd *CmdHelp) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	proto.Message, error) {

	return nil, nil
}
