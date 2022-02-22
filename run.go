package service

import (
	"context"
	"os"

	"github.com/cloneable/service/log"
)

func Run(ctx context.Context, options ...RunOption) {
	ctx, svcCtx := getServiceContext(ctx)
	if !svcCtx.ready {
		log.L(ctx).Fatal("Run() called with wrong context")
	}

	for _, option := range options {
		if option != nil {
			if err := option.apply(ctx, &svcCtx.runConfig); err != nil {
				log.S(ctx).Fatalf("RunOption failed: %v", err)
			}
		}
	}

	// TODO: move into shutdown handler
	for _, fn := range svcCtx.runConfig.shutdownCallbacks {
		if err := fn(ctx); err != nil {
			log.S(ctx).Errorf("shutdown function failed: %v", err)
		}
	}

	// TODO: add termination function, overridable
	os.Exit(0)
}

type runConfig struct {
	shutdownCallbacks []func(context.Context) error
}

type RunOption interface {
	apply(ctx context.Context, cfg *runConfig) error
}

type runOptionFunc func(ctx context.Context, cfg *runConfig) error

func (f runOptionFunc) apply(ctx context.Context, cfg *runConfig) error { return f(ctx, cfg) }

func WithShutdownCallback(fn func(context.Context) error) RunOption {
	return runOptionFunc(func(ctx context.Context, cfg *runConfig) error {
		cfg.shutdownCallbacks = append(cfg.shutdownCallbacks, fn)
		return nil
	})
}
