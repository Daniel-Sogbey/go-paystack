# Paystack Go SDK

## Go SKD for the Paystack API

### Install
> `go get github.com/Daniel-Sogbey/go-paystack`

### Quick Guide

```Go
package main

import (
	"fmt"
	"log"

	"github.com/Daniel-Sogbey/go-paystack/client"
	"github.com/Daniel-Sogbey/go-paystack/pkg/transactions/initialize"
	"github.com/Daniel-Sogbey/go-paystack/pkg/transactions/verify"
)

func main() {

	//SET UP AN HTTP CLIENT WITH AUTHORIZATION KEY AND CONTENT TYPE
	client := &client.Client{
		Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
		ContentType:   "application/json",
	}

	//SET UP INITIALIZE TRANSACTION REQUEST BODY
	initializeTransactionRequest := &initialize.InitializeTransactionRequest{
		Email:  "mathematics06physics@gmail.com",
		Amount: "2000",
	}

	// EXAMPLE OF AN INITIALIZE TRANSACTION REQUEST
	initializeTransactionResponse, err := initialize.InitializeTransaction(client, initializeTransactionRequest)

	if err != nil {
		log.Fatalf("Error making an initialize request to the package api %v", err)
	}

	// SAMPLE JSON RESPONSE FROM THE PAYSTACK INITIALIZE TRANSACTION API
	fmt.Println("JSON RESPONSE : ", initializeTransactionResponse)

	// SET UP VERIFICATION TRANSACTION REQUEST
	verifyTransactionRequest := &verify.VerificationRequest{
		Reference: "",
	}

	// EXAMPLE OF A VERIFY TRANSACTION REQUEST
	verifyTransactionResponse, _ := verify.VerifyTransaction(client, verifyTransactionRequest)

	// SAMPLE JSON RESPONSE FROM THE PAYSTACK INITIALIZE TRANSACTION API
	fmt.Println("JSON RESPONSE : ", verifyTransactionResponse)

}
```
