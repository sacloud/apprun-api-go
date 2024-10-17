package fake

import "fmt"

// ErrorType fakeサーバが扱うエラーの種別
type ErrorType int

const (
	ErrorTypeUnknown        ErrorType = iota // 未知のエラー
	ErrorTypeInvalidRequest                  // Bad Request
	ErrorTypeNotFound                        // Not Found
	ErrorTypeConflict                        // Conflict
)

// String Stringer実装
func (e ErrorType) String() string {
	switch e {
	case ErrorTypeInvalidRequest:
		return "invalid request"
	case ErrorTypeNotFound:
		return "not found"
	case ErrorTypeConflict:
		return "conflict"
	default:
		return "unknown error"
	}
}

// Error fake.Engineが出力するエラー型
type Error struct {
	Type          ErrorType
	Resource      string
	Id            interface{}
	msgFmtAndVars []interface{}
}

func newError(errorType ErrorType, resource string, id interface{}, msgFmtAndVars ...interface{}) *Error {
	return &Error{
		Type:          errorType,
		Resource:      resource,
		Id:            id,
		msgFmtAndVars: msgFmtAndVars,
	}
}

// Error errorインターフェースの実装
func (e *Error) Error() string {
	return fmt.Errorf("%s: %s[%s]%s", e.Type, e.Resource, e.Id, e.message()).Error()
}

func (e *Error) message() string {
	switch len(e.msgFmtAndVars) {
	case 0:
		return ""
	case 1:
		return " " + e.msgFmtAndVars[0].(string)
	default:
		return " " + fmt.Sprintf(e.msgFmtAndVars[0].(string), e.msgFmtAndVars[1:]...)
	}
}
