package chatbot

// CommandsMgr - command manager
type CommandsMgr struct {
	mapCmd map[string]Command
}

// newCommandsMgr - new CommandMgr
func newCommandsMgr() *CommandsMgr {
	return &CommandsMgr{
		mapCmd: make(map[string]Command),
	}
}

// RegisterCommand - add command with cmd
func (m *CommandsMgr) RegisterCommand(cmd string, c Command) {
	m.mapCmd[cmd] = c
}

// // ParseInChat - parse in chat
// func (m *CommandsMgr) ParseInChat(chat *chatbotpb.ChatMsg) (string, proto.Message, error) {
// 	lst := SplitCommandString(chat.Msg)
// 	if len(lst) <= 0 {
// 		return "", nil, chatbotbase.ErrCmdEmptyCmd
// 	}

// 	cmd := lst[0]

// 	c, ok := m.mapCmd[cmd]
// 	if ok {
// 		params, err := c.ParseCommandLine(lst, chat)

// 		return cmd, params, err
// 	}

// 	return "", nil, chatbotbase.ErrCmdNoCmd
// }

// // RunInChat - run func with cmd
// func (m *CommandsMgr) RunInChat(ctx context.Context, cmd string, serv *Serv, params proto.Message,
// 	chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {

// 	c, ok := m.mapCmd[cmd]
// 	if ok {
// 		return c.RunCommand(ctx, serv, params, chat)
// 	}

// 	return nil, chatbotbase.ErrCmdNoCmd
// }

// // HasCommand - has command
// func (m *CommandsMgr) HasCommand(cmd string) bool {
// 	_, ok := m.mapCmd[cmd]
// 	if ok {
// 		return true
// 	}

// 	return false
// }

// GetCommand - get command
func (m *CommandsMgr) GetCommand(cmd string) Command {
	c, ok := m.mapCmd[cmd]
	if ok {
		return c
	}

	return nil
}

var mgrCommands *CommandsMgr

func init() {
	mgrCommands = newCommandsMgr()
}

// RegisterCommand - add command with cmd
func RegisterCommand(cmd string, c Command) {
	mgrCommands.RegisterCommand(cmd, c)
}
