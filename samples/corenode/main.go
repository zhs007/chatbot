package main

import (
	"context"
	"fmt"

	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	"github.com/zhs007/chatbot/basicchatbot"
	chatbotusermgr "github.com/zhs007/chatbot/usermgr"
	"go.uber.org/zap/zapcore"
)

func main() {

	err := basicchatbot.InitBasicChatBot()
	if err != nil {
		fmt.Printf("basicchatbot.InitBasicChatBot %v", err)

		return
	}

	cfg, err := chatbot.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	chatbotbase.InitLogger(zapcore.InfoLevel, true, "./")

	mgr, err := chatbotusermgr.NewUserMgr(cfg.DBPath, "", cfg.DBEngine, nil)
	if err != nil {
		fmt.Printf("NewUserMgr %v", err)

		return
	}

	serv, err := chatbot.NewChatBotServ(cfg, mgr, &chatbot.EmptyServiceCore{})
	if err != nil {
		fmt.Printf("NewChatBotServ %v", err)

		return
	}

	serv.Init(context.Background())

	serv.Start(context.Background())
}
