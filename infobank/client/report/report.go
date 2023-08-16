package report

import (
	"github.com/omni-bizgo/omni-sdk-go/infobank"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"net/http"
)

type ReportManager struct {
	authorization string
	client        *core.HttpClient
}

func NewReportManager(authorization string, httpClient *http.Client) *ReportManager {
	c := core.NewClient(httpClient)
	return &ReportManager{authorization: authorization, client: c}
}

func manageReport(r *ReportManager, method string, subPath string, authorization string, reportId *string) (*http.Response, error) {
	url := infobank.RemoteAddress + subPath
	header := make(map[string]string)
	header["Authorization"] = authorization

	if reportId != nil {
		url += "/" + *reportId
	}

	httpRes, httpErr := r.client.Request(method, url, header, nil)
	if httpErr != nil {
		return nil, *httpErr
	}
	return httpRes, nil

}

func parseReportResponse(r *ReportManager, httpRes *http.Response) (*Response, error) {
	apiRes, parseErr := r.client.ParseHttpResponseBody(httpRes, new(Response))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*Response)
	return res, nil
}
