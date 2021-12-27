package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	appContext := Initialize()
	if appContext.Config.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if appContext.Config.JsonLog {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	}
	logrus.SetReportCaller(true)

	appContext.Echo.Logger.Fatal(appContext.Echo.Start(fmt.Sprintf(":%d", appContext.Config.Port)))
}
