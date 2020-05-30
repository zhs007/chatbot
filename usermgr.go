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

	// GetNoteInfo - get note infomation
	GetNoteInfo(ctx context.Context, name string) (
		*chatbotpb.NoteInfo, error)
	// UpdNoteInfo - update note infomation
	UpdNoteInfo(ctx context.Context, ni *chatbotpb.NoteInfo) error
	// UpdNoteNode - update note node
	UpdNoteNode(ctx context.Context, nn *chatbotpb.NoteNode) error
	// GetNoteNode - get note node
	GetNoteNode(ctx context.Context, nameNote string, noteIndex int64) (
		*chatbotpb.NoteNode, error)
	// DelNoteNode - delete note node
	DelNoteNode(ctx context.Context, nameNote string, noteIndex int64) error
}
