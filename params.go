package checkout

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	// Headers may be used to provide extra header lines on the HTTP request.
	Headers        http.Header `form:"-"`
	IdempotencyKey *string     `form:"-"`
	Provider       *string     `form:"-"`
	SourceID       *string     `form:"-"`
	Authorization  *string     `form:"-"` // Passed as header
}

// GetParams -
func (p *Params) GetParams() *Params {
	return p
}

// SetIdempotencyKey sets a value for the Idempotency-Key header.
func (p *Params) SetIdempotencyKey(val string) {
	p.IdempotencyKey = &val
}

// SetProvider sets a value for the Cko-Provider header for Vault Exchange.
// The name of the third party you would like to forward the request to.
func (p *Params) SetProvider(val string) {
	p.Provider = &val
}

// SetSourceID sets a value for the Cko-Source-Id header for Vault Exchange.
// The ID of the payment source containing the card information you would like to use in the request.
func (p *Params) SetSourceID(val string) {
	p.SourceID = &val
}

// SetAuthorization sets a value for the Cko-Authorization header for Vault Exchange.
func (p *Params) SetAuthorization(val string) {
	p.Authorization = &val
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

// NewIdempotencyKey -
func NewIdempotencyKey() string {
	now := time.Now().UnixNano()
	buf := make([]byte, 4)
	rand.Read(buf)
	return fmt.Sprintf("%v_%v", now, base64.URLEncoding.EncodeToString(buf)[:6])
}
