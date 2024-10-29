package logger

import (
	"log"
	"os"
)

var logger *log.Logger
var LogLevel = LevelInfo

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
)

const LogsDir = "logs"

func init() {
	if _, err := os.ReadDir(LogsDir); err != nil {
		if err := os.Mkdir(LogsDir, 0755); err != nil {
			panic(err)
		}

	}

	out, err := os.OpenFile(LogsDir+"/server.log", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}

	logger = log.New(out, "", log.LstdFlags|log.LUTC)
}

func SetLogLevel(level int) {
	LogLevel = level
}

func logMessage(level int, msg string) {
	if level >= LogLevel {
		logger.Println(msg)
	}
}

func Debug(msg, tag string) {
	logMessage(LevelDebug, "DEBUG:  ["+tag+"] "+msg)
}

func Info(msg, tag string) {
	logMessage(LevelInfo, "INFO: ["+tag+"] "+msg)
}

func Warn(msg, tag string) {
	logMessage(LevelWarn, "WARN: ["+tag+"] "+msg)
}

func Error(msg, tag string, err error) {
	logMessage(LevelError, "ERROR: ["+tag+"] "+msg+" : "+err.Error())
}

func Panic(msg, tag string, err error) {
	logMessage(LevelPanic, "PANIC: ["+tag+"] "+msg+" : "+err.Error())
}
