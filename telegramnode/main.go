package main

import (
	"context"
	"fmt"

	chatbottelegram "github.com/zhs007/chatbot/app/telegram"
	chatbotbase "github.com/zhs007/chatbot/base"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	chatbotbase.InitLogger(zapcore.InfoLevel, true, "./logs")

	chatbotbase.Info("telegram start...",
		zap.String("version", chatbotbase.VERSION))

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
