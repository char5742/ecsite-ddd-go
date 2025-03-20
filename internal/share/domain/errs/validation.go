package shareerrs

type ValidationError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}
type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return "validation errors"
}
