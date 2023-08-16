package regist

import (
	"bytes"
	"github.com/omni-bizgo/omni-sdk-go/infobank"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/send"
	"net/http"
	"reflect"
)

const (
	formSubPath = "/v1/form"
	contentType = "application/json"
)

type MessageForm struct {
	MessageForm []send.OmniMessage `json:"messageForm"`
}

type MessageFormBuilder struct {
	form MessageForm
}

func NewMessageFormBuilder() *MessageFormBuilder {
	return &MessageFormBuilder{form: MessageForm{}}
}

func (m *MessageFormBuilder) Build() MessageForm {
	return m.form
}

func (m *MessageFormBuilder) Message(message interface{}) *MessageFormBuilder {
	msg := send.OmniMessage{}
	switch reflect.TypeOf(message).Name() {
	case reflect.TypeOf(send.OmniSMS{}).Name():
		sms := message.(send.OmniSMS)
		msg.SMS = &sms
		break
	case reflect.TypeOf(send.OmniMMS{}).Name():
		mms := message.(send.OmniMMS)
		msg.MMS = &mms
		break
	case reflect.TypeOf(send.OmniRCS{}).Name():
		rcs := message.(send.OmniRCS)
		msg.RCS = &rcs
		break
	case reflect.TypeOf(send.OmniAlimTalk{}).Name():
		alimtalk := message.(send.OmniAlimTalk)
		msg.AlimTalk = &alimtalk
		break
	case reflect.TypeOf(send.OmniFriendTalk{}).Name():
		friendtalk := message.(send.OmniFriendTalk)
		msg.FriendTalk = &friendtalk
		break
	}
	m.form.MessageForm = append(m.form.MessageForm, msg)
	return m
}

type FormRegister struct {
	authorization string
	client        *core.HttpClient
}

func NewFormRegister(authorization string, httpClient *http.Client) *FormRegister {
	c := core.NewClient(httpClient)
	return &FormRegister{authorization: authorization, client: c}
}

func (f *FormRegister) RegisterForm(req MessageForm) (*FormResponse, error) {
	if response, err := manageForm(f, http.MethodPost, formSubPath, f.authorization, nil, req); err != nil {
		return nil, err
	} else {
		return parseFormResponse(f, response)
	}
}

func (f *FormRegister) InquiryForm(formId string) (*FormResponse, error) {
	if response, err := manageForm(f, http.MethodGet, formSubPath, f.authorization, &formId, nil); err != nil {
		return nil, err
	} else {
		return parseFormResponse(f, response)
	}
}

func (f *FormRegister) UpdateForm(formId string, req MessageForm) (*FormResponse, error) {
	if response, err := manageForm(f, http.MethodPut, formSubPath, f.authorization, &formId, req); err != nil {
		return nil, err
	} else {
		return parseFormResponse(f, response)
	}
}

func (f *FormRegister) RemoveForm(formId string) (*FormResponse, error) {
	if response, err := manageForm(f, http.MethodDelete, formSubPath, f.authorization, &formId, nil); err != nil {
		return nil, err
	} else {
		return parseFormResponse(f, response)
	}
}

func manageForm(f *FormRegister, method string, subPath string, authorization string, formId *string, req interface{}) (*http.Response, error) {
	url := infobank.RemoteAddress + subPath
	var body *bytes.Buffer
	header := make(map[string]string)
	header["Authorization"] = authorization

	if formId != nil {
		url = url + "/" + *formId
		body = nil
	}

	if method == http.MethodPost || method == http.MethodPut {
		header["Content-Type"] = contentType

		parsedBody, parseErr := core.MarshalAndConvertStr(req)
		if parseErr != nil {
			return nil, *parseErr
		}
		body = new(bytes.Buffer)
		if parsedBody != nil {
			body.WriteString(*parsedBody)
		}
	}

	httpRes, httpErr := f.client.Request(method, url, header, body)
	if httpErr != nil {
		return nil, *httpErr
	}

	return httpRes, nil
}

func parseFormResponse(f *FormRegister, httpRes *http.Response) (*FormResponse, error) {
	apiRes, parseErr := f.client.ParseHttpResponseBody(httpRes, new(FormResponse))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*FormResponse)
	return res, nil
}
