package client

import "net/http"

type Client struct {
	Authorization string
	ContentType   string
	Client        *http.Client
}

func NewClient(authorization, contentType string) *Client {
	return &Client{
		Authorization: authorization,
		ContentType:   contentType,
		Client:        &http.Client{},
	}
}
