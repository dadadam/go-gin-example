package main

import (
	"github.com/dadadam/sono-backend/config"
	"github.com/dadadam/sono-backend/db"
	"github.com/dadadam/sono-backend/server"
	log "github.com/sirupsen/logrus"
)

func configureLogger(c *config.Config) {
	logLevel, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		log.Warn("unnable to parse log level, default log level: DEBUG")
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}

func main() {
	config.Init()

	c := config.GetConfig()
	configureLogger(c)
	db.Init()
	server.Init()

	log.Info("Application started")
}
