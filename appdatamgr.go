package chatbot

import (
	"github.com/golang/protobuf/proto"
)

// appDataMgr - app data for chatbot
type appDataMgr struct {
	core ServiceCore
}

// Unmarshal - unmarshal
func (mgr *appDataMgr) Unmarshal(buf []byte) (proto.Message, error) {
	return mgr.core.UnmarshalAppData(buf)
}

// New - new a app data
func (mgr *appDataMgr) New() (proto.Message, error) {
	return mgr.core.NewAppData()
}
