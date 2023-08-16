package send

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

type SimpleSendResponse struct {
	Code   string  `json:"code"`
	Result string  `json:"result"`
	Ref    *string `json:"ref,omitempty"`
	MsgKey *string `json:"msgKey,omitempty"`
}

type OmniSendResponse struct {
	Code   string                `json:"code"`
	Result string                `json:"result"`
	Ref    *string               `json:"ref,omitempty"`
	Data   *OmniSendResponseData `json:"data,omitempty"`
}

type OmniSendResponseData struct {
	Destinations *[]DestinationsResponse `json:"destinations,omitempty"`
}

type DestinationsResponse struct {
	To     string `json:"to"`
	MsgKey string `json:"msgKey"`
	Code   string `json:"code"`
	Result string `json:"result"`
}

func parseSimpleSendResponse(client *core.HttpClient, httpRes *http.Response) (*SimpleSendResponse, error) {
	apiRes, parseErr := client.ParseHttpResponseBody(httpRes, new(SimpleSendResponse))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*SimpleSendResponse)
	return res, nil
}

func parseOmniSendResponse(client *core.HttpClient, httpRes *http.Response) (*OmniSendResponse, error) {
	apiRes, parseErr := client.ParseHttpResponseBody(httpRes, new(OmniSendResponse))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*OmniSendResponse)
	return res, nil
}
