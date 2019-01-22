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
	SendingTimeout time.Duration `split_words:"true" default:"10s"`
	Configured     bool          `default:"false"`
}

// Configure contains implementation for package initialization
func Configure() error {
	return envconfig.Process("MAILGUN", &opts)
}
