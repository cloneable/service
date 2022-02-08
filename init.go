package service

import "context"

func Init(ctx context.Context, options ...InitOption) (context.Context, error) {
	return ctx, nil
}

type InitOption func(cfg *InitConfig) error

type InitConfig struct{}
