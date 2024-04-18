//

package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/oauth2"
	"sigs.k8s.io/yaml"
)

const (
	cacheFileMode  = 0o600
	cacheFileFlags = os.O_RDWR | os.O_CREATE | os.O_TRUNC
)

// CacheToken represents a token that is stored in the cache.
type CacheToken struct {
	// AccessToken is the current access token.
	AccessToken string `json:"access_token,omitempty"`

	// RefreshToken is used to get a new access token.
	RefreshToken string `json:"refresh_token,omitempty"`

	// Expiry is the expiry time of the access token.
	Expiry time.Time `json:"expiry,omitempty"`
}

// NewCacheToken creates a new token from an oauth2.Token.
func NewCacheToken(t *oauth2.Token) *CacheToken {
	return &CacheToken{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
	}
}

// Compare returns true if the tokens are equal.
func (token *CacheToken) Compare(other *CacheToken) bool {
	return token.AccessToken == other.AccessToken &&
		token.RefreshToken == other.RefreshToken &&
		token.Expiry.UTC().Equal(other.Expiry.UTC())
}

// IsOauth2Token returns true if the token is the same as the oauth2.Token.
func (token *CacheToken) IsOauth2Token(other *oauth2.Token) bool {
	return token.AccessToken == other.AccessToken &&
		token.RefreshToken == other.RefreshToken &&
		token.Expiry.UTC().Equal(other.Expiry.UTC())
}

// MarshalJSON marshals the token to JSON. Saving the expiry time as a UTC string.
func (token *CacheToken) MarshalJSON() ([]byte, error) {
	type Alias CacheToken

	value, err := json.Marshal(&struct {
		Expiry string `json:"expiry,omitempty"`
		*Alias
	}{
		Expiry: token.Expiry.UTC().Format(time.RFC3339),
		Alias:  (*Alias)(token),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal token: %w", err)
	}

	return value, nil
}

// UnmarshalJSON unmarshals the token from JSON. Loading the expiry time as a UTC string.
func (token *CacheToken) UnmarshalJSON(data []byte) error {
	type Alias CacheToken

	aux := &struct {
		Expiry string `json:"expiry,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(token),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal token: %w", err)
	}

	expiry, err := time.Parse(time.RFC3339, aux.Expiry)
	if err != nil {
		return fmt.Errorf("failed to parse expiry: %w", err)
	}

	token.Expiry = expiry

	return nil
}

// OAuth2 converts a CacheToken to an oauth2.Token.
func (token *CacheToken) OAuth2() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
}

// SaveToken saves a token to a writer.
func (token *CacheToken) SaveToken(writer io.Writer) error {
	data, err := yaml.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	if _, err := writer.Write(data); err != nil {
		return fmt.Errorf("failed to write token: %w", err)
	}

	return nil
}

// LoadToken loads a token from a the reader.
func LoadToken(r io.Reader) (*CacheToken, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	var token CacheToken
	if err := yaml.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}

	return &token, nil
}

// TokenSource wraps an oauth2.TokenSource and optionally caches the token
// to a file.
type TokenSource struct {
	Source oauth2.TokenSource
	Last   *oauth2.Token
	File   string
	mu     sync.Mutex
}

var _ oauth2.TokenSource = (*TokenSource)(nil)

func WithTokenCacheFile(file string) func(*TokenSource) error {
	return func(source *TokenSource) error {
		if file != "" {
			source.File = file
		}

		return nil
	}
}

// NewTokenSource creates a new TokenSource.
func NewTokenSource(
	source oauth2.TokenSource,
	options ...func(*TokenSource) error,
) (*TokenSource, error) {
	tokenSource := &TokenSource{
		Source: source,
	}

	for _, option := range options {
		if err := option(tokenSource); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	if tokenSource.File != "" {
		if _, err := tokenSource.Token(); err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
	}

	return tokenSource, nil
}

// Token returns a token from the source.
func (source *TokenSource) Token() (*oauth2.Token, error) {
	source.mu.Lock()
	defer source.mu.Unlock()

	if source.Last != nil && source.Last.Valid() {
		return source.Last, nil
	}

	token, err := source.Source.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	if err := source.save(token); err != nil {
		return nil, fmt.Errorf("failed to save token: %w", err)
	}

	source.Last = token

	return token, nil
}

// Save forces a save of the last token to the file.
func (source *TokenSource) Save() error {
	if source.Last == nil {
		return nil
	}

	source.mu.Lock()
	defer source.mu.Unlock()

	if err := source.save(source.Last); err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	return nil
}

// save saves the token to the file.
func (source *TokenSource) save(token *oauth2.Token) error {
	if source.File == "" {
		return nil
	}

	cache := NewCacheToken(token)

	file, err := os.OpenFile(source.File, cacheFileFlags, cacheFileMode)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	if err := cache.SaveToken(file); err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	return nil
}
