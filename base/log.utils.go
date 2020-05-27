package chatbotbase

import "go.uber.org/zap/zapcore"

// ParseLogLevel - string to zapcore.Level
func ParseLogLevel(loglevel string) zapcore.Level {
	if loglevel == "debug" {
		return zapcore.DebugLevel
	}

	if loglevel == "info" {
		return zapcore.InfoLevel
	}

	if loglevel == "warn" {
		return zapcore.WarnLevel
	}

	if loglevel == "error" {
		return zapcore.ErrorLevel
	}

	return zapcore.InfoLevel
}
