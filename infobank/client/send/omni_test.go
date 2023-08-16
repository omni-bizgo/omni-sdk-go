package send

import (
	"errors"
	"fmt"
	"testing"
)

func TestOmniSender_SendMessage(t *testing.T) {
	alimtalk := NewOmniAlimTalkBuilder().Text("").Build()
	flow := NewMessageFlowBuilder().Message(alimtalk).Build()

	destination := NewDestinationBuilder().To("01012345678").Build()
	destinations := append([]Destination{}, destination)
	flowRequest := NewOmniBuilder().MessageFlow(flow).Destinations(destinations).Ref("omni-send-flow-test").Build()
	formRequest := NewOmniBuilder().MessageForm("formId").Destinations(destinations).Ref("omni-send-form-test").Build()

	sender := NewOmniSender("input-your-token", nil)
	flowSendApiResponse, flowSendErr := sender.SendMessage(&flowRequest)
	formSendApiResponse, formSendErr := sender.SendMessage(&formRequest)

	if flowSendErr != nil {
		t.Error("flowSendErr : ", flowSendErr)
		return
	}

	if formSendErr != nil {
		t.Error("formSendErr : ", formSendErr)
		return
	}

	if flowSendApiResponse.Code != "A000" {
		t.Error("flowSendErr : ", errors.New(flowSendApiResponse.Result))
		return
	}
	fmt.Println(*flowSendApiResponse)

	if formSendApiResponse.Code != "A000" {
		t.Error("formSendErr : ", errors.New(formSendApiResponse.Result))
		return
	}
	fmt.Println(*formSendApiResponse)
}
