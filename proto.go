package chatbot

import (
	"bytes"

	"github.com/golang/protobuf/proto"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// BuildUserAppInfo - build UserAppInfo
func BuildUserAppInfo(appType chatbotpb.ChatAppType, usernameAppServ string,
	appuid string, appuname string, lang string) *chatbotpb.UserAppInfo {

	return &chatbotpb.UserAppInfo{
		App:             appType,
		UsernameAppServ: usernameAppServ,
		Appuid:          appuid,
		Appuname:        appuname,
		Lang:            lang,
	}
}

// BuildTextChatMsg - build ChatMsg
func BuildTextChatMsg(msg string, uai *chatbotpb.UserAppInfo, token string,
	sessionid string) *chatbotpb.ChatMsg {

	return &chatbotpb.ChatMsg{
		Msg:       msg,
		Uai:       uai,
		Token:     token,
		SessionID: sessionid,
	}
}

// BuildFileChatMsg - build ChatMsg
func BuildFileChatMsg(msg string, fd *chatbotpb.FileData, uai *chatbotpb.UserAppInfo, token string,
	sessionid string) *chatbotpb.ChatMsg {

	return &chatbotpb.ChatMsg{
		Msg:       msg,
		Uai:       uai,
		Token:     token,
		SessionID: sessionid,
		File:      fd,
	}
}

// BuildErrorChatMsg - build ChatMsg
func BuildErrorChatMsg(err error, uai *chatbotpb.UserAppInfo, token string,
	sessionid string) *chatbotpb.ChatMsg {

	return &chatbotpb.ChatMsg{
		Error:     err.Error(),
		Uai:       uai,
		Token:     token,
		SessionID: sessionid,
	}
}

// BuildChatMsgStream - ChatMsg -> ChatMsgStream
func BuildChatMsgStream(chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsgStream, error) {
	buf, err := proto.Marshal(chat)
	if err != nil {
		return nil, err
	}

	bl := len(buf)
	if bl <= chatbotbase.BigMsgLength {
		stream := &chatbotpb.ChatMsgStream{
			Chat:      chat,
			Token:     chat.Token,
			SessionID: chat.SessionID,
		}

		return []*chatbotpb.ChatMsgStream{stream}, nil
	}

	var lst []*chatbotpb.ChatMsgStream

	st := 0
	for st < bl {
		isend := false
		cl := chatbotbase.BigMsgLength
		if cl > bl-st {
			cl = bl - st
			isend = true
		}

		cb := buf[st:(st + cl)]

		cs := &chatbotpb.ChatMsgStream{
			TotalLength: int32(bl),
			CurStart:    int32(st),
			CurLength:   int32(cl),
			HashData:    chatbotbase.MD5Buffer(cb),
			Data:        cb,
			Token:       chat.Token,
			SessionID:   chat.SessionID,
		}

		st += cl
		if isend {
			cs.TotalHashData = chatbotbase.MD5Buffer(buf)
		}

		lst = append(lst, cs)
	}

	return lst, nil
}

// BuildChatMsg - []ChatMsgStream => ChatMsg
func BuildChatMsg(lst []*chatbotpb.ChatMsgStream) (*chatbotpb.ChatMsg, error) {
	if len(lst) == 1 {
		if lst[0].Chat != nil {
			return lst[0].Chat, nil
		}

		return nil, chatbotbase.ErrStreamNoMsg
	}

	var lstbytes [][]byte
	totalmd5inmsg := ""
	st := 0
	ct := 0
	for i, v := range lst {
		if st != int(v.CurStart) {
			return nil, chatbotbase.ErrInvalidStartInStream
		}

		if len(v.Data) != int(v.CurLength) {
			return nil, chatbotbase.ErrInvalidLengthInStream
		}

		curmd5 := chatbotbase.MD5Buffer(v.Data)
		if curmd5 != v.HashData {
			return nil, chatbotbase.ErrInvalidHashInStream
		}

		lstbytes = append(lstbytes, v.Data)

		st += len(v.Data)
		ct += len(v.Data)

		if i == len(lst)-1 {
			totalmd5inmsg = v.TotalHashData

			if ct != int(v.TotalLength) {
				return nil, chatbotbase.ErrInvalidTotalLengthInStream
			}
		}
	}

	buf := bytes.Join(lstbytes, []byte(""))

	totalmd5 := chatbotbase.MD5Buffer(buf)
	if totalmd5 != totalmd5inmsg {
		return nil, chatbotbase.ErrInvalidTotalHashInStream
	}

	chatmsg := &chatbotpb.ChatMsg{}
	err := proto.Unmarshal(buf, chatmsg)
	if err != nil {
		return nil, err
	}

	return chatmsg, nil
}

// BuildChatMsgList - []ChatMsgStream => []ChatMsg
func BuildChatMsgList(lst []*chatbotpb.ChatMsgStream) ([]*chatbotpb.ChatMsg, error) {
	var lstret []*chatbotpb.ChatMsg
	var curlst []*chatbotpb.ChatMsgStream

	for _, v := range lst {
		if v.Chat != nil {
			cr, err := BuildChatMsg([]*chatbotpb.ChatMsgStream{v})
			if err != nil {
				return nil, err
			}

			if cr != nil {
				lstret = append(lstret, cr)
			}
		} else {
			curlst = append(curlst, v)

			if v.TotalHashData != "" {
				cr, err := BuildChatMsg([]*chatbotpb.ChatMsgStream{v})
				if err != nil {
					return nil, err
				}

				if cr != nil {
					lstret = append(lstret, cr)
				}

				curlst = nil
			}
		}
	}

	return lstret, nil
}

// NewChatMsgWithText - new ChatMsg with TextMgr
func NewChatMsgWithText(l *i18n.Localizer, msgid string, tempParams map[string]interface{}, uai *chatbotpb.UserAppInfo) (
	*chatbotpb.ChatMsg, error) {

	txt, err := l.Localize(&i18n.LocalizeConfig{
		MessageID:    msgid,
		TemplateData: tempParams,
	})
	if err != nil {
		return nil, err
	}

	return &chatbotpb.ChatMsg{
		Msg: txt,
		Uai: uai,
	}, nil
}
