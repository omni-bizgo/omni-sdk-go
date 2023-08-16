package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

const (
	smsSubPath = "/v1/send/sms"
)

type SmsSender struct {
	authorization string
	client        *core.HttpClient
}

type SMS struct {
	From      *string `json:"from"`
	To        *string `json:"to"`
	Text      *string `json:"text"`
	Ref       *string `json:"ref,omitempty"`
	OriginCID *string `json:"originCID,omitempty"`
}

type SmsBuilder struct {
	message SMS
}

func NewSmsBuilder() *SmsBuilder {
	return &SmsBuilder{message: SMS{}}
}

func (b *SmsBuilder) Build() SMS {
	return b.message
}

func (b *SmsBuilder) From(from string) *SmsBuilder {
	b.message.From = &from
	return b
}

func (b *SmsBuilder) To(to string) *SmsBuilder {
	b.message.To = &to
	return b
}

func (b *SmsBuilder) Text(text string) *SmsBuilder {
	b.message.Text = &text
	return b
}

func (b *SmsBuilder) Ref(ref string) *SmsBuilder {
	b.message.Ref = &ref
	return b
}

func (b *SmsBuilder) OriginCID(originCID string) *SmsBuilder {
	b.message.OriginCID = &originCID
	return b
}

func NewSmsSender(authorization string, httpClient *http.Client) *SmsSender {
	c := core.NewClient(httpClient)
	return &SmsSender{authorization: authorization, client: c}
}

func (sender *SmsSender) SendMessage(req *SMS) (*SimpleSendResponse, error) {
	if response, err := sendMessage(sender.client, smsSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseSimpleSendResponse(sender.client, response)
	}
}
