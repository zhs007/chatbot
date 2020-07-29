package chatbotpreprocplugin

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
func (pp *preprocPlugin) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message, scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	for _, v := range pp.cfg.LstRegexp {
		msg, err := procRegexpNode(v, chat)
		if err != nil {
			chatbotbase.Error("preprocPlugin.OnMessage:procRegexpNode",
				zap.Error(err))

			return nil, err
		}

		if msg != nil {
			return nil, nil
		}
	}

	return nil, nil
}

// OnStart - on start
func (pp *preprocPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (pp *preprocPlugin) GetPluginName() string {
	return "preprocessor"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin(fn string) error {
	cfg, err := LoadConfig(fn)
	if err != nil {
		chatbotbase.Error("chatbotpreprocplugin.RegisterPlugin:LoadConfig",
			zap.String("fn", fn),
			zap.Error(err))

		return err
	}

	return chatbot.RegPlugin(&preprocPlugin{cfg: cfg})
}
