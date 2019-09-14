package chatbotdb

import (
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
