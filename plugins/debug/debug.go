package chatbotdebugplugin

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
)

type debugPlugin struct {
}

// OnMessage - get message
func (dbp *debugPlugin) OnMessage(ctx context.Context, msg *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo,
	ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	var lst []*chatbotpb.ChatMsg

	msggetit := &chatbotpb.ChatMsg{
		Msg: "I get it.",
		Uai: msg.Uai,
	}

	lst = append(lst, msggetit)

	strui, err := chatbotbase.JSONFormat(ui)
	if err != nil {
		chatbotbase.Warn("debugPlugin.OnMessage:UI",
			zap.Error(err))
	} else {
		msgui := &chatbotpb.ChatMsg{
			Msg: strui,
			Uai: msg.Uai,
		}

		lst = append(lst, msgui)
	}

	return lst, nil
}

// OnStart - on start
func (dbp *debugPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (dbp *debugPlugin) GetPluginName() string {
	return "debug"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin() error {
	return chatbot.RegPlugin(&debugPlugin{})
}
