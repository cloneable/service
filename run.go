package service

import "context"

func Run(ctx context.Context, options ...RunOption) error {
	return nil
}

type RunOption func(cfg *RunConfig) error

type RunConfig struct{}
