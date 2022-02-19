package service

import "context"

func Init(ctx context.Context, options ...InitOption) (context.Context, error) {
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
