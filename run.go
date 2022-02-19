package service

import "context"

func Run(ctx context.Context, options ...RunOption) error {
	var cfg RunConfig
	for _, option := range options {
		if option != nil {
			if err := option(ctx, &cfg); err != nil {
				return err
			}
		}
	}
	return nil
}

type RunOption func(ctx context.Context, cfg *RunConfig) error

type RunConfig struct{}
