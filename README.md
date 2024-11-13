# fake-useragent

A fake user agent generator for Go language.

## Installation

```bash
go get -u github.com/zengxs/fake-useragent
```

## Usage

```go
package main

import (
    "fmt"
    fakeuseragent "github.com/zengxs/fake-useragent"
)

func main() {
    ua := fakeuseragent.New()
    fmt.Println(ua.Random())
}
```
