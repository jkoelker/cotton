//

package config_test

import (
	"context"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jkoelker/cotton/config"
)

func TestUConfigYamlFile(t *testing.T) {
	t.Parallel()

	path := "/tmp/" + gofakeit.Word() + ".yaml"

	result := config.UConfigYamlFile(path)

	assert.NotNil(t, result.Unmarshal, "Unmarshal function should not be nil")
	assert.Equal(t, path, result.Path, "Path should match the input path")
	assert.True(t, result.Optional, "Optional should be true")
}

//nolint:paralleltest // uconfig does not support parallel tests.
func TestLoadUConfig(t *testing.T) {
	// UConfig `Classic` uses `flag` for parsing so we need to reset the args
	// to avoid conflicts with go test flags.
	origArgs := os.Args
	os.Args = os.Args[:1]

	t.Cleanup(func() {
		os.Args = origArgs
	})

	t.Run("Yaml", func(t *testing.T) {
		type TestConfig struct {
			Name string `yaml:"name"`
			Age  int    `default:"30" yaml:"age"`
		}

		tmpFile, err := os.CreateTemp("", "*.yaml")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		content := "name: John Doe\naddress: 123 Elm St\n"

		_, err = tmpFile.WriteString(content)
		require.NoError(t, err)
		require.NoError(t, tmpFile.Close())

		ctx := context.Background()

		var testConfig TestConfig

		require.NoError(
			t,
			config.LoadUConfig[TestConfig](tmpFile.Name())(ctx, &testConfig),
		)

		assert.Equal(t, "John Doe", testConfig.Name)
		assert.Equal(t, 30, testConfig.Age)
	})

	t.Run("File Not Found", func(t *testing.T) {
		type TestConfig struct {
			Name string `yaml:"name"`
			Age  int    `default:"30" yaml:"age"`
		}

		ctx := context.Background()

		var testConfig TestConfig

		require.NoError(
			t,
			config.LoadUConfig[TestConfig]("does-not-exist.yaml")(ctx, &testConfig),
		)

		assert.Equal(t, "", testConfig.Name)
		assert.Equal(t, 30, testConfig.Age)
	})

	t.Run("Environment Variables", func(t *testing.T) {
		type TestConfig struct {
			Name string `yaml:"name"`
			Age  int    `default:"30" yaml:"age"`
		}

		ctx := context.Background()

		t.Setenv("NAME", "Jane Doe")
		t.Setenv("AGE", "40")

		var testConfig TestConfig

		require.NoError(
			t,
			config.LoadUConfig[TestConfig]()(ctx, &testConfig),
		)

		assert.Equal(t, "Jane Doe", testConfig.Name)
		assert.Equal(t, 40, testConfig.Age)
	})

	t.Run("Flag", func(t *testing.T) {
		type TestConfig struct {
			Name string `yaml:"name"`
			Age  int    `default:"30" yaml:"age"`
		}

		ctx := context.Background()

		orgArgs := os.Args

		t.Cleanup(func() {
			os.Args = orgArgs
		})

		os.Args = append(os.Args[:1], "-name", "John Doe", "-age", "50")

		var testConfig TestConfig

		require.NoError(
			t,
			config.LoadUConfig[TestConfig]()(ctx, &testConfig),
		)

		assert.Equal(t, "John Doe", testConfig.Name)
		assert.Equal(t, 50, testConfig.Age)
	})

	t.Run("Flag Override", func(t *testing.T) {
		type TestConfig struct {
			Name string `yaml:"name"`
			Age  int    `default:"30" yaml:"age"`
		}

		tmpFile, err := os.CreateTemp("", "*.yaml")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		content := "name: Jane Doe\nage: 40\n"

		_, err = tmpFile.WriteString(content)
		require.NoError(t, err)
		require.NoError(t, tmpFile.Close())

		t.Setenv("AGE", "45")

		ctx := context.Background()

		orgArgs := os.Args

		t.Cleanup(func() {
			os.Args = orgArgs
		})

		os.Args = append(os.Args[:1], "-name", "John Doe")

		var testConfig TestConfig

		require.NoError(
			t,
			config.LoadUConfig[TestConfig](tmpFile.Name())(ctx, &testConfig),
		)

		assert.Equal(t, "John Doe", testConfig.Name)
		assert.Equal(t, 45, testConfig.Age)
	})

	t.Run("WithLoaders", func(t *testing.T) {
		type TestConfig struct {
			config.Config
			Name string `yaml:"name"`
			Age  int    `default:"30" yaml:"age"`
		}

		tmpFile, err := os.CreateTemp("", "*.yaml")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		content := "name: Jane Doe\nage: 40\n"

		_, err = tmpFile.WriteString(content)
		require.NoError(t, err)
		require.NoError(t, tmpFile.Close())

		ctx := context.Background()

		var testConfig TestConfig

		require.NoError(
			t,
			config.Load[TestConfig](
				ctx,
				&testConfig,
				config.WithLoaders(config.LoadUConfig[TestConfig](tmpFile.Name())),
			),
		)

		assert.Equal(t, "Jane Doe", testConfig.Name)
		assert.Equal(t, 40, testConfig.Age)
	})
}
