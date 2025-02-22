package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// Logger defines a simple logger
type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	debugLog *log.Logger
}

var instance *Logger

// Initialize sets up the global logger instance
func Initialize(enableDebug bool) {
	instance = New(enableDebug)
}

// New creates a new instance of Logger
func New(enableDebug bool) *Logger {
	flags := log.Ldate | log.Ltime
	return &Logger{
		infoLog:  log.New(os.Stdout, "[INFO] ", flags),
		errorLog: log.New(os.Stderr, "[ERROR] ", flags),
		debugLog: func() *log.Logger {
			if enableDebug {
				return log.New(os.Stdout, "[DEBUG] ", flags)
			}
			return nil
		}(),
	}
}

// Info logs informational messages
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLog.Printf(format, v...)
}

// Error logs error messages
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLog.Printf(format, v...)
}

// Debug logs debug messages (only if enabled)
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.debugLog != nil {
		l.debugLog.Printf(format, v...)
	}
}

// CustomLog logs messages with a custom prefix
func (l *Logger) CustomLog(prefix, format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("[%s] %s", prefix, format), v...)
}

// Info logs informational messages
func Info(format string, v ...interface{}) {
	if instance != nil {
		instance.infoLog.Printf(format, v...)
	}
}

func Panic(v ...any) {
	log.Panic(v...)
}

// Error logs error messages
func Error(format string, v ...interface{}) {
	if instance != nil {
		instance.errorLog.Printf(format, v...)
	}
}

// Debug logs debug messages (only if enabled)
func Debug(format string, v ...interface{}) {
	if instance != nil && instance.debugLog != nil {
		pc, fullPath, line, ok := runtime.Caller(1)
		if ok {
			funcName := runtime.FuncForPC(pc).Name()
			//file := path.Base(fullPath) // Get only the filename
			link := fmt.Sprintf("\033]8;;file://%s\033\\%s:%d\033]8;;\033\\", fullPath, fullPath, line)
			format = fmt.Sprintf("%s [%s] %s", link, funcName, format)
		}
		instance.debugLog.Printf(format, v...)
	}
}

// CustomLog logs messages with a custom prefix
func CustomLog(prefix, format string, v ...interface{}) {
	if instance != nil {
		log.Printf("[%s] "+format, append([]interface{}{prefix}, v...)...)
	}
}
