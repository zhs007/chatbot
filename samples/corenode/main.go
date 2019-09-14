package main

import (
	"context"
	"fmt"

	chatbot "github.com/zhs007/chatbot"
)

func main() {
	cfg, err := chatbot.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	serv, err := chatbot.NewChatBotServ(cfg)
	if err != nil {
		fmt.Printf("NewChatBotServ %v", err)

		return
	}

	serv.Init(context.Background())

	serv.Start(context.Background())
}
