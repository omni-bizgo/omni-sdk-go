package send

type OmniFriendTalk struct {
	SenderKey    *string          `json:"senderKey"`
	Text         *string          `json:"text"`
	KakaoMsgType *KakaoMsgType    `json:"msgType"`
	AdFlag       *string          `json:"adFlag,omitempty"`
	Attachment   *KakaoAttachment `json:"attachment,omitempty"`
}

type OmniFriendTalkBuilder struct {
	message OmniFriendTalk
}

func NewOmniFriendTalkBuilder() *OmniFriendTalkBuilder {
	return &OmniFriendTalkBuilder{message: OmniFriendTalk{}}
}

func (b *OmniFriendTalkBuilder) Build() OmniFriendTalk {
	return b.message
}

func (b *OmniFriendTalkBuilder) SenderKey(senderKey string) *OmniFriendTalkBuilder {
	b.message.SenderKey = &senderKey
	return b
}

func (b *OmniFriendTalkBuilder) Text(text string) *OmniFriendTalkBuilder {
	b.message.Text = &text
	return b
}

func (b *OmniFriendTalkBuilder) KakaoMsgType(kakaoMsgType KakaoMsgType) *OmniFriendTalkBuilder {
	b.message.KakaoMsgType = &kakaoMsgType
	return b
}

func (b *OmniFriendTalkBuilder) AdFlag(adFlag string) *OmniFriendTalkBuilder {
	b.message.AdFlag = &adFlag
	return b
}

func (b *OmniFriendTalkBuilder) Attachment(attachment KakaoAttachment) *OmniFriendTalkBuilder {
	b.message.Attachment = &attachment
	return b
}
