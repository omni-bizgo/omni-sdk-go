package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

const (
	alimtalkSubPath = "/v1/send/alimtalk"
)

type AlimtalkSender struct {
	authorization string
	client        *core.HttpClient
}

type Alimtalk struct {
	SenderKey    *string       `json:"senderKey"`
	To           *string       `json:"to"`
	Text         *string       `json:"text"`
	TemplateCode *string       `json:"templateCode"`
	KakaoMsgType *KakaoMsgType `json:"msgType"`
	Button       []KakaoButton `json:"button,omitempty"`
	Fallback     *Fallback     `json:"fallback,omitempty"`
	Ref          *string       `json:"ref,omitempty"`
}

type AlimtalkBuilder struct {
	message Alimtalk
}

func NewAlimtalkBuilder() *AlimtalkBuilder {
	return &AlimtalkBuilder{message: Alimtalk{}}
}

func (b *AlimtalkBuilder) Build() Alimtalk {
	return b.message
}

func (b *AlimtalkBuilder) SenderKey(senderKey string) *AlimtalkBuilder {
	b.message.SenderKey = &senderKey
	return b
}

func (b *AlimtalkBuilder) To(to string) *AlimtalkBuilder {
	b.message.To = &to
	return b
}

func (b *AlimtalkBuilder) Text(text string) *AlimtalkBuilder {
	b.message.Text = &text
	return b
}

func (b *AlimtalkBuilder) TemplateCode(templateCode string) *AlimtalkBuilder {
	b.message.TemplateCode = &templateCode
	return b
}

func (b *AlimtalkBuilder) KakaoMsgType(kakaoMsgType KakaoMsgType) *AlimtalkBuilder {
	b.message.KakaoMsgType = &kakaoMsgType
	return b
}

func (b *AlimtalkBuilder) Button(button []KakaoButton) *AlimtalkBuilder {
	b.message.Button = button
	return b
}

func (b *AlimtalkBuilder) Fallback(fallback Fallback) *AlimtalkBuilder {
	b.message.Fallback = &fallback
	return b
}

func (b *AlimtalkBuilder) Ref(ref string) *AlimtalkBuilder {
	b.message.Ref = &ref
	return b
}

func NewAlimTalkSender(authorization string, httpClient *http.Client) *AlimtalkSender {
	c := core.NewClient(httpClient)
	return &AlimtalkSender{authorization: authorization, client: c}
}

func (sender *AlimtalkSender) SendMessage(req *Alimtalk) (*SimpleSendResponse, error) {
	if response, err := sendMessage(sender.client, alimtalkSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseSimpleSendResponse(sender.client, response)
	}
}
