package graphql

import (
	"os"

	stackdriver "github.com/icco/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// InitLogging initializes a logger to send things to stackdriver.
func InitLogging() *logrus.Logger {
	log.Formatter = stackdriver.NewFormatter()
	log.SetOutput(os.Stdout)

	// Debug only in dev
	if os.Getenv("NAT_ENV") != "production" {
		log.Level = logrus.DebugLevel
	} else {
		log.Level = logrus.InfoLevel
	}

	log.Debug("Logger successfully initialised!")

	return log
}
