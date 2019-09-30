package chatbot

import (
	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// ServiceCore - chatbot service core
type ServiceCore interface {
	// Unmarshal - unmarshal
	UnmarshalAppData(buf []byte) (proto.Message, error)
	// New - new a app data
	NewAppData() (proto.Message, error)

	// Unmarshal - unmarshal
	UnmarshalUserData(buf []byte) (proto.Message, error)
	// New - new a userdata
	NewUserData(ui *chatbotpb.UserInfo) (proto.Message, error)
}
