package service

import (
	"context"
	"errors"
)

func Run(ctx context.Context, options ...RunOption) error {
	if !checkMarker(ctx) {
		return errors.New("Run() called with wrong context")
	}

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
