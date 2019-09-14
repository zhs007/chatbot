package chatbotdb

import (
	"context"

	"github.com/golang/protobuf/proto"
	ankadb "github.com/zhs007/ankadb"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
)

// AppServDB - appserv database
type AppServDB struct {
	ankaDB ankadb.AnkaDB
}

// NewAppServDB - new AdaCoreDB
func NewAppServDB(dbpath string, httpAddr string, engine string) (*AppServDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   AppServDBName,
		Engine: engine,
		PathDB: AppServDBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		chatbotbase.Error("NewAppServDB", zap.Error(err))

		return nil, err
	}

	chatbotbase.Info("NewAppServDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &AppServDB{
		ankaDB: ankaDB,
	}

	return db, err
}

// GetAppServ - get app service
func (db *AppServDB) GetAppServ(ctx context.Context, token string) (*chatbotpb.AppServInfo, error) {
	buf, err := db.ankaDB.Get(ctx, AppServDBName, makeAppServKey(token))
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	ri := &chatbotpb.AppServInfo{}

	err = proto.Unmarshal(buf, ri)
	if err != nil {
		return nil, err
	}

	return ri, nil
}

// UpdAppServ - update app service
func (db *AppServDB) UpdAppServ(ctx context.Context, cfg *chatbotbase.AppServConfig) (*chatbotpb.AppServInfo, error) {
	casi, err := db.GetAppServ(ctx, cfg.Token)
	if err != nil {
		return nil, err
	}

	t, err := chatbotbase.GetAppServType(cfg.Type)
	if err != nil {
		return nil, err
	}

	asi := &chatbotpb.AppServInfo{
		Token:    cfg.Token,
		AppType:  t,
		Username: cfg.UserName,
	}

	if casi != nil && t == casi.AppType && casi.Username == cfg.UserName {
		asi.Sessionid = casi.Sessionid
	}

	buf, err := proto.Marshal(asi)
	if err != nil {
		return nil, err
	}

	err = db.ankaDB.Set(ctx, AppServDBName, makeAppServKey(cfg.Token), buf)
	if err != nil {
		return nil, err
	}

	return asi, nil
}

// DelAppServ - delete app service
func (db *AppServDB) DelAppServ(ctx context.Context, token string) error {
	err := db.ankaDB.Delete(ctx, AppServDBName, makeAppServKey(token))
	if err != nil {
		return err
	}

	return nil
}

// Init - initial with AppServConfig list
func (db *AppServDB) Init(ctx context.Context, lst []chatbotbase.AppServConfig) error {
	curdb := db.ankaDB.GetDBMgr().GetDB(AppServDBName)
	if curdb == nil {
		return chatbotbase.ErrNoAppServDB
	}

	it := curdb.NewIteratorWithPrefix([]byte(AppServDBKeyPrefix))
	if it.Error() != nil {
		return it.Error()
	}

	for {
		if it.Valid() {
			ri := &chatbotpb.AppServInfo{}
			err := ankadb.GetMsgFromIterator(it, ri)
			if err == nil {
				c := chatbotbase.FindAppServConfig(ri.Token, lst)
				if c == nil {
					db.DelAppServ(ctx, ri.Token)
				}
			}
		}

		if !it.Next() {
			break
		}
	}

	for _, v := range lst {
		_, err := db.UpdAppServ(ctx, &v)
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateSessionID - generate sessionID
func (db *AppServDB) GenerateSessionID(ctx context.Context, asi *chatbotpb.AppServInfo) (
	*chatbotpb.AppServInfo, error) {

	if asi != nil {
		asi.Sessionid = chatbotbase.RandString(32)
	}

	return db.UpdAppServEx(ctx, asi)
}

// UpdAppServEx - update app service
func (db *AppServDB) UpdAppServEx(ctx context.Context, asi *chatbotpb.AppServInfo) (*chatbotpb.AppServInfo, error) {

	buf, err := proto.Marshal(asi)
	if err != nil {
		return nil, err
	}

	err = db.ankaDB.Set(ctx, AppServDBName, makeAppServKey(asi.Token), buf)
	if err != nil {
		return nil, err
	}

	return asi, nil
}
