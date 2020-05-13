package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/pb"
)

// FileProcessor - file processor
type FileProcessor interface {
	// Proc - process
	Proc(ctx context.Context, serv *Serv, chat *chatbotpb.ChatMsg,
		ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error)

	// IsMyFile - is my file
	IsMyFile(chat *chatbotpb.ChatMsg) bool
}
