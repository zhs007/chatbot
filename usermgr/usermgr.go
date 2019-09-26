package chatbotusermgr

import (
	"context"

	chatbotdb "github.com/zhs007/chatbot/db"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// UserMgr - user manager
type UserMgr struct {
	db *chatbotdb.UserDB
}

// NewUserMgr - new UserMgr
func NewUserMgr(dbpath string, httpAddr string, engine string) (*UserMgr, error) {
	db, err := chatbotdb.NewUserDB(dbpath, httpAddr, engine)
	if err != nil {
		return nil, err
	}

	mgr := &UserMgr{
		db: db,
	}

	return mgr, nil
}

// GetAppUserInfo - get user infomation
func (mgr *UserMgr) GetAppUserInfo(ctx context.Context, appToken string, uai *chatbotpb.UserAppInfo) (*chatbotpb.UserInfo, error) {
	ui, err := mgr.db.GetUserInfoEx(ctx, appToken, uai.Appuid)
	if err != nil {
		return nil, err
	}

	if ui == nil {
		uid, err := mgr.db.GetUserID(ctx)
		if err != nil {
			return nil, err
		}

		ui = &chatbotpb.UserInfo{
			Uid:   uid,
			Name:  uai.Appuname,
			Apps:  []*chatbotpb.UserAppInfo{uai},
			Money: 0,
		}

		err = mgr.db.UpdAppUID(ctx, appToken, uai.Appuid, uid)
		if err != nil {
			return nil, err
		}

		err = mgr.db.UpdUser(ctx, ui)
		if err != nil {
			return nil, err
		}

		return ui, nil
	}

	return ui, nil
}
