package chatbotcmdplugin

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

type cmdPlugin struct {
	activeCmd string
}

// OnMessage - get message
func (cp *cmdPlugin) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message, scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	cmd, params, err := serv.Cmds.ParseInChat(chat)
	if err != nil {
		if err != chatbotbase.ErrCmdNoCmd &&
			err != chatbotbase.ErrCmdEmptyCmd {

			return nil, err
		}

		return nil, chatbotbase.ErrPluginItsNotMine
	}

	if cmd != "" {
		isend, lst, err := serv.Cmds.RunInChat(ctx, cmd, serv, params, chat, ui, ud, scs)
		if err != nil {
			if err != chatbotbase.ErrCmdNoCmd {
				return nil, err
			}

			return nil, chatbotbase.ErrPluginItsNotMine
		}

		if !isend {
			cp.activeCmd = cmd
		}
		// cp.activeCmd =

		return lst, nil
	}

	if cp.activeCmd != "" {
		isend, lst, err := serv.Cmds.OnMessage(ctx, cmd, serv, chat, ui, ud, scs)
		if err != nil {
			if err == chatbotbase.ErrCmdItsNotMine {
				return nil, chatbotbase.ErrPluginItsNotMine
			}

			if err != chatbotbase.ErrCmdNoCmd {
				return nil, err
			}

			return nil, nil
		}

		if isend {
			cp.activeCmd = ""
		}

		return lst, nil
	}

	return nil, nil
}

// OnStart - on start
func (cp *cmdPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (cp *cmdPlugin) GetPluginName() string {
	return "command"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin() error {
	return chatbot.RegPlugin(&cmdPlugin{})
}
