package services

import (
	"golang.org/x/xerrors"
	"regexp"
	"strings"
)

func validatePhoneNumber(phoneNumber string) (string, error) {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")

	matched, _ := regexp.MatchString(`^(0[789]0\d{8})$`, phoneNumber)
	if !matched {
		return "", xerrors.Errorf(
			"Invalid phone number",
		)
	}
	return phoneNumber, nil
}

func ConvertToInternationalPhoneNumber(phoneNumber string) (string, error) {
	validatedPhoneNumber, err := validatePhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	internationalPhoneNumber := "+81" + validatedPhoneNumber[1:]
	return internationalPhoneNumber, nil
}
