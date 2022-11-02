package db

import (
	"context"
	"testing"

	"github.com/JILSE7/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: utils.RandomInt(1, 10),
		ToAccountID:   utils.RandomInt(1, 10),
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	return transfer
}

func createRandomTransferById(t *testing.T, fromId, toId int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromId,
		ToAccountID:   toId,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	searchTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NotEmpty(t, transfer)
	require.NoError(t, err)
	require.Equal(t, transfer.Amount, searchTransfer.Amount)
	require.Equal(t, transfer.FromAccountID, searchTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, searchTransfer.ToAccountID)
	require.Equal(t, transfer.CreatedAt, searchTransfer.CreatedAt)
}

func TestListTranfers(t *testing.T) {
	fromId := utils.RandomInt(4, 5)
	toId := utils.RandomInt(5, 10)
	for i := 0; i < 10; i++ {
		createRandomTransferById(t, fromId, toId)
	}

	arg := ListTransfersParams{
		FromAccountID: fromId,
		ToAccountID:   toId,
		Limit:         3,
		Offset:        3,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	}
}
