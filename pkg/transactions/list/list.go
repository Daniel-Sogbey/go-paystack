package listTransactions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/go-paystack/client"
	"github.com/Daniel-Sogbey/go-paystack/internal"
)

type ListTransactionRequest struct {
	Id string `json:"id"`
}

type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
}

type LogHistory struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type Log struct {
	StartTime int          `json:"start_time"`
	TimeSpent int          `json:"time_spent"`
	Attempts  int          `json:"attempts"`
	Errors    int          `json:"errors"`
	Success   bool         `json:"success"`
	Mobile    bool         `json:"mobile"`
	Input     []string     `json:"input"`
	History   []LogHistory `json:"history"`
}

type Authorization struct {
	AuthorizationCode string `json:"authorization_code"`
	Bin               string `json:"bin"`
	Last4             string `json:"last4"`
	ExpMonth          string `json:"exp_month"`
	ExpYear           string `json:"exp_year"`
	Channel           string `json:"channel"`
	CardType          string `json:"card_type"`
	Bank              string `json:"bank"`
	CountryCode       string `json:"country_code"`
	Brand             string `json:"brand"`
	Reusable          bool   `json:"reusable"`
	Signature         string `json:"signature"`
	AccountName       string `json:"account_name"`
}

type Customer struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CustomerCode string `json:"customer_code"`
	Phone        string `json:"phone"`
	Metadata     string `json:"metadata"`
	RiskAction   string `json:"risk_action"`
}

type Metadata struct {
	CustomFields []CustomField `json:"custom_fields"`
	Referrer     string        `json:"referrer"`
}

type TransactionData struct {
	ID              int           `json:"id"`
	Domain          string        `json:"domain"`
	Status          string        `json:"status"`
	Reference       string        `json:"reference"`
	Amount          int           `json:"amount"`
	Message         interface{}   `json:"message"`
	GatewayResponse string        `json:"gateway_response"`
	PaidAt          string        `json:"paid_at"`
	CreatedAt       string        `json:"created_at"`
	Channel         string        `json:"channel"`
	Currency        string        `json:"currency"`
	IPAddress       string        `json:"ip_address"`
	Metadata        string        `json:"metadata"` // Use the Metadata struct
	Log             Log           `json:"log"`
	Fees            int           `json:"fees"`
	Authorization   Authorization `json:"authorization"`
	Customer        Customer      `json:"customer"`
	Plan            interface{}   `json:"plan"`
	Subaccount      interface{}   `json:"subaccount"`
	OrderID         interface{}   `json:"order_id"`
	RequestedAmount int           `json:"requested_amount"`
}

type ListTransactionResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    TransactionData `json:"data"`
}

func ListTransactions(c *client.Client, request *ListTransactionRequest) (ListTransactionResponse, error) {
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
