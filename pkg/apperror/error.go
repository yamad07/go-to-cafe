package apperror

type Code string

const (
	CodeError        Code = "error"
	CodeInvalid      Code = "invalid"
	CodeSuccess      Code = "sucess"
	CodeNotFound     Code = "not found"
	CodeUnauthorized Code = "unauthorized"
)

type Error interface {
	Code() Code
	Error() string
}

type appError struct {
	code    Code
	err     error
	message string
}

func New(code Code, err error) appError {
	return appError{
		code:    code,
		err:     err,
		message: err.Error(),
	}
}

func (e appError) Code() Code {
	return e.code
}

func (e appError) Error() string {
	if e.message != "" {
		return e.message
	}
	return e.err.Error()
}
