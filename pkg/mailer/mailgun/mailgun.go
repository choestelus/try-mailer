// package mailgun defines interoperability to
// mailgun mail sending service
package mailgun

import (
	"time"

	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/kelseyhightower/envconfig"
	mg "github.com/mailgun/mailgun-go"
)

// MailgunService holds information and
// underlying implementation for mailgun service
type MailgunService struct {
	sendingTimeout time.Time
	mailgun        *mg.MailgunImpl
}

// NewMailer returns abstracted mailgun service with mailer interface
func NewMailer(opts MailgunServiceOptions) mailer.Mailer {
	m := mg.NewMailgun(opts.Domain, opts.APIKey)

	return MailgunService{
		sendingTimeout: opts.SendingTimeout,
		mailgun:        m,
	}
}

// Configure loads configuration into declared opts variable.
// The configuration is loaded by package configure by default
// so this function may not be used
func (m *MailgunService) Configure() error {
	return envconfig.Process("MAILER", &opts)
}

// Name returns name of mailer implementation
func (m MailgunService) Name() string {
	return "mailgun"
}

// Version returns version of mailer implementation
func (m MailgunService) Version() string {
	// TODO: use govvv
	return "0.1.0"
}

// Health returns health status of the service
func (m MailgunService) Health() bool {
	return true
}

// Send sends mail content to recipients from msg definition
func (m MailgunService) Send(msg mailer.Message) error {
	return nil
}
