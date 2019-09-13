package chatbot

import (
	chatbotbase "github.com/zhs007/chatbot/base"
)

// pluginsMgr - plugins manager
type PluginsMgr struct {
	mapPlugin map[string]Plugin
}

var mgrPlugins *PluginsMgr

func init() {
	mgrPlugins = &PluginsMgr{
		mapPlugin: make(map[string]Plugin),
	}
}

// RegPlugin - register plugin
func RegPlugin(plugin Plugin) error {
	op := GetStaticPlugin(plugin.GetPluginName())
	if op != nil {
		return chatbotbase.ErrDuplicatePlugin
	}

	mgrPlugins.mapPlugin[plugin.GetPluginName()] = plugin

	return nil
}

// GetStaticPlugin - get static plugin
func GetStaticPlugin(name string) Plugin {
	v, isok := mgrPlugins.mapPlugin[name]
	if !isok {
		return nil
	}

	return v
}
