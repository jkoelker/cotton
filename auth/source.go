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

// WithOutput sets the output writer.
func WithOutput(output io.Writer) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Output = output

		return nil
	}
}

// WithOAuth2Config sets the oauth2.Config.
func WithOAuth2Config(config *oauth2.Config) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config = config

		return nil
	}
}

// WithOAuth2ClientID sets the client ID.
func WithOAuth2ClientID(id string) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config.ClientID = id

		return nil
	}
}

// WithOAuth2ClientSecret sets the client secret.
func WithOAuth2ClientSecret(secret string) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config.ClientSecret = secret

		return nil
	}
}

// WithOAuth2Endpoint sets the endpoint.
func WithOAuth2Endpoint(endpoint oauth2.Endpoint) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config.Endpoint = endpoint

		return nil
	}
}

// WithOAuth2EndpointAuthURL sets the auth URL.
func WithOAuth2EndpointAuthURL(url string) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config.Endpoint.AuthURL = url

		return nil
	}
}

// WithOAuth2EndpointTokenURL sets the token URL.
func WithOAuth2EndpointTokenURL(url string) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config.Endpoint.TokenURL = url

		return nil
	}
}

// WithOAuth2RedirectURL sets the redirect URL.
func WithOAuth2RedirectURL(url string) func(*Auth) error {
	return func(auth *Auth) error {
		auth.Config.RedirectURL = url

		return nil
	}
}

func getEnv(key string, prefixes ...string) string {
	if len(prefixes) == 0 {
		prefixes = []string{"COTTON_SCHWAB_", "COTTON_", "SCHWAB_"}
	}

	for _, prefix := range prefixes {
		if value, ok := os.LookupEnv(prefix + key); ok {
			return value
		}
	}

	return os.Getenv(key)
}

// New creates a new Auth instance. Use `With*` functions to set the options.
// By default it will use the `CLIENT_ID` and `CLIENT_SECRET` (optionally
// prefixed with `COTTON_SCHWAB_`, `COTTON_`, or `SCHWAB_`) environment
// variables.
func New(opts ...func(*Auth) error) (*Auth, error) {
	clientID := getEnv("CLIENT_ID")
	clientSecret := getEnv("CLIENT_SECRET")

	auth := &Auth{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,

			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://api.schwabapi.com/v1/oauth/authorize",
				TokenURL: "https://api.schwabapi.com/v1/oauth/token",
			},

			RedirectURL: "http://127.0.0.1:8080",
		},

		Output: os.Stdout,
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
