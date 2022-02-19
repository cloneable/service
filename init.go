package service

import (
	"context"

	"github.com/cloneable/service/log"
)

func Init(ctx context.Context, options ...InitOption) (context.Context, error) {
	ctx = log.Inject(ctx)
	ctx = initTracing(ctx)

	var cfg InitConfig
	for _, option := range options {
		if option != nil {
			if err := option(ctx, &cfg); err != nil {
				return ctx, err
			}
		}
	}
	return ctx, nil
}

type InitOption func(ctx context.Context, cfg *InitConfig) error

type InitConfig struct{}
