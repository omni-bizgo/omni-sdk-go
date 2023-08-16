package send

type OmniRCS struct {
	From         *string    `json:"from"`
	FormatId     *string    `json:"formatId"`
	Content      RcsContent `json:"content"`
	GroupId      *string    `json:"groupId,omitempty"`
	ExpiryOption *string    `json:"expiryOption,omitempty"`
	CopyAllowed  *string    `json:"copyAllowed,omitempty"`
	Header       *string    `json:"header,omitempty"`
	Footer       *string    `json:"footer,omitempty"`
	Ttl          *string    `json:"ttl,omitempty"`
	BrandId      *string    `json:"brandId,omitempty"`
	BrandKey     *string    `json:"brandKey,omitempty"`
	AgencyId     *string    `json:"agencyId,omitempty"`
	AgencyKey    *string    `json:"agencyKey,omitempty"`
	RcsData      *string    `json:"rcsData,omitempty"`
}

type OmniRcsBuilder struct {
	message OmniRCS
}

func NewOmniRcsBuilder() *OmniRcsBuilder {
	return &OmniRcsBuilder{message: OmniRCS{}}
}

func (b *OmniRcsBuilder) Build() OmniRCS {
	return b.message
}

func (b *OmniRcsBuilder) From(from string) *OmniRcsBuilder {
	b.message.From = &from
	return b
}

func (b *OmniRcsBuilder) FormatId(from string) *OmniRcsBuilder {
	b.message.FormatId = &from
	return b
}

func (b *OmniRcsBuilder) Content(content RcsContent) *OmniRcsBuilder {
	b.message.Content = content
	return b
}

func (b *OmniRcsBuilder) GroupId(groupId string) *OmniRcsBuilder {
	b.message.GroupId = &groupId
	return b
}

func (b *OmniRcsBuilder) ExpiryOption(expiryOption string) *OmniRcsBuilder {
	b.message.ExpiryOption = &expiryOption
	return b
}

func (b *OmniRcsBuilder) Header(copyAllowed string) *OmniRcsBuilder {
	b.message.Header = &copyAllowed
	return b
}

func (b *OmniRcsBuilder) Footer(footer string) *OmniRcsBuilder {
	b.message.Footer = &footer
	return b
}

func (b *OmniRcsBuilder) Ttl(ttl string) *OmniRcsBuilder {
	b.message.Ttl = &ttl
	return b
}

func (b *OmniRcsBuilder) BrandId(brandId string) *OmniRcsBuilder {
	b.message.BrandId = &brandId
	return b
}

func (b *OmniRcsBuilder) BrandKey(brandKey string) *OmniRcsBuilder {
	b.message.BrandKey = &brandKey
	return b
}

func (b *OmniRcsBuilder) AgencyId(agencyId string) *OmniRcsBuilder {
	b.message.AgencyId = &agencyId
	return b
}

func (b *OmniRcsBuilder) AgencyKey(agencyKey string) *OmniRcsBuilder {
	b.message.AgencyKey = &agencyKey
	return b
}

func (b *OmniRcsBuilder) CopyAllowed(copyAllowed string) *OmniRcsBuilder {
	b.message.CopyAllowed = &copyAllowed
	return b
}
