package initialize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Daniel-Sogbey/go-paystack/internal"
	"github.com/Daniel-Sogbey/go-paystack/paystack"
)

type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
}

type MetaData struct {
	CustomFields []CustomField `json:"custom_fields"`
}

type InitializeTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

type Bank struct {
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}

type InitializeTransactionRequest struct {
	Email  string `json:"email"`
	Amount string `json:"amount"`
}

func InitializeTransaction(c *paystack.Client, request *InitializeTransactionRequest) (InitializeTransactionResponse, error) {
	c.Client = &http.Client{}

	apiUrl := "https://api.paystack.co/transaction/initialize"
	request.Email = "1@2.com"

	requestJSON, err := json.Marshal(request)

	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(requestJSON)))

	internal.SetHeaders(req, c.Authorization)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error making request to Paystack Accept Payment API : %v", err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	var initializeTransactionResponse InitializeTransactionResponse

	err = json.Unmarshal(responseBody, &initializeTransactionResponse)

	if err != nil {
		log.Fatalf("Error parsing response to fit format %v", err)
	}

	return InitializeTransactionResponse{
		Status:  initializeTransactionResponse.Status,
		Message: initializeTransactionResponse.Message,
		Data:    initializeTransactionResponse.Data,
	}, nil
}
