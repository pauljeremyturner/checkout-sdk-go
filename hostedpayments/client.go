package hostedpayments

import (
	"encoding/json"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "hosted-payments"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// Request ...
func (c *Client) Request(request *Request, params *checkout.Params) (*Response, error) {
	response, err := c.API.Post("/"+path, request, params)
	resp := &Response{
		StatusResponse: response,
	}
	if err != nil {
		return resp, err
	}
	if response.StatusCode == http.StatusCreated {
		var hostedPaymentResponse HostedPaymentResponse
		err = json.Unmarshal(response.ResponseBody, &hostedPaymentResponse)
		resp.HostedPaymentResponse = &hostedPaymentResponse
	}
	return resp, err
}
