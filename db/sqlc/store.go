package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provide all function to execute db query & transaction
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// nama function dibawah ini pake huruf kecil jadi ga dieksport, hanya bisa dipanggil internal package
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err %v, rb err %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	AkunIDSumber int64 `json:"from_account_id"`
	AkunIDTujuan int64 `json:"to_account_id"`
	Jumlah       int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// var txKey = struct{}{} // it its empty struct & second bracket means we create empty object

// create trf record, add account entries, update account balance
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.AkunIDSumber,
			ToAccountID:   arg.AkunIDTujuan,
			Amount:        arg.Jumlah,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.AkunIDSumber,
			Amount:    -arg.Jumlah,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.AkunIDTujuan, // Corrected parameter
			Amount:    arg.Jumlah,       // Corrected sign
		})

		if err != nil {
			return err
		}

		account1, err := q.GetAccount(ctx, arg.AkunIDSumber)

		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.AkunIDSumber,
			Balance: account1.Balance - arg.Jumlah,
		})

		if err != nil {
			return err
		}

		account2, err := q.GetAccount(ctx, arg.AkunIDTujuan)

		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.AkunIDTujuan,
			Balance: account2.Balance - arg.Jumlah,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
