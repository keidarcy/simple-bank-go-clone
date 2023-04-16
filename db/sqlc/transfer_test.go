package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/keidarcy/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomInt(1, 1000),
	}
	Transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.FromAccountID, Transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, Transfer.ToAccountID)
	require.Equal(t, arg.Amount, Transfer.Amount)

	return Transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)

}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	amount := util.RandomInt(1, 10000)

	updateArg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: amount,
	}

	Transfer2, err := testQueries.UpdateTransfer(context.Background(), updateArg)

	require.NoError(t, err)
	require.Equal(t, transfer.FromAccountID, Transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, Transfer2.ToAccountID)
	require.Equal(t, amount, Transfer2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	err = testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, len(transfers), 5)

}
