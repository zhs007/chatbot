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
	// ErrPluginInvalidServ - invalid serv in plugin
	ErrPluginInvalidServ = errors.New("invalid serv in plugin")
	// ErrPluginInvalidServMgrText - invalid serv.MgrText in plugin
	ErrPluginInvalidServMgrText = errors.New("invalid serv.MgrText in plugin")
	// ErrPluginItsNotMine - It's not mine
	ErrPluginItsNotMine = errors.New("Plugin: It's not mine")

	// ErrNoAppServDB - no appserv database
	ErrNoAppServDB = errors.New("no appserv database")
	// ErrInvalidTokenInAppServDB - invalid token in appserv database
	ErrInvalidTokenInAppServDB = errors.New("invalid token in appserv database")

	// ErrNoServAddr - no servaddr
	ErrNoServAddr = errors.New("no servaddr")
	// ErrNoToken - no token
	ErrNoToken = errors.New("no token")
	// ErrNoUsername - no username
	ErrNoUsername = errors.New("no username")
	// ErrNoTelegramToken - no telegramtoken
	ErrNoTelegramToken = errors.New("no telegramtoken")

	// ErrNoCoolQServURL - no coolq servurl
	ErrNoCoolQServURL = errors.New("no coolq servurl")
	// ErrNoCoolQHttpServAddr - no servaddr
	ErrNoCoolQHttpServAddr = errors.New("no coolq http servaddr")

	// ErrStreamNoMsg - no chat in stream
	ErrStreamNoMsg = errors.New("no chat in stream")
	// ErrInvalidStartInStream - invalid start in stream
	ErrInvalidStartInStream = errors.New("invalid start in stream")
	// ErrInvalidLengthInStream - invalid length in stream
	ErrInvalidLengthInStream = errors.New("invalid length in stream")
	// ErrInvalidHashInStream - invalid hash in stream
	ErrInvalidHashInStream = errors.New("invalid hash in stream")
	// ErrInvalidTotalLengthInStream - invalid totalLength in stream
	ErrInvalidTotalLengthInStream = errors.New("invalid totalLength in stream")
	// ErrInvalidTotalHashInStream - invalid totalHash in stream
	ErrInvalidTotalHashInStream = errors.New("invalid totalHash in stream")
	// ErrInvalidStream - invalid chatstream
	ErrInvalidStream = errors.New("invalid chatstream")

	// ErrServInvalidErr - invalid error in serv
	ErrServInvalidErr = errors.New("invalid error in serv")
	// ErrServInvalidToken - invalid token in serv
	ErrServInvalidToken = errors.New("invalid token in serv")
	// ErrServInvalidUserInfo - invalid userinfo
	ErrServInvalidUserInfo = errors.New("invalid userinfo")

	// ErrCmdNoCmd - no command
	ErrCmdNoCmd = errors.New("no command")
	// ErrCmdEmptyCmd - empty command
	ErrCmdEmptyCmd = errors.New("empty command")
	// ErrCmdInvalidServ - invalid serv in command
	ErrCmdInvalidServ = errors.New("invalid serv in command")
	// ErrCmdInvalidServMgrText - invalid serv.MgrText in command
	ErrCmdInvalidServMgrText = errors.New("invalid serv.MgrText in command")
	// ErrCmdItsNotMine - It's not mine
	ErrCmdItsNotMine = errors.New("Command: It's not mine")

	// ErrUserMgrNoUserData - no userdata
	ErrUserMgrNoUserData = errors.New("no userdata")

	// ErrInvalidFileProcessorMgr - invalid FileProcessorMgr
	ErrInvalidFileProcessorMgr = errors.New("invalid FileProcessorMgr")

	// ErrNoteInvalidName - invalid note name
	ErrNoteInvalidName = errors.New("invalid note name")

	// ErrUnkonow - unknow error
	ErrUnkonow = errors.New("unknow error")
)
