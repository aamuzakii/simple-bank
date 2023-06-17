package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/aamuzakii/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	g := gin.Recovery()
	fmt.Println(g)

	dbString := "host=localhost port=5432 user=postgres password=secret dbname=knex_db sslmode=disable"

	dbConn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	defer dbConn.Close()

	ctx := context.Background()

	createAccountTest(ctx, dbConn)
}

func createAccountTest(ctx context.Context, dbConn *sql.DB) {
	tx, err := dbConn.Begin()
	if err != nil {
		log.Fatal("failed to begin transaction:", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-throw panic after rollback
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Fatal("failed to commit transaction:", err)
			}
		}
	}()

	var account db.Account
	err = tx.QueryRowContext(ctx, "SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = $1 LIMIT 1", 1).Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	if err != nil {
		log.Println("failed to retrieve account:", err)
		return
	}

	fmt.Println("Retrieved account:", account)

	// Other database operations within the transaction can be performed here
	// ...

	// If everything is successful, the transaction will be committed at the end
}
