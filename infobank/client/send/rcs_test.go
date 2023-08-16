package send

import (
	"errors"
	"fmt"
	"testing"
)

func TestRcsSender_SendMessage(t *testing.T) {
	rcsSender := NewRcsSender("input-your-token", nil)
	button := NewRcsButtonBuilder().FallbackUrl("https://www.z").Label("라벨명").Latitude("37.4220041").Longitude("-122.0862515").Name("지도 보여주기").Type("MAP_LOC").Build()
	standalone := NewStandAloneBuilder().Text("rcs-test").Title("title").Button(append([]RcsButton{}, button)).Build()
	content := NewRcsContentBuilder().StandAlone(standalone).Build()
	fallback := NewRcsFallbackBuilder().Type("SMS").Text("rcs-fallback-text").Title("rcs-fallback-title").Text("rcs-fallback-text").OriginCID("originCID").Build()
	message := NewRcsBuilder().FormatId("SS000000").From("0310000000").To("01012345618").Content(content).Fallback(fallback).Ref("rcs-test-ref").Build()

	sub := NewRcsSubContentBuilder().SubTitle("").SubDesc("").Build()
	subList := append([]RcsSubContent{}, sub)
	customMap := make(map[string]interface{})
	customMap["key"] = "value"
	rcsTemplate := NewRcsTemplateBuilder().Description("des").SubContent(subList).CustomMap(customMap).Build()
	content = NewRcsContentBuilder().Template(rcsTemplate).Build()
	message = NewRcsBuilder().FormatId("SS000000").From("0310000000").To("01012345618").Content(content).Fallback(fallback).Ref("rcs-test-ref").Build()
	res, err := rcsSender.SendMessage(&message)
	fmt.Println(res)

	apiResponse, err := rcsSender.SendMessage(&message)
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
