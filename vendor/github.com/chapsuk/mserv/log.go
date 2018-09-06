package mserv

import (
	stdlog "log"
	"os"
)

var log Logger = stdlog.New(os.Stderr, "", stdlog.LstdFlags)

// Logger interface for package things
type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

// SetLogger for package usage
func SetLogger(l Logger) {
	log = l
}
