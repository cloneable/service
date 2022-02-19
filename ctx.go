package service

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type serviceContextKey struct{}

type serviceContextValue struct {
	*serviceContext
}

type serviceContext struct {
	tracerProvider *sdktrace.TracerProvider
	ready          bool
}

func getServiceContext(ctx context.Context) (context.Context, *serviceContext) {
	v, ok := ctx.Value(serviceContextKey{}).(serviceContextValue)
	if !ok || v.serviceContext == nil {
		v = serviceContextValue{new(serviceContext)}
		ctx = context.WithValue(ctx, serviceContextKey{}, v)
	}
	return ctx, v.serviceContext
}
