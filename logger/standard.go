package logger

import "os"

var _ Logger = (*Standard)(nil)
var _ Logger = Standard{}

// Standard logging functions.
//
//	Please always supply human comprehensible logging messages for yourself
//	and others whom may not have worked on your code.
type Standard struct{}

// Dbugf Print a debug message to stdout.
func (sl Standard) Dbugf(message string, vars ...interface{}) {
	verboseF(VerboseLvlDebug, message, vars...)
}

// Errf Print a warning message to stderr.
func (sl Standard) Errf(message string, vars ...interface{}) {
	verboseF(VerboseLvlError, message, vars...)
}

// Fatf Print a fatal message to stderr.
func (sl Standard) Fatf(message string, vars ...interface{}) {
	verboseF(VerboseLvlFatal, message, vars...)
	os.Exit(1)
}

// Infof Print an informational message to stdout.
func (sl Standard) Infof(message string, vars ...interface{}) {
	verboseF(VerboseLvlInfo, message, vars...)
}

// Logf Log a general message, useful for giving the user feedback on progress.
func (sl Standard) Logf(message string, vars ...interface{}) {
	verboseF(VerboseLvlLog, message, vars...)
}

// Panf Log a general message, useful for giving the user feedback on progress.
func (sl Standard) Panf(message string, vars ...interface{}) {
	verboseF(VerboseLvlFatal, message, vars...)
	panic("")
}

// Warnf Print a warning message to stdout.
func (sl Standard) Warnf(message string, vars ...interface{}) {
	verboseF(VerboseLvlWarn, message, vars...)
}
