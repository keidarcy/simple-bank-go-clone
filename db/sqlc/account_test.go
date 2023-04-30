package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/keidarcy/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	a1 := createRandomAccount(t)
	a2, err := testQueries.GetAccount(context.Background(), a1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a1.ID, a2.ID)
	require.Equal(t, a1.Owner, a2.Owner)
	require.Equal(t, a1.Balance, a2.Balance)
	require.Equal(t, a1.Currency, a2.Currency)

	require.WithinDuration(t, a1.CreatedAt, a2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	a1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      a1.ID,
		Balance: util.RandomMoney(),
	}

	a2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a1.ID, a2.ID)
	require.Equal(t, a1.Owner, a2.Owner)
	require.Equal(t, arg.Balance, a2.Balance)
	require.Equal(t, a1.Currency, a2.Currency)

	require.WithinDuration(t, a1.CreatedAt, a2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	a1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), a1.ID)
	require.NoError(t, err)

	a2, err := testQueries.GetAccount(context.Background(), a1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, a2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
