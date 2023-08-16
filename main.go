package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	//SET UP CLIENT
	// client := &verify.Client{
	// 	Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
	// 	ContentType:   "application/json",
	// }

	// client2 := &verify.Client{
	// 	Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
	// 	ContentType:   "application/json",
	// }

	// //SET UP INITIALIZE TRANSACTION REQUEST BODY
	// request := &initialize.InitializeTransactionRequest{
	// 	Email:  "mathematics06physics@gmail.com",
	// 	Amount: "2000",
	// }

	// // EXAMPLE INITIALIZE TRANSACTION REQUEST
	// data, err := client.InitializeTransaction(request)

	// fmt.Println("REQUEST RESPONSE : ", data)

	// if err != nil {
	// 	fmt.Printf("Error making request to paystack api : %v", err)
	// }

	//SET UP VERIFY TRANSACTION REQUEST BODY
	// request2 := &verify.VerificationRequest{
	// 	Reference: "c2z7k6t1i4",
	// }

	// data2, _ := client.VerifyTransaction(request2)

	// fmt.Println("REQUEST RESPONSE : ", data2)

}
