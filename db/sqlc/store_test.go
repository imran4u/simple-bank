package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	store := NewStore(testDb)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)

	// run n concurency transfer tx
	//FIXME: why n=5?
	n := 1

	// Notify can't work on other go routine, so use main gorountine with channel

	errorChan := make(chan error)
	resultChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errorChan <- err
			resultChan <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errorChan
		require.NoError(t, err)

		result := <-resultChan
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)

		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		//check tranfer exist
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check entry - from
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		//Get account entry from DB
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// Check entry - to
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		//Get account entry from DB
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//TODO: check account balance
	}

}
