// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"
	"net/http"
	"slices"

	"github.com/et0and/address-go/internal/apijson"
	"github.com/et0and/address-go/internal/requestconfig"
	"github.com/et0and/address-go/option"
	"github.com/et0and/address-go/packages/param"
	"github.com/et0and/address-go/packages/respjson"
)

// Health and API key onboarding endpoints that do not require authentication.
//
// RequestKeyService contains methods and other services that help with interacting
// with the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestKeyService] method instead.
type RequestKeyService struct {
	options []option.RequestOption
}

// NewRequestKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestKeyService(opts ...option.RequestOption) (r RequestKeyService) {
	r = RequestKeyService{}
	r.options = opts
	return
}

// Submit a proof-of-work solution to obtain an API key. The request must include a
// valid nonce that solves the challenge previously obtained from GET /challenge.
// This prevents automated abuse while allowing legitimate users to access the API.
func (r *RequestKeyService) New(ctx context.Context, body RequestKeyNewParams, opts ...option.RequestOption) (res *RequestKeyNewResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "request-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type RequestKeyNewResponse struct {
	APIKey    string  `json:"apiKey" api:"required"`
	RateLimit float64 `json:"rateLimit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		RateLimit   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RequestKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequestKeyNewParams struct {
	Token     string  `json:"token" api:"required"`
	Challenge string  `json:"challenge" api:"required"`
	Nonce     float64 `json:"nonce" api:"required"`
	paramObj
}

func (r RequestKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RequestKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequestKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
