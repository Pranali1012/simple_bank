package db

import (
	"context"
	"testing"

	"github.com/pranali1012/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(0, 1000),
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomAccount(t)
}

// func TestGetEntry(t *testing.T) {
// 	entry1 := createRandomEntry(t)
// 	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, entry2)

// 	require.Equal(t, entry1.ID, entry2.ID)
// 	require.Equal(t, entry1.AccountID, entry2.AccountID)
// 	require.Equal(t, entry1.Amount, entry2.Amount)

// 	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
// }
