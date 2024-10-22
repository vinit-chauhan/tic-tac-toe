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

func init() {
	out, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY, 0755)
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

func Debug(msg string) {
	logMessage(LevelDebug, "DEBUG: "+msg)
}

func Info(msg string) {
	logMessage(LevelInfo, "INFO: "+msg)
}

func Warn(msg string) {
	logMessage(LevelWarn, "WARN: "+msg)
}

func Error(msg string, err error) {
	logMessage(LevelError, "ERROR: "+msg+" : "+err.Error())
}

func Panic(msg string, err error) {
	logMessage(LevelPanic, "PANIC: "+msg+" : "+err.Error())
}
