// package mailgun defines interoperability to
// mailgun mail sending service
package mailgun

import (
	"context"
	"fmt"
	"time"

	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/kelseyhightower/envconfig"
	mg "github.com/mailgun/mailgun-go"
	"github.com/pkg/errors"
)

// MailgunService holds information and
// underlying implementation for mailgun service
type MailgunService struct {
	sendingTimeout time.Duration
	mailgun        *mg.MailgunImpl
}

// NewMailer returns abstracted mailgun service with mailer interface
func NewMailer() mailer.Mailer {
	m := mg.NewMailgun(opts.Domain, opts.APIKey)

	return MailgunService{
		sendingTimeout: opts.SendingTimeout,
		mailgun:        m,
	}
}

// Configure loads configuration into declared opts variable.
func (m MailgunService) Configure() error {
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
	message := m.mailgun.NewMessage(msg.Sender, msg.Subject, msg.TextMessage, msg.Recipients...)

	// Attach attachments if any
	for _, attachment := range msg.Attachment {
		message.AddBufferAttachment(attachment.Name, attachment.Body)
	}

	// Insert HTML message part
	if msg.HTMLMessage != "" {
		message.SetHtml(msg.HTMLMessage)
	}

	// Add BCC recipients
	for _, recp := range msg.BCC {
		message.AddBCC(recp)
	}

	// Add CC recipients
	for _, recp := range msg.CC {
		message.AddCC(recp)
	}

	ctx, cancel := context.WithTimeout(context.Background(), opts.SendingTimeout)
	defer cancel()

	resp, id, err := m.mailgun.Send(ctx, message)
	if err != nil {
		return errors.Wrap(err, "failed to send mail")
	}

	fmt.Printf("-----------------------mailgun-----------------------\n")
	fmt.Printf("ID:       %v\n", id)
	fmt.Printf("Response: %v\n", resp)
	fmt.Printf("-----------------------mailgun-----------------------\n")

	return nil
}
