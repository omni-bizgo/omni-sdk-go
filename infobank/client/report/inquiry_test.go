package report

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReportManager_InquiryReport(t *testing.T) {
	reportManager := NewReportManager("input-your-token", nil)
	apiResponse, err := reportManager.InquiryReport("20230619081619396PRDR1SM92760200")
	if err != nil {
		t.Error(err)
	} else {
		data, _ := json.Marshal(apiResponse.Data)
		fmt.Println(apiResponse.Code, apiResponse.Data, string(data))
	}
}
