package service

import (
	"context"
)

type serviceContextKey struct{}

type serviceContextValue struct {
	*serviceContext
}

type serviceContext struct {
	initConfig InitConfig
	runConfig  RunConfig

	ready bool
}

func getServiceContext(ctx context.Context) (context.Context, *serviceContext) {
	v, ok := ctx.Value(serviceContextKey{}).(serviceContextValue)
	if !ok || v.serviceContext == nil {
		v = serviceContextValue{new(serviceContext)}
		ctx = context.WithValue(ctx, serviceContextKey{}, v)
	}
	return ctx, v.serviceContext
}
