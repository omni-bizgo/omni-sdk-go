package report

import (
	"net/http"
)

const (
	inquirySubPath = "/v1/report/inquiry"
)

func (r *ReportManager) InquiryReport(reportId string) (*Response, error) {
	if response, err := manageReport(r, http.MethodGet, inquirySubPath, r.authorization, &reportId); err != nil {
		return nil, err
	} else {
		return parseReportResponse(r, response)
	}
}
