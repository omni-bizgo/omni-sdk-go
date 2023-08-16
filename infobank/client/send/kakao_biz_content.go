package send

type KakaoMsgType string

const (
	MSGTYPE_ALIMTALK              KakaoMsgType = "AT"
	MSGTYPE_ALIMTALK_IMAGE        KakaoMsgType = "AI"
	MSGTYPE_FRIENDTALK            KakaoMsgType = "FT"
	MSGTYPE_FRIENDTALK_IMAGE      KakaoMsgType = "FI"
	MSGTYPE_FRIENDTALK_WIDE_IMAGE KakaoMsgType = "FW"
)

type KakaoButtonType string

const (
	BUTTON_WEB_LINK          KakaoButtonType = "WL"
	BUTTON_APP_LINK          KakaoButtonType = "AL"
	BUTTON_BOT_KEYWORD       KakaoButtonType = "BK"
	BUTTON_MESSAGE_DELIVERY  KakaoButtonType = "MD"
	BUTTON_DELIVERY_SCAN     KakaoButtonType = "DS"
	BUTTON_CONSULTBOT_CHANGE KakaoButtonType = "BC"
	BUTTON_CHATBOT_CHANGE    KakaoButtonType = "BT"
	BUTTON_ADD_CHANNEL       KakaoButtonType = "AC"
	BUTTON_BIZ_FORM          KakaoButtonType = "BF"
)

type KakaoAttachment struct {
	Button        []KakaoButton  `json:"button,omitempty"`
	Item          *KakaoItemInfo `json:"item,omitempty"`
	ItemHighlight *KakaoItem     `json:"itemHighlight,omitempty"`
	Image         *KakaoImage    `json:"image,omitempty"`
}

type KakaoSuplement struct {
	QuickReply []KakaoButton `json:"quickReply"`
}

type KakaoButton struct {
	Name          *string          `json:"name"`
	Type          *KakaoButtonType `json:"type"`
	UrlPC         *string          `json:"urlPC,omitempty"`
	UrlMobile     *string          `json:"urlMobile,omitempty"`
	SchemeIOS     *string          `json:"schemeIOS,omitempty"`
	SchemeAndroid *string          `json:"schemeAndroid,omitempty"`
	Target        *string          `json:"target,omitempty"`
	ChatExtra     *string          `json:"chatExtra,omitempty"`
	ChatEvent     *string          `json:"chatEvent,omitempty"`
	BizFormKey    *string          `json:"bizFormKey,omitempty"`
	BizFormId     *string          `json:"bizFormId,omitempty"`
}

type KakaoItemInfo struct {
	List    []KakaoItem `json:"list,omitempty"`
	Summary *KakaoItem  `json:"summary,omitempty"`
}

type KakaoItem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type KakaoImage struct {
	ImgUrl  *string `json:"imgUrl"`
	ImgLink *string `json:"imgLink,omitempty"`
}

type Fallback struct {
	Type      *FallbackServiceType `json:"type"`
	Title     *string              `json:"title,omitempty"`
	Text      *string              `json:"text"`
	FileKey   []string             `json:"fileKey,omitempty"`
	From      *string              `json:"from"`
	OriginCID *string              `json:"originCID,omitempty"`
}

type KakaoAttachmentBuilder struct {
	attachment KakaoAttachment
}

type KakaoSuplementBuilder struct {
	supplement KakaoSuplement
}

type KakaoButtonBuilder struct {
	button KakaoButton
}

type KakaoItemInfoBuilder struct {
	itemInfo KakaoItemInfo
}

type KakaoImageBuilder struct {
	image KakaoImage
}

type KakaoFallbackBuilder struct {
	fallack Fallback
}

func (b *KakaoAttachmentBuilder) NewKakaoAttachmentBuilder() *KakaoAttachmentBuilder {
	return &KakaoAttachmentBuilder{attachment: KakaoAttachment{}}
}

func (b *KakaoAttachmentBuilder) Build() KakaoAttachment {
	return b.attachment
}

func (b *KakaoAttachmentBuilder) Button(button []KakaoButton) *KakaoAttachmentBuilder {
	b.attachment.Button = button
	return b
}

func (b *KakaoAttachmentBuilder) Item(item KakaoItemInfo) *KakaoAttachmentBuilder {
	b.attachment.Item = &item
	return b
}

func (b *KakaoAttachmentBuilder) ItemHighlight(itemHighlight KakaoItem) *KakaoAttachmentBuilder {
	b.attachment.ItemHighlight = &itemHighlight
	return b
}

func (b *KakaoAttachmentBuilder) Image(image KakaoImage) *KakaoAttachmentBuilder {
	b.attachment.Image = &image
	return b
}

func NewKakaoSuplementBuilder() *KakaoSuplementBuilder {
	return &KakaoSuplementBuilder{supplement: KakaoSuplement{}}
}

func (b *KakaoSuplementBuilder) Build() KakaoSuplement {
	return b.supplement
}

func (b *KakaoSuplementBuilder) QuickReply(quickReply []KakaoButton) *KakaoSuplementBuilder {
	b.supplement.QuickReply = quickReply
	return b
}

func NewKakaoButtonBuilder() *KakaoButtonBuilder {
	return &KakaoButtonBuilder{button: KakaoButton{}}
}

func (b *KakaoButtonBuilder) Build() KakaoButton {
	return b.button
}

func (b *KakaoButtonBuilder) Name(name string) *KakaoButtonBuilder {
	b.button.Name = &name
	return b
}

func (b *KakaoButtonBuilder) Type(types KakaoButtonType) *KakaoButtonBuilder {
	b.button.Type = &types
	return b
}

func (b *KakaoButtonBuilder) UrlPC(urlPC string) *KakaoButtonBuilder {
	b.button.UrlPC = &urlPC
	return b
}

func (b *KakaoButtonBuilder) UrlMobile(urlMobile string) *KakaoButtonBuilder {
	b.button.UrlMobile = &urlMobile
	return b
}

func (b *KakaoButtonBuilder) SchemeIOS(schemeIOS string) *KakaoButtonBuilder {
	b.button.SchemeIOS = &schemeIOS
	return b
}

func (b *KakaoButtonBuilder) SchemeAndroid(schemeAndroid string) *KakaoButtonBuilder {
	b.button.SchemeAndroid = &schemeAndroid
	return b
}

func (b *KakaoButtonBuilder) Target(target string) *KakaoButtonBuilder {
	b.button.Target = &target
	return b
}

func (b *KakaoButtonBuilder) ChatExtra(chatExtra string) *KakaoButtonBuilder {
	b.button.ChatExtra = &chatExtra
	return b
}

func (b *KakaoButtonBuilder) ChatEvent(chatEvent string) *KakaoButtonBuilder {
	b.button.ChatEvent = &chatEvent
	return b
}

func (b *KakaoButtonBuilder) BizFormKey(bizFormKey string) *KakaoButtonBuilder {
	b.button.BizFormKey = &bizFormKey
	return b
}

func (b *KakaoButtonBuilder) BizFormId(bizFormId string) *KakaoButtonBuilder {
	b.button.BizFormId = &bizFormId
	return b
}

func NewKakaoItemInfoBuilder() *KakaoItemInfoBuilder {
	return &KakaoItemInfoBuilder{itemInfo: KakaoItemInfo{}}
}

func (b *KakaoItemInfoBuilder) Build() KakaoItemInfo {
	return b.itemInfo
}

func (b *KakaoItemInfoBuilder) List(list []KakaoItem) *KakaoItemInfoBuilder {
	b.itemInfo.List = list
	return b
}

func (b *KakaoItemInfoBuilder) Summary(summary KakaoItem) *KakaoItemInfoBuilder {
	b.itemInfo.Summary = &summary
	return b
}

type KakaoItemBuilder struct {
	item KakaoItem
}

func NewKakaoItemBuilder() *KakaoItemBuilder {
	return &KakaoItemBuilder{item: KakaoItem{}}
}

func (b *KakaoItemBuilder) Build() KakaoItem {
	return b.item
}

func (b *KakaoItemBuilder) Title(title string) *KakaoItemBuilder {
	b.item.Title = &title
	return b
}

func (b *KakaoItemBuilder) Description(description string) *KakaoItemBuilder {
	b.item.Description = &description
	return b
}

func NewKakaoImageBuilder() *KakaoImageBuilder {
	return &KakaoImageBuilder{image: KakaoImage{}}
}

func (b *KakaoImageBuilder) Build() KakaoImage {
	return b.image
}

func (b *KakaoImageBuilder) ImgUrl(imgUrl string) *KakaoImageBuilder {
	b.image.ImgUrl = &imgUrl
	return b
}

func (b *KakaoImageBuilder) ImgLink(imgLink string) *KakaoImageBuilder {
	b.image.ImgLink = &imgLink
	return b
}

func NewKakaoFallbackBuilder() *KakaoFallbackBuilder {
	return &KakaoFallbackBuilder{fallack: Fallback{}}
}

func (b *KakaoFallbackBuilder) Build() Fallback {
	return b.fallack
}

func (b *KakaoFallbackBuilder) Type(types FallbackServiceType) *KakaoFallbackBuilder {
	b.fallack.Type = &types
	return b
}

func (b *KakaoFallbackBuilder) Title(title string) *KakaoFallbackBuilder {
	b.fallack.Title = &title
	return b
}

func (b *KakaoFallbackBuilder) Text(text string) *KakaoFallbackBuilder {
	b.fallack.Text = &text
	return b
}

func (b *KakaoFallbackBuilder) FileKey(fileKey []string) *KakaoFallbackBuilder {
	b.fallack.FileKey = fileKey
	return b
}

func (b *KakaoFallbackBuilder) From(from string) *KakaoFallbackBuilder {
	b.fallack.From = &from
	return b
}

func (b *KakaoFallbackBuilder) OriginCID(originCID string) *KakaoFallbackBuilder {
	b.fallack.OriginCID = &originCID
	return b
}
