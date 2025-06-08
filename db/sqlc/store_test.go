package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	_, account1, _ := createRandomAccount()
	_, account2, _ := createRandomAccount()

	// run n concurrent transfer transactions
	n := 5
	amount := float64(100)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)
		require.NotZero(t, transfer.ID)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.CreatedAt)
		require.NotZero(t, fromEntry.ID)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.CreatedAt)
		require.NotZero(t, toEntry.ID)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
		//check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account1.ID, toAccount.ID)
		//check accounts balances
		diff := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff, diff2)
		require.True(t, diff >= amount)
	}
	// check final balances
	account1After, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2After, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, account1.Balance-float64(n)*amount, account1After.Balance)
	require.Equal(t, account2.Balance+float64(n)*amount, account2After.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	_, account1, _ := createRandomAccount()
	_, account2, _ := createRandomAccount()

	// run n concurrent transfer transactions
	n := 10
	amount := float64(100)

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID
		if i%2 == 0 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
	// check final balances
	account1After, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2After, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, account1.Balance, account1After.Balance)
	require.Equal(t, account2.Balance, account2After.Balance)
}
