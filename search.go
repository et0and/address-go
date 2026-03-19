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
// SearchService contains methods and other services that help with interacting
// with the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.options = opts
	return
}

// Search addresses using full-text search with intelligent query processing.
//
// **Search features:**
//
// - FTS5 full-text search with ranking by relevance
// - Automatic abbreviation expansion (e.g., "st" → "street", "rd" → "road")
// - Fuzzy matching fallback for typos and variations
// - Address component matching (street, suburb, city, postcode)
//
// **Query parameters:**
//
// - `q`: Search query string (required)
// - `limit`: Maximum results (default: 100, max: 1000)
// - `format`: Response format - "full" or "simple"
//
// **Examples:**
//
// - `/v1/search?q=lambton+quay` - Search for addresses on Lambton Quay
// - `/v1/search?q=123+quay+st+auckland` - Search for specific address
// - `/v1/search?q=wlg&limit=20` - Abbreviation expansion
func (r *SearchService) Query(ctx context.Context, query SearchQueryParams, opts ...option.RequestOption) (res *[]SearchQueryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SearchQueryResponse struct {
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
func (r SearchQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryParams struct {
	Q       string            `query:"q" api:"required" json:"-"`
	Bbox    param.Opt[string] `query:"bbox,omitzero" json:"-"`
	Format  param.Opt[string] `query:"format,omitzero" json:"-"`
	Limit   param.Opt[string] `query:"limit,omitzero" json:"-"`
	Polygon param.Opt[string] `query:"polygon,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchQueryParams]'s query parameters as `url.Values`.
func (r SearchQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
