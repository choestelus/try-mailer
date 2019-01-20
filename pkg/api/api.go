package api

import (
	"net/http"

	"github.com/choestelus/try-mailer/pkg/mailer"
	mex "github.com/choestelus/try-mailer/pkg/service/mailexporter"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const (
	RecordHistoryQuery = `INSERT INTO public.histories (email, access_from, status, failed) VALUES(?, ?, ?, ?)`
)

func MailerHandlerFunc(db *pg.DB, me *mex.MailExporter) func(c echo.Context) error {
	return func(c echo.Context) error {
		msg := mailer.Message{}

		if err := c.Bind(&msg); err != nil {
			logrus.Errorf("failed to bind request: [%v]", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "cannot bind request to internal message format",
				"error":   err,
			})
		}

		allRecipients := []string{}
		allRecipients = append(allRecipients, msg.Recipients...)
		allRecipients = append(allRecipients, msg.CC...)
		allRecipients = append(allRecipients, msg.BCC...)

		err := me.SendMail(msg)
		if err != nil {
			logrus.Errorf("failed to send mail: [%v]", err)

			// duplication over abstraction.
			// inefficient, should not commit every insertion, but does the job.
			for _, recp := range allRecipients {
				_, err := db.ExecOne(RecordHistoryQuery, recp, c.RealIP(), err.Error(), true)
				if err != nil {
					logrus.Warnf("failed to record sending status to history: %v", err)
				}
			}

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to send mail",
				"error":   err.Error(),
			})
		}

		// duplication over abstraction.
		// inefficient, should not commit every insertion, but does the job.
		for _, recp := range allRecipients {
			_, err := db.ExecOne(RecordHistoryQuery, recp, c.RealIP(), "", false)
			if err != nil {
				logrus.Warnf("failed to record sending status to history: %v", err)
			}
		}
		return c.NoContent(http.StatusOK)
	}
}
