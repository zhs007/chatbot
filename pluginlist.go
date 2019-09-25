package chatbot

import (
	"context"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// PluginsList - plugins list
type PluginsList struct {
	plugins []Plugin
}

// NewPluginsList - new PluginsList
func NewPluginsList() *PluginsList {
	return &PluginsList{}
}

// AddPlugin - add plugin
func (lst *PluginsList) AddPlugin(name string) error {
	p := GetStaticPlugin(name)
	if p == nil {
		return chatbotbase.ErrNoPlugin
	}

	lst.plugins = append(lst.plugins, p)

	return nil
}

// FindPlugin -
func (lst *PluginsList) FindPlugin(name string) Plugin {
	for _, v := range lst.plugins {
		if v.GetPluginName() == name {
			return v
		}
	}

	return nil
}

// OnMessage - get message
func (lst *PluginsList) OnMessage(ctx context.Context, msg *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {
	for _, v := range lst.plugins {
		lst, err := v.OnMessage(ctx, msg)
		if err != nil {
			return nil, err
		}

		if lst != nil {
			return lst, nil
		}
	}

	return nil, nil
}

// OnStart - on start
func (lst *PluginsList) OnStart(ctx context.Context) error {
	for _, v := range lst.plugins {
		err := v.OnStart(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
