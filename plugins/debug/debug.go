package chatbotdebugplugin

import (
	"context"

	chatbot "github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

type debugPlugin struct {
}

// OnMessage - get message
func (dbp *debugPlugin) OnMessage(ctx context.Context, msg *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo) (
	[]*chatbotpb.ChatMsg, error) {

	var lst []*chatbotpb.ChatMsg

	msggetit := &chatbotpb.ChatMsg{
		Msg: "I get it.",
		Uai: msg.Uai,
	}

	lst = append(lst, msggetit)

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
