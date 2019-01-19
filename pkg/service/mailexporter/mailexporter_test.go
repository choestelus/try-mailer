package mailexporter

import (
	"testing"

	"github.com/choestelus/try-mailer/pkg/mailer/mailgun"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMailgunExporter(t *testing.T) {
	opts := mailgun.MailgunServiceOptions{
		APIKey:         APIKey,
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
