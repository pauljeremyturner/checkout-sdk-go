package common

// Exemption ...
type Exemption string

const (
	// LowValue ...
	LowValue Exemption = "low_value"
	// LowRisk ...
	LowRisk Exemption = "low_risk"
	// SecureCorporatePayment ...
	SecureCorporatePayment Exemption = "secure_corporate_payment"
	// TrustedListing ...
	TrustedListing Exemption = "trusted_listing"
)

func (c Exemption) String() string {
	return string(c)
}
