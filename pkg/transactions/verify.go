package Transactions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/go-paystack/internal"
	"github.com/Daniel-Sogbey/go-paystack/paystack"
)

type VerificationRequest struct {
	Reference string
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
	ID                       int    `json:"id"`
	FirstName                string `json:"first_name"`
	LastName                 string `json:"last_name"`
	Email                    string `json:"email"`
	CustomerCode             string `json:"customer_code"`
	Phone                    string `json:"phone"`
	Metadata                 string `json:"metadata"`
	RiskAction               string `json:"risk_action"`
	InternationalFormatPhone string `json:"international_format_phone"`
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
	Metadata        string        `json:"metadata"`
	Log             Log           `json:"log"`
	Fees            int           `json:"fees"`
	Authorization   Authorization `json:"authorization"`
	Customer        Customer      `json:"customer"`
	Plan            interface{}   `json:"plan"`
	Split           interface{}   `json:"split"`
	OrderID         interface{}   `json:"order_id"`
	// PaidAt             string       `json:"paidAt"`
	// CreatedAt          string       `json:"createdAt"`
	RequestedAmount    int         `json:"requested_amount"`
	PosTransactionData interface{} `json:"pos_transaction_data"`
	Source             interface{} `json:"source"`
	FeesBreakdown      interface{} `json:"fees_breakdown"`
	TransactionDate    string      `json:"transaction_date"`
	PlanObject         interface{} `json:"plan_object"`
	Subaccount         interface{} `json:"subaccount"`
}

type VerificationResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    TransactionData `json:"data"`
}

func Verify(c *paystack.Client, request *VerificationRequest) (VerificationResponse, error) {
	c.Client = &http.Client{}

	apiUrl := "https://api.paystack.co/transaction/verify/" + request.Reference

	req, _ := http.NewRequest("GET", apiUrl, nil)

	internal.SetHeaders(req, c.Authorization)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error making a verification reuqest to Paystack API : %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	if err != nil {
		log.Fatalf("Error reading verification response body into a slice of bytes %v", err)
	}

	var verificationResponse VerificationResponse

	err = json.Unmarshal(responseBody, &verificationResponse)

	if err != nil {
		log.Fatalf("Error unmarshalling verification response body : %v", err)
	}

	return VerificationResponse{
		Status:  verificationResponse.Status,
		Message: verificationResponse.Message,
		Data:    verificationResponse.Data,
	}, nil

}
