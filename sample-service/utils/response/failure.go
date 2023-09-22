package response

import "fmt"

type Code string

type Failure struct {
	Code Code   `json:"code"`
	Desc string `json:"desc"`
}

func New(code Code) Failure {
	return Failure{Code: code}
}

func WithMessage(code Code, desc string) Failure {
	return Failure{Code: code, Desc: desc}
}

func (f Failure) Error() string {
	return fmt.Sprintf("%s: %s", f.Code, f.Desc)
}

const (
	CodeInternal           Code = "Internal"
	CodeUnauthorized       Code = "Unauthorized"
	CodeInvalidRequest     Code = "InvalidRequest"
	CodeEntityAlreadyExist Code = "EntityAlreadyExist"
	CodeEntityNotFound     Code = "EntityNotFound"
)
