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
		log.Println("No .env file found")
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

# ReadMe (hehe)?

I built this for my own use on a project I'm working on and I decided to open source it. <br/>
So I don't even think it's ready for public consumption now. It works fine for me. <br/>
Will make it more generic and add more features as time goes on. <br/>

# Todo

- [] Matrix API
