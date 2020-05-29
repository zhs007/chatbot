package chatbotproprocplugin

import (
	"io/ioutil"
	"os"
	"regexp"

	chatbotbase "github.com/zhs007/chatbot/base"
	"go.uber.org/zap"

	"gopkg.in/yaml.v2"
)

// RegexpNode - RegexpNode
type RegexpNode struct {
	Pattern          string
	Prefix           string
	Mode             string
	ParamArrayPrefix string

	r *regexp.Regexp
}

// Config - config
type Config struct {
	LstRegexp []*RegexpNode
}

// LoadConfig - load config
func LoadConfig(fn string) (*Config, error) {
	fi, err := os.Open(fn)
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

	for _, v := range cfg.LstRegexp {
		r, err := regexp.Compile(v.Pattern)
		if err != nil {
			chatbotbase.Error("chatbotproprocplugin.LoadConfig:Compile",
				zap.String("pattern", v.Pattern),
				zap.Error(err))

			return nil, err
		}

		v.r = r
	}

	return cfg, nil
}
