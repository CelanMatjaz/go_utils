## Validation utils

Utility for validating string fields on structs using reflection and returning readable errors.

### Validation features

 - min/max string lengths
 - specific string lengths
 - required strings
 - email strings
 - password strings

### Usage

Add annotations/tags to fields that you'd like to validate. They can be mixed and matched to get the desired field validation.

```go
type ValidatableStruct struct {
    max string `validate:"max:X"`
    min string `validate:"min:X"`
    length string `validate:"len:X"`
    required string `validate:"required"`
    password string `validate:"password"`
    email string `validate:"email"`
}
```

The tags `max`, `min` and `len` take an extra argument in the form of a number after the `:`.

Next, you can call the `Validate` function to validate the fields on a struct.

```go
func main() {
    obj := ValidatableStruct{ ... }
	errors := Validate(obj)
	if len(errors) != 0 {
		 // Handle errors ...
	}
}
```

The functions for email and password validation can be overridden using setter functions. The function provided in the setter will then be used for every `Validate` call, if there's a valid tag on a field.

Override function definitions, which are also found in [types.go](types.go):
```go
type ValidateEmailFunc = func(value string, fieldName string) string
type ValidatePasswordFunc = func(value string, fieldName string) []string
```

Example overrides:
```go
func main(){
    ...
    // Set function to process the password
    SetValidatePasswordFunc(func (password string, fieldName string) []string {
        ...
    })
    // Use above function from here on out
    Validate(...)
    // Reset function to be used to default
    ResetValidatePasswordFunc()
    // Default function is used from here on out
    Validate(...)
    ...
}
```

### Examples

```go
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

func examples() {
    Validate(EmailStruct{email: "test@test.com"}) // Ok
    Validate(EmailStruct{email: "test.com"})      // Returns error(s)

    Validate(RequiredStruct{required: "required"}) // Ok
    Validate(RequiredStruct{required: ""})         // Returns error(s)

    Validate(PasswordStruct{password: "password"}) // Ok
    Validate(PasswordStruct{password: ""})         // Returns error(s)

    Validate(MinMaxStruct{field1: "AAAA", field2: "AAAA"}) // Ok
    Validate(MinMaxStruct{field1: "AA", field2: "AA"})     // Returns error

    Validate(Mixed{email: "AA", password: "AAAAAAAAAAAAA"}) // Returns errors
    Validate(Mixed{email: "ok@email.com", password: "AAAAAAAa1!"}) // Returns errors
    Validate(Mixed{email: "ok@email.com", password: "AAAAAAAa1!", MinMaxStruct: MinMaxStruct{field1: "AAAA", field2: "AAAA"}}) // Ok
}
```
