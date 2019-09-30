package chatbot

import (
	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/proto"
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
