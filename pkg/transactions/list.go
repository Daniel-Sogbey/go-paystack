package Transactions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/paystack-go-sdk/internal"
	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

type ListTransactionRequest struct {
	Id string `json:"id"`
}

type ListTransactionResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    TransactionData `json:"data"`
}

func List(c *paystack.Client, request *ListTransactionRequest) (ListTransactionResponse, error) {
	c.Client = &http.Client{}
	apiUrl := "https://api.paystack.co/transaction/" + request.Id

	req, _ := http.NewRequest("GET", apiUrl, nil)

	internal.SetHeaders(req, c.Authorization)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error from List Transaction Paystack API endpoint : %v", err)
	}

	responseBody, err := io.ReadAll(response.Body)

	fmt.Println("Response Body : ", string(responseBody))

	if err != nil {
		log.Fatalf("Error reading response body : %v", err)
	}

	var listTransactionResponse ListTransactionResponse

	err = json.Unmarshal(responseBody, &listTransactionResponse)

	if err != nil {
		log.Fatalf("Error unmarshalling response body : %v", err)
	}

	return ListTransactionResponse{
		Status:  listTransactionResponse.Status,
		Message: listTransactionResponse.Message,
		Data:    listTransactionResponse.Data,
	}, nil
}
