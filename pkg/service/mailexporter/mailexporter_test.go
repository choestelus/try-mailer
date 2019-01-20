package mailexporter

import (
	"testing"

	"github.com/choestelus/try-mailer/pkg/mailer/mailgun"
	"github.com/choestelus/try-mailer/pkg/mailer/sendgrid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMailgunExporter(t *testing.T) {
	opts := mailgun.MailgunServiceOptions{
		APIKey:         MailgunAPIKey,
		Domain:         Domain,
		SendingTimeout: Timeout,
		Configured:     true,
	}

	// First mailgun is package reference
	// Second mailgun is mailer
	mailgun.ConfigureFromOptions(opts)
	mailgun := mailgun.NewMailer()

	me := NewMailExporter(MailExporterOptions{
		Logger: logrus.New(),
	})

	me.AddBackend(mailgun)
	err := me.SendMail(msg)

	assert.NoError(t, err)
}

func TestSendgridExporter(t *testing.T) {
	opts := sendgrid.SendgridServiceOptions{
		APIKey:     SendgridAPIKey,
		Configured: true,
	}

	// First sendgrid is package reference
	// Second sendgrid is mailer
	sendgrid.ConfigureFromOptions(opts)
	sendgrid := sendgrid.NewMailer()

	me := NewMailExporter(MailExporterOptions{
		Logger: logrus.New(),
	})

	me.AddBackend(sendgrid)
	err := me.SendMail(msg)

	assert.NoError(t, err)
}
