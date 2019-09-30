package main

import (
	"context"
	"fmt"

	chatbottelegram "github.com/zhs007/chatbot/app/telegram"
	chatbotdebugplugin "github.com/zhs007/chatbot/plugins/debug"
)

func main() {
	err := chatbotdebugplugin.RegisterPlugin()
	if err != nil {
		fmt.Printf("chatbotdebugplugin.RegisterPlugin %v", err)

		return
	}

	cfg, err := chatbottelegram.LoadConfig("./config.yaml")
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
