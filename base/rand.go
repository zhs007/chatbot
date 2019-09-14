package chatbotbase

import (
	"math/rand"
	"time"
)

const letterNormal = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString - random a string
func RandString(length int) (str string) {
	for i := 0; i < length; i++ {
		str += string(letterNormal[rand.Intn(len(letterNormal))])
	}

	return
}

func init() {
	rand.Seed(time.Now().Unix())
}
