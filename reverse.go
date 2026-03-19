// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/et0and/address-go/internal/apijson"
	"github.com/et0and/address-go/internal/apiquery"
	"github.com/et0and/address-go/internal/requestconfig"
	"github.com/et0and/address-go/option"
	"github.com/et0and/address-go/packages/param"
	"github.com/et0and/address-go/packages/respjson"
)

// Full-text address search and reverse geocoding powered by FTS5 with abbreviation
// expansion.
//
// ReverseService contains methods and other services that help with interacting
// with the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReverseService] method instead.
type ReverseService struct {
	options []option.RequestOption
}

// NewReverseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReverseService(opts ...option.RequestOption) (r ReverseService) {
	r = ReverseService{}
	r.options = opts
	return
}

// Find the nearest addresses to given geographic coordinates (reverse geocoding).
//
// **Query parameters:**
//
// - `lat`: Latitude in decimal degrees (required)
// - `lng`: Longitude in decimal degrees (required)
// - `limit`: Maximum number of results (default: 10, max: 100)
// - `format`: Response format - "full" or "simple"
//
// **Distance calculation:** Results are sorted by distance from the provided
// coordinates, calculated using the Haversine formula for spherical distance on
// Earth.
//
// **Example:** `/v1/reverse?lat=-41.2865&lng=174.7762&limit=5`
func (r *ReverseService) Geocode(ctx context.Context, query ReverseGeocodeParams, opts ...option.RequestOption) (res *[]ReverseGeocodeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/reverse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ReverseGeocodeResponse struct {
	AddressID            float64 `json:"addressId" api:"required"`
	FullAddress          string  `json:"fullAddress" api:"required"`
	FullAddressNumber    string  `json:"fullAddressNumber" api:"required"`
	Latitude             float64 `json:"latitude" api:"required"`
	Longitude            float64 `json:"longitude" api:"required"`
	Postcode             string  `json:"postcode" api:"required"`
	Region               string  `json:"region" api:"required"`
	Suburb               string  `json:"suburb" api:"required"`
	TerritorialAuthority string  `json:"territorialAuthority" api:"required"`
	TownCity             string  `json:"townCity" api:"required"`
	FullAddressRoad      string  `json:"fullAddressRoad" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressID            respjson.Field
		FullAddress          respjson.Field
		FullAddressNumber    respjson.Field
		Latitude             respjson.Field
		Longitude            respjson.Field
		Postcode             respjson.Field
		Region               respjson.Field
		Suburb               respjson.Field
		TerritorialAuthority respjson.Field
		TownCity             respjson.Field
		FullAddressRoad      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReverseGeocodeResponse) RawJSON() string { return r.JSON.raw }
func (r *ReverseGeocodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReverseGeocodeParams struct {
	Point  string            `query:"point" api:"required" json:"-"`
	Format param.Opt[string] `query:"format,omitzero" json:"-"`
	Limit  param.Opt[string] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReverseGeocodeParams]'s query parameters as `url.Values`.
func (r ReverseGeocodeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
