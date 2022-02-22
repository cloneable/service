package service

import (
	"context"

	"github.com/cloneable/service/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func Init(ctx context.Context, options ...InitOption) (context.Context, error) {
	ctx, svcCtx := getServiceContext(ctx)
	ctx = log.Inject(ctx)

	for _, option := range options {
		if option != nil {
			if err := option.apply(ctx, &svcCtx.initConfig); err != nil {
				return ctx, err
			}
		}
	}

	ctx, tp := initTracing(ctx)
	svcCtx.initConfig.tracerProvider = tp
	svcCtx.runConfig.ShutdownCallbacks = append(svcCtx.runConfig.ShutdownCallbacks, tp.Shutdown)

	svcCtx.ready = true
	return ctx, nil
}

type initConfig struct {
	tracerProvider *sdktrace.TracerProvider
}

type InitOption interface {
	apply(ctx context.Context, cfg *initConfig) error
}

type initOptionFunc func(ctx context.Context, cfg *runConfig) error

func (f initOptionFunc) apply(ctx context.Context, cfg *runConfig) error { return f(ctx, cfg) }
