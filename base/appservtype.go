package chatbotbase

import (
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// strAppServType - appserv type string
var strAppServType = []string{"telegram", "coolq"}

// GetAppServType - strtype -> chatbotpb.ChatAppType
func GetAppServType(strType string) (chatbotpb.ChatAppType, error) {
	for i, n := range strAppServType {
		if n == strType {
			return chatbotpb.ChatAppType(i), nil
		}
	}

	return -1, ErrInvalidAppServType
}
