package apifs

import "fmt"

type ErrorCode uint32

type Error struct {
	code ErrorCode
	str  string
	err  error
}

const (
	ErrorInternal ErrorCode = iota
	ErrorInvalid
	ErrorDoNotExist
	ErrorVersion
)

var errorCode2Str = map[ErrorCode]string{
	ErrorInternal:   "internal error",
	ErrorInvalid:    "invalid",
	ErrorDoNotExist: "do not exist",
	ErrorVersion:    "version mismatch",
}

func ErrWithError(code ErrorCode, err error) *Error {
	switch err.(type) {
	case *Error:
		return err.(*Error)

	default:
		return &Error{
			code: code,
			err:  err,
		}
	}
}

func ErrWithString(code ErrorCode, s string) *Error {
	return &Error{
		code: code,
		str:  s,
	}
}

func (e *Error) Error() string {
	str, ok := errorCode2Str[e.code]
	if !ok {
		return "unknown error"
	}

	if e.err != nil {
		str = fmt.Sprintf("%s / %s", str, e.err.Error())
	}

	if e.str != "" {
		str = fmt.Sprintf("%s / %s", str, e.str)
	}

	return str
}

func IsErrorCode(err error, code ErrorCode) bool {
	switch err.(type) {
	case *Error:
		myErr := err.(*Error)
		return myErr.code == code

	default:
		return false
	}
}
