package chatbotcmdhelp

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// cmdHelp - command help
type cmdHelp struct {
}

// RunCommand - run command
func (cmd *cmdHelp) RunCommand(ctx context.Context, serv *chatbot.Serv, params interface{},
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	if serv == nil {
		return true, nil, chatbotbase.ErrCmdInvalidServ
	}

	if serv.MgrText == nil {
		return true, nil, chatbotbase.ErrCmdInvalidServMgrText
	}

	lsthelp := serv.Cfg.HelpText
	if lsthelp == nil {
		return true, nil, nil
	}

	lang := serv.GetChatMsgLang(chat)

	mParams, err := serv.BuildBasicParamsMap(chat, ui, lang)
	if err != nil {
		return true, nil, err
	}

	locale, err := serv.MgrText.GetLocalizer(lang)
	if err != nil {
		return true, nil, err
	}

	var lst []*chatbotpb.ChatMsg
	for _, v := range lsthelp {
		txt, err := locale.Localize(&i18n.LocalizeConfig{
			MessageID:    v,
			TemplateData: mParams,
		})
		if err != nil {
			return true, nil, err
		}

		msgtxt := &chatbotpb.ChatMsg{
			Msg: txt,
			Uai: chat.Uai,
		}

		lst = append(lst, msgtxt)
	}

	return true, lst, nil
}

// OnMessage - get message
func (cmd *cmdHelp) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	return true, nil, chatbotbase.ErrCmdItsNotMine
}

// ParseCommandLine - parse command line
func (cmd *cmdHelp) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	interface{}, error) {

	return nil, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand("help", &cmdHelp{})
}
