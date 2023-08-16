package verify

import "testing"

func TestVerifyTransaction(t *testing.T) {

	client := Client{
		Authorization: "sk_test_f572197fbc13951b13afafc0d0f6517ed7ec12eb",
		ContentType:   "application/json",
	}

	sampleVerificationRequest := VerificationRequest{
		Reference: "c2z7k6t1i4",
	}

	response, err := client.VerifyTransaction(&sampleVerificationRequest)

	if err != nil {
		t.Errorf("Test failed with error %v", err)
	}

	if response.Status == false {
		t.Errorf("Expected response of status true, but got a status of %v and an error message that says %v", response.Status, response.Message)
	}
}
