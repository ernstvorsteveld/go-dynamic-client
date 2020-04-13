package client

import (
	"testing"
)

func Test_get_client_type(t *testing.T) {
	clientType := ClientTypeValueOf("public")
	expectedClientType := Public
	if expectedClientType != clientType {
		t.Errorf("The clientType %s is not equals to the expected value %s.\n", clientType, expectedClientType)
	}

	clientType = ClientTypeValueOf("confidential")
	expectedClientType = Confidential
	if expectedClientType != clientType {
		t.Errorf("The clientType %s is not equals to the expected value %s.\n", clientType, expectedClientType)
	}
}

func Test_client_type_to_string(t *testing.T) {
	clientType := Public.String()
	expectedString := "public"

	if expectedString != clientType {
		t.Errorf("The string value of clientType %s is not equals to the expected value %s.\n", clientType, expectedString)
	}
}
