package chatbotdb

import (
	"context"

	ankadb "github.com/zhs007/ankadb"
	chatbotbase "github.com/zhs007/chatbot/base"
	"go.uber.org/zap"
)

// UserDB - user database
type UserDB struct {
	ankaDB ankadb.AnkaDB
}

// NewUserDB - new UserDB
func NewUserDB(dbpath string, httpAddr string, engine string) (*UserDB, error) {
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
		ankaDB: ankaDB,
	}

	return db, err
}

// GetUserID - get userID
func (db *UserDB) GetUserID(ctx context.Context) (int64, error) {
	buf, err := db.ankaDB.Get(ctx, UserDBName, UserIDDBKey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			err = db.updUserID(ctx, 1)
			if err != nil {
				return -1, err
			}

			return 1, nil
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
