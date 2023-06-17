package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   3,
		Amount:        99,
	}

	account, err := testQueries.CreateTransfer(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account, "<<<")
}

func TestGetTransfer(t *testing.T) {

	account, err := testQueries.GetTransfer(context.Background(), 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account, "<<<,,,")
}

func TestGetTransferList(t *testing.T) {
	list, err := testQueries.ListTransfers(context.Background(), ListTransfersParams{
		Limit:  4,
		Offset: 0,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((list), "list")
}

func TestUpdateTransfer(t *testing.T) {
	acc, err := testQueries.UpdateTransfer(context.Background(), UpdateTransferParams{
		ID:     1,
		Amount: 30,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((acc), "list")
}
func TestDeleteTransfer(t *testing.T) {
	err := testQueries.DeleteTransfer(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ok deleted")
	}

}
