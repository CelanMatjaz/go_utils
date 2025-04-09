package validate

type ValidateEmailFunc = func(value string, fieldName string) string
type ValidatePasswordFunc = func(value string, fieldName string) []string
