package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	acount1 := createRandAccount(t)
	acount2 := createRandAccount(t)
	fmt.Println(">> before :", acount1.Balance, acount2.Balance)

	n := 6
	amount := int64(1)
	errs := make(chan error)
	results := make(chan TransferTxResults)
	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: acount1.ID,
				ToAccountID:   acount2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, acount1.ID, transfer.FromAccountID)
		require.Equal(t, acount2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, acount1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, amount, toEntry.Amount)
		require.Equal(t, acount2.ID, toEntry.AccountID)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, acount1.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, acount2.ID)

		fmt.Println(">> tx:,", fromAccount.Balance, toAccount.Balance)
		diff1 := acount1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - acount2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)
		k := int(diff1 / amount)

		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updateFromAccount, err := store.GetAccount(context.Background(), acount1.ID)
	require.NoError(t, err)

	updateToAccount, err := store.GetAccount(context.Background(), acount2.ID)
	require.NoError(t, err)

	fmt.Println(t, updateFromAccount.Balance, updateToAccount.Balance)

	require.Equal(t, updateFromAccount.Balance, acount1.Balance-int64(n)*amount)
	require.Equal(t, updateToAccount.Balance, acount2.Balance+int64(n)*amount)
}

func TestTransferTxDeadLock(t *testing.T) {
	store := NewStore(testDB)

	acount1 := createRandAccount(t)
	acount2 := createRandAccount(t)
	fmt.Println(">> before :", acount1.Balance, acount2.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)
	for i := 0; i < n; i++ {
		fromAccountID := acount1.ID
		toAccountID := acount2.ID

		if i%2 == 1 {
			fromAccountID = acount2.ID
			toAccountID = acount1.ID
		}
		go func() {
			ctx := context.Background()
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	updateFromAccount, err := store.GetAccount(context.Background(), acount1.ID)
	require.NoError(t, err)

	updateToAccount, err := store.GetAccount(context.Background(), acount2.ID)
	require.NoError(t, err)

	fmt.Println(">> after :", updateFromAccount.Balance, updateToAccount.Balance)

	require.Equal(t, updateFromAccount.Balance, acount1.Balance)
	require.Equal(t, updateToAccount.Balance, acount2.Balance)
}
