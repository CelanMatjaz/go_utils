package validate

import (
	"reflect"
	"testing"
)

type TestCaseStruct struct {
	validateString string
	value          string
	shouldFail     bool
}

func TestValidateMin(t *testing.T) {
	t.Parallel()

	testCases := []TestCaseStruct{
		{validateString: "min:0", value: "", shouldFail: false},
		{validateString: "min:1", value: "", shouldFail: true},
		{validateString: "min:10", value: "TestTest", shouldFail: true},
		{validateString: "min:10", value: "TestTestTestTest", shouldFail: false},
		{validateString: "min:", value: "", shouldFail: false},
	}

	for i, tc := range testCases {
		errorString := validateMin(tc.validateString, reflect.ValueOf(tc.value), "Field")
		if (errorString != "") != tc.shouldFail {
			t.Errorf(
				"validateMin() failed for test case with index %d, validation: '%s', value '%s', should fail: %t",
				i, tc.validateString, tc.value, tc.shouldFail)
		}
	}
}

func TestValidateMax(t *testing.T) {
	t.Parallel()

	testCases := []TestCaseStruct{
		{validateString: "max:0", value: "", shouldFail: false},
		{validateString: "max:1", value: "", shouldFail: false},
		{validateString: "max:0", value: "A", shouldFail: true},
		{validateString: "max:1", value: "AA", shouldFail: true},
		{validateString: "max:", value: "", shouldFail: false},
		{validateString: "max:", value: "A", shouldFail: true},
	}

	for i, tc := range testCases {
		errorString := validateMax(tc.validateString, reflect.ValueOf(tc.value), "Field")
		if (errorString != "") != tc.shouldFail {
			t.Errorf(
				"validateMax() failed for test case with index %d, validation: '%s', value '%s', should fail: %t",
				i, tc.validateString, tc.value, tc.shouldFail)
		}
	}
}

func TestValidateLen(t *testing.T) {
	t.Parallel()

	testCases := []TestCaseStruct{
		{validateString: "len:0", value: "", shouldFail: false},
		{validateString: "len:0", value: "A", shouldFail: true},
		{validateString: "len:1", value: "", shouldFail: true},
		{validateString: "len:1", value: "A", shouldFail: false},
		{validateString: "len:", value: "", shouldFail: false},
		{validateString: "len:", value: "A", shouldFail: true},
	}

	for i, tc := range testCases {
		errorString := validateLen(tc.validateString, reflect.ValueOf(tc.value), "Field")
		if (errorString != "") != tc.shouldFail {
			t.Errorf(
				"validateLen() failed for test case with index %d, validation: '%s', value '%s', should fail: %t",
				i, tc.validateString, reflect.ValueOf(tc.value), tc.shouldFail)
		}
	}
}

func TestValidatePassword(t *testing.T) {
	t.Parallel()

	testCases := []TestCaseStruct{
		// Check for number
		{value: "Passwo1!", shouldFail: false},
		{value: "Passwor!", shouldFail: true},

		// Check for special character
		{value: "Passwo1!", shouldFail: false},
		{value: "Passwor1", shouldFail: true},

		// Check for lower case character
		{value: "Passwo1!", shouldFail: false},
		{value: "PASSWO1!", shouldFail: true},

		// Check for upper case character
		{value: "Passwo1!", shouldFail: false},
		{value: "passwo1!", shouldFail: true},
	}

	for i, tc := range testCases {
		errorStrings := validatePassword(reflect.ValueOf(tc.value), "Field")
		if (len(errorStrings) > 0) && !tc.shouldFail {
			t.Errorf(
				"validatePassword() failed for test case with index %d, value '%s', should fail: %t, error message count: %d",
				i, tc.value, tc.shouldFail, len(errorStrings))
		}
	}
}

func TestValidateRequired(t *testing.T) {
	t.Parallel()

	testCases := []TestCaseStruct{
		{value: "", shouldFail: true},
		{value: "Required", shouldFail: false},
	}

	for i, tc := range testCases {
		errorString := validateRequired(reflect.ValueOf(tc.value), "Field")
		if (errorString != "") != tc.shouldFail {
			t.Errorf(
				"validateRequired() failed for test case with index %d, value '%s', should fail: %t",
				i, tc.value, tc.shouldFail)
		}
	}
}

func TestValidateEmail(t *testing.T) {
	t.Parallel()

	testCases := []TestCaseStruct{
		{value: "test@test.com", shouldFail: false},
		{value: "test@test.co", shouldFail: false},
		{value: "test@test.c", shouldFail: true},
		{value: "@test.com", shouldFail: true},
		{value: "test@.com", shouldFail: true},
		{value: "test.com", shouldFail: true},
	}

	for i, tc := range testCases {
		errorString := validateEmail(reflect.ValueOf(tc.value), "Field")
		if (errorString != "") != tc.shouldFail {
			t.Errorf(
				"validateEmail() failed for test case with index %d, value '%s', should fail: %t",
				i, tc.value, tc.shouldFail)
		}
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	type EmailStruct struct {
		email string `validate:"email"`
	}

	type RequiredStruct struct {
		required string `validate:"required"`
	}

	type PasswordStruct struct {
		password string `validate:"password"`
	}

	type MinMaxStruct struct {
		field1 string `validate:"min:1,max:10"`
		field2 string `validate:"min:3,max:6"`
	}

	type Mixed struct {
		email    string `validate:"required,email,min:2,max:100"`
		password string `validate:"required,min:8,max:10"`
		MinMaxStruct
	}

	type ValidateTestCase struct {
		errors     []string
		shouldFail bool
	}

	testCases := []ValidateTestCase{
		{errors: Validate(EmailStruct{email: "test@test.com"}), shouldFail: false},
		{errors: Validate(EmailStruct{email: ""}), shouldFail: true},
		{errors: Validate(RequiredStruct{required: "required"}), shouldFail: false},
		{errors: Validate(RequiredStruct{required: ""}), shouldFail: true},
		{errors: Validate(PasswordStruct{password: ""}), shouldFail: true},
		{errors: Validate(PasswordStruct{password: "password"}), shouldFail: true},
		{errors: Validate(PasswordStruct{password: "Passwo1!"}), shouldFail: false},
		{errors: Validate(MinMaxStruct{field1: "AAAA", field2: "AAAA"}), shouldFail: false},
		{errors: Validate(MinMaxStruct{field1: "AA", field2: "AA"}), shouldFail: true},
		{errors: Validate(Mixed{email: "AA", password: "AAAAAAAAAAAAA"}), shouldFail: true},
		{errors: Validate(Mixed{email: "ok@email.com", password: "AAAAAAAa1!"}), shouldFail: true},
		{errors: Validate(Mixed{email: "ok@email.com", password: "AAAAAAAa1!", MinMaxStruct: MinMaxStruct{field1: "AAAA", field2: "AAAA"}}), shouldFail: false},
	}

	for i, tc := range testCases {
		if (len(tc.errors) > 0) && !tc.shouldFail {
			t.Errorf(
				"Validate() failed for test case with index %d, should fail: %t, error count: %d",
				i, tc.shouldFail, len(tc.errors))
		}
	}
}
