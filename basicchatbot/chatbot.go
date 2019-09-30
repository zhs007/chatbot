package basicchatbot

import (
	chatbotcmdhelp "github.com/zhs007/chatbot/commands/help"
	chatbotcmdstart "github.com/zhs007/chatbot/commands/start"
	chatbotcmdplugin "github.com/zhs007/chatbot/plugins/cmd"
	chatbotdebugplugin "github.com/zhs007/chatbot/plugins/debug"
	chatbotfileplugin "github.com/zhs007/chatbot/plugins/file"
)

// InitBasicChatBot - initial basic chatbot
func InitBasicChatBot() error {
	err := chatbotdebugplugin.RegisterPlugin()
	if err != nil {
		return err
	}

	err = chatbotcmdplugin.RegisterPlugin()
	if err != nil {
		return err
	}

	err = chatbotfileplugin.RegisterPlugin()
	if err != nil {
		return err
	}

	chatbotcmdhelp.RegisterCommand()
	chatbotcmdstart.RegisterCommand()

	return nil
}
