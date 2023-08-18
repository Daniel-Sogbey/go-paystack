package Transactions

import (
	"fmt"
	"testing"

	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

func TestFetch(t *testing.T) {
	client := paystack.NewClient("sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb", "application/json")

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
