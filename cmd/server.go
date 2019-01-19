package main

import (
	"github.com/choestelus/try-mailer/pkg/api"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func newServer(cfg Config, logger *logrus.Logger) *echo.Echo {
	me := mex.NewMailExporter(mex.MailExporterOptions{
		Logger: logger,
	})

	e := echo.New()
	e.POST("/send", api.MailerHandlerFunc(me))
	return nil
}
