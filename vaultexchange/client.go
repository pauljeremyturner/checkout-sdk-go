package vaultexchange

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "exchange/payments"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

type (
	// Response ...
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		ResponseBody   []byte                   `json:"response_body,omitempty"`
	}
)

// NewClient ...
func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// Request ...
func (c *Client) Request(request interface{}, params *checkout.Params) (*Response, error) {
	response, err := c.API.Post("/"+path, request, params)
	resp := &Response{
		StatusResponse: response,
	}
	if err != nil {
		return resp, err
	}
	resp.ResponseBody = response.ResponseBody
	return resp, err
}
