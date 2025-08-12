package models

type FieldError struct {
	Field string
	Msg string
}

func (e *FieldError) Error() string {
	return e.Field + ": " + e.Msg
}