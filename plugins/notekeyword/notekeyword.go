package chatbotnotekeywordplugin

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

type noteKeywordPlugin struct {
}

// OnMessage - get message
func (pp *noteKeywordPlugin) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message, scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	if serv.Cfg.NoteKeyword != "" {
		ni, err := serv.MgrUser.GetNoteInfo(ctx, serv.Cfg.NoteKeyword)
		if err != nil {
			chatbotbase.Error("noteKeywordPlugin.OnMessage:GetNoteInfo",
				zap.Error(err))

			return nil, err
		}

		ks, isok := ni.MapKeys[chat.Msg]
		if isok {
			var lstmsg []*chatbotpb.ChatMsg

			for _, v := range ks.Nodes {
				nn, err := serv.MgrUser.GetNoteNode(ctx, ni.Name, v)
				if err != nil {
					chatbotbase.Error("noteKeywordPlugin.OnMessage:GetNoteNode",
						zap.Error(err))
				}

				if nn != nil {
					if nn.SendAppMsgID == "" {
						msg := &chatbotpb.ChatMsg{
							Uai: chat.Uai,
							Forward: &chatbotpb.ForwardData{
								Uai:      nn.Uai,
								AppMsgID: nn.ForwardAppMsgID,
							},
							IsReplyPrivate: true,
						}

						if msg.Forward.Uai == nil {
							msg.Forward.Uai = chat.Uai
						}

						lstmsg = append(lstmsg, msg)
					} else {
						msg := &chatbotpb.ChatMsg{
							Uai: chat.Uai,
							Forward: &chatbotpb.ForwardData{
								Uai:      nn.SendUai,
								AppMsgID: nn.SendAppMsgID,
							},
							IsReplyPrivate: true,
						}

						lstmsg = append(lstmsg, msg)
					}
				}
			}

			return lstmsg, nil
		}
	}

	// for _, v := range pp.cfg.LstRegexp {
	// 	msg, err := procRegexpNode(v, chat)
	// 	if err != nil {
	// 		chatbotbase.Error("preprocPlugin.OnMessage:procRegexpNode",
	// 			zap.Error(err))

	// 		return nil, err
	// 	}

	// 	if msg != nil {
	// 		return nil, nil
	// 	}
	// }

	return nil, nil
}

// OnStart - on start
func (pp *noteKeywordPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (pp *noteKeywordPlugin) GetPluginName() string {
	return "notekeyword"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin() error {
	// cfg, err := LoadConfig(fn)
	// if err != nil {
	// 	chatbotbase.Error("chatbotproprocplugin.RegisterPlugin:LoadConfig",
	// 		zap.String("fn", fn),
	// 		zap.Error(err))

	// 	return err
	// }

	return chatbot.RegPlugin(&noteKeywordPlugin{})
}
