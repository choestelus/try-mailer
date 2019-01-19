package mailgun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	opts.APIKey = APIKey
	opts.Domain = Domain
	opts.SendingTimeout = Timeout
	mailExporter := NewMailer()

	err := mailExporter.Send(msg)

	assert.NoError(t, err)
}
