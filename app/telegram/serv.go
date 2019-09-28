package chatbottelegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
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
	err := serv.client.RegisterAppService(ctx)
	if err != nil {
		return err
	}

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

			err = serv.onMsg(ctx, &update)
			if err != nil {
				chatbotbase.Warn("chatbottelegram.Serv.Start",
					zap.Error(err))
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

// onMsg - on message
func (serv *Serv) onMsg(ctx context.Context, upd *tgbotapi.Update) error {
	if upd.Message != nil {
		from := upd.Message.From
		uai := chatbot.BuildUserAppInfo(chatbotpb.ChatAppType_CAT_TELEGRAM,
			serv.cfg.Username, ID2Str(from.ID), from.UserName, from.LanguageCode)

		if upd.Message.Text != "" {
			msg := chatbot.BuildTextChatMsg(upd.Message.Text,
				uai, serv.cfg.Token, serv.client.SessionID)

			lstret, err := serv.client.SendChat(ctx, msg)
			if err != nil {
				return err
			}

			for _, v := range lstret {
				err = serv.SendChatMsg(ctx, v)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// SendChatMsg - send a chat message
func (serv *Serv) SendChatMsg(ctx context.Context, msg *chatbotpb.ChatMsg) error {
	i64, err := Str2ID(msg.Uai.Appuid)
	if err != nil {
		return err
	}

	telemsg := tgbotapi.NewMessage(i64, msg.Msg)

	telemsg.DisableWebPagePreview = true

	_, err = serv.bot.Send(telemsg)
	if err != nil {
		return err
	}

	return nil
}
