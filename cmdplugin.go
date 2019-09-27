package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

type cmdPlugin struct {
}

// OnMessage - get message
func (cp *cmdPlugin) OnMessage(ctx context.Context, serv *Serv, msg *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	cmd, params, err := serv.cmds.ParseInChat(msg)
	if err != nil {
		if err != chatbotbase.ErrCmdNoCmd {
			return nil, err
		}

		return nil, nil
	}

	if cmd != "" {
		lst, err := serv.cmds.RunInChat(ctx, cmd, serv, params, msg)
		if err != nil {
			if err != chatbotbase.ErrCmdNoCmd {
				return nil, err
			}

			return nil, nil
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
