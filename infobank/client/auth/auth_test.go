package auth

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {

	clientId := "unit_test_01" //"rcs_dev_socket1"
	password := "12345678"

	authClient := NewAuthenticator(nil, clientId, password)
	err := authClient.Init()
	if err != nil {
		t.Error("err :", err)
	} else {
		token := authClient.GetToken()
		fmt.Println(*token)
	}
}
