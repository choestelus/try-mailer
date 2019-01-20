package main

import (
	"fmt"

	"github.com/go-pg/pg"
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

	// DB.
	DBHost     string `required:"true" envconfig:"db_host"`
	DBPort     int    `required:"true" envconfig:"db_port"`
	DBName     string `required:"true" envconfig:"db_name"`
	DBUser     string `required:"true" envconfig:"db_user"`
	DBPassword string `required:"true" envconfig:"db_password"`
}

func initConfig() Config {
	cfg := Config{}
	envconfig.MustProcess("MAILER", &cfg)
	return cfg
}

func initDB(cfg Config) *pg.DB {
	dbCon := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.DBHost, cfg.DBPort),
		Database: cfg.DBName,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
	})

	return dbCon
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
