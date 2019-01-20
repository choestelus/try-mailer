package api

import (
	"net/http"

	"github.com/choestelus/try-mailer/pkg/mailer"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func MailerHandlerFunc(me *mex.MailExporter) func(c echo.Context) error {
	return func(c echo.Context) error {
		msg := mailer.Message{}

		if err := c.Bind(&msg); err != nil {
			logrus.Errorf("failed to bind request: [%v]", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "cannot bind request to internal message format",
				"error":   err,
			})
		}

		err := me.SendMail(msg)
		if err != nil {
			logrus.Errorf("failed to send mail: [%v]", err)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to send mail",
				"error":   err.Error(),
			})
		}
		return c.NoContent(http.StatusOK)
	}
}
