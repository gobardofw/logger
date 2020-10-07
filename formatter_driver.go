package logger

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

// GregorianFormatter gregorian date formatter
func GregorianFormatter(t time.Time, format string) string {
	return t.Format(format)
}

// JalaliFormatter jalali date formatter
func JalaliFormatter(t time.Time, format string) string {
	return ptime.Unix(t.Unix(), t.UnixNano(), ptime.Iran()).TimeFormat(format)
}
