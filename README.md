# errpile

A minimal Go library for error wrapping with function trace context.

## Overview

`errpile` is a lightweight Go package designed to enhance error handling by providing additional context about where an error originated. By capturing the function name, and optionally the file name and line number, `errpile` aims to make debugging easier and more intuitive.

## Features

- **Minimal Setup**: Easy to integrate into any Go project without overhead.
- **Function Trace**: Automatically adds the calling function's name to error messages.
- **Optional Complete Trace**: Choose to include the original or each function in the call chain.
- **Customizable Formatting**: Allows developers to define their error message format.

## Installation

To install `errpile`, use `go get`:

```bash
go get github.com/zealsprince/errpile
```

Then import it in your Go code:

```go
import "github.com/zealsprince/errpile"
```

## Usage Example

Here's a concise example of how to use `errpile` in a chain of function calls:

```go
package main

import (
 "errors"
 "fmt"
 "github.com/zealsprince/errpile"
)

// Functions that return wrapped errors using errpile.
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
  fmt.Println(err) // Output will include function name and line number.
 }
}
```

With this setup, if an error occurs, `errpile` will automatically trace back to the relevant function and provide context, making it easier to pinpoint issues.

Running the main example included in the repository will output:

    main.riskiestOperation->something went wrong (/[...]/errpile/cmd/example/main.go:13)

## Why use this?

Oftentimes you'll find yourself sending errors up the call stack without any context about where they originated. This can make debugging difficult, especially in larger codebases or when creating concurrent applications. By using `errpile`, you can add context to your errors without much overhead, making it easier to identify where an error occurred.

To summarize:

- **Debugging**: Quickly identify where errors are originating in complex systems.
- **Logging**: Enhance log messages with trace context to provide better insights into runtime issues.

## Customizing the Error Pile

`errpile` provides flexibility through its configurable `ErrorPile` structure, allowing you to tailor the error wrapping to suit your needs.

### Custom Options

You can create a custom `ErrorPile` instance with the following parameters:

- **Template**: A string format that specifies how error messages are composed. It can include placeholders like `%s` for the function name and `%v` for the error message.
  
- **Show Line Number**: A boolean to determine whether the line number and file name should be included in the error message.
  
- **Track Full Chain**: A boolean that decides if the error trace should include updates at each function in the call chain or only the initial occurrence.

### Creating a Custom ErrorPile

Here's a quick example demonstrating how to create and use a custom `ErrorPile`:

```go
package main

import (
 "errors"
 "fmt"
 "github.com/zealsprince/errpile"
)

// Create a customized ErrorPile instance.
var customPile = errpile.New("%[1]s encountered error: %[2]v", true, true)

func anotherRiskyOperation() error {
 err := errors.New("operation failed")
 return customPile.Error(err)
}

func main() {
 err := anotherRiskyOperation()
 if err != nil {
  fmt.Println(err) // Output will adhere to the customized format and options.
 }
}
```

## License

This project is licensed under the MIT [License](LICENSE).
