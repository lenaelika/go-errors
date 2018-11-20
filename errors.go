// Primitives for cascading error handling w/o stacktrace.
package errors

import "fmt"

// basic is a simple text error.
type basic struct {
	msg string
}

// wrap is an annotated error.
type wrap struct {
	cause error
	msg   string
}

// causer can determine an underlying error.
type causer interface {
	Cause() error
}

// New creates an error with the given message.
func New(msg string) error {
	return &basic{msg}
}

// Errorf creates an error with a message according to a format specifier.
func Errorf(format string, args ...interface{}) error {
	return &basic{fmt.Sprintf(format, args...)}
}

// Error returns the error message.
func (e *basic) Error() string {
	return e.msg
}

// Error returns the annotated error message.
func (e *wrap) Error() string {
	return e.msg + ": " + e.cause.Error()
}

// Cause returns the underlying cause of the error.
func (e *wrap) Cause() error {
	return e.cause
}

// Wrap annotates an error. If err is nil, Wrap returns nil.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &wrap{err, msg}
}

// Wrapf annotates an error using a format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrap{err, fmt.Sprintf(format, args...)}
}

// Cause returns an underlying cause of the error, if possible.
// If err does not implement Cause, it will be returned itself.
func Cause(err error) error {
	for err != nil {
		e, ok := err.(causer)
		if !ok {
			break
		}
		err = e.Cause()
	}
	return err
}
