package main

import (
	"context"
	"fmt"

	chatbotcoolq "github.com/zhs007/chatbot/app/coolq"
	chatbotbase "github.com/zhs007/chatbot/base"
	"go.uber.org/zap/zapcore"
)

func main() {
	chatbotbase.InitLogger(zapcore.InfoLevel, true, "./logs")

	cfg, err := chatbotcoolq.LoadConfig("./cfg/coolq.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	serv, err := chatbotcoolq.NewServ(cfg)
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
