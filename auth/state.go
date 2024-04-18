//

package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

var ErrInvalidLength = errors.New("invalid length")

// NewState generates a new random state string.
func NewState(length ...int) (string, error) {
	if len(length) == 0 {
		length = []int{16}
	}

	if length[0] <= 0 {
		return "", fmt.Errorf("%w: %d", ErrInvalidLength, length[0])
	}

	buff := make([]byte, length[0])

	if _, err := rand.Read(buff); err != nil {
		return "", fmt.Errorf("failed to generate random state: %w", err)
	}

	return hex.EncodeToString(buff), nil
}
