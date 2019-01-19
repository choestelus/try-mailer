// package mailgun defines interoperability to
// mailgun mail sending service
package mailgun

import (
	"time"

	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/kelseyhightower/envconfig"
	mg "github.com/mailgun/mailgun-go"
)

var opts = MailgunServiceOptions{}

// MailgunServiceOptions contains configuration
// used for initialize mailgun service
type MailgunServiceOptions struct {
	Domain         string    `required:"true"`
	APIKey         string    `required:"true" envconfig:"api_key"`
	SendingTimeout time.Time `default:"10s"`
}

func Configure([]byte) error {
	return envconfig.Process("MAILGUN", &opts)
}

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

func (m MailgunService) Name() string {
	return "mailgun"
}

func (m MailgunService) Version() string {
	// TODO: use govvv
	return "0.1.0"
}
