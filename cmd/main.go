package main

import (
	"fmt"
)

func main() {
	cfg := initConfig()
	log := initLog(cfg)

	log.Infof("server is starting...")

	apiServer := newServer(cfg)

	log.Fatal(apiServer.Start(fmt.Sprintf("%v:%v", cfg.APIHost, cfg.APIPort)))

}
