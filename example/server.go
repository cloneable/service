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

	if err := service.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
