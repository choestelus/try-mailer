package mailer

type Header struct {
	Sender     string
	Recipients []string
	BCC        []string
	CC         []string
	Subject    string
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
	Configure(interface{}) error
	Health() bool
	Send(Mail) error
}
