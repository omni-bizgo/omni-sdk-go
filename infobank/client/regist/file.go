package regist

import (
	"bytes"
	"fmt"
	"github.com/omni-bizgo/omni-sdk-go/infobank"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/core"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strings"
)

const (
	fileSubPath = "/v1/file"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

type FileServiceTypeEnum string

// FileServiceTypeEnum 상수들
const (
	FILE_SERVICE_TYPE_MMS        FileServiceTypeEnum = "MMS"
	FILE_SERVICE_TYPE_RCS        FileServiceTypeEnum = "RCS"
	FILE_SERVICE_TYPE_FRIENDTALK FileServiceTypeEnum = "friendtalk"
)

type FileMsgTypeEnum string

// FileMsgTypeEnum 상수들
const (
	FILE_MSG_TYPE_FRIENDTALK_IMAGE      FileMsgTypeEnum = "FI"
	FILE_MSG_TYPE_FRIENDTALK_WIDE_IMAGE FileMsgTypeEnum = "FW"
)

type FileUploader struct {
	authorization string
	client        *core.HttpClient
}

func NewFileUploader(authorization string, httpClient *http.Client) *FileUploader {
	c := core.NewClient(httpClient)
	return &FileUploader{authorization: authorization, client: c}
}

func (f *FileUploader) UploadFile(serviceType *FileServiceTypeEnum, msgType *FileMsgTypeEnum, fileDir string, fileName string) (*FileResponse, error) {
	subPath := fileSubPath
	if serviceType != nil {
		subPath += "/" + string(*serviceType)
		if msgType != nil {
			subPath += "/" + string(*msgType)
		}
	}
	filePath := path.Join(fileDir, fileName)

	//file open
	file, fileOpenErr := os.Open(filePath)
	if fileOpenErr != nil {
		return nil, fileOpenErr
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if part, err := writer.CreatePart(getMIMEHeader(file.Name(), filePath)); err != nil {
		return nil, err
	} else {
		if _, fileCopyErr := io.Copy(part, file); fileCopyErr != nil {
			return nil, fileCopyErr
		} else {
			contentType := writer.FormDataContentType()
			_ = writer.Close()
			if response, err := uploadFile(f, subPath, contentType, body); err != nil {
				return nil, err
			} else {
				return parseFileResponse(f, response)
			}
		}
	}
}

func getMIMEHeader(fileName string, filePath string) textproto.MIMEHeader {
	mime := mime.TypeByExtension(path.Ext(filePath))
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes("file"), escapeQuotes(fileName)))
	h.Set("Content-Type", mime)
	return h
}

func uploadFile(f *FileUploader, subPath string, contentType string, body *bytes.Buffer) (*http.Response, error) {
	url := infobank.RemoteAddress + subPath
	header := make(map[string]string)
	header["Content-Type"] = contentType
	header["Authorization"] = f.authorization

	httpRes, httpErr := f.client.Request(http.MethodPost, url, header, body)
	if httpErr != nil {
		return nil, *httpErr
	}

	return httpRes, nil
}

func parseFileResponse(f *FileUploader, httpRes *http.Response) (*FileResponse, error) {
	apiRes, parseErr := f.client.ParseHttpResponseBody(httpRes, new(FileResponse))
	if parseErr != nil {
		return nil, parseErr
	}

	res := apiRes.(*FileResponse)
	return res, nil
}
