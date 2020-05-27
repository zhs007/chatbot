package basicchatbot

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	basicchatbotpb "github.com/zhs007/chatbot/basicchatbotpb"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// // DebugExcelColumnType - excel column type
// type DebugExcelColumnType struct {
// 	Name      string
// 	Type      string
// 	TypeAuto  string
// 	Separator string
// }

// ServiceCore - chatbot service core
type ServiceCore struct {
}

// UnmarshalAppData - unmarshal
func (core *ServiceCore) UnmarshalAppData(buf []byte) (proto.Message, error) {
	ad := &basicchatbotpb.ChatBotData{}

	err := proto.Unmarshal(buf, ad)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

// NewAppData - new a app data
func (core *ServiceCore) NewAppData() (proto.Message, error) {
	return &basicchatbotpb.ChatBotData{}, nil
}

// UnmarshalUserData - unmarshal
func (core *ServiceCore) UnmarshalUserData(buf []byte) (proto.Message, error) {
	ud := &basicchatbotpb.UserData{}

	err := proto.Unmarshal(buf, ud)
	if err != nil {
		return nil, err
	}

	return ud, nil
}

// NewUserData - new a userdata
func (core *ServiceCore) NewUserData(ui *chatbotpb.UserInfo) (proto.Message, error) {
	return &basicchatbotpb.UserData{}, nil
}

// OnDebug - call in plugin.debug
func (core *ServiceCore) OnDebug(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	return nil, nil
}
