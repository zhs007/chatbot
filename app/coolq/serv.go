package chatbotcoolq

import (
	"context"

	qqbotapi "github.com/catsworld/qq-bot-api"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
)

// Serv - serv
type Serv struct {
	cfg    *Config
	bot    *qqbotapi.BotAPI
	client *chatbot.Client
}

// NewServ - new a service
func NewServ(cfg *Config) (*Serv, error) {
	bot, err := qqbotapi.NewBotAPI(cfg.CoolQToken, cfg.CoolQServURL, cfg.CoolQSecret)
	if err != nil {
		return nil, err
	}

	client := chatbot.NewClient(cfg.ServAddr,
		chatbotpb.ChatAppType_CAT_COOLQ,
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
	err := serv.client.RegisterAppService(ctx)
	if err != nil {
		return err
	}

	serv.bot.Debug = serv.cfg.Debug

	u := qqbotapi.NewUpdate(0)
	u.PreloadUserInfo = true
	updates, err := serv.bot.GetUpdatesChan(u)

	for {
		isend := false

		select {
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			chatbotbase.Info(update.Message.Text,
				zap.Int64("ID", update.Message.Chat.ID),
				zap.String("Type", update.Message.Chat.Type))

			// err = serv.onMsg(ctx, &update)
			// if err != nil {
			// 	chatbotbase.Warn("chatbottelegram.Serv.Start",
			// 		zap.Error(err))
			// }
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
