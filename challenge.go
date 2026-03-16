// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/address-go/internal/apijson"
	"github.com/stainless-sdks/address-go/internal/requestconfig"
	"github.com/stainless-sdks/address-go/option"
	"github.com/stainless-sdks/address-go/packages/respjson"
)

// Health, API information, and API key onboarding endpoints that do not require
// authentication.
//
// ChallengeService contains methods and other services that help with interacting
// with the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChallengeService] method instead.
type ChallengeService struct {
	options []option.RequestOption
}

// NewChallengeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChallengeService(opts ...option.RequestOption) (r ChallengeService) {
	r = ChallengeService{}
	r.options = opts
	return
}

// Returns a cryptographic challenge for proof-of-work based API key registration.
// The challenge must be solved by finding a nonce that, when combined with the
// challenge data, produces a hash below the difficulty threshold. Use this
// challenge with the POST /request-key endpoint to obtain an API key.
func (r *ChallengeService) Get(ctx context.Context, opts ...option.RequestOption) (res *ChallengeGetResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "challenge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ChallengeGetResponse struct {
	Token      string  `json:"token" api:"required"`
	Challenge  string  `json:"challenge" api:"required"`
	Difficulty float64 `json:"difficulty" api:"required"`
	ExpiresAt  float64 `json:"expiresAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Challenge   respjson.Field
		Difficulty  respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChallengeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChallengeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
