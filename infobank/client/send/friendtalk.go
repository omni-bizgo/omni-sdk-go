package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

const (
	friendTalkSubPath = "/v1/send/friendtalk"
)

type FriendtalkSender struct {
	authorization string
	client        *core.HttpClient
}

type Friendtalk struct {
	SenderKey    *string       `json:"senderKey"`
	To           *string       `json:"to"`
	Text         *string       `json:"text"`
	Button       []KakaoButton `json:"button,omitempty"`
	KakaoMsgType *KakaoMsgType `json:"msgType"`
	ImgUrl       *string       `json:"imgUrl,omitempty"`
	Fallback     *Fallback     `json:"fallback,omitempty"`
	Ref          *string       `json:"ref,omitempty"`
}

type FriendtalkBuilder struct {
	message Friendtalk
}

func NewFriendtalkBuilder() *FriendtalkBuilder {
	return &FriendtalkBuilder{message: Friendtalk{}}
}

func (b *FriendtalkBuilder) Build() Friendtalk {
	return b.message
}

func (b *FriendtalkBuilder) SenderKey(senderKey string) *FriendtalkBuilder {
	b.message.SenderKey = &senderKey
	return b
}

func (b *FriendtalkBuilder) To(to string) *FriendtalkBuilder {
	b.message.To = &to
	return b
}

func (b *FriendtalkBuilder) Text(text string) *FriendtalkBuilder {
	b.message.Text = &text
	return b
}

func (b *FriendtalkBuilder) Button(button []KakaoButton) *FriendtalkBuilder {
	b.message.Button = button
	return b
}

func (b *FriendtalkBuilder) KakaoMsgType(kakaoMsgType KakaoMsgType) *FriendtalkBuilder {
	b.message.KakaoMsgType = &kakaoMsgType
	return b
}

func (b *FriendtalkBuilder) ImgUrl(imgUrl string) *FriendtalkBuilder {
	b.message.ImgUrl = &imgUrl
	return b
}

func (b *FriendtalkBuilder) Fallback(fallback Fallback) *FriendtalkBuilder {
	b.message.Fallback = &fallback
	return b
}

func (b *FriendtalkBuilder) Ref(ref string) *FriendtalkBuilder {
	b.message.Ref = &ref
	return b
}

func NewFriendTalkSender(authorization string, httpClient *http.Client) *FriendtalkSender {
	c := core.NewClient(httpClient)
	return &FriendtalkSender{authorization: authorization, client: c}
}

func (sender *FriendtalkSender) SendMessage(req *Friendtalk) (*SimpleSendResponse, error) {
	if response, err := sendMessage(sender.client, friendTalkSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseSimpleSendResponse(sender.client, response)
	}
}
