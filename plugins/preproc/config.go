package chatbotproprocplugin

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// RegexpNode - RegexpNode
type RegexpNode struct {
	Pattern          string
	Prefix           string
	Mode             string
	ParamArrayPrefix string
}

// Config - config
type Config struct {
	LstRegexp []RegexpNode
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

	return cfg, nil
}
