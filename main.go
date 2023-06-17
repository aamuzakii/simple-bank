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

	db_string := "host=localhost port=5432 user=postgres password=secret dbname=knex_db sslmode=disable"

	db, err := sql.Open("postgres", db_string)
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	defer db.Close()

	ctx := context.Background()

	createAccountTest(ctx, db)

}

func createAccountTest(ctx context.Context, db *db.Queries) {
	params := db.CreateAccountParams{
		Owner:    "John Doe",
		Balance:  1000,
		Currency: "USD",
	}

	account, err := db.CreateAccount(ctx, params)
	if err != nil {
		log.Println("failed to create account:", err)
		return
	}

	fmt.Println("Created account:", account)
}
