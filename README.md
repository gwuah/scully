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
I'm working on a tool using go and couldn't find a golang mapbox library that suited me.
So I set out to create one for my usage. Hopefully it helps you too.
It returns pure json, unparsed, it's up to you to decode the byte and use it.

# Todo
- [] Matrix API
