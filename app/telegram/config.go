package chatbottelegram

import (
	"io/ioutil"
	"os"

	chatbotbase "github.com/zhs007/chatbot/base"
	"gopkg.in/yaml.v2"
)

// Config - configuate
type Config struct {
	TelegramToken string
	Token         string
	ServAddr      string
	Username      string
}

func checkConfig(cfg *Config) error {
	if cfg.ServAddr == "" {
		return chatbotbase.ErrNoServAddr
	}

	if cfg.Username == "" {
		return chatbotbase.ErrNoUsername
	}

	if cfg.ServAddr == "" {
		return chatbotbase.ErrNoServAddr
	}

	if cfg.TelegramToken == "" {
		return chatbotbase.ErrNoTelegramToken
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
