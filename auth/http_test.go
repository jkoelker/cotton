//

package auth_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jkoelker/cotton/auth"
)

func TestGenerateSelfSignedCert(t *testing.T) {
	t.Parallel()

	cert, err := auth.GenerateSelfSignedCert()
	require.NoError(t, err)
	assert.NotEmpty(t, cert)
}

func TestWaitForCode(t *testing.T) {
	t.Parallel()

	state := gofakeit.UUID()
	code := gofakeit.UUID()

	// Find a free port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	err = listener.Close()
	require.NoError(t, err)

	results := make(chan struct {
		code string
		err  error
	})

	go func() {
		code, err := auth.WaitForCode(listener.Addr().String(), state)
		results <- struct {
			code string
			err  error
		}{code, err}
	}()

	// Give the server time to start
	time.Sleep(500 * time.Millisecond)

	client := &http.Client{
		Timeout: 1 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //nolint:gosec // This is a self signed cert.
			},
		},
	}

	serverURL := fmt.Sprintf(
		"https://%s/?state=%s&code=%s",
		listener.Addr().String(),
		url.QueryEscape(state),
		url.QueryEscape(code),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	t.Cleanup(cancel)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	select {
	case res := <-results:
		require.NoError(t, res.err)
		assert.Equal(t, code, res.code)

	case <-time.After(1 * time.Second):
		t.Fatal("Test timed out waiting for code")
	}
}
