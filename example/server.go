package main

import (
	"context"

	"github.com/cloneable/service"
	"github.com/cloneable/service/log"
)

func main() {
	ctx, err := service.Init(context.Background())
	if err != nil {
		log.S(ctx).Fatal(err)
	}

	// Set up server(s)...

	if err := service.Run(ctx); err != nil {
		log.S(ctx).Fatal(err)
	}
}
