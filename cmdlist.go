package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

// CommondsList - command list
type CommondsList struct {
	cmds map[string]Command
}

// CommandData - command data
type CommandData struct {
	Cmd    string
	Params interface{}
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

// MultiParseInChat - parse in chat
func (cmds *CommondsList) MultiParseInChat(chat *chatbotpb.ChatMsg) ([]*CommandData, error) {
	lstlines := SplitMultiCommandString(chat.Msg)
	if len(lstlines) <= 0 {
		return nil, chatbotbase.ErrCmdEmptyCmd
	}

	lstcmds := []*CommandData{}
	for _, v := range lstlines {
		lst := SplitCommandString(v)
		if len(lst) <= 0 {
			chatbotbase.Error("CommondsList.MultiParseInChat.SplitCommandString",
				zap.String("cmd", v),
				zap.Error(chatbotbase.ErrCmdEmptyCmd))

			continue
		}

		cmd := lst[0]

		c, ok := cmds.cmds[cmd]
		if ok {
			params, err := c.ParseCommandLine(lst, chat)
			if err != nil {
				chatbotbase.Error("CommondsList.MultiParseInChat.ParseCommandLine",
					zap.String("cmd", cmd),
					zap.String("cmd string", v),
					zap.Error(err))

				continue
			}

			lstcmds = append(lstcmds, &CommandData{
				Cmd:    cmd,
				Params: params,
			})
		}
	}

	return lstcmds, chatbotbase.ErrCmdNoCmd
}

// ParseInChat - parse in chat
func (cmds *CommondsList) ParseInChat(chat *chatbotpb.ChatMsg) (string, interface{}, error) {
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
func (cmds *CommondsList) RunInChat(ctx context.Context, cmd string, serv *Serv, params interface{},
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	c, ok := cmds.cmds[cmd]
	if ok {
		return c.RunCommand(ctx, serv, params, chat, ui, ud, scs)
	}

	return true, nil, chatbotbase.ErrCmdNoCmd
}

// OnMessage - get message
func (cmds *CommondsList) OnMessage(ctx context.Context, cmd string, serv *Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	c, ok := cmds.cmds[cmd]
	if ok {
		return c.OnMessage(ctx, serv, chat, ui, ud, scs)
	}

	return true, nil, chatbotbase.ErrCmdNoCmd
}

// HasCommand - has command
func (cmds *CommondsList) HasCommand(cmd string) bool {
	_, ok := cmds.cmds[cmd]
	if ok {
		return true
	}

	return false
}
