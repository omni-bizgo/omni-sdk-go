package send

type OmniMMS struct {
	From      *string  `json:"from"`
	Text      *string  `json:"text"`
	Title     *string  `json:"title,omitempty"`
	FileKey   []string `json:"fileKey,omitempty"`
	Ttl       *string  `json:"ttl,omitempty"`
	OriginCID *string  `json:"originCID,omitempty"`
}

type OmniMmsBuilder struct {
	message OmniMMS
}

func NewOmniMmsBuilderBuilder() *OmniMmsBuilder {
	return &OmniMmsBuilder{message: OmniMMS{}}
}

func (b *OmniMmsBuilder) Build() OmniMMS {
	return b.message
}

func (b *OmniMmsBuilder) From(from string) *OmniMmsBuilder {
	b.message.From = &from
	return b
}

func (b *OmniMmsBuilder) Text(text string) *OmniMmsBuilder {
	b.message.Text = &text
	return b
}

func (b *OmniMmsBuilder) Title(title string) *OmniMmsBuilder {
	b.message.Title = &title
	return b
}

func (b *OmniMmsBuilder) FileKey(fileKey []string) *OmniMmsBuilder {
	b.message.FileKey = fileKey
	return b
}

func (b *OmniMmsBuilder) Ttl(ttl string) *OmniMmsBuilder {
	b.message.Ttl = &ttl
	return b
}

func (b *OmniMmsBuilder) OriginCID(originCID string) *OmniMmsBuilder {
	b.message.OriginCID = &originCID
	return b
}
