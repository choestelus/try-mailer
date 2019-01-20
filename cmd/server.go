package main

import (
	"github.com/choestelus/try-mailer/pkg/api"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/labstack/echo"
)

func newServer(cfg Config, me *mex.MailExporter) *echo.Echo {
	e := echo.New()
	e.POST("/send", api.MailerHandlerFunc(me))
	return e
}
