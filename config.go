package chatbot

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	chatbotbase "github.com/zhs007/chatbot/base"
)

// TelegramConfig - config for telegram
type TelegramConfig struct {
	Channels []string
}

// Config - config
type Config struct {

	//------------------------------------------------------------------
	// appserv

	AppServ []chatbotbase.AppServConfig

	//------------------------------------------------------------------
	// Database Config

	BindAddr string
	DBPath   string
	DBEngine string

	//------------------------------------------------------------------
	// plugins

	PluginsPreprocess []string
	Plugins           []string
	PluginsSecondLine []string

	//------------------------------------------------------------------
	// connands

	Commands []string

	//------------------------------------------------------------------
	// config

	Language        string
	LangPath        string
	StartText       []string
	HelpText        []string
	ChatBotNameText string

	//------------------------------------------------------------------
	// plugin config

	Preprocessor string
	NoteKeyword  string

	//------------------------------------------------------------------
	// telegram

	Telegram *TelegramConfig
}

func checkConfig(cfg *Config) error {
	for _, v := range cfg.AppServ {
		err := chatbotbase.CheckAppServConfig(&v)
		if err != nil {
			return err
		}
	}

	if cfg.BindAddr == "" {
		return chatbotbase.ErrNoBindAddr
	}

	if cfg.DBPath == "" {
		return chatbotbase.ErrNoDBPath
	}

	if cfg.DBEngine == "" {
		return chatbotbase.ErrNoDBEngine
	}

	if cfg.Language == "" {
		cfg.Language = "en-US"
	}

	if cfg.LangPath == "" {
		cfg.LangPath = "./lang"
	}

	return nil
}

// LoadConfig - load config
func LoadConfig(filename string) (*Config, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return nil, err
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
