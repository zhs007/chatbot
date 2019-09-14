package chatbotbase

// AppServConfig - app serv
type AppServConfig struct {
	Type     string
	Token    string
	UserName string
	// typeAppServ chatbotpb.ChatAppType
	// sessionID   string
}

// CheckAppServConfig - check app service config
func CheckAppServConfig(cfg *AppServConfig) error {
	if cfg.Type == "" {
		return ErrNoAppServType
	}

	if cfg.Token == "" {
		return ErrNoAppServToken
	}

	if cfg.UserName == "" {
		return ErrNoAppServUserName
	}

	_, err := GetAppServType(cfg.Type)
	if err != nil {
		return err
	}

	// cfg.typeAppServ = t

	return nil
}

// FindAppServConfig - find a app service config
func FindAppServConfig(token string, lst []AppServConfig) *AppServConfig {
	for _, v := range lst {
		if v.Token == token {
			return &v
		}
	}

	return nil
}
