package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/jilse17/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) (Account, error) {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account, err
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	// create account
	account, err := createRandomAccount(t)

	// Get account
	searchAccount, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NotEmpty(t, searchAccount)
	require.NoError(t, err)

	require.Equal(t, account.ID, searchAccount.ID)
	require.Equal(t, account.Balance, searchAccount.Balance)
	require.Equal(t, account.Currency, searchAccount.Currency)
	require.Equal(t, account.Owner, searchAccount.Owner)
	require.Equal(t, account.CreatedAt, searchAccount.CreatedAt)
	require.NotZero(t, searchAccount.CreatedAt)

}

func TestUpdateAccount(t *testing.T) {
	// create account
	account, err := createRandomAccount(t)

	arg := UpdateAccountParams{
		Balance: 100,
		ID:      account.ID,
	}
	searchAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, searchAccount)
	require.Equal(t, account.ID, searchAccount.ID)
	require.NotEqual(t, account.Balance, searchAccount.Balance)
	fmt.Println(account.Balance, searchAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	// create account
	account, err := createRandomAccount(t)
	err = testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount, _ = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
