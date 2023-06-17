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

	// Initialize the db.Queries struct
	queries := db.New(dbConn)

	createAccountTest(ctx, queries)
}

func createAccountTest(ctx context.Context, queries *db.Queries) {
	params := db.CreateAccountParams{
		Owner:    "John Doe",
		Balance:  1000,
		Currency: "USD",
	}

	account, err := queries.CreateAccount(ctx, params)
	if err != nil {
		log.Println("failed to create account:", err)
		return
	}

	fmt.Println("Created account:", account)
}
