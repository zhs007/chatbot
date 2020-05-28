package chatbottelegram

import (
	"bytes"
	"context"
	"net/http"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

// Serv - serv
type Serv struct {
	cfg    *Config
	bot    *tgbotapi.BotAPI
	client *chatbot.Client
	ticker *time.Ticker
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
		ticker: time.NewTicker(time.Second),
	}

	return serv, nil
}

// Start - start telegram bot
func (serv *Serv) Start(ctx context.Context) error {
	defer serv.ticker.Stop()

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

	curseconds := 0

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
		case <-serv.ticker.C:
			curseconds++

			if curseconds >= 5 {
				curseconds = 0

				serv.procTicker(ctx)
			}
		}

		if isend {
			break
		}
	}

	return nil
}

func (serv *Serv) buildChatMsg(msg *tgbotapi.Message) (*chatbotpb.ChatMsg, error) {
	var cmsg *chatbotpb.ChatMsg

	from := msg.From
	uai := chatbot.BuildUserAppInfo(chatbotpb.ChatAppType_CAT_TELEGRAM,
		serv.cfg.Username, chatbotbase.ID2Str(from.ID), from.UserName, from.LanguageCode)

	if msg.Document != nil {
		str := chatbotbase.FormatCommand(msg.Caption)

		if msg.ForwardDate > 0 {
			cmsg = chatbot.BuildTextChatMsg(str, uai, serv.cfg.Token, serv.client.SessionID)
		} else {
			fd, err := serv.getFileDataWithDocument(msg.Document)
			if err != nil {
				return nil, err
			}

			cmsg = chatbot.BuildFileChatMsg(str, fd, uai, serv.cfg.Token, serv.client.SessionID)
		}
	} else if msg.Photo != nil {
		str := chatbotbase.FormatCommand(msg.Caption)

		// fd, err := serv.getFileDataWithDocument(msg.Document)
		// if err != nil {
		// 	return nil, err
		// }

		cmsg = chatbot.BuildTextChatMsg(str, uai, serv.cfg.Token, serv.client.SessionID)

		// cmsg = chatbot.BuildFileChatMsg(str, fd, uai, serv.cfg.Token, serv.client.SessionID)
	} else {
		str := chatbotbase.FormatCommand(msg.Text)

		cmsg = chatbot.BuildTextChatMsg(str, uai, serv.cfg.Token, serv.client.SessionID)
	}

	cmsg.AppMsgID = strconv.Itoa(msg.MessageID)

	if msg.ForwardDate > 0 {
		cmsg.Forward = &chatbotpb.ForwardData{
			Date: int64(msg.ForwardDate),
		}

		if msg.ForwardFrom != nil {
			cmsg.Forward.Uai = chatbot.BuildUserAppInfo(chatbotpb.ChatAppType_CAT_TELEGRAM,
				serv.cfg.Username, chatbotbase.ID2Str(msg.ForwardFrom.ID), msg.ForwardFrom.UserName, msg.ForwardFrom.LanguageCode)
		}
	}

	return cmsg, nil
}

// onMsg - on message
func (serv *Serv) onMsg(ctx context.Context, upd *tgbotapi.Update) error {
	if upd.Message != nil {
		chatbotbase.Debug("Serv:onMsg",
			chatbotbase.JSON("msg", upd.Message))

		// from := upd.Message.From
		// uai := chatbot.BuildUserAppInfo(chatbotpb.ChatAppType_CAT_TELEGRAM,
		// 	serv.cfg.Username, chatbotbase.ID2Str(from.ID), from.UserName, from.LanguageCode)

		msg, err := serv.buildChatMsg(upd.Message)
		if err != nil {
			chatbotbase.Error("Serv:onMsg:buildChatMsg",
				zap.Error(err))

			return err
		}

		if msg != nil {
			// if msg.Forward != nil {
			// 	serv.ForwardMsg(ctx, &chatbotpb.ChatMsg{
			// 		Uai: msg.Uai,
			// 		Forward: &chatbotpb.ForwardData{
			// 			Uai:      msg.Uai,
			// 			AppMsgID: msg.AppMsgID,
			// 		},
			// 	})
			// }

			return serv.procChatMsg(ctx, msg)
		}

		// str := chatbotbase.FormatCommand(upd.Message.Text)

		// chatbotbase.Info("onMsg",
		// 	zap.String("Text", upd.Message.Text),
		// 	zap.String("lang", from.LanguageCode))

		// if upd.Message.Document != nil {
		// 	fd, err := serv.getFileDataWithDocument(upd.Message.Document)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	msg := chatbot.BuildFileChatMsg(str, fd, uai, serv.cfg.Token, serv.client.SessionID)

		// 	return serv.procChatMsg(ctx, msg)
		// }

		// if isValidMsg(upd.Message) {

		// }
		// if str != "" && isValidMsg(upd.Message) {
		// 	msg := chatbot.BuildTextChatMsg(str,
		// 		uai, serv.cfg.Token, serv.client.SessionID)

		// 	return serv.procChatMsg(ctx, msg)
		// }
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

	telemsg.DisableWebPagePreview = !serv.cfg.PreviewWebPage

	_, err = serv.bot.Send(telemsg)
	if err != nil {
		return err
	}

	return nil
}

// ForwardMsg - forward a chat message
func (serv *Serv) ForwardMsg(ctx context.Context, chat *chatbotpb.ChatMsg) error {
	if chat.Forward == nil {
		return serv.SendChatMsg(ctx, chat)
	}

	i64, err := chatbotbase.Str2ID64(chat.Uai.Appuid)
	if err != nil {
		return err
	}

	fromi64, err := chatbotbase.Str2ID64(chat.Forward.Uai.Appuid)
	if err != nil {
		return err
	}

	msgid, err := strconv.Atoi(chat.Forward.AppMsgID)

	telemsg := tgbotapi.NewForward(i64, fromi64, msgid)

	_, err = serv.bot.Send(telemsg)
	if err != nil {
		chatbotbase.Error("Serv.ForwardMsg:Send",
			chatbotbase.JSON("forward", telemsg),
			zap.Error(err))

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

	chatbotbase.Info("getFileDataWithDocument",
		zap.String("url", url))

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

	chatbotbase.Info("getFileDataWithDocument",
		zap.Int("length", len(fd.FileData)),
		zap.Int("FileType", len(fd.FileType)))

	return fd, nil
}

func (serv *Serv) procChatMsg(ctx context.Context, chat *chatbotpb.ChatMsg) error {
	lstret, err := serv.client.SendChat(ctx, chat)
	if err != nil {
		chatbotbase.Error("procChatMsg:SendChat",
			zap.Error(err))

		return err
	}

	for _, v := range lstret {
		if v.Forward != nil {
			err = serv.ForwardMsg(ctx, v)

			if err != nil {
				chatbotbase.Error("procChatMsg:ForwardMsg",
					zap.Error(err))

				return err
			}
		} else {
			err = serv.SendChatMsg(ctx, v)
			if err != nil {
				chatbotbase.Error("procChatMsg:SendChatMsg",
					zap.Error(err))

				return err
			}
		}
	}

	return nil
}

func (serv *Serv) procTicker(ctx context.Context) error {
	lstret, err := serv.client.RequestChat(ctx)
	if err != nil {
		chatbotbase.Error("procTicker:RequestChat",
			zap.Error(err))

		return err
	}

	for _, v := range lstret {
		err = serv.SendChatMsg(ctx, v)
		if err != nil {
			chatbotbase.Error("procTicker:SendChatMsg",
				zap.Error(err))

			return err
		}
	}

	return nil
}
