package client

import (
	"errors"
	"fmt"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/auth"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/regist"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/report"
	"github.com/omni-bizgo/omni-sdk-go/infobank/client/send"
	"net/http"
	"reflect"
)

type OmniClient struct {
	auth             auth.Authenticator
	sms              *send.SmsSender
	mms              *send.MmsSender
	rcs              *send.RcsSender
	internationalSMS *send.InternationalSmsSender
	alimtalk         *send.AlimtalkSender
	friendtalk       *send.FriendtalkSender
	omni             *send.OmniSender
	report           *report.ReportManager
	file             *regist.FileUploader
	form             *regist.FormRegister
	httpClient       *http.Client
}

func NewOmniClient(httpClient *http.Client, clientId string, password string) *OmniClient {
	c := &OmniClient{}

	if httpClient != nil {
		c.httpClient = httpClient
	}

	a := auth.NewAuthenticator(httpClient, clientId, password)
	c.auth = *a
	_ = c.auth.Init()
	return c
}

func (c *OmniClient) SendMessage(req interface{}) (interface{}, error) {
	var res interface{}
	var err error
	var isTokenExpired bool

	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err = c.auth.Refresh()
		if err != nil {
			return auth.Response{Code: "A100", Result: "Invalid authentication information"}, err
		}
	}
	token := c.auth.GetToken()

	if req == nil {
		return send.SimpleSendResponse{Code: "A300", Result: "Wrong request"}, err
	}

	switch reflect.TypeOf(req).String() {
	case reflect.TypeOf(&send.SMS{}).String():
		if c.sms == nil || isTokenExpired {
			c.sms = send.NewSmsSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.sms.SendMessage(req.(*send.SMS))
		break
	case reflect.TypeOf(&send.MMS{}).String():
		if c.mms == nil || isTokenExpired {
			c.mms = send.NewMmsSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.mms.SendMessage(req.(*send.MMS))
		break
	case reflect.TypeOf(&send.InternationalSMS{}).String():
		if c.internationalSMS == nil || isTokenExpired {
			c.internationalSMS = send.NewInternationalSmsSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.internationalSMS.SendMessage(req.(*send.InternationalSMS))
		break
	case reflect.TypeOf(&send.RCS{}).String():
		if c.rcs == nil || isTokenExpired {
			c.rcs = send.NewRcsSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.rcs.SendMessage(req.(*send.RCS))
		break
	case reflect.TypeOf(&send.Alimtalk{}).String():
		if c.alimtalk == nil || isTokenExpired {
			c.alimtalk = send.NewAlimTalkSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.alimtalk.SendMessage(req.(*send.Alimtalk))
		break
	case reflect.TypeOf(&send.Friendtalk{}).String():
		if c.friendtalk == nil || isTokenExpired {
			c.friendtalk = send.NewFriendTalkSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.friendtalk.SendMessage(req.(*send.Friendtalk))
		break
	case reflect.TypeOf(&send.Omni{}).String():
		if c.omni == nil || isTokenExpired {
			c.omni = send.NewOmniSender(*token, cloneHttpClient(c.httpClient))
		}
		res, err = c.omni.SendMessage(req.(*send.Omni))
		break
	default:
		res = nil
		err = errors.New(fmt.Sprintf("invalid request, type : %s", reflect.TypeOf(req).Name()))
		break
	}
	return res, err
}

func (c *OmniClient) UploadFile(serviceType *regist.FileServiceTypeEnum, msgType *regist.FileMsgTypeEnum, fileDir string, fileName string) (*regist.FileResponse, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.file == nil || isTokenExpired {
		c.file = regist.NewFileUploader(*token, cloneHttpClient(c.httpClient))
	}
	return c.file.UploadFile(serviceType, msgType, fileDir, fileName)
}

func (c *OmniClient) RegisterForm(req regist.MessageForm) (*regist.FormResponse, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.form == nil || isTokenExpired {
		c.form = regist.NewFormRegister(*token, cloneHttpClient(c.httpClient))
	}
	return c.form.RegisterForm(req)
}

func (c *OmniClient) InquiryForm(formId string) (*regist.FormResponse, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.form == nil || isTokenExpired {
		c.form = regist.NewFormRegister(*token, cloneHttpClient(c.httpClient))
	}
	return c.form.InquiryForm(formId)
}

func (c *OmniClient) UpdateForm(formId string, req regist.MessageForm) (*regist.FormResponse, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.form == nil || isTokenExpired {
		c.form = regist.NewFormRegister(*token, cloneHttpClient(c.httpClient))
	}
	return c.form.UpdateForm(formId, req)
}

func (c *OmniClient) RemoveForm(formId string) (*regist.FormResponse, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.form == nil || isTokenExpired {
		c.form = regist.NewFormRegister(*token, cloneHttpClient(c.httpClient))
	}
	return c.form.RemoveForm(formId)
}

func (c *OmniClient) InquiryReport(reportId string) (*report.Response, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.report == nil || isTokenExpired {
		c.report = report.NewReportManager(*token, cloneHttpClient(c.httpClient))
	}
	return c.report.InquiryReport(reportId)
}

func (c *OmniClient) PollingReport() (*report.Response, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.report == nil || isTokenExpired {
		c.report = report.NewReportManager(*token, cloneHttpClient(c.httpClient))
	}
	return c.report.PollingReport()
}

func (c *OmniClient) DeleteReport(reportId string) (*report.Response, error) {
	var isTokenExpired bool
	if isTokenExpired = c.auth.VerifyInValidToken(); isTokenExpired {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}
	token := c.auth.GetToken()
	if c.report == nil || isTokenExpired {
		c.report = report.NewReportManager(*token, cloneHttpClient(c.httpClient))
	}
	return c.report.DeleteReport(reportId)
}

func cloneHttpClient(httpClient *http.Client) *http.Client {
	if httpClient != nil {
		return &http.Client{
			Transport:     httpClient.Transport,
			CheckRedirect: httpClient.CheckRedirect,
			Jar:           httpClient.Jar,
			Timeout:       httpClient.Timeout,
		}
	} else {
		return http.DefaultClient
	}
}
