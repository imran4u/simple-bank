package db

import (
	"context"
	"database/sql"
	"fmt"
)

// to do all db operation and transaction( multiple operation with automicity)
type store struct {
	*Queries // composition to use as inharitance
	db       *sql.DB
}

// create New store object
func NewStore(db *sql.DB) *store {
	return &store{
		db:      db,
		Queries: New(db),
	}
}

func (s *store) execTx(context context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(context, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("tx error %v , rollback error %v", err, errRb)
		}
		return err
	}

	return tx.Commit()

}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		//TODO:

		// if arg.FromAccountID < arg.ToAccountID {
		// 	result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		// } else {
		// 	result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		// }

		return err
	})

	return result, err
}
