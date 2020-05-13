package chatbotcmdstart

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/pb"
)

// cmdStart - command start
type cmdStart struct {
}

// RunCommand - run command
func (cmd *cmdStart) RunCommand(ctx context.Context, serv *chatbot.Serv, params proto.Message,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {
	if serv == nil {
		return nil, chatbotbase.ErrCmdInvalidServ
	}

	if serv.MgrText == nil {
		return nil, chatbotbase.ErrCmdInvalidServMgrText
	}

	lststart := serv.Cfg.StartText
	if lststart == nil {
		return nil, nil
	}

	lang := serv.GetChatMsgLang(chat)

	mParams, err := serv.BuildBasicParamsMap(chat, ui, lang)
	if err != nil {
		return nil, err
	}

	locale, err := serv.MgrText.GetLocalizer(lang)
	if err != nil {
		return nil, err
	}

	var lst []*chatbotpb.ChatMsg
	for _, v := range lststart {
		txt, err := locale.Localize(&i18n.LocalizeConfig{
			MessageID:    v,
			TemplateData: mParams,
		})
		if err != nil {
			return nil, err
		}

		msgtxt := &chatbotpb.ChatMsg{
			Msg: txt,
			Uai: chat.Uai,
		}

		lst = append(lst, msgtxt)
	}

	return lst, nil
}

// ParseCommandLine - parse command line
func (cmd *cmdStart) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	proto.Message, error) {

	return nil, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand("start", &cmdStart{})
}
