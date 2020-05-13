package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/pb"
)

// ServiceCore - chatbot service core
type ServiceCore interface {
	// UnmarshalAppData - unmarshal
	UnmarshalAppData(buf []byte) (proto.Message, error)
	// NewAppData - new a app data
	NewAppData() (proto.Message, error)

	// UnmarshalUserData - unmarshal
	UnmarshalUserData(buf []byte) (proto.Message, error)
	// NewUserData - new a userdata
	NewUserData(ui *chatbotpb.UserInfo) (proto.Message, error)

	// OnDebug - call in plugin.debug
	OnDebug(ctx context.Context, serv *Serv, chat *chatbotpb.ChatMsg,
		ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error)
}

// EmptyServiceCore - chatbot service core
type EmptyServiceCore struct {
}

// UnmarshalAppData - unmarshal
func (core *EmptyServiceCore) UnmarshalAppData(buf []byte) (proto.Message, error) {
	return nil, nil
}

// NewAppData - new a app data
func (core *EmptyServiceCore) NewAppData() (proto.Message, error) {
	return nil, nil
}

// UnmarshalUserData - unmarshal
func (core *EmptyServiceCore) UnmarshalUserData(buf []byte) (proto.Message, error) {
	return nil, nil
}

// NewUserData - new a userdata
func (core *EmptyServiceCore) NewUserData(ui *chatbotpb.UserInfo) (proto.Message, error) {
	return nil, nil
}

// OnDebug - call in plugin.debug
func (core *EmptyServiceCore) OnDebug(ctx context.Context, serv *Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {
	return nil, nil
}
