package chatbottelegram

import (
	"bytes"
	"context"
	"net/http"

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
			serv.cfg.Username, chatbotbase.ID2Str(from.ID), from.UserName, from.LanguageCode)

		str := chatbotbase.FormatCommand(upd.Message.Text)

		chatbotbase.Info("onMsg",
			zap.String("Text", upd.Message.Text))

		if upd.Message.Document != nil {
			fd, err := serv.getFileDataWithDocument(upd.Message.Document)
			if err != nil {
				return err
			}

			msg := chatbot.BuildFileChatMsg(str, fd, uai, serv.cfg.Token, serv.client.SessionID)

			return serv.procChatMsg(ctx, msg)
		}

		if str != "" {
			msg := chatbot.BuildTextChatMsg(str,
				uai, serv.cfg.Token, serv.client.SessionID)

			return serv.procChatMsg(ctx, msg)
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

	telemsg := tgbotapi.NewMessage(i64, chat.Msg)

	telemsg.DisableWebPagePreview = true

	_, err = serv.bot.Send(telemsg)
	if err != nil {
		return err
	}

	return nil
}

func (serv *Serv) getFileDataWithDocument(doc *tgbotapi.Document) (*chatbotpb.FileData, error) {
	chatbotbase.Info("getFileDataWithDocument",
		zap.String("FileID", doc.FileID))

	file, err := serv.bot.GetFile(tgbotapi.FileConfig{
		FileID: doc.FileID,
	})
	if err != nil {
		chatbotbase.Error("getFileDataWithDocument",
			zap.Error(err),
			zap.String("FileID", doc.FileID))

		return nil, err
	}

	url := file.Link(serv.bot.Token)

	res, err := http.Get(url)
	if err != nil {
		chatbotbase.Error("getFileDataWithDocument",
			zap.Error(err),
			zap.String("url", url))

		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	fd := &chatbotpb.FileData{
		Filename: doc.FileName,
		FileData: buf.Bytes(),
		FileType: doc.MimeType,
	}

	return fd, nil
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
