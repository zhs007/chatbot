package chatbotpreprocplugin

import (
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

func procRegexpNode(rn *RegexpNode, chat *chatbotpb.ChatMsg) (*chatbotpb.ChatMsg, error) {
	arr := rn.r.FindAllStringSubmatchIndex(chat.Msg, -1)
	if len(arr) > 0 {
		chatbotbase.Debug("chatbotpreprocplugin.procRegexpNode:FindAllStringSubmatchIndex",
			zap.String("pattern", rn.Pattern),
			zap.String("msg", chat.Msg),
			chatbotbase.JSON("ret", arr))

		if len(arr) == 1 && len(arr[0]) == 4 {
			str := rn.Prefix

			ss := chat.Msg[arr[0][2]:arr[0][3]]
			if ss != "" {
				if rn.Mode == "paramarray" {
					strarr := chatbot.SplitCommandString(ss)
					if len(strarr) == 0 {
						return nil, nil
					}

					// strarr := strings.Fields(ss)
					for _, v := range strarr {
						str += rn.ParamArrayPrefix
						str += "\""
						str += v
						str += "\""
					}

				} else if rn.Mode == "nodata" {
				}

				chat.Msg = str

				return chat, nil
			}
		}

	}

	return nil, nil
}
