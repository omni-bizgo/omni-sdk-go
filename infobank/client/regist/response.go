package regist

import "github.com/omni-bizgo/omni-sdk-go/infobank/client/send"

type FileResponse struct {
	Code   string            `json:"code"`
	Result string            `json:"result"`
	Data   *FileResponseData `json:"data,omitempty"`
}

type FileResponseData struct {
	ImgUrl  *string `json:"imgUrl"`
	FileKey *string `json:"fileKey"`
	Media   *string `json:"media"`
	Expired *string `json:"expired"`
}

type FormResponse struct {
	Code   string            `json:"code"`
	Result string            `json:"result"`
	Data   *FormResponseData `json:"data,omitempty"`
}

type FormResponseData struct {
	FormId      *string            `json:"formId,omitempty"`
	Expired     *string            `json:"expired,omitempty"`
	MessageForm []send.OmniMessage `json:"messageForm"`
}
