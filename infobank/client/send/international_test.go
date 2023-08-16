package send

import (
	"errors"
	"fmt"
	"testing"
)

func TestInternationalSmsSender_SendMessage(t *testing.T) {
	sender := NewInternationalSmsSender("input-your-token", nil)
	message := NewInternationalSmsBuilder().From("0310000000").To("+821012345618").Text("omni-sdk-go text").Ref("omni-sdk-go ref").Build()

	apiResponse, err := sender.SendMessage(&message)

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
