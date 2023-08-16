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

## Run Sample Test

```Go
go test
```

```GO
package initialize

import (
	"testing"
	"os"
	"github.com/Daniel-Sogbey/go-paystack/client"
)

func TestInitializeTransactions(t *testing.T) {

	client := &client.Client{
		Authorization: os.Getenv("PAYSTACK_API_KEY"),
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

```


## TODOS

### Transactions
- [x] Initialize Transaction
- [x] Verify Transaction
- [ ] List Transactions
- [ ] Fetch Transaction
- [ ] Charge Authorization
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
