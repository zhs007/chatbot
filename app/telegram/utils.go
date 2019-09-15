package chatbottelegram

import "strconv"

// ID2Str - id -> string
func ID2Str(id int) string {
	return strconv.Itoa(id)
}

// Str2ID - string -> id
func Str2ID(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
