package service

import (
	"context"
	"errors"

	"github.com/cloneable/service/log"
)

func Run(ctx context.Context, options ...RunOption) error {
	ctx, svcCtx := getServiceContext(ctx)
	if !svcCtx.ready {
		return errors.New("Run() called with wrong context")
	}

	for _, option := range options {
		if option != nil {
			if err := option(ctx, &svcCtx.runConfig); err != nil {
				return err
			}
		}
	}

	// TODO: move into shutdown handler
	for _, fn := range svcCtx.runConfig.ShutdownCallbacks {
		if err := fn(ctx); err != nil {
			log.S(ctx).Errorf("shutdown function failed: %v", err)
		}
	}
	return nil
}

type RunOption func(ctx context.Context, cfg *RunConfig) error

type RunConfig struct {
	ShutdownCallbacks []func(context.Context) error
}

func WithShutdownCallback(fn func(context.Context) error) RunOption {
	return func(ctx context.Context, cfg *RunConfig) error {
		cfg.ShutdownCallbacks = append(cfg.ShutdownCallbacks, fn)
		return nil
	}
}
