package chatbottelegram

import (
	"testing"
)

func Test_LoadConfig001(t *testing.T) {
	cfg, err := LoadConfig("../../testingdata/telegram001.yaml")
	if err != nil {
		t.Fatalf("Test_LoadConfig001 LoadConfig telegram001.yaml %v", err)
	}

	if cfg != nil {
		if cfg.TelegramToken != "1234567" {
			t.Fatalf("Test_LoadConfig001 LoadConfig telegram001.yaml TelegramToken %v", cfg.TelegramToken)
		}

		if cfg.Token != "1234567" {
			t.Fatalf("Test_LoadConfig001 LoadConfig telegram001.yaml Token %v", cfg.Token)
		}

		if cfg.ServAddr != "127.0.0.1:123" {
			t.Fatalf("Test_LoadConfig001 LoadConfig telegram001.yaml ServAddr %v", cfg.ServAddr)
		}

		if cfg.Username != "ada" {
			t.Fatalf("Test_LoadConfig001 LoadConfig telegram001.yaml Username %v", cfg.Username)
		}

		if cfg.PreviewWebPage != false {
			t.Fatalf("Test_LoadConfig001 LoadConfig telegram001.yaml PreviewWebPage %v", cfg.PreviewWebPage)
		}
	}

	t.Log("Test_LoadConfig001 OK")
}

func Test_LoadConfig002(t *testing.T) {
	cfg, err := LoadConfig("../../testingdata/telegram002.yaml")
	if err != nil {
		t.Fatalf("Test_LoadConfig002 LoadConfig telegram002.yaml %v", err)
	}

	if cfg != nil {
		if cfg.TelegramToken != "1234567" {
			t.Fatalf("Test_LoadConfig002 LoadConfig telegram002.yaml TelegramToken %v", cfg.TelegramToken)
		}

		if cfg.Token != "12345678" {
			t.Fatalf("Test_LoadConfig002 LoadConfig telegram002.yaml Token %v", cfg.Token)
		}

		if cfg.ServAddr != "127.0.0.1:123" {
			t.Fatalf("Test_LoadConfig002 LoadConfig telegram002.yaml ServAddr %v", cfg.ServAddr)
		}

		if cfg.Username != "ada" {
			t.Fatalf("Test_LoadConfig002 LoadConfig telegram002.yaml Username %v", cfg.Username)
		}

		if cfg.PreviewWebPage != true {
			t.Fatalf("Test_LoadConfig002 LoadConfig telegram002.yaml PreviewWebPage %v", cfg.PreviewWebPage)
		}
	}

	t.Log("Test_LoadConfig002 OK")
}
