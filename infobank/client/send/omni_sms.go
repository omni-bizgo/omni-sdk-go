package send

type OmniSMS struct {
	From      *string `json:"from"`
	Text      *string `json:"text"`
	Ttl       *string `json:"ttl,omitempty"`
	OriginCID *string `json:"originCID,omitempty"`
}

type OmniSmsBuilder struct {
	message OmniSMS
}

func NewOmniSmsBuilder() *OmniSmsBuilder {
	return &OmniSmsBuilder{message: OmniSMS{}}
}

func (b *OmniSmsBuilder) Build() OmniSMS {
	return b.message
}

func (b *OmniSmsBuilder) From(from string) *OmniSmsBuilder {
	b.message.From = &from
	return b
}

func (b *OmniSmsBuilder) Text(test string) *OmniSmsBuilder {
	b.message.Text = &test
	return b
}

func (b *OmniSmsBuilder) Ttl(ttl string) *OmniSmsBuilder {
	b.message.Ttl = &ttl
	return b
}

func (b *OmniSmsBuilder) OriginCID(originCID string) *OmniSmsBuilder {
	b.message.OriginCID = &originCID
	return b
}
