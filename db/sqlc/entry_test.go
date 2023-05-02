package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ryanmiranda/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, arg CreateEntryParams) Entry {
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotZero(t, entry.ID)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.WithinDuration(t, entry.CreatedAt, time.Now(), time.Second)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	createRandomEntry(t, arg)
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry := createRandomEntry(t, arg)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry, err = testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.Empty(t, entry)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	createdEntry := createRandomEntry(t, arg)

	entry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.ID, createdEntry.ID)
	require.Equal(t, entry.AccountID, createdEntry.AccountID)
	require.Equal(t, entry.Amount, createdEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, createdEntry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry := createRandomEntry(t, arg)

	updateArg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	err := testQueries.UpdateEntry(context.Background(), updateArg)
	require.NoError(t, err)

	updatedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, entry.ID, updatedEntry.ID)
	require.Equal(t, entry.AccountID, updatedEntry.AccountID)
	require.NotEqual(t, entry.Amount, updatedEntry.Amount)
}

func TestListAllEntriesForAccount(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}

		testQueries.CreateEntry(context.Background(), arg)
	}

	arg := ListAllEntriesForAccountParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    2,
	}
	entries, err := testQueries.ListAllEntriesForAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Len(t, entries, 5)
}

func TestListAllEntries(t *testing.T) {
	for i := 0; i < 5; i++ {
		account := createRandomAccount(t)
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		}
		testQueries.CreateEntry(context.Background(), arg)
	}

	arg := ListAllEntriesParams{
		Limit:  4,
		Offset: 2,
	}
	entries, err := testQueries.ListAllEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)

	// TODO: Check how length is 4 even after offset
	require.Len(t, entries, 4)
}
