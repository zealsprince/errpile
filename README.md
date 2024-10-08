# errpile

A minimal Go library for error wrapping with function trace context.

## Overview

`errpile` is a lightweight Go package designed to enhance error handling by providing additional context about where an error originated. By capturing the function name, and optionally the file name and line number, `errpile` aims to make debugging easier and more intuitive.

## Features

- **Minimal Setup**: Easy to integrate into any Go project without overhead.
- **Function Trace**: Automatically adds the calling function's name to error messages.
- **Optional Line Number**: Can include file name and line number for even more context.
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

## Usage

Here's a quick example of how to use `errpile` in your project:

```go
package main

import (
 "errors"
 "fmt"
 "github.com/zealsprince/errpile"
)

func riskyOperation() error {
 return errors.New("something went wrong")
}

func main() {
 err := riskyOperation()

 if err != nil {
  // Use the default Error function to wrap the error
  wrappedErr := errpile.Error(err)
  fmt.Println(wrappedErr) // Output: main.riskyOperation encountered an error: something went wrong
 }
}
```

### Custom Error Pile

You can also create a custom `ErrorPile` with specific configuration options:

```go
func main() {
 customPile := errpile.New("%s encountered an error: %v", 1, true)

 err := riskyOperation()
 if err != nil {
  wrappedErr := customPile.Error(err)
  fmt.Println(wrappedErr)
 }
}
```

## Why use this?

Oftentimes you'll find yourself sending errors up the call stack without any context about where they originated. This can make debugging difficult, especially in larger codebases or when creating concurrent applications. By using `errpile`, you can add context to your errors without much overhead, making it easier to identify where an error occurred.

## License

This project is licensed under the MIT [License](LICENSE).
