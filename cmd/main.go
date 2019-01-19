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

	for backend, option := range configure.Mailers {
		mailer := option.Mailer()
		err := mailer.Configure()
		if err != nil {
			log.Errorf("failed to initialize [%v] backend service")
			continue
		}
		log.Infof("initialized [%v] backend service", backend)
	}

	me := mex.NewMailExporter(mex.MailExporterOptions{
		Logger: log,
	})

	apiServer := newServer(cfg, log)

	log.Fatal(apiServer.Start(fmt.Sprintf("%v:%v", cfg.APIHost, cfg.APIPort)))
}
