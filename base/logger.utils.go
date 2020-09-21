package chatbotbase

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ParseLogLevel - string to zapcore.Level
func ParseLogLevel(str string) zapcore.Level {
	if str == "debug" {
		return zapcore.DebugLevel
	}

	if str == "info" {
		return zapcore.InfoLevel
	}

	if str == "warn" {
		return zapcore.WarnLevel
	}

	if str == "error" {
		return zapcore.ErrorLevel
	}

	if str == "dpanic" {
		return zapcore.DPanicLevel
	}

	if str == "panic" {
		return zapcore.PanicLevel
	}

	if str == "fatal" {
		return zapcore.FatalLevel
	}

	return zapcore.WarnLevel
}

// JSON - It's like zap.String(name, str)
func JSON(name string, jobj interface{}) zap.Field {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	b, err := json.Marshal(jobj)
	if err != nil {
		Warn("chatbotbase.JSON",
			zap.Error(err))

		return zap.String(name, err.Error())
	}

	return zap.String(name, string(b))
}
