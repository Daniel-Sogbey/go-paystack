package main

import (
	"fmt"

	payment "github.com/Daniel-Sogbey/paystack/payments"
)

func main() {
	fmt.Println("Hello, World")

	//SET UP CLIENT

	client := &payment.Client{
		Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
		ContentType:   "application/json",
	}

	//SET UP ACCEPT PAYMENT REQUEST BODY

	request := &payment.AcceptPaymentRequest{
		Email:  "mathematics06physics@gmail.com",
		Amount: "2000",
	}

	// EXAMPLE ACCEPT PAYMENT REQUEST
	data, err := client.AcceptPayment(request)

	if err != nil {
		fmt.Printf("Error making request to paystack api : %v", err)
	}

	fmt.Println("REQUEST RESPONSE : ", data)

}
