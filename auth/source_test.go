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

	authInstance, err := auth.New(auth.WithOAuth2Config(config), auth.WithOutput(output))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, config, authInstance.Config)
	assert.Equal(t, output, authInstance.Output)
}

func TestWithOAuth2ClientID(t *testing.T) {
	t.Parallel()

	clientID := "test-client-id"

	authInstance, err := auth.New(auth.WithOAuth2ClientID(clientID))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, clientID, authInstance.Config.ClientID)
}

func TestWithOAuth2ClientSecret(t *testing.T) {
	t.Parallel()

	clientSecret := "test-client-secret"

	authInstance, err := auth.New(auth.WithOAuth2ClientSecret(clientSecret))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, clientSecret, authInstance.Config.ClientSecret)
}

func TestWithOAuth2Endpoint(t *testing.T) {
	t.Parallel()

	endpoint := oauth2.Endpoint{
		AuthURL:  "https://auth.example.com",
		TokenURL: "https://token.example.com",
	}

	authInstance, err := auth.New(auth.WithOAuth2Endpoint(endpoint))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, endpoint, authInstance.Config.Endpoint)
}

func TestWithOAuth2EndpointAuthURL(t *testing.T) {
	t.Parallel()

	authURL := "https://auth.example.com"

	authInstance, err := auth.New(auth.WithOAuth2EndpointAuthURL(authURL))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, authURL, authInstance.Config.Endpoint.AuthURL)
}

func TestWithOAuth2EndpointTokenURL(t *testing.T) {
	t.Parallel()

	url := "https://token.example.com"

	authInstance, err := auth.New(auth.WithOAuth2EndpointTokenURL(url))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, url, authInstance.Config.Endpoint.TokenURL)
}

func TestWithOAuth2RedirectURL(t *testing.T) {
	t.Parallel()

	redirectURL := "https://redirect.example.com"

	authInstance, err := auth.New(auth.WithOAuth2RedirectURL(redirectURL))
	require.NoError(t, err)
	require.NotNil(t, authInstance)

	assert.Equal(t, redirectURL, authInstance.Config.RedirectURL)
}

func TestEnvClientID(t *testing.T) {
	t.Run("CLIENT_ID", func(t *testing.T) {
		clientID := gofakeit.UUID()
		t.Setenv("CLIENT_ID", clientID)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientID, authInstance.Config.ClientID)
	})

	t.Run("SCHWAB_CLIENT_ID", func(t *testing.T) {
		clientID := gofakeit.UUID()
		t.Setenv("SCHWAB_CLIENT_ID", clientID)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientID, authInstance.Config.ClientID)
	})

	t.Run("COTTON_CLIENT_ID", func(t *testing.T) {
		clientID := gofakeit.UUID()
		t.Setenv("COTTON_CLIENT_ID", clientID)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientID, authInstance.Config.ClientID)
	})

	t.Run("COTTON_SCHWAB_CLIENT_ID", func(t *testing.T) {
		clientID := gofakeit.UUID()
		t.Setenv("COTTON_SCHWAB_CLIENT_ID", clientID)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientID, authInstance.Config.ClientID)
	})
}

func TestEnvClientSecret(t *testing.T) {
	t.Run("CLIENT_SECRET", func(t *testing.T) {
		clientSecret := gofakeit.UUID()
		t.Setenv("CLIENT_SECRET", clientSecret)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientSecret, authInstance.Config.ClientSecret)
	})

	t.Run("SCHWAB_CLIENT_SECRET", func(t *testing.T) {
		clientSecret := gofakeit.UUID()
		t.Setenv("SCHWAB_CLIENT_SECRET", clientSecret)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientSecret, authInstance.Config.ClientSecret)
	})

	t.Run("COTTON_CLIENT_SECRET", func(t *testing.T) {
		clientSecret := gofakeit.UUID()
		t.Setenv("COTTON_CLIENT_SECRET", clientSecret)

		t.Logf("COTTON_CLIENT_SECRET: %s", os.Getenv("COTTON_CLIENT_SECRET"))

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientSecret, authInstance.Config.ClientSecret)
	})

	t.Run("COTTON_SCHWAB_CLIENT_SECRET", func(t *testing.T) {
		clientSecret := gofakeit.UUID()
		t.Setenv("COTTON_SCHWAB_CLIENT_SECRET", clientSecret)

		authInstance, err := auth.New()
		require.NoError(t, err)
		require.NotNil(t, authInstance)

		assert.Equal(t, clientSecret, authInstance.Config.ClientSecret)
	})
}
