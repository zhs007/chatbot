package chatbotdb

import (
	"strconv"
	"unicode"

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

// AppDataDBKey - This is AppDataDBKey
const AppDataDBKey = "appdata"

// UserDBKeyPrefix - This is the prefix of UserDBKey
const UserDBKeyPrefix = "ui:"

// makeUserKey - Generate a database key via uid
func makeUserKey(uid int64) string {
	return chatbotbase.AppendString(UserDBKeyPrefix, strconv.FormatInt(uid, 10))
}

// UserDataDBKeyPrefix - This is the prefix of UserDataDBKey
const UserDataDBKeyPrefix = "ud:"

// makeUserDataKey - Generate a database key via uid
func makeUserDataKey(uid int64) string {
	return chatbotbase.AppendString(UserDataDBKeyPrefix, strconv.FormatInt(uid, 10))
}

// AppUIDDBKeyPrefix - This is the prefix of AppUIDDBKey
const AppUIDDBKeyPrefix = "at:"

// makeAppUID - Generate a database key via apptoken and appuid
func makeAppUID(appToken string, appUID string) string {
	return chatbotbase.AppendString(AppUIDDBKeyPrefix, appToken, ":", appUID)
}

// NoteInfoDBKeyPrefix - This is the prefix of NoteInfo
const NoteInfoDBKeyPrefix = "ni:"

// makeNoteInfoKey - Generate a database key via note name
func makeNoteInfoKey(name string) string {
	return chatbotbase.AppendString(NoteInfoDBKeyPrefix, name)
}

// IsValidNoteName - is valid note name
func IsValidNoteName(name string) bool {
	if name == "" {
		return false
	}

	for _, v := range name {
		if !(unicode.IsLetter(v) || unicode.IsDigit(v)) {
			return false
		}
	}

	return true
}
