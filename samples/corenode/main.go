package main

import (
	"context"
	"fmt"

	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	"github.com/zhs007/chatbot/basicchatbot"
)

func main() {

	cfg, err := chatbot.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	chatbotbase.InitLogger("chatbot", chatbotbase.VERSION, "info", true, "./")

	err = basicchatbot.InitBasicChatBot(*cfg)
	if err != nil {
		fmt.Printf("basicchatbot.InitBasicChatBot %v", err)

		return
	}

	serv, err := chatbot.NewSimpleChatBotServ(cfg, &chatbot.EmptyServiceCore{})
	if err != nil {
		fmt.Printf("NewChatBotServ %v", err)

		return
	}

	serv.Init(context.Background())

	serv.Start(context.Background())
}
