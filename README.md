# Paystack Go SDK 🥤

## Robust Go SDK for the Paystack API

> A well tested and maintained Go sdk for the paystack api

### Install

```Go
go get github.com/Daniel-Sogbey/go-paystack

```

### Quick Guide

>

```Go
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


```

## Run Sample Test

```Go
go test
```

```GO
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



```

## TODOS

### Transactions

- [x] Initialize Transaction
- [x] Verify Transaction
- [x] List Transactions
- [x] Fetch Transaction
- [x] Charge Authorization
- [ ] View Transaction Timeline
- [ ] Transaction Totals
- [ ] Export Transactions
- [ ] Partial Debit

### Transaction Splits

- [ ] Create Split
- [ ] List/Search Splits
- [ ] Fetch Split
- [ ] Update Split
- [ ] Add/Update Split
- [ ] Subaccount
- [ ] Remove Subaccount from Split

### Refunds

- [ ] Create Refund
- [ ] List Refunds
- [ ] Fetch Refunds

## CONTRIBUTIONS

> Contributions are very welcomed from the community🙏
