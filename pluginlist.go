package chatbot

import (
	chatbotbase "github.com/zhs007/chatbot/base"
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
