package db

import (
	"context"
	"testing"
	"time"

	"github.com/natnael-meresa/simplbank/utils"
	"github.com/stretchr/testify/require"
)

func CreateTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		Amount:        utils.RandomAmount(),
		ToAccountID:   utils.RandomID(),
		FromAccountID: utils.RandomID(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotEmpty(t, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer, transfer2)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := CreateTransfer(t)
	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: 80,
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, arg.Amount, transfer2.Amount)

}

func TestDeleteTransfer(t *testing.T) {
	transfer := CreateTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
}
