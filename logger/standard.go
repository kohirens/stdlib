package logger

import "os"

var _ Logger = (*StdLogger)(nil)
var _ Logger = StdLogger{}

// StdLogger Standard logging functions.
//
//	Please always supply human comprehensible logging messages for yourself
//	and others whom may not have worked on your code.
type StdLogger struct {
}

// Dbugf Print a debug message to stdout.
func (sl StdLogger) Dbugf(message string, vars ...interface{}) {
	Dbugf(message, vars...)
}

// Errf Print a warning message to stderr.
func (sl StdLogger) Errf(message string, vars ...interface{}) {
	Errf(message, vars...)
}

// Fatf Print a fatal message to stderr then exit 1.
func (sl StdLogger) Fatf(message string, vars ...interface{}) {
	Fatf(message, vars...)
	os.Exit(1)
}

// Infof Print an informational message to stdout.
func (sl StdLogger) Infof(message string, vars ...interface{}) {
	Infof(message, vars...)
}

// Logf Log a general message, useful for giving the user feedback on progress.
func (sl StdLogger) Logf(message string, vars ...interface{}) {
	Logf(message, vars...)
}

// Panf Panic printing a message to stderr before exiting.
func (sl StdLogger) Panf(message string, vars ...interface{}) {
	Panf(message, vars...)
}
