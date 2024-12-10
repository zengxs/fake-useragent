# fake-useragent

A fake user agent generator for Go language.

## Installation

```bash
go get -u github.com/zengxs/fake-useragent
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    fakeua "github.com/zengxs/fake-useragent"
)

func main() {
    ua := fakeua.New()
    fmt.Println(ua.Random())
}
```

### With Filters

```go
import (
    "fmt"
    fakeua "github.com/zengxs/fake-useragent"
)

ua := fakeua.NewUserAgentGenerator(
    fakeua.WithTagFilters(
        fakeua.TagDesktop,
        fakeua.TagChrome,
        fakeua.TagSafari,
        fakeua.TagFirefox,
        fakeua.TagMSEdge,
    ),
    fakeua.WithFilterMode(fakeua.FilterModeInclude),  // only user-agents with specified tags will be included
)

fmt.Println(ua.Random())
```
