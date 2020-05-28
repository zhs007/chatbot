package chatbotdb

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	ankadb "github.com/zhs007/ankadb"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
	"go.uber.org/zap"
)

// UserDB - user database
type UserDB struct {
	ankaDB      ankadb.AnkaDB
	MgrUserData UserDataMgr
}

// NewUserDB - new UserDB
func NewUserDB(dbpath string, httpAddr string, engine string, mgrUserData UserDataMgr) (*UserDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   UserDBName,
		Engine: engine,
		PathDB: UserDBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		chatbotbase.Error("NewUserDB", zap.Error(err))

		return nil, err
	}

	chatbotbase.Info("NewAppServDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &UserDB{
		ankaDB:      ankaDB,
		MgrUserData: mgrUserData,
	}

	return db, err
}

// GetUserID - get userID
func (db *UserDB) GetUserID(ctx context.Context) (int64, error) {
	buf, err := db.ankaDB.Get(ctx, UserDBName, UserIDDBKey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			err = db.updUserID(ctx, MinUID)
			if err != nil {
				return -1, err
			}

			return MinUID, nil
		}

		return -1, err
	}

	uid := chatbotbase.BytesToInt64(buf)

	uid++
	err = db.updUserID(ctx, uid)
	if err != nil {
		return -1, err
	}

	return uid, nil
}

// updUserID - update userID
func (db *UserDB) updUserID(ctx context.Context, uid int64) error {
	buf := chatbotbase.Int64ToBytes(uid)

	err := db.ankaDB.Set(ctx, UserDBName, UserIDDBKey, buf)
	if err != nil {
		return err
	}

	return nil
}

// UpdUser - update user
func (db *UserDB) UpdUser(ctx context.Context, ui *chatbotpb.UserInfo) error {
	buf, err := proto.Marshal(ui)
	if err != nil {
		return err
	}

	err = db.ankaDB.Set(ctx, UserDBName, makeUserKey(ui.Uid), buf)
	if err != nil {
		return err
	}

	return nil
}

// UpdAppUID - update app uid
func (db *UserDB) UpdAppUID(ctx context.Context, appToken string, appUID string, uid int64) error {
	buf := chatbotbase.Int64ToBytes(uid)

	err := db.ankaDB.Set(ctx, UserDBName, makeAppUID(appToken, appUID), buf)
	if err != nil {
		return err
	}

	return nil
}

// GetUserIDFromApp - get userid from app
func (db *UserDB) GetUserIDFromApp(ctx context.Context, appToken string, appUID string) (int64, error) {
	buf, err := db.ankaDB.Get(ctx, UserDBName, makeAppUID(appToken, appUID))
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return -1, nil
		}

		return -1, err
	}

	uid := chatbotbase.BytesToInt64(buf)

	return uid, nil
}

// GetUserInfo - get user infomation
func (db *UserDB) GetUserInfo(ctx context.Context, uid int64) (*chatbotpb.UserInfo, error) {
	k := makeUserKey(uid)

	buf, err := db.ankaDB.Get(ctx, UserDBName, k)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	ui := &chatbotpb.UserInfo{}

	err = proto.Unmarshal(buf, ui)
	if err != nil {
		return nil, err
	}

	return ui, nil
}

// GetUserInfoEx - get user infomation
func (db *UserDB) GetUserInfoEx(ctx context.Context, appToken string, appUID string) (*chatbotpb.UserInfo, error) {
	uid, err := db.GetUserIDFromApp(ctx, appToken, appUID)
	if err != nil {
		return nil, err
	}

	if uid > 0 {
		return db.GetUserInfo(ctx, uid)
	}

	return nil, nil
}

// UpdUserData - update user
func (db *UserDB) UpdUserData(ctx context.Context, uid int64, ud proto.Message) error {
	buf, err := proto.Marshal(ud)
	if err != nil {
		return err
	}

	err = db.ankaDB.Set(ctx, UserDBName, makeUserDataKey(uid), buf)
	if err != nil {
		return err
	}

	return nil
}

// GetUserData - get user data
func (db *UserDB) GetUserData(ctx context.Context, uid int64) (proto.Message, error) {
	if db.MgrUserData == nil {
		return nil, nil
	}

	k := makeUserDataKey(uid)

	buf, err := db.ankaDB.Get(ctx, UserDBName, k)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	return db.MgrUserData.Unmarshal(buf)
}

// GetNoteInfo - get note infomation
func (db *UserDB) GetNoteInfo(ctx context.Context, name string) (*chatbotpb.NoteInfo, error) {
	if !IsValidNoteName(name) {
		return nil, chatbotbase.ErrNoteInvalidName
	}

	name = strings.ToLower(name)

	k := makeNoteInfoKey(name)
	buf, err := db.ankaDB.Get(ctx, UserDBName, k)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	ni := &chatbotpb.NoteInfo{}

	err = proto.Unmarshal(buf, ni)
	if err != nil {
		return nil, err
	}

	return ni, nil
}

// UpdNoteInfo - update note infomation
func (db *UserDB) UpdNoteInfo(ctx context.Context, ni *chatbotpb.NoteInfo) error {
	if !IsValidNoteName(ni.Name) {
		return chatbotbase.ErrNoteInvalidName
	}

	ni.Name = strings.ToLower(ni.Name)

	k := makeNoteInfoKey(ni.Name)

	buf, err := proto.Marshal(ni)
	if err != nil {
		return err
	}

	err = db.ankaDB.Set(ctx, UserDBName, k, buf)
	if err != nil {
		return err
	}

	return nil
}

// UpdNoteNode - update note node
func (db *UserDB) UpdNoteNode(ctx context.Context, nn *chatbotpb.NoteNode) error {
	ni, err := db.GetNoteInfo(ctx, nn.Name)
	if err != nil {
		return err
	}

	if ni == nil {
		return chatbotbase.ErrNoteNone
	}

	nn.NoteIndex = ni.NoteNums

	ni.NoteNums++
	ni.Keys = MergeKeys(ni.Keys, nn.Keys)
	InsMapKeys(ni, nn.Keys, nn.NoteIndex)

	err = db.UpdNoteInfo(ctx, ni)
	if err != nil {
		return err
	}

	k := makeNoteNodeKey(ni.Name, nn.NoteIndex)

	buf, err := proto.Marshal(nn)
	if err != nil {
		return err
	}

	err = db.ankaDB.Set(ctx, UserDBName, k, buf)
	if err != nil {
		return err
	}

	return nil
}

// GetNoteNode - get note node
func (db *UserDB) GetNoteNode(ctx context.Context, nameNote string, noteIndex int64) (*chatbotpb.NoteNode, error) {
	k := makeNoteNodeKey(nameNote, noteIndex)

	buf, err := db.ankaDB.Get(ctx, UserDBName, k)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	nn := &chatbotpb.NoteNode{}

	err = proto.Unmarshal(buf, nn)
	if err != nil {
		return nil, err
	}

	return nn, nil
}
