package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func Init(logPath string) error {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

func Info(msg string, keysAndValues ...interface{}) {
	infoLogger.Println(append([]interface{}{msg}, keysAndValues...)...)
}

func Error(msg string, keysAndValues ...interface{}) {
	errorLogger.Println(append([]interface{}{msg}, keysAndValues...)...)
}
