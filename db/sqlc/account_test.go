package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"goBank/util"
	"goBank/util/test-helper"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount() (CreateAccountParams, Account, error) {
	arg := CreateAccountParams{
		Owner:    test_helper.RandomOwner(),
		Balance:  test_helper.RandomMoney(),
		Currency: Currency(test_helper.RandomCurrency()),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	return arg, account, err
}
func TestCreateAccount(t *testing.T) {
	arg, account, err := createRandomAccount()
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	_, account, err := createRandomAccount()
	require.NoError(t, err)
	createdAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NotEmpty(t, createdAccount)

	require.NoError(t, err)
	require.Equal(t, account.ID, createdAccount.ID)
	require.Equal(t, account.Owner, createdAccount.Owner)
	require.Equal(t, account.Balance, createdAccount.Balance)
	require.Equal(t, account.Currency, createdAccount.Currency)
	require.WithinDuration(t, util.SafeTime(account.CreatedAt), util.SafeTime(createdAccount.CreatedAt), time.Second)
}

func TestUpdateAccount(t *testing.T) {
	_, account, err := createRandomAccount()
	require.NoError(t, err)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: test_helper.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.WithinDuration(t, util.SafeTime(account.CreatedAt), util.SafeTime(updatedAccount.CreatedAt), time.Second)
}

func TestDeleteAccount(t *testing.T) {
	_, account, err := createRandomAccount()
	require.NoError(t, err)

	err = testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	dbAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, dbAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount()
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
