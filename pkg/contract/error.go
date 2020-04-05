package contract

import (
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	HTTPCode    int    `json:"code"`
	Message     string `json:"mesage"`
	Description string `json:"description"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[error] %s::%s", e.Message, e.Description)
}

func (e *Error) SetDescription(d string) {
	e.Description += " " + d
}

func newError(code int, msg ...string) func(d ...string) *Error {
	return func(d ...string) *Error {
		var message string
		if len(msg) == 0 {
			message = http.StatusText(code)
		} else {
			message = msg[0]
		}
		e := &Error{
			HTTPCode: code,
			Message:  message,
		}
		e.SetDescription(strings.Join(d, " "))
		return e
	}
}

var (
	ErrMethodNotAllowed    = newError(http.StatusMethodNotAllowed)
	ErrBadRequest          = newError(http.StatusBadRequest)
	ErrBadParam            = newError(http.StatusBadRequest, "invalid parameter")
	ErrInternalServerError = newError(http.StatusInternalServerError)
)
