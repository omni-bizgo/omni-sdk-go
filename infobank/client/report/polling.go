package report

import (
	"net/http"
)

const (
	pollingSubPath = "/v1/report/polling"
)

func (r *ReportManager) PollingReport() (*Response, error) {
	if response, err := manageReport(r, http.MethodGet, pollingSubPath, r.authorization, nil); err != nil {
		return nil, err
	} else {
		return parseReportResponse(r, response)
	}
}

func (r *ReportManager) DeleteReport(reportId string) (*Response, error) {
	if response, err := manageReport(r, http.MethodDelete, pollingSubPath, r.authorization, &reportId); err != nil {
		return nil, err
	} else {
		return parseReportResponse(r, response)
	}
}
