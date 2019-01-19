package mailgun

import (
	"testing"

	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	opts.APIKey = APIKey
	opts.Domain = Domain
	opts.SendingTimeout = Timeout
	mailExporter := NewMailer()

	msg := mailer.Message{
		Content: mailer.Content{
			HTMLMessage: "<h3>Hello h3</h3>",
			TextMessage: "Hello plaintext",
			Attachment: []mailer.Attachment{mailer.Attachment{
				Name: "test_file1.txt",
				Body: []byte("test file 1 content"),
			}},
		},
		Header: mailer.Header{
			Sender:     "mg@chochoe.net",
			Recipients: []string{"nattapong@chochoe.net"},
			BCC:        []string{"choestelus@gmail.com"},
			CC:         []string{"sprintf.null@gmail.com"},
			Subject:    "Test sending mail from go mailer - mailgun",
		},
	}

	err := mailExporter.Send(msg)

	assert.NoError(t, err)
}
