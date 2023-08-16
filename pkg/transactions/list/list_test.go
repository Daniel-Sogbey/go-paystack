package listTransactions

import (
	"testing"

	"github.com/Daniel-Sogbey/go-paystack/client"
)

func TestListTransactions(t *testing.T) {
	client := &client.Client{
		Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
		ContentType:   "application/json",
	}

	listTransactionRequest := &ListTransactionRequest{
		Id: "3030558719",
	}

	response, err := ListTransactions(client, listTransactionRequest)

	if err != nil {
		t.Errorf("Test failed , go an error that says : %v", err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}

}
