package chatbotdb

import (
	"github.com/golang/protobuf/proto"
)

// AppDataMgr - app data for chatbot
type AppDataMgr interface {
	// Unmarshal - unmarshal
	Unmarshal(buf []byte) (proto.Message, error)
	// New - new a app data
	New() (proto.Message, error)
}
