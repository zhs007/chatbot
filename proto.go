package chatbot

import (
	"bytes"

	"github.com/golang/protobuf/proto"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// BuildUserAppInfo - build UserAppInfo
func BuildUserAppInfo(appType chatbotpb.ChatAppType, usernameAppServ string,
	appuid string, appuname string) *chatbotpb.UserAppInfo {

	return &chatbotpb.UserAppInfo{
		App:             appType,
		UsernameAppServ: usernameAppServ,
		Appuid:          appuid,
		Appuname:        appuname,
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
func BuildChatMsgStream(msg *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsgStream, error) {
	buf, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}

	bl := len(buf)
	if bl <= chatbotbase.BigMsgLength {
		stream := &chatbotpb.ChatMsgStream{
			Chat:      msg,
			Token:     msg.Token,
			SessionID: msg.SessionID,
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
			Token:       msg.Token,
			SessionID:   msg.SessionID,
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
