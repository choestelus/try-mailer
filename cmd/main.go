package main

import (
	"fmt"

	"github.com/choestelus/try-mailer/pkg/mailer/configure"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
)

func main() {
	cfg := initConfig()
	log := initLog(cfg)

	log.Infof("server is starting...")

	me := mex.NewMailExporter(mex.MailExporterOptions{
		Logger: log,
	})

	for backend, option := range configure.Mailers {
		mailer := option.Mailer()
		err := option.Configurator()
		if err != nil {
			log.Panicf("mailer misconfigured, instantiate config again and retry: %v", err)
		}
		log.Infof("registered [%v] backend service", backend)
		configuredMailer, err := mailer.Configure()
		me.AddBackend(configuredMailer)
	}

	apiServer := newServer(cfg, me)

	log.Fatal(apiServer.Start(fmt.Sprintf("%v:%v", cfg.APIHost, cfg.APIPort)))
}
