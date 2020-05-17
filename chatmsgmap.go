package chatbot

import (
	"sync"

	chatbotpb "github.com/zhs007/chatbot/pb"
)

// FuncOnChatMsgs - function
type FuncOnChatMsgs func(lst []*chatbotpb.ChatMsg)

// chatMsgList - ChatMsg List
type chatMsgList struct {
	lst  []*chatbotpb.ChatMsg
	lock sync.Mutex
}

// chatMsgMap - ChatMsg Map
type chatMsgMap struct {
	mapMsgs map[string]*chatMsgList
}

func newChatMsgMap() *chatMsgMap {
	return &chatMsgMap{
		mapMsgs: make(map[string]*chatMsgList),
	}
}

// addChatMsgs - add chat msgs
func (msgmap *chatMsgMap) addChatMsgs(token string, lst []*chatbotpb.ChatMsg) {
	msgs, isok := msgmap.mapMsgs[token]
	if !isok {
		msgs = &chatMsgList{}
		msgmap.mapMsgs[token] = msgs
	}

	msgs.lock.Lock()
	msgs.lst = append(msgs.lst, lst...)
	msgs.lock.Unlock()
}

// addChatMap - push list
func (msgmap *chatMsgMap) getChatMsgs(token string, onChatMsgs FuncOnChatMsgs) {
	msgs, isok := msgmap.mapMsgs[token]
	if !isok {
		onChatMsgs(nil)

		return
	}

	msgs.lock.Lock()
	onChatMsgs(msgs.lst)
	msgs.lst = nil
	msgs.lock.Unlock()
}
