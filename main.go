package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	// Ping the database to test the connection
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("failed to ping the database:", err)
	}

	fmt.Println("Connected to the database!")

}
