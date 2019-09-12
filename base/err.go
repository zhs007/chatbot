package chatbotbase

import "errors"

var (
	// ErrNoAppServType - no app serv type
	ErrNoAppServType = errors.New("no app serv type")
	// ErrNoAppServToken - no app serv token
	ErrNoAppServToken = errors.New("no app serv token")
	// ErrNoAppServUserName - no app serv username
	ErrNoAppServUserName = errors.New("no app serv username")

	// ErrInvalidAppServType - invalid appservtype
	ErrInvalidAppServType = errors.New("invalid appservtype")

	// ErrUnkonow - unknow error
	ErrUnkonow = errors.New("unknow error")
)
