package main

import (
	"github.com/choestelus/try-mailer/pkg/api"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newServer(db *pg.DB, me *mex.MailExporter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/send", api.MailerHandlerFunc(db, me))
	e.GET("/history/:mail", api.HistoryHandlerFunc(db))
	return e
}
