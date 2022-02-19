package service

import (
	"context"

	"github.com/cloneable/service/log"
)

func Init(ctx context.Context, options ...InitOption) (context.Context, error) {
	ctx, svcCtx := getServiceContext(ctx)
	ctx = log.Inject(ctx)
	ctx, tp := initTracing(ctx)

	svcCtx.tracerProvider = tp

	var cfg InitConfig
	for _, option := range options {
		if option != nil {
			if err := option(ctx, &cfg); err != nil {
				return ctx, err
			}
		}
	}

	svcCtx.ready = true
	return ctx, nil
}

type InitOption func(ctx context.Context, cfg *InitConfig) error

type InitConfig struct{}
