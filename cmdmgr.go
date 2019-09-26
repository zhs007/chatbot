package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// CommandMgr - command manager
type CommandMgr struct {
	mapCmd map[string]Command
}

// newCommandMgr - new CommandMgr
func newCommandMgr() *CommandMgr {
	return &CommandMgr{
		mapCmd: make(map[string]Command),
	}
}

// RegisterCommand - add command with cmd
func (m *CommandMgr) RegisterCommand(cmd string, c Command) {
	m.mapCmd[cmd] = c
}

// ParseInChat - parse in chat
func (m *CommandMgr) ParseInChat(chat *chatbotpb.ChatMsg) (string, proto.Message, error) {
	lst := SplitCommandString(chat.Msg)
	if len(lst) <= 0 {
		return "", nil, chatbotbase.ErrCmdEmptyCmd
	}

	cmd := lst[0]

	c, ok := m.mapCmd[cmd]
	if ok {
		params, err := c.ParseCommandLine(lst, chat)

		return cmd, params, err
	}

	return "", nil, chatbotbase.ErrCmdNoCmd
}

// RunInChat - run func with cmd
func (m *CommandMgr) RunInChat(ctx context.Context, cmd string, params proto.Message,
	chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {

	c, ok := m.mapCmd[cmd]
	if ok {
		return c.RunCommand(ctx, params, chat)
	}

	return nil, chatbotbase.ErrCmdNoCmd
}

// HasCommand - has command
func (m *CommandMgr) HasCommand(cmd string) bool {
	_, ok := m.mapCmd[cmd]
	if ok {
		return true
	}

	return false
}

var mgrCommand *CommandMgr

func init() {
	mgrCommand = newCommandMgr()
}

// RegisterCommand - add command with cmd
func RegisterCommand(cmd string, c Command) {
	mgrCommand.RegisterCommand(cmd, c)
}
