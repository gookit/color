package main

// Logger console logger
type Logger struct {
	style  string
	fields map[string]string
}

// log level to cli color style
var LogLevel2tag = map[string]string{
	"info":    "info",
	"warn":    "warning",
	"warning": "warning",
	"debug":   "cyan",
	"notice":  "notice",
	"error":   "error",
}

func NewLog(fields map[string]string) *Logger {
	return &Logger{"info", fields}
}

// Info log message
func (l *Logger) Info(args ...interface{}) {

}

// Log message
func (l *Logger) Log(args ...interface{}) {

}
