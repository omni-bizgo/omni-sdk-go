package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

const (
	mmsSubPath = "/v1/send/mms"
)

type MmsSender struct {
	authorization string
	client        *core.HttpClient
}

type MMS struct {
	From      *string  `json:"from"`
	To        *string  `json:"to"`
	Text      *string  `json:"text"`
	Title     *string  `json:"title,omitempty"`
	FileKey   []string `json:"fileKey,omitempty"`
	Ref       *string  `json:"ref,omitempty"`
	OriginCID *string  `json:"originCID,omitempty"`
}

type MmsBuilder struct {
	message MMS
}

func NewMmsBuilder() *MmsBuilder {
	return &MmsBuilder{message: MMS{}}
}

func (b *MmsBuilder) Build() MMS {
	return b.message
}

func (b *MmsBuilder) To(To string) *MmsBuilder {
	b.message.To = &To
	return b
}

func (b *MmsBuilder) Text(text string) *MmsBuilder {
	b.message.Text = &text
	return b
}

func (b *MmsBuilder) Title(title string) *MmsBuilder {
	b.message.Title = &title
	return b
}

func (b *MmsBuilder) FileKey(fileKey []string) *MmsBuilder {
	b.message.FileKey = fileKey
	return b
}

func (b *MmsBuilder) From(from string) *MmsBuilder {
	b.message.From = &from
	return b
}

func (b *MmsBuilder) Ref(ref string) *MmsBuilder {
	b.message.Ref = &ref
	return b
}

func (b *MmsBuilder) OriginCID(originCID string) *MmsBuilder {
	b.message.OriginCID = &originCID
	return b
}

func NewMmsSender(authorization string, httpClient *http.Client) *MmsSender {
	c := core.NewClient(httpClient)
	return &MmsSender{authorization: authorization, client: c}
}

func (sender *MmsSender) SendMessage(req *MMS) (*SimpleSendResponse, error) {
	if response, err := sendMessage(sender.client, mmsSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseSimpleSendResponse(sender.client, response)
	}
}
