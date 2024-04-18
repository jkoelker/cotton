//

package auth_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"

	"github.com/jkoelker/cotton/auth"
)

func TestNewCacheToken(t *testing.T) {
	t.Parallel()

	access := gofakeit.UUID()
	refresh := gofakeit.UUID()
	expiry := time.Now().Add(1 * time.Hour)

	oauthToken := &oauth2.Token{
		AccessToken:  access,
		RefreshToken: refresh,
		Expiry:       expiry,
	}

	cacheToken := auth.NewCacheToken(oauthToken)
	require.NotNil(t, cacheToken)

	assert.Equal(t, access, cacheToken.AccessToken)
	assert.Equal(t, refresh, cacheToken.RefreshToken)
	assert.EqualValues(t, expiry, cacheToken.Expiry)
}

func TestSaveLoadToken(t *testing.T) {
	t.Parallel()

	token := auth.CacheToken{
		AccessToken:  gofakeit.UUID(),
		RefreshToken: gofakeit.UUID(),
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	var buffer bytes.Buffer

	err := token.SaveToken(&buffer)
	require.NoError(t, err)

	loaded, err := auth.LoadToken(&buffer)
	require.NoError(t, err)

	assert.Equal(t, token.AccessToken, loaded.AccessToken)
	assert.Equal(t, token.RefreshToken, loaded.RefreshToken)
	assert.WithinDuration(t, token.Expiry, loaded.Expiry, 1*time.Second)
}

func TestTokenSource_Token(t *testing.T) {
	t.Parallel()

	token := &oauth2.Token{
		AccessToken:  gofakeit.UUID(),
		RefreshToken: gofakeit.UUID(),
		Expiry:       time.Now().Add(1 * time.Hour),
	}

	temp, err := os.CreateTemp("", "cache")
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := os.Remove(temp.Name()); err != nil {
			t.Fatal(err)
		}
	})

	source, err := auth.NewTokenSource(
		oauth2.StaticTokenSource(token),
		auth.WithTokenCacheFile(temp.Name()),
	)
	require.NoError(t, err)

	next, err := source.Token()
	require.NoError(t, err)

	assert.Equal(t, token.AccessToken, next.AccessToken)
	assert.Equal(t, token.RefreshToken, next.RefreshToken)
	assert.EqualValues(t, token.Expiry, next.Expiry)

	content, err := os.ReadFile(temp.Name())
	require.NoError(t, err)
	assert.Contains(t, string(content), token.AccessToken)
	assert.Contains(t, string(content), token.RefreshToken)
	assert.Contains(t, string(content), token.Expiry.Format(time.RFC3339))
}
