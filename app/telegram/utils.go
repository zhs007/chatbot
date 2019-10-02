package chatbottelegram

import (
	"strconv"
	"strings"
)

// ID2Str - id -> string
func ID2Str(id int) string {
	return strconv.Itoa(id)
}

// Str2ID - string -> id
func Str2ID(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// FormatCommand - format command, /start => start
func FormatCommand(str string) string {
	str = strings.TrimSpace(str)

	if len(str) <= 1 {
		return str
	}

	if str[0] == '/' && strings.IndexByte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", str[1]) >= 0 {
		return str[1:]
	}

	return str
}
