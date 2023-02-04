package stdlib

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

// Dbugf Print a debug message to stdout.`
//
// Deprecated: This was moved to the log package, please use log.Dbugf instead.
// This will be removed in the next major release.
func Dbugf(message string, vars ...interface{}) {
	verboseF(VerboseLvlDebug, message, vars...)
}

// Errf Print a warning message to stderr.
//
// Deprecated: This was moved to the log package, please use log.Errf instead.
// This will be removed in the next major release.
func Errf(message string, vars ...interface{}) {
	verboseF(VerboseLvlError, message, vars...)
}

// Fatf Print a fatal message to stderr.
//
// Deprecated: This was moved to the log package, please use log.Fatf instead.
// This will be removed in the next major release.
func Fatf(message string, vars ...interface{}) {
	verboseF(VerboseLvlFatal, message, vars...)
}

// Infof Print an informational message to stdout.
//
// Deprecated: This was moved to the log package, please use log.Infof instead.
// This will be removed in the next major release.
func Infof(message string, vars ...interface{}) {
	verboseF(VerboseLvlInfo, message, vars...)
}

// Logf Log a general message, useful for giving the user feedback on progress.
//
// Deprecated: This was moved to the log package, please use log.Logf instead
// This will be removed in the next major release.
func Logf(message string, vars ...interface{}) {
	verboseF(VerboseLvlLog, message, vars...)
}

// Warnf Print a warning message to stdout.
//
// Deprecated: This was moved to the log package, please use log.Warnf instead.
// This will be removed in the next major release.
func Warnf(message string, vars ...interface{}) {
	verboseF(VerboseLvlWarn, message, vars...)
}

// verboseF Print log message based on the verbosity level. Prints a
// newline after every message.
func verboseF(lvl int, messageTmpl string, vars ...interface{}) {
	if lvl == VerboseLvlError || lvl == VerboseLvlFatal {
		_, _ = fmt.Fprintf(os.Stderr, messageTmpl, vars...)
	} else if VerbosityLevel >= lvl {
		fmt.Printf(messageTmpl, vars...)
		fmt.Println()
	}
}
