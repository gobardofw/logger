package logger

import (
	"io"
)

// NewLog Create new log message instance
func NewLog(tf string, w io.Writer, f TimeFormatter, typ string) Log {
	log := new(logDriver)
	log.init(tf, w, f)
	log.Type(typ)
	return log
}

// NewLogger create a new logger instance
func NewLogger(tf string, w io.Writer, f TimeFormatter) Logger {
	lgr := new(loggerDriver)
	lgr.init(tf, w, f)
	return lgr
}

// NewFileLogger create new file logger writer
func NewFileLogger(path string, prefix string, tf string, f TimeFormatter) io.Writer {
	fLogger := new(fileLogger)
	fLogger.init(path, prefix, tf, f)
	return fLogger
}
