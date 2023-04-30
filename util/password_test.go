package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(6)
	hashedPassword1, err := HashPassword(password)

	require.NoError(t, err)

	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	wrongPassword := RandomString(8)
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.Error(t, err)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)

}
