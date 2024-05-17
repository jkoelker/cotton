//

package config_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jkoelker/cotton/config"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		type AppConfig struct {
			config.Config
			AppName string
		}

		defaultAppName := gofakeit.Word()

		appConfig, err := config.New[AppConfig](
			config.WithDefaults(
				func(cfg *AppConfig) error {
					cfg.AppName = defaultAppName

					return nil
				},
			),
		)
		require.NoError(t, err)
		assert.Equal(t, defaultAppName, appConfig.AppName)
	})

	t.Run("DefaultsError", func(t *testing.T) {
		t.Parallel()

		type AppConfig struct {
			config.Config
			AppName string
		}

		_, err := config.New[AppConfig](
			config.WithDefaults(
				func(_ *AppConfig) error {
					return assert.AnError
				},
			),
		)
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("NotConfigurable", func(t *testing.T) {
		t.Parallel()

		type SimpleConfig struct {
			AppName string
		}

		_, err := config.New[SimpleConfig]()
		assert.ErrorIs(t, err, config.ErrNotConfigurable)
	})
}

func TestLoad(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		type AppConfig struct {
			config.Config
			AppName string
		}

		appConfig := &AppConfig{}

		loader := func(_ context.Context, cfg *AppConfig) error {
			cfg.AppName = gofakeit.Word()

			return nil
		}

		err := config.Load[AppConfig](context.Background(), appConfig, config.WithLoaders(loader))
		require.NoError(t, err)

		assert.NotEmpty(t, appConfig.AppName)
	})

	t.Run("Failure", func(t *testing.T) {
		t.Parallel()

		type AppConfig struct {
			config.Config
			AppName string
		}

		appConfig := &AppConfig{}

		loader := func(_ context.Context, _ *AppConfig) error {
			return assert.AnError
		}

		err := config.Load[AppConfig](context.Background(), appConfig, config.WithLoaders(loader))
		assert.ErrorIs(t, err, assert.AnError)
	})
}
