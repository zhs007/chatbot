package chatbot

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	chatbotbase "github.com/zhs007/chatbot/base"
)

// AppServConfig - app serv
type AppServConfig struct {
	Type     string
	Token    string
	UserName string
}

// Config - config
type Config struct {
	//------------------------------------------------------------------
	// appserv

	AppServ []AppServConfig
}

func checkAppServConfig(cfg *AppServConfig) error {
	if cfg.Type == "" {
		return chatbotbase.ErrNoAppServType
	}

	if cfg.Token == "" {
		return chatbotbase.ErrNoAppServToken
	}

	if cfg.UserName == "" {
		return chatbotbase.ErrNoAppServUserName
	}

	_, err := chatbotbase.GetAppServType(cfg.Type)
	if err != nil {
		return err
	}

	return nil
}

func checkConfig(cfg *Config) error {
	for _, v := range cfg.AppServ {
		err := checkAppServConfig(&v)
		if err != nil {
			return err
		}
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
