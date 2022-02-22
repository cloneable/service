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
			if err := option.apply(ctx, &svcCtx.runConfig); err != nil {
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

type runConfig struct {
	ShutdownCallbacks []func(context.Context) error
}

type RunOption interface {
	apply(ctx context.Context, cfg *runConfig) error
}

type runOptionFunc func(ctx context.Context, cfg *runConfig) error

func (f runOptionFunc) apply(ctx context.Context, cfg *runConfig) error { return f(ctx, cfg) }

func WithShutdownCallback(fn func(context.Context) error) RunOption {
	return runOptionFunc(func(ctx context.Context, cfg *runConfig) error {
		cfg.ShutdownCallbacks = append(cfg.ShutdownCallbacks, fn)
		return nil
	})
}
