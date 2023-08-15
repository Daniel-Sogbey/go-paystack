package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CustomField struct {
	DisplayName  string `json:"display_name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
}

type MetaData struct {
	CustomFields []CustomField `json:"custom_fields"`
}

type ResponseData struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type AcceptPaymentResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type Bank struct {
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}

type AcceptPaymentRequest struct {
	Email  string `json:"email"`
	Amount string `json:"amount"`
}

type Client struct {
	Authorization string
	ContentType   string
	Client        *http.Client
}

func setHeaders(request *http.Request, authorization string) {
	request.Header.Set("authorization", "Bearer "+authorization)
	request.Header.Set("content-type", "application/json")
}

func (c *Client) AcceptPayment(request *AcceptPaymentRequest) (AcceptPaymentResponse, error) {

	c.Client = &http.Client{}

	apiUrl := "https://api.paystack.co/transaction/initialize"
	request.Email = "1@2.com"

	requestJSON, err := json.Marshal(request)

	fmt.Println("REQUEST BODY", string(requestJSON))

	req, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(requestJSON)))

	setHeaders(req, c.Authorization)

	response, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("Error making request to Paystack Accept Payment API : %v", err)
	}

	defer response.Body.Close()

	fmt.Println("RESPONSE STATUS CODE : ", response.StatusCode)

	responseBody, err := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	var acceptedResponse AcceptPaymentResponse

	err = json.Unmarshal(responseBody, &acceptedResponse)

	if err != nil {
		log.Fatalf("Error parsing response to fit format %v", err)
	}

	return AcceptPaymentResponse{
		Status:  acceptedResponse.Status,
		Message: acceptedResponse.Message,
		Data:    acceptedResponse.Data,
	}, nil
}
