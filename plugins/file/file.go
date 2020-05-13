package chatbotfileplugin

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/pb"
)

type filePlugin struct {
}

// OnMessage - get message
func (dbp *filePlugin) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	if chat.File == nil {
		return nil, nil
	}

	if serv == nil {
		return nil, chatbotbase.ErrPluginInvalidServ
	}

	if serv.MgrFile == nil {
		return nil, chatbotbase.ErrInvalidFileProcessorMgr
	}

	return serv.MgrFile.ProcFile(ctx, serv, chat, ui, ud)
}

// OnStart - on start
func (dbp *filePlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (dbp *filePlugin) GetPluginName() string {
	return "fileproc"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin() error {
	return chatbot.RegPlugin(&filePlugin{})
}
