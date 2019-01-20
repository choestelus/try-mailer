package main

import (
	"github.com/choestelus/try-mailer/pkg/api"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

func newServer(db *pg.DB, me *mex.MailExporter) *echo.Echo {
	e := echo.New()
	e.POST("/send", api.MailerHandlerFunc(db, me))
	return e
}
