// Package logger Provide simple logging to os.Stdout and os.Stderr
package logger

import (
	"fmt"
	"os"
)

type Logger interface {
	Dbugf(message string, vars ...interface{})
	Errf(message string, vars ...interface{})
	Fatf(message string, vars ...interface{})
	Infof(message string, vars ...interface{})
	Logf(message string, vars ...interface{})
	Panf(message string, vars ...interface{})
}

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

// verboseF Print log message based on the verbosity level. Prints a newline
// after every message.
func verboseF(lvl int, messageTmpl string, vars ...interface{}) {
	var e error
	if lvl == VerboseLvlError || lvl == VerboseLvlFatal {
		_, e = fmt.Fprintf(os.Stderr, messageTmpl, vars...)
		fmt.Println()
	} else if VerbosityLevel >= lvl {
		_, e = fmt.Fprintf(os.Stdout, messageTmpl, vars...)
		fmt.Println()
	}

	if e != nil {
		msg := fmt.Errorf("an error occured while loggin an error, which is cause to panic:\n%w", e)
		panic(msg)
	}
}
