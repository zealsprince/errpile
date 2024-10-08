package errpile

import (
	"fmt"
	"runtime"
)

// ErrorPile provides options for customizing error messages.
type ErrorPile struct {
	template       string
	callerDepth    int
	showLineNumber bool
}

// New creates a new instance of ErrorPile with customizable options.
func New(template string, callerDepth int, showLineNumber bool) *ErrorPile {
	return &ErrorPile{
		template:       template,
		callerDepth:    callerDepth,
		showLineNumber: showLineNumber,
	}
}

// Error wraps an error with the calling function's name and other context.
func (ep *ErrorPile) Error(err error) error {
	if err == nil {
		return nil
	}

	counter, file, line, ok := runtime.Caller(ep.callerDepth)
	if !ok {
		return err
	}

	caller := runtime.FuncForPC(counter).Name()
	if ep.showLineNumber {
		return fmt.Errorf(ep.template+" (%s:%d)", caller, err, file, line)
	}

	return fmt.Errorf(ep.template, caller, err)
}

// DefaultErrorPile provides a default configuration.
var DefaultErrorPile = New("%s->%v", 1, true)

// Error is a shortcut function using the default ErrorPile configuration.
func Error(err error) error {
	return DefaultErrorPile.Error(err)
}
