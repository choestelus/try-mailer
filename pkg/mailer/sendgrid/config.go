package sendgrid

import (
	"github.com/kelseyhightower/envconfig"
)

var opts = SendgridServiceOptions{}

// MailgunServiceOptions contains configuration
// used for initialize mailgun service
type SendgridServiceOptions struct {
	APIKey     string `required:"true" envconfig:"api_key"`
	Configured bool   `default:"true"`
}

// Configure contains implementation for package initialization
func Configure([]byte) error {
	return envconfig.Process("SENDGRID", &opts)
}
