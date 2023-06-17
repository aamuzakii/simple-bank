package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestCreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 3,
		Amount:    -750,
	}

	account, err := testQueries.CreateEntry(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account, "<<<")
}

func TestGetEntry(t *testing.T) {

	account, err := testQueries.GetEntry(context.Background(), 4)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account, "<<<,,,")
}

func TestGetEntryList(t *testing.T) {
	list, err := testQueries.ListEntrys(context.Background(), ListEntrysParams{
		Limit:  4,
		Offset: 0,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((list), "list")
}

func TestUpdateEntry(t *testing.T) {
	acc, err := testQueries.UpdateEntry(context.Background(), UpdateEntryParams{
		ID:     3,
		Amount: 30,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((acc), "list")
}
func TestDeleteEntry(t *testing.T) {
	err := testQueries.DeleteEntry(context.Background(), 3)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ok deleted")
	}

}
