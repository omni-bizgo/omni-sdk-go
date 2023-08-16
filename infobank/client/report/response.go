package report

type Response struct {
	Code   string        `json:"code"`
	Result string        `json:"result"`
	Data   *ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	ReportId *string  `json:"reportId,omitempty"`
	Report   []Report `json:"report"`
}

type Report struct {
	MsgKey      string `json:"msgKey"`
	ServiceType string `json:"serviceType"`
	MsgType     string `json:"msgType"`
	ReportType  string `json:"reportType"`
	ReportCode  string `json:"reportCode"`
	ReportText  string `json:"reportText,omitempty"`
	ReportTime  string `json:"reportTime"`
	Carrier     string `json:"carrier,omitempty"`
	Ref         string `json:"ref,omitempty"`
	ResCnt      string `json:"resCnt,omitempty"`
}
