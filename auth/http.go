//

package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"
)

const (
	defaultReadTimeout = 5 * time.Second
	defaultSerialBits  = 128

	oneDay = 24 * time.Hour
)

func GenerateSelfSignedCert() (tls.Certificate, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to generate private key: %w", err)
	}

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), defaultSerialBits))
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to generate serial number: %w", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Temporary Self-Signed"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(oneDay),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to create certificate: %w", err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	privBytes, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to marshal EC private key: %w", err)
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privBytes})

	pair, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to create key pair: %w", err)
	}

	return pair, nil
}

func WaitForCode(addr string, state string) (string, error) {
	cert, err := GenerateSelfSignedCert()
	if err != nil {
		return "", fmt.Errorf("failed to generate self-signed certificate: %w", err)
	}

	srv := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: defaultReadTimeout,
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
		},
	}

	codeChan := make(chan string)
	errChan := make(chan error)

	http.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		if query.Get("state") != state {
			http.Error(resp, "state does not match", http.StatusBadRequest)

			return
		}

		if _, err := resp.Write([]byte("Success! You may now close this window.")); err != nil {
			http.Error(resp, "failed to write response", http.StatusInternalServerError)
		}

		codeChan <- query.Get("code")
	})

	go func() {
		if err := srv.ListenAndServeTLS("", ""); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("failed to start HTTPS server: %w", err)
		}

		close(codeChan)
		close(errChan)
	}()

	select {
	case code := <-codeChan:
		if err := srv.Close(); err != nil {
			return code, fmt.Errorf("failed to close server: %w", err)
		}

		return code, nil

	case err := <-errChan:
		return "", fmt.Errorf("failed to get code: %w", err)
	}
}
