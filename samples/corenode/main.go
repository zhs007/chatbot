package main

import (
	"context"
	"fmt"

	chatbot "github.com/zhs007/chatbot"
	chatbotdebugplugin "github.com/zhs007/chatbot/plugins/debug"
	chatbotusermgr "github.com/zhs007/chatbot/usermgr"
)

func main() {
	err := chatbotdebugplugin.RegisterPlugin()
	if err != nil {
		fmt.Printf("chatbotdebugplugin.RegisterPlugin %v", err)

		return
	}

	cfg, err := chatbot.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	mgr, err := chatbotusermgr.NewUserMgr(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		fmt.Printf("NewUserMgr %v", err)

		return
	}

	serv, err := chatbot.NewChatBotServ(cfg, mgr)
	if err != nil {
		fmt.Printf("NewChatBotServ %v", err)

		return
	}

	serv.Init(context.Background())

	serv.Start(context.Background())
}
