package main

import (
	"errors"
	"fmt"

	"github.com/zealsprince/errpile"
)

func riskiestOperation() error {
	err := errors.New("something went wrong")

	return errpile.Error(err)
}

func riskierOperation() error {
	return errpile.Error(riskiestOperation())
}

func riskyOperation() error {
	return errpile.Error(riskierOperation())
}

func main() {
	err := riskyOperation()

	if err != nil {
		// Use the default Error function to wrap the error
		wrappedErr := errpile.Error(err)
		fmt.Println(wrappedErr)
	}
}
