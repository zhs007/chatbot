package chatbotcmdnote

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// CmdName - cmd name
const CmdName = "note"

type paramsCmd struct {
	keys []string
}

// cmdNote - command note
type cmdNote struct {
	IsTakingNotes bool
}

// RunCommand - run command
func (cmd *cmdNote) RunCommand(ctx context.Context, serv *chatbot.Serv, params interface{},
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	if serv == nil {
		return true, nil, chatbotbase.ErrCmdInvalidServ
	}

	if params == nil {
		return true, nil, ErrCmdNoParams
	}

	return true, nil, nil
}

// OnMessage - get message
func (cmd *cmdNote) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	return true, nil, chatbotbase.ErrCmdItsNotMine
}

// ParseCommandLine - parse command line
func (cmd *cmdNote) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	interface{}, error) {

	flagset := pflag.NewFlagSet(CmdName, pflag.ContinueOnError)

	keys := flagset.StringArrayP("keys", "k", []string{}, "key")

	err := flagset.Parse(cmdline[1:])
	if err != nil {
		return nil, err
	}

	return paramsCmd{keys: *keys}, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand(CmdName, &cmdNote{})
}
