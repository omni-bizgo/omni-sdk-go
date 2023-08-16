package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
	"reflect"
)

const (
	omniSubPath = "/v1/send/omni"
)

type OmniSender struct {
	authorization string
	client        *core.HttpClient
}

type Omni struct {
	Destinations []Destination `json:"destinations"`
	MessageFlow  []OmniMessage `json:"messageFlow,omitempty"`
	MessageForm  *string       `json:"messageForm,omitempty"`
	PaymentCode  *string       `json:"paymentCode,omitempty"`
	Ref          *string       `json:"ref,omitempty"`
}

type Destination struct {
	To           *string           `json:"to"`
	ClientKey    *string           `json:"clientKey,omitempty"`
	ReplaceWords map[string]string `json:"replaceWords,omitempty"`
}

type OmniMessage struct {
	SMS        *OmniSMS        `json:"sms,omitempty"`
	MMS        *OmniMMS        `json:"mms,omitempty"`
	RCS        *OmniRCS        `json:"rcs,omitempty"`
	AlimTalk   *OmniAlimTalk   `json:"alimtalk,omitempty"`
	FriendTalk *OmniFriendTalk `json:"friendtalk,omitempty"`
}

type OmniBuilder struct {
	request Omni
}

type DestinationBuilder struct {
	destination Destination
}

type MessageFlowBuilder struct {
	flow []OmniMessage
}

func NewOmniBuilder() *OmniBuilder {
	return &OmniBuilder{request: Omni{}}
}

func (b *OmniBuilder) Build() Omni {
	return b.request
}

func (b *OmniBuilder) Destinations(destinations []Destination) *OmniBuilder {
	b.request.Destinations = destinations
	return b
}

func (b *OmniBuilder) MessageFlow(flow []OmniMessage) *OmniBuilder {
	b.request.MessageFlow = flow
	return b
}

func (b *OmniBuilder) MessageForm(messageForm string) *OmniBuilder {
	b.request.MessageForm = &messageForm
	return b
}

func (b *OmniBuilder) PaymentCode(paymentCode string) *OmniBuilder {
	b.request.PaymentCode = &paymentCode
	return b
}

func (b *OmniBuilder) Ref(ref string) *OmniBuilder {
	b.request.Ref = &ref
	return b
}

func NewDestinationBuilder() *DestinationBuilder {
	return &DestinationBuilder{destination: Destination{}}
}

func (b *DestinationBuilder) Build() Destination {
	return b.destination
}

func (b *DestinationBuilder) To(to string) *DestinationBuilder {
	b.destination.To = &to
	return b
}

func (b *DestinationBuilder) ClientKey(clientKey string) *DestinationBuilder {
	b.destination.ClientKey = &clientKey
	return b
}

func (b *DestinationBuilder) ReplaceWords(replaceWords map[string]string) *DestinationBuilder {
	b.destination.ReplaceWords = replaceWords
	return b
}

func NewMessageFlowBuilder() *MessageFlowBuilder {
	return &MessageFlowBuilder{flow: []OmniMessage{}}
}

func (b *MessageFlowBuilder) Build() []OmniMessage {
	return b.flow
}

func (b *MessageFlowBuilder) Message(message interface{}) *MessageFlowBuilder {
	msg := OmniMessage{}
	switch reflect.TypeOf(message).Name() {
	case reflect.TypeOf(OmniSMS{}).Name():
		sms := message.(OmniSMS)
		msg.SMS = &sms
		break
	case reflect.TypeOf(OmniMMS{}).Name():
		mms := message.(OmniMMS)
		msg.MMS = &mms
		break
	case reflect.TypeOf(OmniRCS{}).Name():
		rcs := message.(OmniRCS)
		msg.RCS = &rcs
		break
	case reflect.TypeOf(OmniAlimTalk{}).Name():
		alimtalk := message.(OmniAlimTalk)
		msg.AlimTalk = &alimtalk
		break
	case reflect.TypeOf(OmniFriendTalk{}).Name():
		friendtalk := message.(OmniFriendTalk)
		msg.FriendTalk = &friendtalk
		break
	}
	b.flow = append(b.flow, msg)
	return b
}

func NewOmniSender(authorization string, httpClient *http.Client) *OmniSender {
	c := core.NewClient(httpClient)
	return &OmniSender{authorization: authorization, client: c}
}

func (sender *OmniSender) SendMessage(req *Omni) (*OmniSendResponse, error) {
	if response, err := sendMessage(sender.client, omniSubPath, sender.authorization, req); err != nil {
		return nil, err
	} else {
		return parseOmniSendResponse(sender.client, response)
	}
}
