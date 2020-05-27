package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// UserMgr - user maqnager
type UserMgr interface {
	// GetAppUserInfo - get user infomation
	GetAppUserInfo(ctx context.Context, appToken string, uai *chatbotpb.UserAppInfo) (
		*chatbotpb.UserInfo, proto.Message, error)
}
