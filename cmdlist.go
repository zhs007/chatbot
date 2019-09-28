package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// CommondsList - command list
type CommondsList struct {
	cmds map[string]Command
}

// NewCommondsList - new CommondsList
func NewCommondsList() *CommondsList {
	return &CommondsList{
		cmds: make(map[string]Command),
	}
}

// AddCommand - add command
func (cmds *CommondsList) AddCommand(cmd string) error {
	c := mgrCommands.GetCommand(cmd)
	if c == nil {
		return chatbotbase.ErrCmdNoCmd
	}

	cmds.cmds[cmd] = c

	return nil
}

// ParseInChat - parse in chat
func (cmds *CommondsList) ParseInChat(chat *chatbotpb.ChatMsg) (string, proto.Message, error) {
	lst := SplitCommandString(chat.Msg)
	if len(lst) <= 0 {
		return "", nil, chatbotbase.ErrCmdEmptyCmd
	}

	cmd := lst[0]

	c, ok := cmds.cmds[cmd]
	if ok {
		params, err := c.ParseCommandLine(lst, chat)

		return cmd, params, err
	}

	return "", nil, chatbotbase.ErrCmdNoCmd
}

// RunInChat - run func with cmd
func (cmds *CommondsList) RunInChat(ctx context.Context, cmd string, serv *Serv, params proto.Message,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	c, ok := cmds.cmds[cmd]
	if ok {
		return c.RunCommand(ctx, serv, params, chat, ui, ud)
	}

	return nil, chatbotbase.ErrCmdNoCmd
}

// HasCommand - has command
func (cmds *CommondsList) HasCommand(cmd string) bool {
	_, ok := cmds.cmds[cmd]
	if ok {
		return true
	}

	return false
}
