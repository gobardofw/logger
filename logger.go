package logger

// Logger is the interface for logger drivers.
type Logger interface {
	Log() Log
	Error() Log
	Warning() Log
	Divider(divider string, count uint8, title string)
	Raw(format string, params ...interface{})
}
