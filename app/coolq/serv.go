package chatbotcoolq

import (
	"context"

	qqbotapi "github.com/catsworld/qq-bot-api"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/pb"
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

			if serv.cfg.Debug {
				chatbotbase.Info("recv message",
					chatbotbase.JSON("msg", &update))
			}

			// chatbotbase.Info(update.Message.Text,
			// 	zap.Int64("ID", update.Message.Chat.ID),
			// 	zap.String("Type", update.Message.Chat.Type))

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
func (serv *Serv) onMsg(ctx context.Context, upd *qqbotapi.Update) error {
	if upd.Message != nil {
		strID := chatbotbase.ID642Str(upd.Message.Chat.ID)

		if upd.Message.Chat.IsPrivate() {
			uai := chatbot.BuildUserAppInfo(chatbotpb.ChatAppType_CAT_COOLQ,
				serv.cfg.Username, strID, strID, "zh-hans")

			str := chatbotbase.FormatCommand(upd.Message.Text)

			if str != "" {
				msg := chatbot.BuildTextChatMsg(str,
					uai, serv.cfg.Token, serv.client.SessionID)

				return serv.procChatMsg(ctx, msg)
			}
		}
	}

	return nil
}

// SendChatMsg - send a chat message
func (serv *Serv) SendChatMsg(ctx context.Context, chat *chatbotpb.ChatMsg) error {
	i64, err := chatbotbase.Str2ID64(chat.Uai.Appuid)
	if err != nil {
		return err
	}

	_, err = serv.bot.SendMessage(i64, "private", chat.Msg)
	if err != nil {
		return err
	}

	return nil
}

func (serv *Serv) procChatMsg(ctx context.Context, chat *chatbotpb.ChatMsg) error {
	lstret, err := serv.client.SendChat(ctx, chat)
	if err != nil {
		return err
	}

	for _, v := range lstret {
		err = serv.SendChatMsg(ctx, v)
		if err != nil {
			return err
		}
	}

	return nil
}
