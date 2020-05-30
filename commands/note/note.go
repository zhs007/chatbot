package chatbotcmdnote

import (
	"context"
	"strings"

	chatbotdb "github.com/zhs007/chatbot/db"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/pflag"
	"github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// CmdName - cmd name
const CmdName = "note"

// NoteMode - note mode
type NoteMode int

const (
	// NoteModeNone - none
	NoteModeNone NoteMode = 0
	// NoteModeNew - new note
	NoteModeNew NoteMode = 1
	// NoteModeForward - taking notes for forward
	NoteModeForward NoteMode = 2
	// NoteModeInfo - get infomation for note
	NoteModeInfo NoteMode = 3
	// NoteModeSearch - search note with key
	NoteModeSearch NoteMode = 4
	// NoteModeKeys - show keys
	NoteModeKeys NoteMode = 5
	// NoteModeRemoveKeys - remove keys
	NoteModeRemoveKeys NoteMode = 6
)

// ParseNoteMode - string => NoteMode
func ParseNoteMode(str string) NoteMode {
	if str == "new" {
		return NoteModeNew
	} else if str == "forward" {
		return NoteModeForward
	} else if str == "info" {
		return NoteModeInfo
	} else if str == "search" {
		return NoteModeSearch
	} else if str == "keys" {
		return NoteModeKeys
	} else if str == "rmkeys" {
		return NoteModeRemoveKeys
	}

	return NoteModeNone
}

type paramsCmd struct {
	mode NoteMode
	name string
	keys []string
}

// cmdNote - command note
type cmdNote struct {
	isRunning bool
	keys      []string
	noteName  string
}

// onNew - run command
func (cmd *cmdNote) onNew(ctx context.Context, serv *chatbot.Serv, params paramsCmd,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	if !chatbotdb.IsValidNoteName(params.name) {
		msg := &chatbotpb.ChatMsg{
			Msg: "Please input a valid name for note, it consists only of lowercase letters and numbers.",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	ni := &chatbotpb.NoteInfo{
		Name:     strings.ToLower(params.name),
		Masters:  []int64{ui.Uid},
		IsPublic: true,
	}

	err := serv.MgrUser.UpdNoteInfo(ctx, ni)
	if err != nil {
		chatbotbase.Error("cmdNote.onNew:UpdNoteInfo",
			zap.Error(err))

		return nil, err
	}

	msg := &chatbotpb.ChatMsg{
		Msg: "The note (" + params.name + ") is created.",
		Uai: chat.Uai,
	}

	return []*chatbotpb.ChatMsg{msg}, nil
}

// onInfo - run command
func (cmd *cmdNote) onInfo(ctx context.Context, serv *chatbot.Serv, params paramsCmd,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	if !chatbotdb.IsValidNoteName(params.name) {
		msg := &chatbotpb.ChatMsg{
			Msg: "Please input a valid name for note, it consists only of lowercase letters and numbers.",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	params.name = strings.ToLower(params.name)

	ni, err := serv.MgrUser.GetNoteInfo(ctx, params.name)
	if err != nil {
		chatbotbase.Error("cmdNote.onInfo:GetNoteInfo",
			zap.Error(err))

		return nil, err
	}

	if ni == nil {
		msg := &chatbotpb.ChatMsg{
			Msg: "Sorry, I can't find the note (" + params.name + ")",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	strni, err := chatbotbase.JSONFormat(ni)
	if err != nil {
		chatbotbase.Error("cmdNote.onInfo:JSONFormat",
			zap.Error(err))

		return nil, err
	}

	msg := &chatbotpb.ChatMsg{
		Msg: "The note (" + params.name + ") is \n " + strni + ".",
		Uai: chat.Uai,
	}

	return []*chatbotpb.ChatMsg{msg}, nil
}

// onForward - run command
func (cmd *cmdNote) onForward(ctx context.Context, serv *chatbot.Serv, params paramsCmd,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	if !chatbotdb.IsValidNoteName(params.name) {
		msg := &chatbotpb.ChatMsg{
			Msg: "Please input a valid name for note, it consists only of lowercase letters and numbers.",
			Uai: chat.Uai,
		}

		return true, []*chatbotpb.ChatMsg{msg}, nil
	}

	cmd.isRunning = true
	cmd.keys = params.keys
	cmd.noteName = params.name

	msg := &chatbotpb.ChatMsg{
		Msg: "note starts ...",
		Uai: chat.Uai,
	}

	return false, []*chatbotpb.ChatMsg{msg}, nil
}

// onSearch - run command
func (cmd *cmdNote) onSearch(ctx context.Context, serv *chatbot.Serv, params paramsCmd,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	if !chatbotdb.IsValidNoteName(params.name) {
		msg := &chatbotpb.ChatMsg{
			Msg: "Please input a valid name for note, it consists only of lowercase letters and numbers.",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	params.name = strings.ToLower(params.name)

	ni, err := serv.MgrUser.GetNoteInfo(ctx, params.name)
	if err != nil {
		chatbotbase.Error("cmdNote.onSearch:GetNoteInfo",
			zap.Error(err))

		return nil, err
	}

	if ni == nil {
		msg := &chatbotpb.ChatMsg{
			Msg: "Sorry, I can't find the note (" + params.name + ")",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	lsti64 := chatbotdb.SearchKeys(ni, params.keys)
	var lstmsg []*chatbotpb.ChatMsg

	for _, v := range lsti64 {
		nn, err := serv.MgrUser.GetNoteNode(ctx, ni.Name, v)
		if err != nil {
			chatbotbase.Error("cmdNote.onSearch:GetNoteNode",
				zap.Error(err))
		}

		if nn != nil {
			msg := &chatbotpb.ChatMsg{
				Uai: chat.Uai,
				Forward: &chatbotpb.ForwardData{
					Uai:      nn.Uai,
					AppMsgID: nn.ForwardAppMsgID,
				},
			}

			if msg.Forward.Uai == nil {
				msg.Forward.Uai = chat.Uai
			}

			lstmsg = append(lstmsg, msg)
		}
	}

	return lstmsg, nil
}

// onKeys - run command
func (cmd *cmdNote) onKeys(ctx context.Context, serv *chatbot.Serv, params paramsCmd,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	if !chatbotdb.IsValidNoteName(params.name) {
		msg := &chatbotpb.ChatMsg{
			Msg: "Please input a valid name for note, it consists only of lowercase letters and numbers.",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	params.name = strings.ToLower(params.name)

	ni, err := serv.MgrUser.GetNoteInfo(ctx, params.name)
	if err != nil {
		chatbotbase.Error("cmdNote.onSearch:GetNoteInfo",
			zap.Error(err))

		return nil, err
	}

	if ni == nil {
		msg := &chatbotpb.ChatMsg{
			Msg: "Sorry, I can't find the note (" + params.name + ")",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	str, err := chatbotbase.JSONFormat(ni.Keys)

	msg := &chatbotpb.ChatMsg{
		Msg: str,
		Uai: chat.Uai,
	}

	return []*chatbotpb.ChatMsg{msg}, nil
}

// onRemoveKeys - run command
func (cmd *cmdNote) onRemoveKeys(ctx context.Context, serv *chatbot.Serv, params paramsCmd,
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) ([]*chatbotpb.ChatMsg, error) {

	if !chatbotdb.IsValidNoteName(params.name) {
		msg := &chatbotpb.ChatMsg{
			Msg: "Please input a valid name for note, it consists only of lowercase letters and numbers.",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	params.name = strings.ToLower(params.name)

	ni, err := serv.MgrUser.GetNoteInfo(ctx, params.name)
	if err != nil {
		chatbotbase.Error("cmdNote.onSearch:GetNoteInfo",
			zap.Error(err))

		return nil, err
	}

	if ni == nil {
		msg := &chatbotpb.ChatMsg{
			Msg: "Sorry, I can't find the note (" + params.name + ")",
			Uai: chat.Uai,
		}

		return []*chatbotpb.ChatMsg{msg}, nil
	}

	ni = chatbotdb.RemoveKeys(ni, params.keys)

	str, err := chatbotbase.JSONFormat(ni.Keys)

	msg := &chatbotpb.ChatMsg{
		Msg: str,
		Uai: chat.Uai,
	}

	return []*chatbotpb.ChatMsg{msg}, nil
}

// RunCommand - run command
func (cmd *cmdNote) RunCommand(ctx context.Context, serv *chatbot.Serv, params interface{},
	chat *chatbotpb.ChatMsg, ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	cmd.isRunning = false

	if serv == nil {
		return true, nil, chatbotbase.ErrCmdInvalidServ
	}

	if params == nil {
		return true, nil, ErrCmdNoParams
	}

	cp, isok := params.(paramsCmd)
	if !isok {
		return true, nil, ErrCmdInvalidParams
	}

	if cp.mode == NoteModeNone {
		return true, nil, nil
	}

	if cp.mode == NoteModeNew {
		lst, err := cmd.onNew(ctx, serv, cp, chat, ui, ud, scs)
		return true, lst, err
	}

	if cp.mode == NoteModeInfo {
		lst, err := cmd.onInfo(ctx, serv, cp, chat, ui, ud, scs)
		return true, lst, err
	}

	if cp.mode == NoteModeForward {
		return cmd.onForward(ctx, serv, cp, chat, ui, ud, scs)
	}

	if cp.mode == NoteModeSearch {
		lst, err := cmd.onSearch(ctx, serv, cp, chat, ui, ud, scs)
		return true, lst, err
	}

	if cp.mode == NoteModeKeys {
		lst, err := cmd.onKeys(ctx, serv, cp, chat, ui, ud, scs)
		return true, lst, err
	}

	if cp.mode == NoteModeRemoveKeys {
		lst, err := cmd.onRemoveKeys(ctx, serv, cp, chat, ui, ud, scs)
		return true, lst, err
	}

	return true, nil, ErrCmdInvalidNoteMode
}

// OnMessage - get message
func (cmd *cmdNote) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message,
	scs chatbotpb.ChatBotService_SendChatServer) (bool, []*chatbotpb.ChatMsg, error) {

	if !cmd.isRunning {
		return true, nil, chatbotbase.ErrCmdItsNotMine
	}

	if chat.Forward == nil {
		msg := &chatbotpb.ChatMsg{
			Msg: "note end.",
			Uai: chat.Uai,
		}

		return true, []*chatbotpb.ChatMsg{msg}, chatbotbase.ErrCmdItsNotMine
	}

	text := chat.Msg
	if text == "" {
		text = chat.Caption
	}

	err := serv.MgrUser.UpdNoteNode(ctx, &chatbotpb.NoteNode{
		Keys:            cmd.keys,
		Name:            cmd.noteName,
		ForwardAppMsgID: chat.AppMsgID,
		Text:            text,
		Uai:             chat.Uai,
	})
	if err != nil {
		chatbotbase.Error("cmdNote.OnMessage:UpdNoteNode",
			zap.Error(err))

		return false, nil, err
	}

	msg := &chatbotpb.ChatMsg{
		Uai: chat.Uai,
		Forward: &chatbotpb.ForwardData{
			Uai:      chat.Uai,
			AppMsgID: chat.AppMsgID,
		},
	}

	return false, []*chatbotpb.ChatMsg{msg}, nil
}

// ParseCommandLine - parse command line
func (cmd *cmdNote) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	interface{}, error) {

	flagset := pflag.NewFlagSet(CmdName, pflag.ContinueOnError)

	strMode := flagset.StringP("mode", "m", "none", "mode")
	strName := flagset.StringP("name", "n", "", "name")
	keys := flagset.StringArrayP("keys", "k", []string{}, "key")

	err := flagset.Parse(cmdline[1:])
	if err != nil {
		return nil, err
	}

	return paramsCmd{
		mode: ParseNoteMode(*strMode),
		name: *strName,
		keys: *keys,
	}, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand(CmdName, &cmdNote{})
}
