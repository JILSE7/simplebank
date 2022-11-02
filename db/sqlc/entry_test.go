package db

import (
	"context"
	"testing"

	"github.com/jilse17/simplebank/utils"
	"github.com/stretchr/testify/require"
)

const ID = 17

func createRandomEntry(t *testing.T, id int64) Entry {
	arg := CreateEntryParams{
		AccountID: id,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func TestCreateEntry(t *testing.T) {
	// id := utils.RandomInt(0, 50)
	createRandomEntry(t, 14)

}

func TestGetEntry(t *testing.T) {

	// id := utils.RandomInt(0, 50)
	entry := createRandomEntry(t, 14)

	searchEntry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, searchEntry)

	require.Equal(t, entry.Amount, searchEntry.Amount)
	require.Equal(t, entry.ID, searchEntry.ID)
	require.Equal(t, entry.CreatedAt, searchEntry.CreatedAt)
	require.NotZero(t, searchEntry.ID)
	require.NotZero(t, searchEntry.CreatedAt)
}

func TestListEntries(t *testing.T) {
	account, _ := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account.ID)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
