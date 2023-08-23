package Transactions

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Daniel-Sogbey/paystack-go-sdk/internal"
	"github.com/Daniel-Sogbey/paystack-go-sdk/paystack"
)

type ChargeAuthorizationRequest struct {
	Email             string `json:"email"`
	Amount            string `json:"amount"`
	AuthorizationCode string `json:"authorization_code"`
}

type Data struct {
	Amount          int           `json:"amount"`
	Currency        string        `json:"currency"`
	TransactionDate string        `json:"transaction_date"`
	Status          string        `json:"status"`
	Reference       string        `json:"reference"`
	Domain          string        `json:"domain"`
	Metadata        string        `json:"metadata"`
	GatewayResponse string        `json:"gateway_response"`
	Message         interface{}   `json:"message"`
	Channel         string        `json:"channel"`
	IPAddress       interface{}   `json:"ip_address"`
	Log             interface{}   `json:"log"`
	Fees            int           `json:"fees"`
	Authorization   Authorization `json:"authorization"`
	Customer        Customer      `json:"customer"`
	Plan            interface{}   `json:"plan"`
	ID              int           `json:"id"`
}

type ChargeAuthorizationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func ChargeAuthorization(c *paystack.Client, request *ChargeAuthorizationRequest) (ChargeAuthorizationResponse, error) {
	c.Client = &http.Client{}

	apiUrl := "https://api.paystack.co/transaction/charge_authorization"

	requestJson, err := json.Marshal(request)

	if err != nil {

	}

	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestJson))

	internal.SetHeaders(req, c.Authorization)

	response, err := c.Client.Do(req)

	responseBody, err := io.ReadAll(response.Body)

	var chargeAuthorizationResponse ChargeAuthorizationResponse

	err = json.Unmarshal(responseBody, &chargeAuthorizationResponse)

	if err != nil {

	}

	return ChargeAuthorizationResponse{
		Status:  chargeAuthorizationResponse.Status,
		Message: chargeAuthorizationResponse.Message,
		Data:    chargeAuthorizationResponse.Data,
	}, nil

}
