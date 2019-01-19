// package mailexporter provides mailing services and APIs.
package mailexporter

import (
	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type MailExporter struct {
	mailers []mailer.Mailer
	logger  *logrus.Logger
}

type MailExporterOptions struct {
	Logger *logrus.Logger
}

func NewMailExporter(opts MailExporterOptions) *MailExporter {
	return &MailExporter{
		logger: opts.Logger,
	}
}

func (me *MailExporter) AddBackend(mailer mailer.Mailer) {
	me.mailers = append(me.mailers, mailer)
}

// SendMail initialize mailer backend
func (me *MailExporter) SendMail(msg mailer.Message) error {
	mailer := me.mailers[0]
	err := mailer.Configure()
	if err != nil {
		return errors.Wrap(err, "failed to initialize mailer")
	}

	err = mailer.Send(msg)
	if err != nil {
		return errors.Wrap(err, "mail-exporter: failed to send mail")
	}

	return nil
}
