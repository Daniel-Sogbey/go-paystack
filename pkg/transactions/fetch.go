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

type FetchTransactionsResquest struct {
	Id string
}

type Metadata struct {
	CustomFields []CustomField `json:"custom_fields"`
	Referrer     string        `json:"referrer"`
}

type FetchTransactionsResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    TransactionData `json:"data"`
}

func Fetch(c *paystack.Client, request *FetchTransactionsResquest) (FetchTransactionsResponse, error) {
	c.Client = &http.Client{}

	apiUrl := "https://api.paystack.co/transaction/" + request.Id

	req, _ := http.NewRequest("GET", apiUrl, nil)

	internal.SetHeaders(req, c.Authorization)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error attempting to make request to paystack API : %v", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	fmt.Println("RESPONSE BODY : ", string(responseBody))

	var fetchTransactionsJson FetchTransactionsResponse

	err = json.Unmarshal(responseBody, &fetchTransactionsJson)

	if err != nil {
		log.Fatalf("Error attempting to unmarshall response json : %v", err)
	}

	return FetchTransactionsResponse{
		Status:  fetchTransactionsJson.Status,
		Message: fetchTransactionsJson.Message,
		Data:    fetchTransactionsJson.Data,
	}, nil

}
