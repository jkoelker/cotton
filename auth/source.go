//

package auth

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

const cacheDirectoryMode = 0o700

// Auth wraps the oauth2.Config and provides a method to get the and persist
// the token / refresh token.
type Auth struct {
	Config *oauth2.Config
	File   string
	Output io.Writer
}

// WithCacheFile sets the file to cache the token.
func WithCacheFile(file string) func(*Auth) error {
	return func(auth *Auth) error {
		abs, err := filepath.Abs(file)
		if err != nil {
			return fmt.Errorf("failed to get absolute path: %w", err)
		}

		path := filepath.Dir(abs)
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(path, cacheDirectoryMode); err != nil {
				return fmt.Errorf("failed to create auth cache directory: %w", err)
			}
		}

		auth.File = abs

		return nil
	}
}

// New creates a new Auth instance.
func New(config *oauth2.Config, output io.Writer, opts ...func(*Auth) error) (*Auth, error) {
	auth := &Auth{
		Config: config,
		Output: output,
	}

	for _, opt := range opts {
		if err := opt(auth); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return auth, nil
}

// TokenSource returns the token source.
func (auth *Auth) TokenSource(ctx context.Context) (*TokenSource, error) {
	if auth.File == "" {
		return auth.initialToken(ctx)
	}

	file, err := os.Open(auth.File)
	if errors.Is(err, os.ErrNotExist) {
		return auth.initialToken(ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	cached, err := LoadToken(file)
	if err != nil {
		return nil, fmt.Errorf("failed to load token: %w", err)
	}

	token := cached.OAuth2()
	if !token.Valid() {
		return auth.refreshToken(ctx, token)
	}

	source, err := NewTokenSource(
		auth.Config.TokenSource(ctx, token),
		WithTokenCacheFile(auth.File),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create token source: %w", err)
	}

	return source, nil
}

// initialToken performs the oauth2 flow and saves the token.
func (auth *Auth) initialToken(ctx context.Context) (*TokenSource, error) {
	verifier := oauth2.GenerateVerifier()

	state, err := NewState()
	if err != nil {
		return nil, fmt.Errorf("failed to create state: %w", err)
	}

	url := auth.Config.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))

	if _, err := io.WriteString(
		auth.Output,
		fmt.Sprintf("Visit the URL for the auth dialog: %s\n", url),
	); err != nil {
		return nil, fmt.Errorf("failed to write to output: %w", err)
	}

	code, err := WaitForCode(":8080", state)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for code: %w", err)
	}

	token, err := auth.Config.Exchange(ctx, code, oauth2.VerifierOption(verifier))
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	source, err := NewTokenSource(
		auth.Config.TokenSource(ctx, token),
		WithTokenCacheFile(auth.File),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create token source: %w", err)
	}

	return source, nil
}

// refreshToken refreshes the token and if its successful, saves the token.
// Otherwise it runs the initialToken flow.
func (auth *Auth) refreshToken(ctx context.Context, token *oauth2.Token) (*TokenSource, error) {
	tokenSource := auth.Config.TokenSource(ctx, token)

	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	if !token.Valid() {
		return auth.initialToken(ctx)
	}

	source, err := NewTokenSource(
		tokenSource,
		WithTokenCacheFile(auth.File),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create token source: %w", err)
	}

	return source, nil
}
