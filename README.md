# scully

A golang wrapper around mapbox apis.

# Usage

```go
package main

import (
	  "fmt"
    "github.com/gwuah/scully"
    "github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		og.Println("No .env file found")
		return
	}

	mapbox, err := scully.New(os.Getenv("ACCESS_TOKEN"))

	if err != nil {
		log.Println(err)
		return
	}

	response, err := mapbox.Matrix.GetMatrix(os.Getenv("TEST_POINTS"))

	if err != nil {
		log.Println(err)
    		return
	}
	log.Println(response)

}
```

# Why?

Built for my own use in a tool T'm working on and I decided to open source it.

# Todo

- [] Matrix API
