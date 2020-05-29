package chatbotproprocplugin

import (
	"strings"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

func procRegexpNode(rn *RegexpNode, chat *chatbotpb.ChatMsg) (*chatbotpb.ChatMsg, error) {
	arr := rn.r.FindAllStringSubmatchIndex(chat.Msg, -1)
	if len(arr) > 0 {
		chatbotbase.Debug("chatbotproprocplugin.procRegexpNode:FindAllStringSubmatchIndex",
			zap.String("pattern", rn.Pattern),
			zap.String("msg", chat.Msg),
			chatbotbase.JSON("ret", arr))

		if len(arr) == 1 && len(arr[0]) == 4 {
			str := rn.Prefix

			ss := chat.Msg[arr[0][2]:arr[0][3]]

			strarr := strings.Fields(ss)
			for _, v := range strarr {
				str += rn.ParamArrayPrefix
				str += v
			}

			chat.Msg = str

			return chat, nil
		}

	}

	return nil, nil
}
