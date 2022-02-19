package main

import (
	"context"

	"github.com/cloneable/service"
	"github.com/cloneable/service/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func main() {
	ctx, err := service.Init(context.Background())
	if err != nil {
		log.S(ctx).Fatal(err)
	}
	tracer = otel.Tracer("example")

	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	log.L(ctx).Info("Hallo!")

	doSomething(ctx)

	if err := service.Run(ctx); err != nil {
		log.S(ctx).Fatal(err)
	}
}

func doSomething(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "doSomething")
	defer span.End()

	span.AddEvent("Oh!")

	log.L(ctx).Info("Doing something...")
}
