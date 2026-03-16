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

// Dataset version, ingestion status, and service metadata.
//
// MetaService contains methods and other services that help with interacting with
// the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMetaService] method instead.
type MetaService struct {
	options []option.RequestOption
}

// NewMetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMetaService(opts ...option.RequestOption) (r MetaService) {
	r = MetaService{}
	r.options = opts
	return
}

// Returns metadata about the LINZ NZ Addresses dataset including:
//
// - Dataset version identifier
// - Last ingestion timestamp
// - Total record count
// - Data source information
//
// This endpoint is useful for client applications that need to verify data
// currency or display dataset attribution.
func (r *MetaService) Get(ctx context.Context, opts ...option.RequestOption) (res *MetaGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/meta"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MetaGetResponse struct {
	LastUpdated    string  `json:"lastUpdated" api:"required"`
	TotalAddresses float64 `json:"totalAddresses" api:"required"`
	Version        string  `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdated    respjson.Field
		TotalAddresses respjson.Field
		Version        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetaGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MetaGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
