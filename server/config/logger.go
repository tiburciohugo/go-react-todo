package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warningLogger *log.Logger
	writer        io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debugLogger:   log.New(writer, "DEBUG: ", logger.Flags()),
		infoLogger:    log.New(writer, "INFO: ", logger.Flags()),
		errorLogger:   log.New(writer, "ERROR: ", logger.Flags()),
		warningLogger: log.New(writer, "WARNING: ", logger.Flags()),
		writer:        writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debugLogger.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warningLogger.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warningLogger.Printf(format, v...)
}
