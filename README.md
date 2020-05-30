# scully

A golang wrapper around mapbox apis.

# Usage

```go
package main

import (
  "fmt"
  "https://github.com/gwuah/scully"
)

func main() {
  mapbox := scully.New()
  fmt.Println(mapbox.Matrix.Greet())
  // this should print "Hello World"
}
```

# Why?

Built for my own use in a tool T'm working on and I decided to open source it.

# Todo

- [] Matrix API
