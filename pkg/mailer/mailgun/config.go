package mailgun

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var opts = MailgunServiceOptions{}

// MailgunServiceOptions contains configuration
// used for initialize mailgun service
type MailgunServiceOptions struct {
	Domain         string        `required:"true"`
	APIKey         string        `required:"true" envconfig:"api_key"`
	SendingTimeout time.Duration `default:"10s"`
	Configured     bool          `default:"true"`
}

// Configure contains implementation for package initialization
func Configure([]byte) error {
	return envconfig.Process("MAILGUN", &opts)
}
