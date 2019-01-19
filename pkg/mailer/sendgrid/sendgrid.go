// package sendgrid defines interoperability to
// sendgrid mail sending service
package sendgrid

import (
	"fmt"

	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendgridService holds information and
// underlying implementation for sendgrid service
type SendgridService struct {
	sendgrid *sendgrid.Client
}

// NewMailer returns abstracted mailgun service with mailer interface
func NewMailer() mailer.Mailer {
	sg := sendgrid.NewSendClient(opts.APIKey)

	return SendgridService{
		sendgrid: sg,
	}
}

// Configure loads configuration into declared opts variable.
func (s SendgridService) Configure() error {
	return envconfig.Process("MAILER", &opts)
}

func (m SendgridService) Configured() bool {
	return opts.Configured

}

// ConfigureFromOptions sets package local configuration from initialized options
func ConfigureFromOptions(so SendgridServiceOptions) {
	opts = so
}

// Name returns name of mailer implementation
func (s SendgridService) Name() string {
	return "sendgrid"
}

// Version returns version of mailer implementation
func (s SendgridService) Version() string {
	return sendgrid.Version
}

// Health returns health status of the service
func (s SendgridService) Health() bool {
	return true
}

// Send sends mail content to recipients from msg definition
func (s SendgridService) Send(msg mailer.Message) error {

	message := mail.NewV3Mail()
	from := mail.NewEmail(msg.Sender, msg.Sender)

	message.SetFrom(from)
	message.Subject = msg.Subject
	p := mail.NewPersonalization()

	for _, recp := range msg.Recipients {
		to := mail.NewEmail(recp, recp)
		p.AddTos(to)
	}

	for _, recp := range msg.CC {
		cc := mail.NewEmail(recp, recp)
		p.AddCCs(cc)
	}

	for _, recp := range msg.BCC {
		bcc := mail.NewEmail(recp, recp)
		p.AddBCCs(bcc)
	}

	message.AddPersonalizations(p)

	plainText := mail.NewContent("text/plain", msg.TextMessage)
	html := mail.NewContent("text/html", msg.HTMLMessage)

	message.AddContent(plainText, html)

	for _, file := range msg.Attachment {
		am := mail.NewAttachment()

		am.SetContent(string(file.Body))
		am.SetFilename(file.Name)
		am.SetType(file.ContentType)

		message.AddAttachment(am)
	}

	resp, err := s.sendgrid.Send(message)
	if err != nil {
		return errors.Wrap(err, "failed to send mail")
	}

	fmt.Printf("-----------------------sendgrid-----------------------\n")
	fmt.Printf("StatusCode:       %v\n", resp.StatusCode)
	fmt.Printf("ResponseBody:     %v\n", resp.Body)
	fmt.Printf("ResponseHeader    %v\n", resp.Headers)
	fmt.Printf("-----------------------sendgrid-----------------------\n")

	return nil
}
