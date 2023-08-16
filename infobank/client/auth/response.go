package auth

type Response struct {
	Code   string        `json:"code"`
	Result string        `json:"result"`
	Data   *ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	Token   string `json:"token"`
	Schema  string `json:"schema"`
	Expired string `json:"expired"`
}
