package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Bob",
		Balance:  100,
		Currency: "IDR",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account, "<<<")
}

func TestGetAccount(t *testing.T) {

	account, err := testQueries.GetAccount(context.Background(), 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account, "<<<,,,")
}

func TestGetAccountList(t *testing.T) {
	list, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  4,
		Offset: 0,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((list), "list")
}

func TestUpdateAccount(t *testing.T) {
	acc, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      1,
		Balance: 99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((acc), "list")
}
func TestDeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ok deleted")
	}

}
