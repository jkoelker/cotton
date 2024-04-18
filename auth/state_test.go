//

package auth_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jkoelker/cotton/auth"
)

func TestNewState(t *testing.T) {
	t.Parallel()

	hexRegex := regexp.MustCompile("^[a-f0-9]+$")

	t.Run("Default length", func(t *testing.T) {
		t.Parallel()

		state, err := auth.NewState()
		require.NoError(t, err)

		assert.Len(t, state, 32)
		assert.Regexp(t, hexRegex, state)
	})

	t.Run("Custom length", func(t *testing.T) {
		t.Parallel()

		state, err := auth.NewState(20)
		require.NoError(t, err)

		assert.Len(t, state, 40)
		assert.Regexp(t, hexRegex, state)
	})

	t.Run("Negative length", func(t *testing.T) {
		t.Parallel()

		_, err := auth.NewState(-5)
		require.ErrorIs(t, err, auth.ErrInvalidLength)
	})

	t.Run("Zero length", func(t *testing.T) {
		t.Parallel()

		_, err := auth.NewState(0)
		require.ErrorIs(t, err, auth.ErrInvalidLength)
	})
}
