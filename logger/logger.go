// logger/logger.go
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log is the global logger
var Log = logrus.New()

func init() {
	// Configure logger output to stdout and set log level to Info
	Log.Out = os.Stdout
	Log.SetLevel(logrus.InfoLevel)

	// Set logger to use JSON formatter
	Log.SetFormatter(&logrus.JSONFormatter{})
}
