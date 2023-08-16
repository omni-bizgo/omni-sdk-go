package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

const (
	rcsSubPath = "/v1/send/rcs"
)

type RcsSender struct {
	authorization string
	client        *core.HttpClient
}

type RCS struct {
	From         *string      `json:"from"`
	To           *string      `json:"to"`
	FormatId     *string      `json:"formatId"`
	BrandKey     *string      `json:"brandKey"`
	BrandId      *string      `json:"brandId"`
	Content      *RcsContent  `json:"content"`
	ExpiryOption *string      `json:"expiryOption,omitempty"`
	Header       *string      `json:"header,omitempty"`
	Footer       *string      `json:"footer,omitempty"`
	Fallback     *RcsFallback `json:"fallback,omitempty"`
	Ref          *string      `json:"ref,omitempty"`
	RcsData      *string      `json:"rcsData,omitempty"`
}

type FallbackServiceType string

const (
	FALLBACK_SMS FallbackServiceType = "SMS"
	FALLBACK_MMS FallbackServiceType = "MMS"
)

type RcsFallback struct {
	Type      *FallbackServiceType `json:"type"`
	Title     *string              `json:"title,omitempty"`
	Text      *string              `json:"text"`
	FileKey   []string             `json:"fileKey,omitempty"`
	OriginCID *string              `json:"originCID,omitempty"`
}

type RcsBuilder struct {
	message RCS
}

func NewRcsBuilder() *RcsBuilder {
	return &RcsBuilder{message: RCS{}}
}

func (b *RcsBuilder) Build() RCS {
	return b.message
}

func (b *RcsBuilder) From(from string) *RcsBuilder {
	b.message.From = &from
	return b
}

func (b *RcsBuilder) To(from string) *RcsBuilder {
	b.message.To = &from
	return b
}

func (b *RcsBuilder) FormatId(from string) *RcsBuilder {
	b.message.FormatId = &from
	return b
}

func (b *RcsBuilder) Content(content RcsContent) *RcsBuilder {
	b.message.Content = &content
	return b
}

func (b *RcsBuilder) ExpiryOption(expiryOption string) *RcsBuilder {
	b.message.ExpiryOption = &expiryOption
	return b
}

func (b *RcsBuilder) Header(header string) *RcsBuilder {
	b.message.Header = &header
	return b
}

func (b *RcsBuilder) Footer(footer string) *RcsBuilder {
	b.message.Footer = &footer
	return b
}

func (b *RcsBuilder) Fallback(fallback RcsFallback) *RcsBuilder {
	b.message.Fallback = &fallback
	return b
}

func (b *RcsBuilder) Ref(ref string) *RcsBuilder {
	b.message.Ref = &ref
	return b
}

type RcsFallbackBuilder struct {
	fallback RcsFallback
}

func NewRcsFallbackBuilder() *RcsFallbackBuilder {
	return &RcsFallbackBuilder{fallback: RcsFallback{}}
}

func (b *RcsFallbackBuilder) Build() RcsFallback {
	return b.fallback
}

func (b *RcsFallbackBuilder) Type(types FallbackServiceType) *RcsFallbackBuilder {
	b.fallback.Type = &types
	return b
}

func (b *RcsFallbackBuilder) Title(title string) *RcsFallbackBuilder {
	b.fallback.Title = &title
	return b
}

func (b *RcsFallbackBuilder) Text(text string) *RcsFallbackBuilder {
	b.fallback.Text = &text
	return b
}

func (b *RcsFallbackBuilder) FileKey(fileKey []string) *RcsFallbackBuilder {
	b.fallback.FileKey = fileKey
	return b
}

func (b *RcsFallbackBuilder) OriginCID(originCID string) *RcsFallbackBuilder {
	b.fallback.OriginCID = &originCID
	return b
}

func NewRcsSender(authorization string, httpClient *http.Client) *RcsSender {
	c := core.NewClient(httpClient)
	return &RcsSender{authorization: authorization, client: c}
}

func (sender *RcsSender) SendMessage(req *RCS) (*SimpleSendResponse, error) {
	if response, err := sendMessage(sender.client, rcsSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseSimpleSendResponse(sender.client, response)
	}
}
