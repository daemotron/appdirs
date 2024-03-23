# Appdirs for Go

![Build and Test State](https://github.com/daemotron/appdirs/actions/workflows/go.yml/badge.svg?event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/daemotron/appdirs.svg)](https://pkg.go.dev/github.com/daemotron/appdirs)

This is my implementation of [ActiveState's appdirs](https://github.com/ActiveState/appdirs) for Go.

## Installation

Use `go get` to retrieve the module and add it to your `GOPATH` workspace,
or project's Go module dependencies:

```shell
go get github.com/daemotron/appdirs
```

To update the SDK, use `go get -U` to retrieve the latest version of the module:

```shell
go get -U github.com/daemotron/appdirs
```

This module has been tested with Go 1.18 and 1.22.

## Documentation

### Usage Example

```go
package main

import (
    "log"
    "github.com/daemotron/appdirs"
)

func main() {
    // create application context
    app := appdirs.AppConf{Name: "my-app"}

    // get the user-specific configuration directory and print it
    userDir, err := app.UserConfigDir()
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(userDir)
}
```

### API Reference

Full documentation for this package is available via the
[Go Package Index](https://pkg.go.dev/github.com/daemotron/appdirs)
