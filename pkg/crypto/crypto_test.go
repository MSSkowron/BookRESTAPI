package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCrypto(t *testing.T) {
	data := []struct {
		name          string
		inputPassword string
	}{
		{
			name:          "normal password",
			inputPassword: "password123",
		},
		{
			name:          "empty password",
			inputPassword: "",
		},
		{
			name:          "short password",
			inputPassword: "123",
		},
		{
			name:          "long password",
			inputPassword: "averylongpasswordthatexceedsthelimitof72charactersandshouldcausethe",
		},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.inputPassword)
			require.NoError(t, err)
			require.NotEmpty(t, hash)

			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(tt.inputPassword))
			require.NoError(t, err)

			err = CheckPassword(tt.inputPassword, string(hash))
			require.NoError(t, err)

			err = CheckPassword("wrongPassword123@", string(hash))
			require.ErrorIs(t, ErrInvalidCredentials, err)
		})
	}
}
