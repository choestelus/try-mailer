// package mailexporter provides mailing services and APIs.
package mailexporter

import (
	"fmt"

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

// SendMail attempt to send mail with backend registered in MailExporter
// it'll try to use backend in mailer list to send, if failed, will try on
// next backend until out of mailer. if no mailer success to send, returns error
func (me *MailExporter) SendMail(msg mailer.Message) error {
	for _, mailer := range me.mailers {
		if !mailer.Configured() {
			err := mailer.Configure()
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed to initialize %v mailer", mailer.Name()))
			}
		}

		err := mailer.Send(msg)
		if err != nil {
			me.logger.Warnf("mail-exporter: failed to send mail using %v", mailer.Name())
			continue
		}
		return nil
	}

	return fmt.Errorf("failed to send mail - out of usable backend")
}
