package util

import "github.com/omni-bizgo/omni-sdk-go/infobank/client/send"

func MakeButtonURL(name string, url string) send.RcsButton {
	types := send.ButtonURL
	return send.RcsButton{Type: &types, Name: &name, Url: &url}
}

func MakeButtonMapLoc(name string, label string, latitude string, longitude string, fallbackUrl string) send.RcsButton {
	types := send.ButtonMapLoc
	return send.RcsButton{Type: &types, Name: &name, Label: &label, Latitude: &latitude, Longitude: &longitude, FallbackUrl: &fallbackUrl}
}

func MakeButtonMapQry(name string, query string, fallbackUrl string) send.RcsButton {
	types := send.ButtonMapQry
	return send.RcsButton{Type: &types, Name: &name, Query: &query, FallbackUrl: &fallbackUrl}
}

func MakeButtonMapSend(name string) send.RcsButton {
	types := send.ButtonMapSend
	return send.RcsButton{Type: &types, Name: &name}
}

func MakeButtonCalendar(name string, startTime string, endTime string, title string, description string) send.RcsButton {
	types := send.ButtonCalendar
	return send.RcsButton{Type: &types, Name: &name, StartTime: &startTime, EndTime: &endTime, Title: &title, Description: &description}
}

func MakeButtonCopy(name string, text string) send.RcsButton {
	types := send.ButtonCopy
	return send.RcsButton{Type: &types, Name: &name, Text: &text}
}

func MakeButtonCOMT(name string, phoneNumber string, text string) send.RcsButton {
	types := send.ButtonCOMT
	return send.RcsButton{Type: &types, Name: &name, PhoneNumber: &phoneNumber, Text: &text}
}

func MakeButtonCOMV(name string, phoneNumber string) send.RcsButton {
	types := send.ButtonCOMV
	return send.RcsButton{Type: &types, Name: &name, PhoneNumber: &phoneNumber}
}

func MakeButtonDial(name string, phoneNumber string) send.RcsButton {
	types := send.ButtonDial
	return send.RcsButton{Type: &types, Name: &name, PhoneNumber: &phoneNumber}
}
