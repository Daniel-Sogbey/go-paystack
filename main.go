package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Daniel-Sogbey/paystack-go-sdk/internal"
	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
	Transactions "github.com/Daniel-Sogbey/paystack-go-sdk/pkg/transactions"
)

func main() {
	internal.LoadEnv()

	client := paystack.NewClient(os.Getenv("API_KEY"), "application/json")

	//SET UP INITIALIZE TRANSACTION REQUEST BODY
	initializeTransactionRequest := &Transactions.InitializeTransactionRequest{
		Email:  "mathematics06physics@gmail.com",
		Amount: "2000",
	}

	// EXAMPLE OF AN INITIALIZE TRANSACTION REQUEST
	initializeTransactionResponse, err := Transactions.Initialize(client, initializeTransactionRequest)

	if err != nil {
		log.Fatalf("Error making an initialize request to the package api %v", err)
	}

	// SAMPLE JSON RESPONSE FROM THE PAYSTACK INITIALIZE TRANSACTION API
	fmt.Println("JSON RESPONSE : ", initializeTransactionResponse)

	// SET UP VERIFICATION TRANSACTION REQUEST
	verifyTransactionRequest := &Transactions.VerificationRequest{
		Reference: "",
	}

	// EXAMPLE OF A VERIFY TRANSACTION REQUEST
	verifyTransactionResponse, _ := Transactions.Verify(client, verifyTransactionRequest)

	// SAMPLE JSON RESPONSE FROM THE PAYSTACK INITIALIZE TRANSACTION API
	fmt.Println("JSON RESPONSE : ", verifyTransactionResponse)

	fetchTransactionRequest := Transactions.FetchTransactionsResquest{
		Id: "3030558719",
	}

	fetchTransactionResponse, _ := Transactions.Fetch(client, &fetchTransactionRequest)

	fmt.Println("JSON RESPONSE : ", fetchTransactionResponse)
}
