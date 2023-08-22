package Transactions

import (
	"os"
	"testing"

	"github.com/Daniel-Sogbey/paystack-go-sdk/internal"
	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

func TestVerify(t *testing.T) {
	internal.LoadEnv()

	client := paystack.NewClient(os.Getenv("API_KEY"), "application/json")

	sampleVerificationRequest := &VerificationRequest{
		Reference: "c2z7k6t1i4",
	}

	response, err := Verify(client, sampleVerificationRequest)

	if err != nil {
		t.Errorf("Test failed with error %v", err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}
}
