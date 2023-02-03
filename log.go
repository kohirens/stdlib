package stdlib

import "fmt"

const (
	VerboseLvlLog   = 0
	VerboseLvlInfo  = 1
	VerboseLvlWarn  = 2
	VerboseLvlError = 3
	VerboseLvlFatal = 4
	VerboseLvlDebug = 5
)

var VerbosityLevel = VerboseLvlLog

func Dbugf(message string, vars ...interface{}) {
	verboseF(VerboseLvlDebug, message, vars...)
}

func Errf(message string, vars ...interface{}) {
	verboseF(VerboseLvlError, message, vars...)
}

func Fatf(message string, vars ...interface{}) {
	verboseF(VerboseLvlFatal, message, vars...)
}

func Infof(message string, vars ...interface{}) {
	verboseF(VerboseLvlInfo, message, vars...)
}

// logf Log all the time, useful for giving the user feedback on progress.
func logf(message string, vars ...interface{}) {
	verboseF(VerboseLvlLog, message, vars...)
}

// Show additional logging based on the verbosity level. Prints a newline after every message.
func verboseF(lvl int, message string, a ...interface{}) {
	if VerbosityLevel >= lvl {
		fmt.Printf(message, a...)
		fmt.Println()
	}
}

func warnf(message string, vars ...interface{}) {
	verboseF(VerboseLvlWarn, message, vars...)
}
