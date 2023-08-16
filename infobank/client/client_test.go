package client

import (
	"fmt"
	"github.com/omni-bizgo/omni-sdk-go/infobank"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/regist"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/send"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/util"
	"net/http"
	"testing"
)

var clientId string
var clientPassword string
var httpClient http.Client

func init() {
	//테스트 환경(stg) 설정, 별도로 입력하지 않으면 prod 환경으로 발송
	infobank.RemoteAddress = "https://stg-omni.ibapi.kr"

	//omni-client에 필요한 정보 초기화
	clientId = "gw_omni_real01"
	clientPassword = "12345678"
	//clientId 및 password는 omni 시스템에 등록된 계정정보
	httpClient = http.Client{}
}

func TestSendSms(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	sms := send.NewSmsBuilder().Text("test-text").To("yournumber").From("021234567").OriginCID("originCid").Ref("ref").Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&sms); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendMms(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	var fileKey []string
	fileKey = append(fileKey, "filekey.jpg")
	sms := send.NewMmsBuilder().Title("title").FileKey(fileKey).Text("test-text").To("yournumber").From("021234567").OriginCID("originCid").Ref("ref").Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&sms); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendInternationSMS(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	internationalSms := send.NewInternationalSmsBuilder().Text("international-test").From("021234567").To("yournumber").Ref("ref").Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&internationalSms); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendRcsStandAlone(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	var buttons []send.RcsButton
	//tip. rcs 버튼은 util 패키지를 통해 쉽게 만들 수 있다.
	buttons = append(buttons, util.MakeButtonCOMV("button-test", "01012345678"))
	standAlone := send.NewStandAloneBuilder().Title("rcs-standalone-title").Text("rcs-standalone-title").Button(buttons).Build()
	content := send.NewRcsContentBuilder().StandAlone(standAlone).Build()
	fallback := send.NewRcsFallbackBuilder().Title("rcs-fallback-title").Text("rcs-fallback-text").Type(send.FALLBACK_SMS).OriginCID("originCid").Build()
	rcs := send.NewRcsBuilder().Content(content).From("0212345678").To("010337711801").FormatId("SS000001").Fallback(fallback).Ref("ref").Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&rcs); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendRcsCarousel(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	var carousels []send.Carousel
	carousels = append(carousels, send.NewCarouselBuilder().Title("rcs-carousel-title1").Text("rcs-carousel-title1").Build())
	carousels = append(carousels, send.NewCarouselBuilder().Title("rcs-carousel-title2").Text("rcs-carousel-title2").Build())
	carousels = append(carousels, send.NewCarouselBuilder().Title("rcs-carousel-title3").Text("rcs-carousel-title3").Build())
	content := send.NewRcsContentBuilder().Carousel(carousels).Build()
	fallback := send.NewRcsFallbackBuilder().Title("rcs-fallback-title").Text("rcs-fallback-text").Type(send.FALLBACK_SMS).OriginCID("originCid").Build()
	rcs := send.NewRcsBuilder().Content(content).From("0212345678").To("010337711801").FormatId("SS000001").Fallback(fallback).Ref("ref").Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&rcs); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendRcsTemplate(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	template := send.NewRcsTemplateBuilder().Description("rcs-template").CustomMap(nil).Build()
	content := send.NewRcsContentBuilder().Template(template).Build()
	fallback := send.NewRcsFallbackBuilder().Title("rcs-fallback-title").Text("rcs-fallback-text").Type(send.FALLBACK_SMS).OriginCID("originCid").Build()
	rcs := send.NewRcsBuilder().Content(content).From("0212345678").To("010337711801").FormatId("SS000001").Fallback(fallback).Ref("ref").Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&rcs); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendAlimtalk(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	var buttons []send.KakaoButton
	buttons = append(buttons, send.NewKakaoButtonBuilder().Name("web-test").Type(send.BUTTON_WEB_LINK).Build())
	fallback := send.NewKakaoFallbackBuilder().Title("rcs-fallback-title").Text("rcs-fallback-text").Type(send.FALLBACK_SMS).OriginCID("originCid").Build()
	alimtalk := send.NewAlimtalkBuilder().To("01012345678").Text("alimtalk-test").SenderKey("sendekey").TemplateCode("templateCode").KakaoMsgType(send.MSGTYPE_ALIMTALK).Button(buttons).Fallback(fallback).Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&alimtalk); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendFriendtalk(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 작성
	var buttons []send.KakaoButton
	buttons = append(buttons, send.NewKakaoButtonBuilder().Name("web-test").Type(send.BUTTON_WEB_LINK).Build())
	fallback := send.NewKakaoFallbackBuilder().Title("rcs-fallback-title").Text("rcs-fallback-text").Type(send.FALLBACK_SMS).OriginCID("originCid").Build()
	friendtalk := send.NewFriendtalkBuilder().To("01012345678").Text("alimtalk-test").SenderKey("sendekey").KakaoMsgType(send.MSGTYPE_FRIENDTALK).Button(buttons).Fallback(fallback).Build()

	//step3. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&friendtalk); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendOmniFlow(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	///step2. 메시지 flow 작성
	alimtalk := send.NewOmniAlimTalkBuilder().Text("").Build()
	var buttons []send.RcsButton
	buttons = append(buttons, util.MakeButtonCOMV("button-test", "01012345678"))
	standAlone := send.NewStandAloneBuilder().Title("rcs-standalone-title").Text("rcs-standalone-title").Button(buttons).Build()
	content := send.NewRcsContentBuilder().StandAlone(standAlone).Build()
	rcs := send.NewOmniRcsBuilder().Content(content).Build()
	sms := send.NewOmniSmsBuilder().Build()
	flow := send.NewMessageFlowBuilder().Message(alimtalk).Message(rcs).Message(sms).Build()

	//step3. 대상자 작성
	destination := send.NewDestinationBuilder().To("01012345678").Build()
	destinations := append([]send.Destination{}, destination)

	//step4. omni 메시지 작성
	omni := send.NewOmniBuilder().MessageFlow(flow).Destinations(destinations).Ref("omni-send-flow-test").Build()

	//step5. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&omni); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestSendOmniForm(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	//step2. 대상자 작성
	destination := send.NewDestinationBuilder().To("01012345678").Build()
	destinations := append([]send.Destination{}, destination)

	//step3. omni 메시지 작성(이미 등록된 form id 작성 필수)
	omni := send.NewOmniBuilder().MessageForm("formId").Destinations(destinations).Ref("omni-send-flow-test").Build()

	//step4. 메시지 발송 및 발송결과 확인
	if res, err := omniClient.SendMessage(&omni); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestReportInquiry(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	//step2. 리포트 조회
	if res, err := omniClient.InquiryReport("msgKey"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestReportPollingAndDelete(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	//step2. 리포트 폴링
	if res, err := omniClient.PollingReport(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
		//...
		reportId := res.Data.ReportId
		//step3. 만약 리포트가 있다면, 비즈니스 로직을 처리한 뒤, 아래와 같이 리포트 확인처리를 권장.
		if res, err := omniClient.DeleteReport(*reportId); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res)
		}
	}
}

func TestFormRegist(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	//step2. form 작성
	alimtalk := send.NewOmniAlimTalkBuilder().Text("").Build()
	var buttons []send.RcsButton
	buttons = append(buttons, util.MakeButtonCOMV("button-test", "01012345678"))
	standAlone := send.NewStandAloneBuilder().Title("rcs-standalone-title").Text("rcs-standalone-title").Button(buttons).Build()
	content := send.NewRcsContentBuilder().StandAlone(standAlone).Build()
	rcs := send.NewOmniRcsBuilder().Content(content).Build()
	sms := send.NewOmniSmsBuilder().Build()

	form := regist.NewMessageFormBuilder().Message(alimtalk).Message(rcs).Message(sms).Build()

	//step3. form 등록
	if res, err := omniClient.RegisterForm(form); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func TestFormUpdateAndInquiry(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	//step2. form 작성
	alimtalk := send.NewOmniAlimTalkBuilder().Text("").Build()
	var buttons []send.RcsButton
	buttons = append(buttons, util.MakeButtonCOMV("button-test", "01012345678"))
	standAlone := send.NewStandAloneBuilder().Title("rcs-standalone-title").Text("rcs-standalone-title").Button(buttons).Build()
	content := send.NewRcsContentBuilder().StandAlone(standAlone).Build()
	rcs := send.NewOmniRcsBuilder().Content(content).Build()
	sms := send.NewOmniSmsBuilder().Build()

	form := regist.NewMessageFormBuilder().Message(alimtalk).Message(rcs).Message(sms).Build()

	//step3. form 수정
	if res, err := omniClient.UpdateForm("formId", form); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)

		formId := res.Data.FormId
		//step3. 등록된 폼을 조회하기 위함
		if res, err := omniClient.InquiryForm(*formId); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res)
		}
	}
}

func TestFormDeleteAndInquiry(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)
	//httpClient == nil 이면, 기본 httpClient 생성

	formId := "formId"

	//step2. omni-form 삭제
	if res, err := omniClient.RemoveForm(formId); err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)

		//step3. 정말 폼이 잘 삭제되었는지 조회
		if res, err := omniClient.InquiryForm(formId); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res)
		}
	}
}

func TestFileUpload(t *testing.T) {
	//step1. omni-client 생성
	omniClient := NewOmniClient(&httpClient, clientId, clientPassword)

	serviceType := regist.FILE_SERVICE_TYPE_MMS
	response, err := omniClient.UploadFile(&serviceType, nil, "D:\\", "2.jpg")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(response)
	}
}
