package chatbot

import (
	"context"

	chatbotpb "github.com/zhs007/chatbot/proto"
)

// Plugin - chat bot plugin interface
type Plugin interface {
	// OnMessage - get message
	OnMessage(ctx context.Context, msg *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo) ([]*chatbotpb.ChatMsg, error)

	// OnStart - on start
	OnStart(ctx context.Context) error

	// GetPluginName - get plugin name
	GetPluginName() string
}
