package chatbotproprocplugin

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

type preprocPlugin struct {
	cfg *Config
}

// OnMessage - get message
func (cp *preprocPlugin) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message, scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	return nil, nil
}

// OnStart - on start
func (cp *preprocPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (cp *preprocPlugin) GetPluginName() string {
	return "preprocessor"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin(fn string) error {
	cfg, err := LoadConfig(fn)
	if err != nil {
		chatbotbase.Error("chatbotproprocplugin.RegisterPlugin:LoadConfig",
			zap.String("fn", fn),
			zap.Error(err))

		return err
	}

	return chatbot.RegPlugin(&preprocPlugin{cfg: cfg})
}
