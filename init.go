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
			if err := option(ctx, &svcCtx.initConfig); err != nil {
				return ctx, err
			}
		}
	}

	ctx, tp := initTracing(ctx)
	svcCtx.initConfig.TracerProvider = tp
	svcCtx.runConfig.ShutdownCallbacks = append(svcCtx.runConfig.ShutdownCallbacks, tp.Shutdown)

	svcCtx.ready = true
	return ctx, nil
}

type InitOption func(ctx context.Context, cfg *InitConfig) error

type InitConfig struct {
	TracerProvider *sdktrace.TracerProvider
}
