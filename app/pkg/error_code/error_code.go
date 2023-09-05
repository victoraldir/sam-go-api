package errorcode

type ErrorType string

const (
	InvalidUsername    ErrorType = "INVALID_USERNAME"
	InvalidDateOfBirth ErrorType = "INVALID_DATE_OF_BIRTH"
	NotFound           ErrorType = "NOT_FOUND"
)
