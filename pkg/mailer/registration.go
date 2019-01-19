package mailer

// Configurator defines function signature for mailer configure function
type Configurator func([]byte) error

// NewMailer defines function signature for mailer creations
type NewMailer func() (Mailer, error)

// RegisterOptions contains mailer creation directive and its configure procedure
// for use when register into executable
type RegisterOptions struct {
	Mailer       NewMailer
	Configurator Configurator
}

// RegisterFunc register mailer implementation and its options into
// global map for main executable usage
func RegisterFunc(mailers map[string]RegisterOptions) func(name string, opts RegisterOptions) {
	return func(name string, opts RegisterOptions) {
		mailers[name] = opts
	}
}
