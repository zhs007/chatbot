package chatbot

import (
	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// userDataMgr - user data for chatbot
type userDataMgr struct {
	core ServiceCore
}

// Unmarshal - unmarshal
func (mgr *userDataMgr) Unmarshal(buf []byte) (proto.Message, error) {
	return mgr.core.UnmarshalUserData(buf)
}

// New - new a app data
func (mgr *userDataMgr) New(ui *chatbotpb.UserInfo) (proto.Message, error) {
	return mgr.core.NewUserData(ui)
}
