package validate

import (
	"fmt"
	"regexp"
)

var validatePasswordFunc = defaultValidatePasswordFunc
var validateEmailFunc = defaultValidateEmailFunc

func SetValidatePasswordFunc(function ValidatePasswordFunc) {
	validatePasswordFunc = function
}

func ResetValidatePasswordFunc() {
	validatePasswordFunc = defaultValidatePasswordFunc
}

func SetValidateEmailFunc(function ValidateEmailFunc) {
	validateEmailFunc = function
}

func ResetValidateEmailFunc() {
	validateEmailFunc = defaultValidateEmailFunc
}

func test() {
}

var defaultValidatePasswordFunc ValidatePasswordFunc = func(password string, fieldName string) []string {
	errors := make([]string, 0)

	number := false
	specialCharacter := false
	upperCase := false
	lowerCase := false

	for _, c := range []byte(password) {
		if IsNumber(c) {
			number = true
		} else if IsSpecialCharacter(c) {
			specialCharacter = true
		} else if c >= 'a' && c <= 'z' {
			lowerCase = true
		} else if c >= 'A' && c <= 'Z' {
			upperCase = true
		}
	}

	if !number {
		errors = append(errors, fmt.Sprintf("Field '%s' requires at least one digit", fieldName))
	}
	if !specialCharacter {
		errors = append(errors, fmt.Sprintf("Field '%s' requires at least one special character", fieldName))
	}
	if !upperCase {
		errors = append(errors, fmt.Sprintf("Field '%s' requires at least one upper case letter", fieldName))
	}
	if !lowerCase {
		errors = append(errors, fmt.Sprintf("Field '%s' requires at least one lower case letter", fieldName))
	}

	return errors
}

var defaultValidateEmailFunc ValidateEmailFunc = func(email string, fieldName string) string {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if ok := regexp.MustCompile(regex).MatchString(email); !ok {
		return fmt.Sprintf("Field '%s' is not a valid email", fieldName)
	}
	return ""
}

func IsNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func IsSpecialCharacter(c byte) bool {
	if (c >= '!' && c <= '/') ||
		(c >= ':' && c <= '@') ||
		(c >= '[' && c <= '^') ||
		(c >= '{' && c <= '~') {
		return true
	}

	return false
}
