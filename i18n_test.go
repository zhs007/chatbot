package chatbot

import (
	"testing"

	chatbotbase "github.com/zhs007/chatbot/base"
)

func Test_loadLangKeys(t *testing.T) {
	arr, err := loadLangKeys("./lang/chatbot.en.yaml")
	if err != nil {
		t.Fatalf("Test_loadLangKeys Err %v", err)
	}

	if len(arr) != 13 {
		t.Fatalf("Test_loadLangKeys len %v", len(arr))
	}

	arr0 := []string{
		"chatbotname",
		"help001",
		"help002",
		"help003",
		"start001",
		"start002",
		"start003",
		"igetit",
		"yousaid",
		"igetfile",
		"notesearch001",
		"notesearch002",
		"notesearchnone",
	}

	for _, v := range arr {
		if chatbotbase.IndexOfArrayString(arr0, v) < 0 {
			t.Fatalf("Test_loadLangKeys %v", v)
		}
	}

	t.Logf("Test_loadLangKeys OK")
}
