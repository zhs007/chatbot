package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/pb"
)

// Plugin - chat bot plugin interface
type Plugin interface {
	// OnMessage - get message
	OnMessage(ctx context.Context, serv *Serv, chat *chatbotpb.ChatMsg,
		ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error)

	// OnStart - on start
	OnStart(ctx context.Context) error

	// GetPluginName - get plugin name
	GetPluginName() string
}
