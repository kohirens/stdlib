// Package log Provide simple logging to os.Stdout and os.Stderr
package log

import (
	"fmt"
	"os"
)

const (
	VerboseLvlFatal = 1
	VerboseLvlError = 2
	VerboseLvlWarn  = 3
	VerboseLvlLog   = 4
	VerboseLvlInfo  = 5
	VerboseLvlDebug = 6
)

// VerbosityLevel Logging level (default=3); 0=fatal,1=error,2=warning,3=general,4=info,5=debug
var VerbosityLevel = VerboseLvlLog

// Dbugf Print a debug message to stdout.
func Dbugf(message string, vars ...interface{}) {
	verboseF(VerboseLvlDebug, message, vars...)
}

// Errf Print a warning message to stderr.
func Errf(message string, vars ...interface{}) {
	verboseF(VerboseLvlError, message, vars...)
}

// Fatf Print a fatal message to stderr.
func Fatf(message string, vars ...interface{}) {
	verboseF(VerboseLvlFatal, message, vars...)
	os.Exit(1)
}

// Infof Print an informational message to stdout.
func Infof(message string, vars ...interface{}) {
	verboseF(VerboseLvlInfo, message, vars...)
}

// Logf Log a general message, useful for giving the user feedback on progress.
func Logf(message string, vars ...interface{}) {
	verboseF(VerboseLvlLog, message, vars...)
}

// Panf Log a general message, useful for giving the user feedback on progress.
func Panf(message string, vars ...interface{}) {
	verboseF(VerboseLvlFatal, message, vars...)
	panic("")
}

// Warnf Print a warning message to stdout.
func Warnf(message string, vars ...interface{}) {
	verboseF(VerboseLvlWarn, message, vars...)
}

// verboseF Print log message based on the verbosity level. Prints a
// newline after every message.
func verboseF(lvl int, messageTmpl string, vars ...interface{}) {
	var err1 error
	if lvl == VerboseLvlError || lvl == VerboseLvlFatal {
		_, err1 = fmt.Fprintf(os.Stderr, messageTmpl, vars...)
		fmt.Println()
	} else if VerbosityLevel >= lvl {
		_, err1 = fmt.Fprintf(os.Stdout, messageTmpl, vars...)
		fmt.Println()
	}

	if err1 != nil {
		panic(err1)
	}
}

type Logger interface {
	Dbugf(message string, vars ...interface{})
	Errf(message string, vars ...interface{})
	Fatf(message string, vars ...interface{})
	Infof(message string, vars ...interface{})
	Logf(message string, vars ...interface{})
	Panf(message string, vars ...interface{})
}

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

// Fatf Print a fatal message to stderr.
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

// Panf Log a general message, useful for giving the user feedback on progress.
func (sl StdLogger) Panf(message string, vars ...interface{}) {
	Panf(message, vars...)
}
