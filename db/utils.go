package chatbotdb

import (
	"strconv"

	chatbotbase "github.com/zhs007/chatbot/base"
)

// AppServDBKeyPrefix - This is the prefix of AppServDBKey
const AppServDBKeyPrefix = "as:"

// makeAppServKey - Generate a database key via token
func makeAppServKey(token string) string {
	return chatbotbase.AppendString(AppServDBKeyPrefix, token)
}

// UserIDDBKey - This is UserIDDBKey
const UserIDDBKey = "uid"

// UserDBKeyPrefix - This is the prefix of UserDBKey
const UserDBKeyPrefix = "ui:"

// makeUserKey - Generate a database key via uid
func makeUserKey(uid int64) string {
	return chatbotbase.AppendString(AppServDBKeyPrefix, strconv.FormatInt(uid, 10))
}

// AppUIDDBKeyPrefix - This is the prefix of AppUIDDBKey
const AppUIDDBKeyPrefix = "at:"

// makeAppUID - Generate a database key via apptoken and appuid
func makeAppUID(appToken string, appUID string) string {
	return chatbotbase.AppendString(AppUIDDBKeyPrefix, appToken, ":", appUID)
}
