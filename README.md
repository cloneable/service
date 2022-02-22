# github.com/cloneable/service &mdash; A Go Microservice Chassis

## Usage

```go
package main

import (
	"context"
	"log"

	"github.com/cloneable/service"
)

func main() {
	ctx, err := service.Init(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Set up server(s)...

	// Run() comes last. It never returns.
	service.Run(ctx)
}
```

## `service.Init()`

## `service.Run()`
