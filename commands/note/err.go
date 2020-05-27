package chatbotcmdnote

import "errors"

var (
	// ErrCmdNoParams - no command params
	ErrCmdNoParams = errors.New("no command params in note")
)
