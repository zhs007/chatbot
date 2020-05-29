package basicchatbot

import (
	chatbot "github.com/zhs007/chatbot"
	chatbotcmdhelp "github.com/zhs007/chatbot/commands/help"
	chatbotcmdnote "github.com/zhs007/chatbot/commands/note"
	chatbotcmdstart "github.com/zhs007/chatbot/commands/start"
	chatbotcmdplugin "github.com/zhs007/chatbot/plugins/cmd"
	chatbotdebugplugin "github.com/zhs007/chatbot/plugins/debug"
	chatbotfileplugin "github.com/zhs007/chatbot/plugins/file"
	chatbotpreprocplugin "github.com/zhs007/chatbot/plugins/preproc"
)

// InitBasicChatBot - initial basic chatbot
func InitBasicChatBot(cfg chatbot.Config) error {
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

	if cfg.Preprocessor != "" {
		err = chatbotpreprocplugin.RegisterPlugin(cfg.Preprocessor)
		if err != nil {
			return err
		}
	}

	chatbotcmdhelp.RegisterCommand()
	chatbotcmdstart.RegisterCommand()
	chatbotcmdnote.RegisterCommand()

	return nil
}
