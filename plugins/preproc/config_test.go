package chatbotproprocplugin

import (
	"testing"
)

func Test_LoadConfig001(t *testing.T) {
	cfg, err := LoadConfig("../../testingdata/preprocessor001.yaml")
	if err != nil {
		t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml %v", err)
	}

	if cfg != nil {
		if len(cfg.LstRegexp) != 2 {
			t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml LstRegexp %v",
				len(cfg.LstRegexp))
		}

		if cfg.LstRegexp[0].Pattern != "找漫画(.*)" {
			t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml LstRegexp[0].Pattern %v",
				cfg.LstRegexp[0].Pattern)
		}

		if cfg.LstRegexp[0].Prefix != "note -m search -n comic " {
			t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml LstRegexp[0].Prefix %v",
				cfg.LstRegexp[0].Prefix)
		}

		if cfg.LstRegexp[0].Mode != "paramarray" {
			t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml LstRegexp[0].Mode %v",
				cfg.LstRegexp[0].Mode)
		}

		if cfg.LstRegexp[0].ParamArrayPrefix != " -k " {
			t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml LstRegexp[0].ParamArrayPrefix %v",
				cfg.LstRegexp[0].ParamArrayPrefix)
		}

		if cfg.LstRegexp[1].Pattern != "找(.*)" {
			t.Fatalf("Test_LoadConfig001 LoadConfig preprocessor001.yaml LstRegexp[1].Pattern %v",
				cfg.LstRegexp[0].Pattern)
		}

		// arr := cfg.LstRegexp[0].r.FindAllStringSubmatchIndex("找漫画北斗神拳", -1)
		// // t.Fatalf("%v", arr)

		// arr = cfg.LstRegexp[1].r.FindAllStringSubmatchIndex("找漫画北斗神拳", -1)
		// t.Fatalf("%v", arr)
	}

	t.Log("Test_LoadConfig001 OK")
}
