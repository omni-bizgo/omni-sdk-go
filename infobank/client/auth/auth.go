package auth

import (
	"errors"
	"github.com/omni-bizgo/omni-sdk-go/infobank"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
	"time"
)

const (
	subPath    = "/v1/auth/token"
	timeFormat = "2006-01-02T15:04:05+07:00"
)

type Authenticator struct {
	clientId      string
	password      string
	authorization *ResponseData
	client        *core.HttpClient
}

func NewAuthenticator(httpClient *http.Client, clientId string, password string) *Authenticator {
	c := core.NewClient(httpClient)
	return &Authenticator{
		clientId: clientId,
		password: password,
		client:   c,
	}
}

func (auth *Authenticator) Init() error {
	response, err := getToken(auth, subPath)
	if err != nil {
		return err
	} else if response.Code != "A000" {
		errs := errors.New(response.Result)
		return errs
	}

	auth.authorization = response.Data
	return err
}

func (auth *Authenticator) Refresh() error {
	response, err := getToken(auth, subPath)
	if err != nil {
		return err
	} else if response.Code != "A000" {
		errs := errors.New(response.Result)
		return errs
	}

	auth.authorization = response.Data
	return err
}

func (auth *Authenticator) GetToken() *string {
	if auth.authorization == nil {
		return nil
	} else {
		token := auth.authorization.Schema + " " + auth.authorization.Token
		return &token
	}
}

func (auth *Authenticator) VerifyInValidToken() bool {

	currentTime := time.Now().Add(time.Minute * 10)
	if auth.authorization == nil {
		return true
	}

	expiredAt, err := time.Parse(timeFormat, auth.authorization.Expired)
	if err != nil {
		return true
	}

	if currentTime.After(expiredAt) {
		return true
	}

	return false
}

func getToken(auth *Authenticator, subPath string) (*Response, error) {
	url := infobank.RemoteAddress + subPath
	contentType := "application/json"
	header := make(map[string]string)
	header["X-IB-Client-Id"] = auth.clientId
	header["X-IB-Client-Passwd"] = auth.password
	header["Content-Type"] = contentType

	httpRes, httpErr := auth.client.Request(http.MethodPost, url, header, nil)
	if httpErr != nil {
		return nil, *httpErr
	}

	apiRes, parseErr := auth.client.ParseHttpResponseBody(httpRes, new(Response))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*Response)
	return res, nil
}
