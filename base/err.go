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
	// ErrNoDBPath - no dbpath
	ErrNoDBPath = errors.New("no dbpath")
	// ErrNoDBEngine - no dbengine
	ErrNoDBEngine = errors.New("no dbengine")

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

	// ErrDuplicatePlugin - duplicate plugin
	ErrDuplicatePlugin = errors.New("duplicate plugin")
	// ErrNoPlugin - no plugin
	ErrNoPlugin = errors.New("no plugin")

	// ErrNoAppServDB - no appserv database
	ErrNoAppServDB = errors.New("no appserv database")

	// ErrUnkonow - unknow error
	ErrUnkonow = errors.New("unknow error")
)
