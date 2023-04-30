// Package log with formatted messages to an output based on type.
// Defaults to stdout and Err
// stderr.
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
}

// Infof Print an informational message to stdout.
func Infof(message string, vars ...interface{}) {
	verboseF(VerboseLvlInfo, message, vars...)
}

// Logf Log a general message, useful for giving the user feedback on progress.
func Logf(message string, vars ...interface{}) {
	verboseF(VerboseLvlLog, message, vars...)
}

// Warnf Print a warning message to stdout.
func Warnf(message string, vars ...interface{}) {
	verboseF(VerboseLvlWarn, message, vars...)
}

// verboseF Print log message based on the verbosity level. Prints a
// newline after every message.
func verboseF(lvl int, messageTmpl string, vars ...interface{}) {
	if lvl == VerboseLvlError || lvl == VerboseLvlFatal {
		_, _ = fmt.Fprintf(os.Stderr, messageTmpl, vars...)
		fmt.Println()
	} else if VerbosityLevel >= lvl {
		fmt.Printf(messageTmpl, vars...)
		fmt.Println()
	}
}
