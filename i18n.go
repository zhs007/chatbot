package chatbot

import (
	"io/ioutil"
	"path"
	"strings"
	"sync"

	"golang.org/x/text/language"

	"gopkg.in/yaml.v2"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// TextMgr - text manager
type TextMgr struct {
	bundle *i18n.Bundle
	mapL   sync.Map
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

	for _, fi := range dir {
		if !fi.IsDir() {
			ok := strings.HasSuffix(fi.Name(), ".yaml")
			if ok {
				fn := path.Join(cfg.LangPath, fi.Name())
				_, err = b.LoadMessageFile(fn)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return &TextMgr{
		bundle: b,
	}, nil
}

// GetLocalizer - get Localizer
func (mgr *TextMgr) GetLocalizer(lang string) (*i18n.Localizer, error) {
	_, err := language.Parse(lang)
	if err != nil {
		return nil, err
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
