package hostedpayments

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	// Request ...
	Request struct {
		Amount     uint64         `json:"amount,omitempty"`
		Currency   string         `json:"currency"`
		Reference  string         `json:"reference,omitempty"`
		Billing    common.Address `json:"billing,omitempty"`
		SuccessURL string         `json:"success_url,omitempty,omitempty"`
		FailureURL string         `json:"failure_url,omitempty,omitempty"`
		CancelURL  string         `json:"cancel_url,omitempty,omitempty"`
	}
)

type (
	// Response ...
	Response struct {
		StatusResponse        *checkout.StatusResponse `json:"api_response,omitempty"`
		HostedPaymentResponse *HostedPaymentResponse   `json:"hostedpayment_response,omitempty"`
	}

	// HostedPaymentResponse ...
	HostedPaymentResponse struct {
		Reference string                 `json:"reference"`
		Links     map[string]common.Link `json:"_links"`
	}
)
