package sendgrid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	opts.APIKey = APIKey
	mailExporter := NewMailer()

	err := mailExporter.Send(msg)

	assert.NoError(t, err)
}
