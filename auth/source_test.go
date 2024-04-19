//

package auth_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"

	"github.com/jkoelker/cotton/auth"
)

func TestWithCacheFile(t *testing.T) {
	t.Parallel()

	testDir := t.TempDir()
	t.Cleanup(func() {
		err := os.RemoveAll(testDir)
		require.NoError(t, err)
	})

	subdir := gofakeit.UUID()
	testFile := filepath.Join(testDir, subdir, "token_cache.yaml")

	authInstance, err := auth.New(
		&oauth2.Config{},
		os.Stdout,
		auth.WithCacheFile(testFile),
	)
	require.NoError(t, err)

	assert.Equal(t, testFile, authInstance.File)

	_, err = os.Stat(filepath.Join(testDir, subdir))
	require.NoError(t, err)
}

func TestAuth_New(t *testing.T) {
	t.Parallel()

	config := &oauth2.Config{}
	output := new(bytes.Buffer)

	authInstance, err := auth.New(config, output)
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, config, authInstance.Config)
	assert.Equal(t, output, authInstance.Output)
}
