package chatbotdb

import (
	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// UserDataMgr - user data for chatbot
type UserDataMgr interface {
	// Unmarshal - unmarshal
	Unmarshal(buf []byte) (proto.Message, error)
	// New - new a userdata
	New(ui *chatbotpb.UserInfo) (proto.Message, error)
}
