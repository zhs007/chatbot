package chatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// FileProcessorMgr - file processor manager
type FileProcessorMgr struct {
	lst []FileProcessor
}

// RegisterFileProcessor - register file processor
func (m *FileProcessorMgr) RegisterFileProcessor(fp FileProcessor) {
	m.lst = append(m.lst, fp)
}

// ProcFile - process a file
func (m *FileProcessorMgr) ProcFile(ctx context.Context, serv *Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	for _, v := range m.lst {
		if v.IsMyFile(chat) {
			return v.Proc(ctx, serv, chat, ui, ud)
		}
	}

	return nil, nil
}
