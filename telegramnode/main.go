package main

import (
	"context"
	"fmt"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbottelegram "github.com/zhs007/chatbot/app/telegram"
	"go.uber.org/zap/zapcore"
)

func main() {
	chatbotbase.InitLogger(zapcore.InfoLevel, true, "./logs")

	cfg, err := chatbottelegram.LoadConfig("./cfg/telegram.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	serv, err := chatbottelegram.NewServ(cfg)
	if err != nil {
		fmt.Printf("NewChatBotServ %v", err)

		return
	}

	err = serv.Start(context.Background())
	if err != nil {
		fmt.Printf("Start %v", err)

		return
	}
}
