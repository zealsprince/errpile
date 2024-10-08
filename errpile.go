package errpile

import (
	"fmt"
	"runtime"
)

// CustomError is a struct that holds an error message and tracing context.
type CustomError struct {
	err error  // original error
	msg string // wrapped message
}

// Error implements the error interface for CustomError.
func (ce *CustomError) Error() string {
	return ce.msg
}

// ErrorPile provides options for customizing error messages.
type ErrorPile struct {
	template       string
	showLineNumber bool
	trackFullChain bool
}

// New creates a new instance of ErrorPile with customizable options.
func New(template string, showLineNumber, trackFullChain bool) *ErrorPile {
	return &ErrorPile{
		template:       template,
		showLineNumber: showLineNumber,
		trackFullChain: trackFullChain,
	}
}

// Error wraps an error with the calling function's name and context information.
func (ep *ErrorPile) Error(err error) error {
	if err == nil {
		return nil
	}

	counter, file, line, ok := runtime.Caller(2)
	if !ok {
		return err
	}

	caller := runtime.FuncForPC(counter).Name()

	// Check if tracking full chain is disabled and the error is already a CustomError.
	if !ep.trackFullChain {
		if existing, ok := err.(*CustomError); ok {
			return existing // Return the existing error without modifying it.
		}
	}

	var msg string
	if ep.showLineNumber {
		msg = fmt.Sprintf(ep.template+" (%s:%d)", caller, err, file, line)
	} else {
		msg = fmt.Sprintf(ep.template, caller, err)
	}

	return &CustomError{
		err: err,
		msg: msg,
	}
}

// DefaultErrorPile provides a default configuration.
var DefaultErrorPile = New("%[1]s->%[2]v", true, false)

// Error is a shortcut function using the default ErrorPile configuration.
func Error(err error) error {
	return DefaultErrorPile.Error(err)
}
