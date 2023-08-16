package initialize

import (
	"testing"

	"github.com/Daniel-Sogbey/go-paystack/paystack"
)

func TestInitializeTransactions(t *testing.T) {

	client := &paystack.Client{
		Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
		ContentType:   "application/json",
	}

	sampleInitializeTransactionRequest := &InitializeTransactionRequest{
		Email:  "1@2.com",
		Amount: "20",
	}

	response, err := InitializeTransaction(client, sampleInitializeTransactionRequest)

	if err != nil {
		t.Errorf("Expected response of status %v, but got a status of %v and an error that says %v", response.Status, response.Status, err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}

}
