package oauth2client

import (
	"testing"
)

func Test_get_registration_error_type_as_string(t *testing.T) {
	registrationErrorType := Invalid_redirect_uri_registration.String()
	expectedValue := "invalid_redirect_uri"
	if expectedValue != registrationErrorType {
		t.Errorf("The string value of RegistrationErrorType %s is not equals to expected value %s.\n", registrationErrorType, expectedValue)
	}

	registrationErrorType = Uapproved_software_statement.String()
	expectedValue = "uapproved_software"
	if expectedValue != registrationErrorType {
		t.Errorf("The string value of RegistrationErrorType %s is not equals to expected value %s.\n", registrationErrorType, expectedValue)
	}

}
