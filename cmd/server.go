package main

import (
	"github.com/choestelus/try-mailer/pkg/api"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// insert key at compile time
var authKey = ""

func newServer(db *pg.DB, me *mex.MailExporter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == authKey, nil
	}))
	e.POST("/send", api.MailerHandlerFunc(db, me))
	e.GET("/history/:mail", api.HistoryHandlerFunc(db))
	return e
}
