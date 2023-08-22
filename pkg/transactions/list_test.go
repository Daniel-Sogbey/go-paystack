package Transactions

import (
	"testing"

	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

func TestListTransactions(t *testing.T) {
	client := paystack.NewClient("sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb", "application/json")

	listTransactionRequest := &ListTransactionRequest{
		Id: "3030558719",
	}

	response, err := List(client, listTransactionRequest)

	if err != nil {
		t.Errorf("Test failed , go an error that says : %v", err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}

}
