//

package config

import (
	"context"
	"errors"
	"fmt"
)

// ErrNotConfigurable is returned when the configuration struct does not
// embed a `*Config` field.
var ErrNotConfigurable = errors.New("config: not configurable")

// Config is a struct that should be embedded in your applications's
// configuration struct.
type Config struct{}

// Base returns the `*Config` field of the configuration struct.
func (c *Config) Base() *Config {
	return c
}

// Configurable is an interface for types that embed a `*Config` field.
type Configurable interface {
	Base() *Config
}

// New returns a new configuration struct with the default values set from the
// provided `defaults` functions.
func New[C any](defaults ...func(Configurable) error) (*C, error) {
	config := new(C)

	configurable, ok := any(config).(Configurable)
	if !ok || configurable == nil {
		return nil, fmt.Errorf("%w: found %T", ErrNotConfigurable, config)
	}

	*configurable.Base() = Config{}

	for _, defaultFunc := range defaults {
		if err := defaultFunc(configurable); err != nil {
			return nil, fmt.Errorf("error setting default configuration: %w", err)
		}
	}

	return config, nil
}

// Load loads the configuration from the provided loaders. Each loader can override
// the configuration values set by the previous loaders.
func Load[C any](
	ctx context.Context,
	configurable Configurable,
	loaders ...func(context.Context, Configurable) error,
) error {
	for _, loader := range loaders {
		if err := loader(ctx, configurable); err != nil {
			return fmt.Errorf("error loading configuration: %w", err)
		}
	}

	return nil
}

// WithDefaults returns a function that sets the default configuration values
// for the configuration struct.
func WithDefaults[C any](defaults ...func(*C) error) func(Configurable) error {
	return func(configurable Configurable) error {
		config, ok := (any)(configurable).(*C)
		if !ok {
			return fmt.Errorf("%w: want %T found %T", ErrNotConfigurable, config, configurable)
		}

		for _, defaultFunc := range defaults {
			if err := defaultFunc(config); err != nil {
				return fmt.Errorf("error setting default configuration: %w", err)
			}
		}

		return nil
	}
}

// WithLoaders returns a function that loads the configuration from the
// loaders.
func WithLoaders[C any](
	loaders ...func(context.Context, *C) error,
) func(context.Context, Configurable) error {
	return func(ctx context.Context, configurable Configurable) error {
		config, ok := (any)(configurable).(*C)
		if !ok {
			return fmt.Errorf("%w: want %T found %T", ErrNotConfigurable, config, configurable)
		}

		for _, loader := range loaders {
			if err := loader(ctx, config); err != nil {
				return fmt.Errorf("error loading configuration: %w", err)
			}
		}

		return nil
	}
}
