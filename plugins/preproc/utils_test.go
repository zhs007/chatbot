package chatbotproprocplugin

import (
	"regexp"
	"testing"

	"github.com/zhs007/chatbot/chatbotpb"
)

func Test_procRegexpNode(t *testing.T) {
	rn := &RegexpNode{
		Pattern:          "找漫画(.*)",
		Prefix:           "note -m search -n comic ",
		Mode:             "paramarray",
		ParamArrayPrefix: " -k ",
	}

	r, err := regexp.Compile(rn.Pattern)
	if err != nil {
		t.Fatalf("Test_procRegexpNode Compile %v", err)
	}

	rn.r = r

	chat := &chatbotpb.ChatMsg{
		Msg: "找漫画北斗神拳",
	}

	ret, err := procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret.Msg != "note -m search -n comic  -k 北斗神拳" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "找漫画北斗神拳 原哲夫",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret.Msg != "note -m search -n comic  -k 北斗神拳 -k 原哲夫" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	t.Log("procRegexpNode OK")
}
