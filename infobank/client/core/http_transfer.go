package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	apiHttpStatusCode = map[int]bool{
		http.StatusOK:                   true,
		http.StatusBadRequest:           true,
		http.StatusUnauthorized:         true,
		http.StatusForbidden:            true,
		http.StatusNotFound:             true,
		http.StatusUnsupportedMediaType: true,
		http.StatusTooManyRequests:      true,
		http.StatusInternalServerError:  true,
		http.StatusServiceUnavailable:   true,
	}
)

type HttpClient struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *HttpClient {
	client := new(HttpClient)

	if httpClient == nil {
		client.httpClient = http.DefaultClient
	} else {
		client.httpClient = httpClient
	}

	return client
}

func (c *HttpClient) Request(method string, url string, header map[string]string, data *bytes.Buffer) (*http.Response, *error) {

	var request *http.Request
	var err error

	if data == nil {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, data)
	}

	if err != nil {
		return nil, &err
	} else {
		if header != nil && len(header) != 0 {
			for key, value := range header {
				request.Header.Set(key, value)
			}
		}

		if response, err := c.httpClient.Do(request); err != nil {
			return nil, &err
		} else {
			if _, exists := apiHttpStatusCode[response.StatusCode]; !exists {
				err = errors.New(fmt.Sprintf("httpStatusCode := %d, msg := %s", response.StatusCode, response.Status))
				errors.Join()
				return nil, &err
			}

			return response, nil
		}
	}
}

func (c *HttpClient) ParseHttpResponseBody(res *http.Response, vo interface{}) (interface{}, error) {
	defer res.Body.Close()
	body, responseReadErr := io.ReadAll(res.Body)
	if responseReadErr != nil {
		return nil, responseReadErr
	}

	if marshalErr := Unmarshal(body, &vo); marshalErr != nil {
		return nil, marshalErr
	}

	return vo, nil
}
