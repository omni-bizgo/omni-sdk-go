package send

import (
	"errors"
	"fmt"
	"testing"
)

func TestFriendTalkSender_SendMessage(t *testing.T) {
	friendtalkSender := NewFriendTalkSender("input-your-token", nil)
	button1 := NewKakaoButtonBuilder().Type("WL").Name("미리 주문하기").UrlMobile("http://www.kakao.com").Build()
	button2 := NewKakaoButtonBuilder().Type("MD").Name("상담원 연결하기").Build()
	button3 := NewKakaoButtonBuilder().Type("AL").Name("방송 알림 설정 보기").SchemeIOS("daumapps://open").SchemeAndroid("daumapps://open").Build()
	buttonList := append([]KakaoButton{}, button1)
	buttonList = append(buttonList, button2)
	buttonList = append(buttonList, button3)
	fallback := NewKakaoFallbackBuilder().Type("SMS").Title("fallback-title").Text("text").From("0310000000").OriginCID("0000").Build()
	message := NewFriendtalkBuilder().KakaoMsgType("AT").SenderKey("senderKey").To("01012345678").Text("친구톡테스트").Ref("ref").Fallback(fallback).Build()
	apiResponse, err := friendtalkSender.SendMessage(&message)
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
