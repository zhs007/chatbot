package chatbot

import (
	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/pb"
	chatbotusermgr "github.com/zhs007/chatbot/usermgr"
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

// NewLocalUserMgr - new local UserMgr
func NewLocalUserMgr(dbpath string, httpAddr string, engine string, core ServiceCore) (
	*chatbotusermgr.UserMgr, error) {

	return chatbotusermgr.NewUserMgr(dbpath, httpAddr, engine,
		&userDataMgr{core: core})
}
