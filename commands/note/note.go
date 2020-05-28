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

// NoteMode - note mode
type NoteMode int

const (
	// NoteModeNone - none
	NoteModeNone NoteMode = 0
	// NoteModeNew - new note
	NoteModeNew NoteMode = 1
	// NoteModeForward - taking notes for forward
	NoteModeForward NoteMode = 2
)

// ParseNoteMode - string => NoteMode
func ParseNoteMode(str string) NoteMode {
	if str == "new" {
		return NoteModeNew
	} else if str == "forward" {
		return NoteModeForward
	}

	return NoteModeNone
}

type paramsCmd struct {
	mode NoteMode
	name string
	keys []string
}

// cmdNote - command note
type cmdNote struct {
	IsTakingNotes bool
	Keys          []string
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

	cp, isok := params.(paramsCmd)
	if !isok {
		return true, nil, ErrCmdInvalidParams
	}

	cmd.IsTakingNotes = true
	cmd.Keys = cp.keys

	msg := &chatbotpb.ChatMsg{
		Msg: "note start...",
		Uai: chat.Uai,
	}

	return false, []*chatbotpb.ChatMsg{msg}, nil
}

// OnMessage - get message
func (cmd *cmdNote) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	if !cmd.IsTakingNotes {
		return true, nil, chatbotbase.ErrCmdItsNotMine
	}

	if chat.Forward == nil {
		msg := &chatbotpb.ChatMsg{
			Msg: "note end.",
			Uai: chat.Uai,
		}

		return true, []*chatbotpb.ChatMsg{msg}, chatbotbase.ErrCmdItsNotMine
	}

	msg := &chatbotpb.ChatMsg{
		Uai: chat.Uai,
		Forward: &chatbotpb.ForwardData{
			Uai:      chat.Uai,
			AppMsgID: chat.AppMsgID,
		},
	}

	return false, []*chatbotpb.ChatMsg{msg}, nil
}

// ParseCommandLine - parse command line
func (cmd *cmdNote) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	interface{}, error) {

	flagset := pflag.NewFlagSet(CmdName, pflag.ContinueOnError)

	strMode := flagset.StringP("mode", "m", "none", "mode")
	strName := flagset.StringP("name", "n", "", "name")
	keys := flagset.StringArrayP("keys", "k", []string{}, "key")

	err := flagset.Parse(cmdline[1:])
	if err != nil {
		return nil, err
	}

	return paramsCmd{
		mode: ParseNoteMode(*strMode),
		name: *strName,
		keys: *keys,
	}, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand(CmdName, &cmdNote{})
}
