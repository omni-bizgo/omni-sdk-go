package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

const (
	internationalSmsSubPath = "/v1/send/international"
	regex                   = `^\+[0-9]+$`
)

type InternationalSmsSender struct {
	authorization string
	client        *core.HttpClient
}

type InternationalSMS struct {
	From *string `json:"from"`
	To   *string `json:"to"`
	Text *string `json:"text"`
	Ref  *string `json:"ref,omitempty"`
}

type InternationalSmsBuilder struct {
	message InternationalSMS
}

func NewInternationalSmsBuilder() *InternationalSmsBuilder {
	return &InternationalSmsBuilder{message: InternationalSMS{}}
}

func (b *InternationalSmsBuilder) Build() InternationalSMS {
	return b.message
}

func (b *InternationalSmsBuilder) From(from string) *InternationalSmsBuilder {
	b.message.From = &from
	return b
}

func (b *InternationalSmsBuilder) To(to string) *InternationalSmsBuilder {
	b.message.To = &to
	return b
}

func (b *InternationalSmsBuilder) Text(text string) *InternationalSmsBuilder {
	b.message.Text = &text
	return b
}

func (b *InternationalSmsBuilder) Ref(ref string) *InternationalSmsBuilder {
	b.message.Ref = &ref
	return b
}

func NewInternationalSmsSender(authorization string, httpClient *http.Client) *InternationalSmsSender {
	c := core.NewClient(httpClient)
	return &InternationalSmsSender{authorization: authorization, client: c}
}

func (sender *InternationalSmsSender) SendMessage(req *InternationalSMS) (*SimpleSendResponse, error) {
	if response, err := sendMessage(sender.client, internationalSmsSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseSimpleSendResponse(sender.client, response)
	}
}
