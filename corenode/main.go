package main

import (
	"context"
	"fmt"

	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	"github.com/zhs007/chatbot/basicchatbot"
	"go.uber.org/zap/zapcore"
)

func main() {

	err := basicchatbot.InitBasicChatBot()
	if err != nil {
		fmt.Printf("basicchatbot.InitBasicChatBot %v", err)

		return
	}

	cfg, err := chatbot.LoadConfig("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	chatbotbase.InitLogger(zapcore.InfoLevel, true, "./")

	serv, err := chatbot.NewSimpleChatBotServ(cfg, &basicchatbot.ServiceCore{})
	if err != nil {
		fmt.Printf("NewChatBotServ %v", err)

		return
	}

	serv.Init(context.Background())

	serv.Start(context.Background())
}
