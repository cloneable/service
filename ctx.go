package service

import (
	"context"
)

type serviceContextKey struct{}

func injectMarker(ctx context.Context) context.Context {
	return context.WithValue(ctx, serviceContextKey{}, "marker")
}

func checkMarker(ctx context.Context) bool {
	v, ok := ctx.Value(serviceContextKey{}).(string)
	return ok && v == "marker"
}
