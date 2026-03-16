// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"
	"net/http"
	"slices"

	"github.com/et0and/address-go/internal/apijson"
	"github.com/et0and/address-go/internal/requestconfig"
	"github.com/et0and/address-go/option"
	"github.com/et0and/address-go/packages/respjson"
)

// Health, API information, and API key onboarding endpoints that do not require
// authentication.
//
// HealthService contains methods and other services that help with interacting
// with the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.options = opts
	return
}

// Returns a simple OK response to verify the API is operational. This endpoint
// does not require authentication and is useful for monitoring and load balancer
// health checks.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.options, opts)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type HealthCheckResponse struct {
	Status    string `json:"status" api:"required"`
	Timestamp string `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
