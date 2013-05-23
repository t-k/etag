etag
====

[![Build Status](https://travis-ci.org/t-k/etag.png?branch=master)](https://travis-ci.org/t-k/etag)

## Provide etag support for Go web server.

## How to install

```
go get github.com/t-k/etag
```

## Usage

Install the package with `go get` and use `import` to include it in your project.

```
import "github.com/t-k/etag"
```

GoDoc: http://godoc.org/github.com/t-k/etag

##Example

```go
package main

import (
  "github.com/t-k/etag"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if etag.IsStale(w, r, "etag-key") {
      // normal response processing
      w.Write([]byte("Hello world!"))
    } else {
      // you don't need to do anything.
    }
  })
  http.ListenAndServe(":8000", nil)
}

```

## Tests
```
go test
```
