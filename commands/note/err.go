package chatbotcmdnote

import "errors"

var (
	// ErrCmdNoParams - no command params
	ErrCmdNoParams = errors.New("no command params in note")
	// ErrCmdInvalidParams - invalid command params
	ErrCmdInvalidParams = errors.New("invalid command params in note")
)