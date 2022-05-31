package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"msg"`
	Internal error                  `json:"internal"`
	Metadata map[string]interface{} `json:"metadata"`
}

func New(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func NewErrorWithInternal(code int, msg string, err error) *Error {
	return &Error{
		Code:     code,
		Msg:      msg,
		Internal: err,
	}
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); As(err, &se) {
		return se
	}

	return NewErrorWithInternal(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d msg = %s", e.Code, e.Msg)
}

func (e *Error) WithInternal(err error) *Error {
	return &Error{
		Code:     e.Code,
		Msg:      e.Msg,
		Internal: err,
		Metadata: e.Metadata,
	}
}

func (e *Error) WithMetadata(md map[string]interface{}) *Error {
	return &Error{
		Code:     e.Code,
		Msg:      e.Msg,
		Internal: e.Internal,
		Metadata: md,
	}
}
