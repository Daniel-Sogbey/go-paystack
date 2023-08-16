package internal

import "net/http"

func SetHeaders(request *http.Request, authorization string) {
	request.Header.Set("authorization", "Bearer "+authorization)

	request.Header.Set("content-type", "application/json")
}
