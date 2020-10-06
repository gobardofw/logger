package logger

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// logDriver standard log message
type logDriver struct {
	typ        string
	tags       []string
	timeFormat string
	writer     io.Writer
	formatter  TimeFormatter
}

func (log *logDriver) init(tf string, w io.Writer, f TimeFormatter) {
	log.timeFormat = tf
	log.writer = w
	log.formatter = f
}

// Type Set message type
func (log *logDriver) Type(t string) Log {
	log.typ = t
	return log
}

// Tags add tags to message
func (log *logDriver) Tags(tags ...string) Log {
	for _, tag := range tags {
		log.tags = append(log.tags, tag)
	}
	return log
}

// Print print message to writer
func (log *logDriver) Print(format string, params ...interface{}) {
	// Datetime
	log.writer.Write([]byte(log.formatter(time.Now().UTC(), log.timeFormat)))
	// Type
	t := []rune(strings.ToUpper(log.typ))
	if len(t) >= 5 {
		t = t[0:5]
	}
	log.writer.Write([]byte(fmt.Sprintf("%5s ", string(t))))
	// Message
	log.writer.Write([]byte(fmt.Sprintf(strings.ReplaceAll(format, "\n", ""), params...)))
	// Tags
	for _, tag := range log.tags {
		log.writer.Write([]byte(fmt.Sprintf(" [%s]", tag)))
	}
	log.writer.Write([]byte("\n"))
}
