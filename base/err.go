package chatbotbase

import "errors"

var (
	// ErrNoAppServType - no app serv type
	ErrNoAppServType = errors.New("no app serv type")
	// ErrNoAppServToken - no app serv token
	ErrNoAppServToken = errors.New("no app serv token")
	// ErrNoAppServUserName - no app serv username
	ErrNoAppServUserName = errors.New("no app serv username")

	// ErrNoBindAddr - no bindaddr
	ErrNoBindAddr = errors.New("no bindaddr")

	// ErrNoConfig - no config
	ErrNoConfig = errors.New("no config")

	// ErrInvalidAppServType - invalid appservtype
	ErrInvalidAppServType = errors.New("invalid appservtype")
	// ErrInvalidAppServToken - invalid appservtoken
	ErrInvalidAppServToken = errors.New("invalid appservtoken")
	// ErrInvalidAppServUserName - invalid appservusername
	ErrInvalidAppServUserName = errors.New("invalid appservusername")

	// ErrNoServAddrInClient - no servaddr in client
	ErrNoServAddrInClient = errors.New("no servaddr in client")
	// ErrNoTokenInClient - no token in client
	ErrNoTokenInClient = errors.New("no token in client")

	// ErrUnkonow - unknow error
	ErrUnkonow = errors.New("unknow error")
)
