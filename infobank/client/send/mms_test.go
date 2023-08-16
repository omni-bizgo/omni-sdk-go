package send

import (
	"errors"
	"fmt"
	"testing"
)

func TestMmsSender_SendMessage(t *testing.T) {
	mmsSender := NewMmsSender("input-your-token", nil)
	message := NewMmsBuilder().From("0310000000").To("01012345618").Text("omni-sdk-go text").Ref("omni-sdk-go ref").OriginCID("0000").Build()

	apiResponse, err := mmsSender.SendMessage(&message)

	if err != nil {
		t.Error("err : ", err)
		return
	}

	if apiResponse.Code != "A000" {
		t.Error("err : ", errors.New(apiResponse.Result))
		return
	}
	fmt.Println(*apiResponse)
}
