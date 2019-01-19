// Package configure provides mailer configuration and
// initialize registered mailer implementations
package configure

import (
	"github.com/choestelus/try-mailer/pkg/mailer"
	"github.com/choestelus/try-mailer/pkg/mailer/mailgun"
)

var Mailers = map[string]mailer.RegisterOptions{}

func init() {
	register := mailer.RegisterFunc(Mailers)
	register("mailgun", mailer.RegisterOptions{
		Mailer:       mailgun.NewMailer,
		Configurator: mailgun.Configure,
	})
}
