package db

import (
	"context"
	"testing"
	"time"

	"github.com/keidarcy/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          util.RandomOwner(),
		FullName:       util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreateAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	u1 := createRandomUser(t)
	u2, err := testQueries.GetUser(context.Background(), u1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, u2)

	require.Equal(t, u1.Username, u2.Username)
	require.Equal(t, u1.HashedPassword, u2.HashedPassword)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.FullName, u2.FullName)
	require.WithinDuration(t, u1.PasswordChangedAt, u2.PasswordChangedAt, time.Second)
}
