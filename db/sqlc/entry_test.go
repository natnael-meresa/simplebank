package db

import (
	"context"
	"testing"
	"time"

	"github.com/natnael-meresa/simplbank/utils"
	"github.com/stretchr/testify/require"
)

func CreateEntry(t *testing.T) Entry {
	account := createRandAccount(t)
	arg := CreateEntryParams{
		Amount:    utils.RandomAmount(),
		AccountID: account.ID,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotEmpty(t, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := CreateEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry, entry2)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestUpdateEntrie(t *testing.T) {
	entry := CreateEntry(t)
	arg := UpdateEntrieParams{
		ID:     entry.ID,
		Amount: 80,
	}
	entry2, err := testQueries.UpdateEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, arg.Amount, entry2.Amount)
}
func TestDeleteEntrie(t *testing.T) {
	entry := CreateEntry(t)
	err := testQueries.DeleteEntrie(context.Background(), entry.ID)
	require.NoError(t, err)
}

func TestListentries(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateEntry(t)
	}

	arg := ListentriesParams{
		Limit:  5,
		Offset: 0,
	}

	entries, err := testQueries.Listentries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
}
