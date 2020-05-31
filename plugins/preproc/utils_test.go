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

	if ret.Msg != "note -m search -n comic  -k \"北斗神拳\"" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "找漫画北斗神拳 原哲夫",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret.Msg != "note -m search -n comic  -k \"北斗神拳\" -k \"原哲夫\"" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "很难找",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	t.Log("Test_procRegexpNode OK")
}

func Test_procRegexpNode2(t *testing.T) {
	rn := &RegexpNode{
		Pattern:          ".*(漫画列表)",
		Prefix:           "note -m keys -n comic",
		Mode:             "nodata",
		ParamArrayPrefix: "",
	}

	r, err := regexp.Compile(rn.Pattern)
	if err != nil {
		t.Fatalf("Test_procRegexpNode2 Compile %v", err)
	}

	rn.r = r

	chat := &chatbotpb.ChatMsg{
		Msg: "我想查漫画列表",
	}

	ret, err := procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode2 procRegexpNode %v", err)
	}

	if ret.Msg != "note -m keys -n comic" {
		t.Fatalf("Test_procRegexpNode2 procRegexpNode fail. %v", ret.Msg)
	}

	t.Log("Test_procRegexpNode2 OK")
}

func Test_procRegexpNode3(t *testing.T) {
	rn := &RegexpNode{
		Pattern:          "找+(.*)",
		Prefix:           "note -m search -n comic ",
		Mode:             "paramarray",
		ParamArrayPrefix: " -k ",
	}

	r, err := regexp.Compile(rn.Pattern)
	if err != nil {
		t.Fatalf("Test_procRegexpNode3 Compile %v", err)
	}

	rn.r = r

	chat := &chatbotpb.ChatMsg{
		Msg: "很难找",
	}

	ret, err := procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode3 procRegexpNode %v", err)
	}

	if ret != nil {
		t.Fatalf("Test_procRegexpNode3 procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "找漫画北斗神拳 原哲夫",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret.Msg != "note -m search -n comic  -k \"漫画北斗神拳\" -k \"原哲夫\"" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "找找漫画北斗神拳 原哲夫",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret.Msg != "note -m search -n comic  -k \"漫画北斗神拳\" -k \"原哲夫\"" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "给我好好找找  ",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	chat = &chatbotpb.ChatMsg{
		Msg: "找找\"漫画北斗神拳 原\"哲夫",
	}

	ret, err = procRegexpNode(rn, chat)
	if err != nil {
		t.Fatalf("Test_procRegexpNode procRegexpNode %v", err)
	}

	if ret.Msg != "note -m search -n comic  -k \"漫画北斗神拳 原\" -k \"哲夫\"" {
		t.Fatalf("Test_procRegexpNode procRegexpNode fail. %v", ret.Msg)
	}

	t.Log("Test_procRegexpNode3 OK")
}
