package chatbottelegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func isValidMsg(msg *tgbotapi.Message) bool {
	if msg.Text != "" || msg.Caption != "" || msg.Document != nil || msg.Photo != nil {
		return true
	}

	return false
}
