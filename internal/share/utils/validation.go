package shareutils

// ValidationError はバリデーションエラーを表す構造体です。
// フィールド名とエラーを保持します。
type validationError struct {
	Field string
	Err   error
}

// ValidationErrors は複数のバリデーションエラーを保持するためのスライスです。
type ValidationErrors []validationError

func (errs ValidationErrors) Error() string {
	if len(errs) == 0 {
		return ""
	}

	var result string
	for _, err := range errs {
		result += err.Field + ": " + err.Err.Error() + "\n"
	}
	return result
}
func (errs ValidationErrors) IsEmpty() bool {
	return len(errs) == 0
}
func (errs ValidationErrors) Add(
	field string,
	err error,
) ValidationErrors {
	return append(errs, validationError{Field: field, Err: err})
}
