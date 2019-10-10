package chatbotcoolq

import (
	qqbotapi "github.com/catsworld/qq-bot-api"
	"github.com/zhs007/chatbot"
)

// Serv - serv
type Serv struct {
	bot    *qqbotapi.BotAPI
	client *chatbot.Client
}
