// package mailer provides common interface and
// abstraction layer for mailing infrastructure
package mailer

// Header contains necessary information about mail header
type Header struct {
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	BCC        []string `json:"bcc"`
	CC         []string `json:"cc"`
	Subject    string   `json:"subject"`
}

// Content contains messages both plaintext and HTML format
// with attachments if any
type Content struct {
	TextMessage string       `json:"text_message"`
	HTMLMessage string       `json:"html_message"`
	Attachment  []Attachment `json:"attachments"`
}

// Attachment contains file information to attach
type Attachment struct {
	Name        string `json:"name"`
	Body        []byte `json:"body"`
	ContentType string `json:"content_type"`
}

// Message compose various mail information to be sent
type Message struct {
	Header
	Content
}

func (msg *Message) SetSender(s string) {
	msg.Sender = s
}

func (msg *Message) AddRecipient(recp string) {
	msg.Recipients = append(msg.Recipients, recp)
}

func (msg *Message) AddBCC(recp string) {
	msg.BCC = append(msg.BCC, recp)
}

func (msg *Message) AddCC(recp string) {
	msg.CC = append(msg.CC, recp)
}

func (msg *Message) SetSubject(s string) {
	msg.Subject = s
}

func (msg *Message) SetTextMessage(txt string) {
	msg.TextMessage = txt
}

func (msg *Message) SetHTMLMessage(html string) {
	msg.HTMLMessage = html
}

func (msg *Message) AddAttachment(name string, content []byte) {
	msg.Attachment = append(msg.Attachment, Attachment{Name: name, Body: content})
}

// Mailer defines interface for mail exporter backend implementation
type Mailer interface {
	Name() string
	Version() string
	Configure() (Mailer, error)
	Configured() bool
	Health() bool
	Send(Message) error
}
