package regist

import (
	"fmt"
	"testing"
)

func TestFileUploader_UploadFile(t *testing.T) {
	fileUploader := NewFileUploader("input-your-token", nil)
	serviceType := FILE_SERVICE_TYPE_MMS
	response, err := fileUploader.UploadFile(&serviceType, nil, "D:\\", "2.jpg")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(response)
	}
}
