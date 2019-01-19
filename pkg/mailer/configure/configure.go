// Package configure provides mailer configuration and
// initialize registered mailer implementations
package configure

import (
	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/choestelus/try-mailer/pkg/mailer/mailgun"
)

var mailers = map[string]mailer.RegisterOptions{}

func init() {
	register := mailer.RegisterFunc(mailers)
	register("mailgun", mailer.RegisterOptions{
		Mailer:       mailgun.NewMailer,
		Configurator: mailgun.Configure,
	})
}
