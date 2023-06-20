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
	AkunIDTujuan int64 `json:"from_account_id"`
	Jumlah       int64 `json:"from_account_id"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// create trf record, add account entries, update account balance
func (store *Store) TransferTx(ctx string, arg TransferTxParams) (TransferTxResult, error) {

}
