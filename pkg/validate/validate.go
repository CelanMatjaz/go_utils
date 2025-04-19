package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Validate(value any) []string {
	errors := make([]string, 0)
	v := reflect.ValueOf(value)

	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		if field.Kind() == reflect.Struct {
			errors = append(errors, Validate(field.Interface())...)
			continue
		}

		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

		fieldName := fieldType.Tag.Get("json")
		if fieldName == "" {
			fieldName = fieldType.Name
		}

		errors = append(errors, validateField(field, fieldName, tag)...)
	}

	return errors
}

func validateField(val reflect.Value, fieldName string, tag string) []string {
	errors := make([]string, 0)
	validations := strings.SplitSeq(tag, ",")

	for validation := range validations {
		if strings.HasPrefix(validation, "min:") {
			if errorString := validateMin(validation, val, fieldName); errorString != "" {
				errors = append(errors, errorString)
			}
		} else if strings.HasPrefix(validation, "max:") {
			if errorString := validateMax(validation, val, fieldName); errorString != "" {
				errors = append(errors, errorString)
			}
		} else if strings.HasPrefix(validation, "len:") {
			if errorString := validateMax(validation, val, fieldName); errorString != "" {
				errors = append(errors, errorString)
			}
		} else if validation == "required" {
			if errorString := validateRequired(val, fieldName); errorString != "" {
				errors = append(errors, errorString)
			}
		} else if validation == "password" {
			if errorStrings := validatePassword(val, fieldName); len(errorStrings) > 0 {
				errors = append(errors, errorStrings...)
			}
		} else if validation == "email" {
			if errorString := validateEmail(val, fieldName); errorString != "" {
				errors = append(errors, errorString)
			}
		}
	}

	return errors
}

func validateMin(validate string, val reflect.Value, fieldName string) string {
	min, _ := strconv.Atoi(strings.TrimPrefix(validate, "min:"))
	if len(val.String()) < min {
		return fmt.Sprintf("Field '%s' must be at least %d characters long", fieldName, min)
	}
	return ""
}

func validateMax(validate string, val reflect.Value, field string) string {
	max, _ := strconv.Atoi(strings.TrimPrefix(validate, "max:"))
	if len(val.String()) > max {
		return fmt.Sprintf("Field '%s' must be at most %d characters long", field, max)
	}
	return ""
}

func validateLen(validate string, val reflect.Value, field string) string {
	length, _ := strconv.Atoi(strings.TrimPrefix(validate, "len:"))
	if len(val.String()) != length {
		return fmt.Sprintf("Field '%s' must be at most %d characters long", field, length)
	}
	return ""
}

func validateRequired(val reflect.Value, fieldName string) string {
	if val.String() == "" {
		return fmt.Sprintf("Field '%s' is required", fieldName)
	}
	return ""
}

func validateEmail(val reflect.Value, fieldName string) string {
	return validateEmailFunc(val.String(), fieldName)
}

func validatePassword(val reflect.Value, fieldName string) []string {
	return validatePasswordFunc(val.String(), fieldName)
}
