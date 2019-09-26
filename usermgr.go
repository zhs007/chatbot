package chatbot

import (
	"context"

	chatbotpb "github.com/zhs007/chatbot/proto"
)

// UserMgr - user maqnager
type UserMgr interface {
	// GetAppUserInfo - get user infomation
	GetAppUserInfo(ctx context.Context, appToken string, uai *chatbotpb.UserAppInfo) (*chatbotpb.UserInfo, error)
}
