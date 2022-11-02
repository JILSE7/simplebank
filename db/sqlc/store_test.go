package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/* func TestTranferTx(t *testing.T) {
	store := NewStore(testDB)

	account1, _ := createRandomAccount(t)
	account2, _ := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// RUN N CONCURRENT TRANSFER TRANSACTIONS
	n := 2
	amount := int64(10)

	errors := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(ctx, TranferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errors <- err
			results <- result

		}()
	}

	// check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errors
		require.NoError(t, err)

		// check transfer
		result := <-results

		require.Equal(t, result.Transfer.FromAccountID, account1.ID)
		require.Equal(t, result.Transfer.ToAccountID, account2.ID)
		require.Equal(t, result.Transfer.Amount, amount)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		//check entry
		require.Equal(t, result.FromEntry.AccountID, account1.ID)
		require.Equal(t, result.FromEntry.Amount, -amount)
		require.NotZero(t, result.FromEntry.ID)
		require.NotZero(t, result.FromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)

		require.Equal(t, result.ToEntry.AccountID, account2.ID)
		require.Equal(t, result.ToEntry.Amount, amount)
		require.NotZero(t, result.ToEntry.ID)
		require.NotZero(t, result.ToEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)

		// Check accounts
		fmt.Println(result.FromAccount.ID, "dfj")
		require.Equal(t, result.FromAccount.ID, account1.ID)
		require.Equal(t, result.ToAccount.ID, account2.ID)

		// TODO: CHECK BALANCE
		fmt.Println(">> tx:", result.FromAccount.Balance, result.ToAccount.Balance)
		diff1 := account1.Balance - result.FromAccount.Balance
		diff2 := result.ToAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		// require.True(t, k >= 1 && k <= n)
		//require.NotContains(t, existed, k)
		existed[k] = true

	}

	// check the final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

} */

func TestTranferTxDeadLook(t *testing.T) {
	store := NewStore(testDB)

	account1, _ := createRandomAccount(t)
	account2, _ := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// RUN N CONCURRENT TRANSFER TRANSACTIONS
	n := 10
	amount := int64(10)

	errors := make(chan error)

	for i := 0; i < n; i++ {
		FromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			FromAccountID = account2.ID
			toAccountID = account1.ID
		}
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			_, err := store.TransferTx(ctx, TranferTxParams{
				FromAccountID: FromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errors <- err

		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errors
		require.NoError(t, err)
	}

	// check the final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
