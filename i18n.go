package chatbot

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"go.uber.org/zap"
	"golang.org/x/text/language"

	"gopkg.in/yaml.v2"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	chatbotbase "github.com/zhs007/chatbot/base"
)

// TextMgr - text manager
type TextMgr struct {
	bundle *i18n.Bundle
	mapL   sync.Map // map[lang]*i18n.Localizer
	lang   string
	keys   []string
}

func loadLangKeys(fn string) ([]string, error) {
	fi, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return nil, err
	}

	mapkeys := make(map[string]interface{})

	err = yaml.Unmarshal(fd, &mapkeys)
	if err != nil {
		return nil, err
	}

	var keys []string

	for k := range mapkeys {
		keys = append(keys, k)
	}

	return keys, nil
}

// NewTextMgr - new TextMgr
func NewTextMgr(cfg *Config) (*TextMgr, error) {
	lt, err := language.Parse(cfg.Language)
	if err != nil {
		return nil, err
	}

	b := i18n.NewBundle(lt)
	b.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	dir, err := ioutil.ReadDir(cfg.LangPath)
	if err != nil {
		return nil, err
	}

	mgr := &TextMgr{
		bundle: b,
		lang:   cfg.Language,
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			ok := strings.HasSuffix(fi.Name(), ".yaml")
			if ok {
				fn := path.Join(cfg.LangPath, fi.Name())
				_, err = b.LoadMessageFile(fn)
				if err != nil {
					return nil, err
				}

				if mgr.keys != nil {
					arr, err := loadLangKeys(fn)
					if err != nil {
						return nil, err
					}

					mgr.keys = arr
				}
			}
		}
	}

	return mgr, nil
}

// GetLocalizer - get Localizer
func (mgr *TextMgr) GetLocalizer(lang string) (*i18n.Localizer, error) {
	_, err := language.Parse(lang)
	if err != nil {
		chatbotbase.Warn("TextMgr.GetLocalizer",
			zap.Error(err),
			zap.String("lang", lang))

		lang = mgr.lang
		// return nil, err
	}

	li, isok := mgr.mapL.Load(lang)
	if !isok {
		l := i18n.NewLocalizer(mgr.bundle, lang)

		mgr.mapL.Store(lang, l)

		return l, nil
	}

	l, isok := li.(*i18n.Localizer)
	if !isok || l == nil {
		l = i18n.NewLocalizer(mgr.bundle, lang)

		mgr.mapL.Store(lang, l)

		return l, nil
	}

	return l, nil
}

// FindKeys - find keys
func (mgr *TextMgr) FindKeys(prefix string) []string {
	var keys []string

	for _, v := range mgr.keys {
		if strings.Index(v, prefix) == 0 {
			keys = append(keys, v)
		}
	}

	return keys
}
