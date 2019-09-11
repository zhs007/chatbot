package chatbotbase

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var onceLogger sync.Once
var logPath string

// var curtime int64
var panicFile *os.File
var logSubName string

func initPanicFile() error {
	file, err := os.OpenFile(
		path.Join(logPath, BuildLogFilename("panic", logSubName)),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Warn("initPanicFile:OpenFile",
			zap.Error(err))

		return err
	}

	panicFile = file

	if err = syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd())); err != nil {
		Warn("initPanicFile:Dup2",
			zap.Error(err))

		return err
	}

	return nil
}

func initLogger(level zapcore.Level, isConsole bool, logpath string) (*zap.Logger, error) {

	// logSubName = subName

	// curtime = time.Now().Unix()

	logPath = logpath

	loglevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	if isConsole {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebugging := zapcore.Lock(os.Stdout)
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleDebugging, loglevel),
		)

		cl := zap.New(core)
		// defer cl.Sync()

		return cl, nil
	}

	cfg := &zap.Config{}

	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.OutputPaths = []string{path.Join(logpath, BuildLogFilename("output", logSubName))}
	cfg.ErrorOutputPaths = []string{path.Join(logpath, BuildLogFilename("error", logSubName))}
	cfg.Encoding = "json"
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:     "T",
		LevelKey:    "L",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		MessageKey:  "msg",
	}

	cl, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	err = initPanicFile()
	if err != nil {
		return nil, err
	}

	return cl, nil
}

// InitLogger - initializes a thread-safe singleton logger
func InitLogger(level zapcore.Level, isConsole bool, logpath string) {

	// once ensures the singleton is initialized only once
	onceLogger.Do(func() {
		cl, err := initLogger(level, isConsole, logpath)
		if err != nil {
			fmt.Printf("initLogger error! %v \n", err)

			os.Exit(-1)
		}

		logger = cl
	})

	return
}

// // Log a message at the given level with given fields
// func Log(level zap.Level, message string, fields ...zap.Field) {
// 	singleton.Log(level, message, fields...)
// }

// Debug logs a debug message with the given fields
func Debug(message string, fields ...zap.Field) {
	if logger == nil {
		return
	}

	logger.Debug(message, fields...)
}

// Info logs a debug message with the given fields
func Info(message string, fields ...zap.Field) {
	if logger == nil {
		return
	}

	logger.Info(message, fields...)
}

// Warn logs a debug message with the given fields
func Warn(message string, fields ...zap.Field) {
	if logger == nil {
		return
	}

	logger.Warn(message, fields...)
}

// Error logs a debug message with the given fields
func Error(message string, fields ...zap.Field) {
	if logger == nil {
		return
	}

	logger.Error(message, fields...)
}

// Fatal logs a message than calls os.Exit(1)
func Fatal(message string, fields ...zap.Field) {
	if logger == nil {
		return
	}

	logger.Fatal(message, fields...)
}

// SyncLogger - sync logger
func SyncLogger() {
	logger.Sync()
}

// ClearLogs - clear logs
func ClearLogs() error {
	if logPath != "" {
		fn := path.Join(logPath, "*.log")
		lst, err := filepath.Glob(fn)
		if err != nil {
			return err
		}

		panicfile := BuildLogFilename("panic", logSubName)
		outputfile := BuildLogFilename("output", logSubName)
		errorfile := BuildLogFilename("error", logSubName)

		for _, v := range lst {
			cfn := filepath.Base(v)
			if cfn != panicfile && cfn != outputfile && cfn != errorfile {
				os.Remove(v)
			}
		}
	}

	return nil
}
