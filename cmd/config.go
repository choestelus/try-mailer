package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// Config contains general configuration for running API server
type Config struct {
	APIHost string `required:"false" envconfig:"api_host" default:"localhost"`
	APIPort int    `required:"false" envconfig:"api_port" default:"40000"`

	// debug flag and log level will be associated with logrus log level
	// when debug flag is set to true, log level will alwayls be "DEBUG" level
	Debug    bool   `default:"true"`
	LogLevel string `default:"DEBUG"`
}

func initConfig() Config {
	cfg := Config{}
	envconfig.MustProcess("MAILER", &cfg)
	return cfg
}

func initLog(cfg Config) *logrus.Logger {
	log := logrus.New()

	if cfg.Debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		switch cfg.LogLevel {
		case "PANIC":
			log.SetLevel(logrus.PanicLevel)
		case "FATAL":
			log.SetLevel(logrus.FatalLevel)
		case "ERROR":
			log.SetLevel(logrus.ErrorLevel)
		case "WARN":
			log.SetLevel(logrus.WarnLevel)
		case "INFO":
			log.SetLevel(logrus.InfoLevel)
		case "DEBUG":
			log.SetLevel(logrus.DebugLevel)
		default:
			log.SetLevel(logrus.InfoLevel)
		}
	}
	return log
}
