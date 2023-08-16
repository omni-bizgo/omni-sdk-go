package regist

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/send"
	"testing"
)

func TestFormRegister_RegisterForm(t *testing.T) {

	formRegister := NewFormRegister("input-your-token", nil)
	msg := send.NewOmniSmsBuilder().Text("text").From("01012345678").From("01012345678").Build()
	form := NewMessageFormBuilder().Message(msg).Build()
	apiResponse, err := formRegister.RegisterForm(form)

	if err != nil {
		t.Error(err)
	} else {
		if apiResponse.Code != "A000" {
			t.Error(errors.New(apiResponse.Result))
		} else {
			fmt.Println(apiResponse.Result, apiResponse.Code, *apiResponse.Data.FormId, *apiResponse.Data.Expired)
		}
	}
}

func TestFormRegister_InquiryForm(t *testing.T) {
	formId := "20230710111716094DEVF10000000146"
	formRegister := NewFormRegister("input-your-token", nil)
	apiResponse, err := formRegister.InquiryForm(formId)
	if err != nil {
		t.Error(err)
	} else {

		data, _ := json.Marshal(apiResponse.Data.MessageForm)
		fmt.Println(apiResponse, *apiResponse.Data.FormId, *apiResponse.Data.Expired, string(data))
	}
}

func TestFormRegister_UpdateForm(t *testing.T) {
	formId := "20230710111716094DEVF10000000146"
	formRegister := NewFormRegister("input-your-token", nil)
	msg := send.NewOmniSmsBuilder().Text("text").From("01012345678").From("01012345678").Build()
	msg2 := send.NewMmsBuilder().Text("text").From("01012345678").From("01012345678").Build()
	form := NewMessageFormBuilder().Message(msg).Message(msg2).Build()
	apiResponse, err := formRegister.UpdateForm(formId, form)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(apiResponse, *apiResponse.Data.FormId, *apiResponse.Data.Expired)
	}
}

func TestFormRegister_RemoveForm(t *testing.T) {
	formId := "20230710111716094DEVF10000000146"
	formRegister := NewFormRegister("input-your-token", nil)
	apiResponse, err := formRegister.RemoveForm(formId)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(apiResponse)
	}
}
