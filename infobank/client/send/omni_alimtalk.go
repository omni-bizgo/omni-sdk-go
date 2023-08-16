package send

type OmniAlimTalk struct {
	SenderKey    *string          `json:"senderKey"`
	Text         *string          `json:"text"`
	TemplateCode *string          `json:"templateCode"`
	KakaoMsgType *KakaoMsgType    `json:"msgType"`
	Title        *string          `json:"title,omitempty"`
	Attachment   *KakaoAttachment `json:"attachment,omitempty"`
	Price        *string          `json:"price,omitempty"`
	CurrencyType *string          `json:"currencyType,omitempty"`
	Supplement   *KakaoSuplement  `json:"supplement,omitempty"`
}

type OmniAlimTalkBuilder struct {
	message OmniAlimTalk
}

func NewOmniAlimTalkBuilder() *OmniAlimTalkBuilder {
	return &OmniAlimTalkBuilder{message: OmniAlimTalk{}}
}

func (b *OmniAlimTalkBuilder) Build() OmniAlimTalk {
	return b.message
}

func (b *OmniAlimTalkBuilder) SenderKey(senderKey string) *OmniAlimTalkBuilder {
	b.message.SenderKey = &senderKey
	return b
}

func (b *OmniAlimTalkBuilder) Text(text string) *OmniAlimTalkBuilder {
	b.message.Text = &text
	return b
}

func (b *OmniAlimTalkBuilder) TemplateCode(templateCode string) *OmniAlimTalkBuilder {
	b.message.TemplateCode = &templateCode
	return b
}

func (b *OmniAlimTalkBuilder) KakaoMsgType(kakaoMsgType KakaoMsgType) *OmniAlimTalkBuilder {
	b.message.KakaoMsgType = &kakaoMsgType
	return b
}

func (b *OmniAlimTalkBuilder) Title(title string) *OmniAlimTalkBuilder {
	b.message.Title = &title
	return b
}

func (b *OmniAlimTalkBuilder) Attachment(attachment KakaoAttachment) *OmniAlimTalkBuilder {
	b.message.Attachment = &attachment
	return b
}

func (b *OmniAlimTalkBuilder) Price(price string) *OmniAlimTalkBuilder {
	b.message.Price = &price
	return b
}

func (b *OmniAlimTalkBuilder) CurrencyType(currencyType string) *OmniAlimTalkBuilder {
	b.message.CurrencyType = &currencyType
	return b
}

func (b *OmniAlimTalkBuilder) Supplement(supplement KakaoSuplement) *OmniAlimTalkBuilder {
	b.message.Supplement = &supplement
	return b
}
