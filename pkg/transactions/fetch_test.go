package Transactions

import (
	"fmt"
	"os"
	"testing"

	"github.com/Daniel-Sogbey/paystack-go-sdk/internal"
	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

func TestFetch(t *testing.T) {

	internal.LoadEnv()

	client := paystack.NewClient(os.Getenv("API_KEY"), "application/json")

	fetchTransactionsResquest := FetchTransactionsResquest{
		Id: "3030558719",
	}

	response, err := Fetch(client, &fetchTransactionsResquest)

	if err != nil {
		t.Errorf("Test failed with error %v", err)
	}

	if response.Status == false {
		fmt.Println(response.Message)
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)

	}
}
