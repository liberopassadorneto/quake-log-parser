package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Log is the Global Logger
var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
}
