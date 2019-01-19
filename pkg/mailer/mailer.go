// package mailer provides common interface and
// abstraction layer for mailing infrastructure
package mailer

type Header struct {
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	BCC        []string `json:"bcc"`
	CC         []string `json:"cc"`
	Subject    string   `json:"subject"`
}

type Content struct {
	TextMessage string
	HTMLMessage string
	Attachment  []Attachment
}

type Attachment struct {
	Name string
	Body []byte
}

type Message struct {
	Header
	Content
}

type Messager interface {
	AddHeader(header string, value string)
	AddAttachment([]byte)
	AddBCC(...string)
	AddCC(...string)
	AddRecipient(...string)
	SetTextContent(string)
	SetHTMLContent(string)
}

type Mailer interface {
	Name() string
	Version() string
	Configure() error
	Health() bool
	Send(Message) error
}
