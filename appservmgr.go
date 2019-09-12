package chatbot

import "sync"

// AppServMgr - appserv manager
type AppServMgr struct {
	MapAppServ sync.Map
}

// Init - init with config
func (mgr *AppServMgr) Init(cfg *Config) error {
	for _, v := range cfg.AppServ {
		mgr.setAppServ(&v)
	}

	return nil
}

func (mgr *AppServMgr) setAppServ(cfgAppServ *AppServConfig) {
	mgr.MapAppServ.Store(cfgAppServ.Token, cfgAppServ)
}

func (mgr *AppServMgr) getAppServ(token string) *AppServConfig {
	v, isok := mgr.MapAppServ.Load(token)
	if !isok {
		return nil
	}

	cfgAppServ, isok := v.(*AppServConfig)
	if !isok {
		return nil
	}

	return cfgAppServ
}
