package chatbottelegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	chatbot "github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// Serv - serv
type Serv struct {
	cfg    *Config
	bot    *tgbotapi.BotAPI
	client *chatbot.Client
}

// NewServ - new a service
func NewServ(cfg *Config) (*Serv, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	client := chatbot.NewClient(cfg.ServAddr,
		chatbotpb.ChatAppType_CAT_TELEGRAM,
		cfg.Token,
		cfg.Username)

	serv := &Serv{
		cfg:    cfg,
		bot:    bot,
		client: client,
	}

	return serv, nil
}

// Start - start telegram bot
func (serv *Serv) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := serv.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for {
		isend := false

		select {
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}
		case <-ctx.Done():
			isend = true

			break
		}

		if isend {
			break
		}
	}

	return nil
}
