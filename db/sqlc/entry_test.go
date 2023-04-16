package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/keidarcy/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {

	arg := CreateEntryParams{
		AccountID: util.RandomInt(1, 10),
		Amount:    util.RandomInt(1, 1000),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)

}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	amount := util.RandomInt(1, 10000)

	updateArg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: amount,
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), updateArg)

	require.NoError(t, err)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, amount, entry2.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	err = testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, len(entries), 5)

}
