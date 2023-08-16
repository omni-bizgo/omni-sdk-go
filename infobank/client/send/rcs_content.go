package send

type RcsContent struct {
	StandAlone *StandAlone            `json:"standalone,omitempty"`
	Carousel   []Carousel             `json:"carousel,omitempty"`
	Template   map[string]interface{} `json:"template,omitempty"`
}

type StandAlone struct {
	Title      *string         `json:"title,omitempty"`
	Text       *string         `json:"text,omitempty"`
	Media      *string         `json:"media,omitempty"`
	MediaUrl   *string         `json:"mediaUrl,omitempty"`
	Button     []RcsButton     `json:"button,omitempty"`
	SubContent []RcsSubContent `json:"subContent,omitempty"`
}

type Carousel struct {
	Title    *string     `json:"title,omitempty"`
	Text     *string     `json:"text,omitempty"`
	Media    *string     `json:"media,omitempty"`
	MediaUrl *string     `json:"mediaUrl,omitempty"`
	FileKey  *string     `json:"fileKey,omitempty"`
	Button   []RcsButton `json:"button,omitempty"`
}

type RcsTemplate struct {
	Description *string
	SubContent  []RcsSubContent
	CustomMap   map[string]interface{}
}

type RcsButtonType string

const (
	ButtonURL      RcsButtonType = "URL"
	ButtonMapLoc   RcsButtonType = "MAP_LOC"
	ButtonMapQry   RcsButtonType = "MAP_QRY"
	ButtonMapSend  RcsButtonType = "MAP_SEND"
	ButtonCalendar RcsButtonType = "CALENDAR"
	ButtonCopy     RcsButtonType = "COPY"
	ButtonCOMT     RcsButtonType = "COM_T"
	ButtonCOMV     RcsButtonType = "COM_V"
	ButtonDial     RcsButtonType = "DIAL"
)

type RcsButton struct {
	Type        *RcsButtonType `json:"type"`
	Name        *string        `json:"name,omitempty"`
	Url         *string        `json:"url,omitempty"`
	Label       *string        `json:"label,omitempty"`
	Latitude    *string        `json:"latitude,omitempty"`
	Longitude   *string        `json:"longitude,omitempty"`
	FallbackUrl *string        `json:"fallbackUrl,omitempty"`
	Query       *string        `json:"query,omitempty"`
	StartTime   *string        `json:"startTime,omitempty"`
	EndTime     *string        `json:"endTime,omitempty"`
	Title       *string        `json:"title,omitempty"`
	Description *string        `json:"description,omitempty"`
	Text        *string        `json:"text,omitempty"`
	PhoneNumber *string        `json:"phoneNumber,omitempty"`
}

type RcsSubContent struct {
	SubTitle    *string `json:"subTitle,omitempty"`
	SubDesc     *string `json:"subDesc,omitempty"`
	SubMedia    *string `json:"subMedia,omitempty"`
	SubMediaUrl *string `json:"subMediaUrl,omitempty"`
}

type RcsContentBuilder struct {
	message [1]RcsContent
}

func NewRcsContentBuilder() *RcsContentBuilder {
	var message [1]RcsContent
	return &RcsContentBuilder{message: message}
}

func (b *RcsContentBuilder) Build() RcsContent {
	return b.message[0]
}

func (b *RcsContentBuilder) StandAlone(standAlone StandAlone) *RcsContentBuilder {
	b.message[0] = RcsContent{}
	b.message[0].StandAlone = &standAlone
	return b
}

func (b *RcsContentBuilder) Carousel(carousel []Carousel) *RcsContentBuilder {
	b.message[0] = RcsContent{}
	b.message[0].Carousel = carousel
	return b
}

func (b *RcsContentBuilder) Template(template RcsTemplate) *RcsContentBuilder {
	b.message[0] = RcsContent{}
	b.message[0].Template = make(map[string]interface{})
	if template.Description != nil {
		b.message[0].Template["description"] = template.Description
	}
	if template.SubContent != nil {
		b.message[0].Template["subContent"] = template.SubContent
	}
	if template.CustomMap != nil {
		for key, value := range template.CustomMap {
			b.message[0].Template[key] = value
		}
	}
	return b
}

type StandAloneBuilder struct {
	content StandAlone
}

func NewStandAloneBuilder() *StandAloneBuilder {
	return &StandAloneBuilder{content: StandAlone{}}
}

func (b *StandAloneBuilder) Build() StandAlone {
	return b.content
}

func (b *StandAloneBuilder) Title(title string) *StandAloneBuilder {
	b.content.Title = &title
	return b
}

func (b *StandAloneBuilder) Text(text string) *StandAloneBuilder {
	b.content.Text = &text
	return b
}

func (b *StandAloneBuilder) Media(media string) *StandAloneBuilder {
	b.content.Media = &media
	return b
}

func (b *StandAloneBuilder) MediaUrl(mediaUrl string) *StandAloneBuilder {
	b.content.MediaUrl = &mediaUrl
	return b
}

func (b *StandAloneBuilder) Button(button []RcsButton) *StandAloneBuilder {
	b.content.Button = button
	return b
}

func (b *StandAloneBuilder) SubContent(subContent []RcsSubContent) *StandAloneBuilder {
	b.content.SubContent = subContent
	return b
}

type CarouselBuilder struct {
	content Carousel
}

func NewCarouselBuilder() *CarouselBuilder {
	return &CarouselBuilder{content: Carousel{}}
}

func (b *CarouselBuilder) Build() Carousel {
	return b.content
}

func (b *CarouselBuilder) Title(title string) *CarouselBuilder {
	b.content.Title = &title
	return b
}

func (b *CarouselBuilder) Text(text string) *CarouselBuilder {
	b.content.Text = &text
	return b
}

func (b *CarouselBuilder) Media(media string) *CarouselBuilder {
	b.content.Media = &media
	return b
}

func (b *CarouselBuilder) MediaUrl(mediaUrl string) *CarouselBuilder {
	b.content.MediaUrl = &mediaUrl
	return b
}

func (b *CarouselBuilder) Button(button []RcsButton) *CarouselBuilder {
	b.content.Button = button
	return b
}

func (b *CarouselBuilder) FileKey(fileKey string) *CarouselBuilder {
	b.content.FileKey = &fileKey
	return b
}

type RcsTemplateBuilder struct {
	content RcsTemplate
}

func NewRcsTemplateBuilder() *RcsTemplateBuilder {
	return &RcsTemplateBuilder{content: RcsTemplate{}}
}

func (b *RcsTemplateBuilder) Build() RcsTemplate {
	return b.content
}

func (b *RcsTemplateBuilder) Description(description string) *RcsTemplateBuilder {
	b.content.Description = &description
	return b
}

func (b *RcsTemplateBuilder) SubContent(subContent []RcsSubContent) *RcsTemplateBuilder {
	b.content.SubContent = subContent
	return b
}

func (b *RcsTemplateBuilder) CustomMap(customMap map[string]interface{}) *RcsTemplateBuilder {
	b.content.CustomMap = customMap
	return b
}

type RcsButtonBuilder struct {
	button RcsButton
}

func NewRcsButtonBuilder() *RcsButtonBuilder {
	return &RcsButtonBuilder{RcsButton{}}
}

func (b *RcsButtonBuilder) Build() RcsButton {
	return b.button
}

func (b *RcsButtonBuilder) Type(types RcsButtonType) *RcsButtonBuilder {
	b.button.Type = &types
	return b
}

func (b *RcsButtonBuilder) Name(name string) *RcsButtonBuilder {
	b.button.Name = &name
	return b
}

func (b *RcsButtonBuilder) Url(url string) *RcsButtonBuilder {
	b.button.Url = &url
	return b
}

func (b *RcsButtonBuilder) Label(label string) *RcsButtonBuilder {
	b.button.Label = &label
	return b
}

func (b *RcsButtonBuilder) Latitude(latitude string) *RcsButtonBuilder {
	b.button.Latitude = &latitude
	return b
}

func (b *RcsButtonBuilder) Longitude(longitude string) *RcsButtonBuilder {
	b.button.Longitude = &longitude
	return b
}

func (b *RcsButtonBuilder) FallbackUrl(fallbackUrl string) *RcsButtonBuilder {
	b.button.FallbackUrl = &fallbackUrl
	return b
}

func (b *RcsButtonBuilder) Query(query string) *RcsButtonBuilder {
	b.button.Query = &query
	return b
}

func (b *RcsButtonBuilder) StartTime(startTime string) *RcsButtonBuilder {
	b.button.StartTime = &startTime
	return b
}

func (b *RcsButtonBuilder) EndTime(endTime string) *RcsButtonBuilder {
	b.button.EndTime = &endTime
	return b
}

func (b *RcsButtonBuilder) Title(title string) *RcsButtonBuilder {
	b.button.Title = &title
	return b
}

func (b *RcsButtonBuilder) Description(description string) *RcsButtonBuilder {
	b.button.Description = &description
	return b
}

func (b *RcsButtonBuilder) Text(text string) *RcsButtonBuilder {
	b.button.Text = &text
	return b
}

func (b *RcsButtonBuilder) PhoneNumber(phoneNumber string) *RcsButtonBuilder {
	b.button.PhoneNumber = &phoneNumber
	return b
}

type RcsSubContentBuilder struct {
	subContent RcsSubContent
}

func NewRcsSubContentBuilder() *RcsSubContentBuilder {
	return &RcsSubContentBuilder{subContent: RcsSubContent{}}
}

func (b *RcsSubContentBuilder) Build() RcsSubContent {
	return b.subContent
}

func (b *RcsSubContentBuilder) SubTitle(subTitle string) *RcsSubContentBuilder {
	b.subContent.SubTitle = &subTitle
	return b
}

func (b *RcsSubContentBuilder) SubDesc(subDesc string) *RcsSubContentBuilder {
	b.subContent.SubDesc = &subDesc
	return b
}

func (b *RcsSubContentBuilder) SubMedia(subMedia string) *RcsSubContentBuilder {
	b.subContent.SubMedia = &subMedia
	return b
}

func (b *RcsSubContentBuilder) SubMediaUrl(subMediaUrl string) *RcsSubContentBuilder {
	b.subContent.SubMediaUrl = &subMediaUrl
	return b
}
