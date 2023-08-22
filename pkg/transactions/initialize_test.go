package Transactions

import (
	"os"
	"testing"

	"github.com/Daniel-Sogbey/paystack-go-sdk/internal"
	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

func TestInitialize(t *testing.T) {
	internal.LoadEnv()

	client := paystack.NewClient(os.Getenv("API_KEY"), "application/json")

	sampleInitializeTransactionRequest := &InitializeTransactionRequest{
		Email:  "1@2.com",
		Amount: "20",
	}

	response, err := Initialize(client, sampleInitializeTransactionRequest)

	if err != nil {
		t.Errorf("Expected response of status %v, but got a status of %v and an error that says %v", response.Status, response.Status, err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}

}
