package chatbotusermgr

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotdb "github.com/zhs007/chatbot/db"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// UserMgr - user manager
type UserMgr struct {
	db *chatbotdb.UserDB
}

// NewUserMgr - new UserMgr
func NewUserMgr(dbpath string, httpAddr string, engine string, mgrUserData chatbotdb.UserDataMgr) (*UserMgr, error) {
	db, err := chatbotdb.NewUserDB(dbpath, httpAddr, engine, mgrUserData)
	if err != nil {
		return nil, err
	}

	mgr := &UserMgr{
		db: db,
	}

	return mgr, nil
}

// GetAppUserInfo - get user infomation
func (mgr *UserMgr) GetAppUserInfo(ctx context.Context, appToken string, uai *chatbotpb.UserAppInfo) (
	*chatbotpb.UserInfo, proto.Message, error) {

	ui, err := mgr.db.GetUserInfoEx(ctx, appToken, uai.Appuid)
	if err != nil {
		return nil, nil, err
	}

	if ui == nil {
		uid, err := mgr.db.GetUserID(ctx)
		if err != nil {
			return nil, nil, err
		}

		ui = &chatbotpb.UserInfo{
			Uid:   uid,
			Name:  uai.Appuname,
			Apps:  []*chatbotpb.UserAppInfo{uai},
			Money: 0,
		}

		err = mgr.db.UpdAppUID(ctx, appToken, uai.Appuid, uid)
		if err != nil {
			return nil, nil, err
		}

		err = mgr.db.UpdUser(ctx, ui)
		if err != nil {
			return nil, nil, err
		}

		if mgr.db.MgrUserData != nil {
			ud, err := mgr.db.MgrUserData.New(ui)
			if err != nil {
				return nil, nil, err
			}

			if ud == nil {
				return nil, nil, chatbotbase.ErrUserMgrNoUserData
			}

			err = mgr.db.UpdUserData(ctx, uid, ud)
			if err != nil {
				return nil, nil, err
			}

			return ui, ud, nil
		}

		return ui, nil, nil
	}

	if mgr.db.MgrUserData != nil {
		ud, err := mgr.db.MgrUserData.New(ui)
		if err != nil {
			return nil, nil, err
		}

		if ud == nil {
			return nil, nil, chatbotbase.ErrUserMgrNoUserData
		}

		err = mgr.db.UpdUserData(ctx, ui.Uid, ud)
		if err != nil {
			return nil, nil, err
		}

		return ui, ud, nil
	}

	return ui, nil, nil
}
