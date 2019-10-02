package main

import (
	"context"
	"fmt"

	chatbottelegram "github.com/zhs007/chatbot/app/telegram"
)

func main() {
	cfg, err := chatbottelegram.LoadConfig("./cfg/config.yaml")
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
