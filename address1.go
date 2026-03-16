// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/address-go/internal/apijson"
	"github.com/stainless-sdks/address-go/internal/apiquery"
	"github.com/stainless-sdks/address-go/internal/requestconfig"
	"github.com/stainless-sdks/address-go/option"
	"github.com/stainless-sdks/address-go/packages/param"
	"github.com/stainless-sdks/address-go/packages/respjson"
)

// Look up and list NZ addresses with filtering, pagination, and address ID lookup.
//
// AddressService contains methods and other services that help with interacting
// with the address API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	options []option.RequestOption
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r AddressService) {
	r = AddressService{}
	r.options = opts
	return
}

// Retrieve a single address record by its LINZ address_id.
//
// The address_id is a unique identifier assigned by Land Information New Zealand
// (LINZ). This is the canonical way to retrieve a specific address when you know
// its ID.
//
// **Response formats:**
//
// - Default: Full address object with all LINZ attributes
// - Simple (format=simple): Compact representation with essential fields only
//
// **Example:** `/v1/addresses/123456?format=simple`
func (r *AddressService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AddressGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/addresses/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a paginated list of addresses with optional filtering by location.
//
// **Filtering:**
//
//   - `town_city`: Filter by town/city name (e.g., "Wellington")
//   - `suburb_locality`: Filter by suburb/locality (e.g., "Te Aro")
//   - `road_name`: Filter by road/street name (e.g., "Lambton Quay")
//   - `bbox`: Bounding box filter as comma-separated coordinates
//     (min_lon,min_lat,max_lon,max_lat)
//
// **Pagination:**
//
// - `limit`: Maximum number of results (default: 100, max: 1000)
// - `offset`: Number of results to skip
//
// **Example:** `/v1/addresses?town_city=Wellington&limit=50`
func (r *AddressService) List(ctx context.Context, query AddressListParams, opts ...option.RequestOption) (res *[]AddressListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AddressGetResponse struct {
	AddressID            float64 `json:"addressId" api:"required"`
	FullAddress          string  `json:"fullAddress" api:"required"`
	FullAddressNumber    string  `json:"fullAddressNumber" api:"required"`
	FullAddressRoad      string  `json:"fullAddressRoad" api:"required"`
	Latitude             float64 `json:"latitude" api:"required"`
	Longitude            float64 `json:"longitude" api:"required"`
	Postcode             string  `json:"postcode" api:"required"`
	Region               string  `json:"region" api:"required"`
	Suburb               string  `json:"suburb" api:"required"`
	TerritorialAuthority string  `json:"territorialAuthority" api:"required"`
	TownCity             string  `json:"townCity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressID            respjson.Field
		FullAddress          respjson.Field
		FullAddressNumber    respjson.Field
		FullAddressRoad      respjson.Field
		Latitude             respjson.Field
		Longitude            respjson.Field
		Postcode             respjson.Field
		Region               respjson.Field
		Suburb               respjson.Field
		TerritorialAuthority respjson.Field
		TownCity             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressListResponse struct {
	AddressID            float64 `json:"addressId" api:"required"`
	FullAddress          string  `json:"fullAddress" api:"required"`
	FullAddressNumber    string  `json:"fullAddressNumber" api:"required"`
	FullAddressRoad      string  `json:"fullAddressRoad" api:"required"`
	Latitude             float64 `json:"latitude" api:"required"`
	Longitude            float64 `json:"longitude" api:"required"`
	Postcode             string  `json:"postcode" api:"required"`
	Region               string  `json:"region" api:"required"`
	Suburb               string  `json:"suburb" api:"required"`
	TerritorialAuthority string  `json:"territorialAuthority" api:"required"`
	TownCity             string  `json:"townCity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressID            respjson.Field
		FullAddress          respjson.Field
		FullAddressNumber    respjson.Field
		FullAddressRoad      respjson.Field
		Latitude             respjson.Field
		Longitude            respjson.Field
		Postcode             respjson.Field
		Region               respjson.Field
		Suburb               respjson.Field
		TerritorialAuthority respjson.Field
		TownCity             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressListResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressListParams struct {
	Bbox           param.Opt[string] `query:"bbox,omitzero" json:"-"`
	Format         param.Opt[string] `query:"format,omitzero" json:"-"`
	Limit          param.Opt[string] `query:"limit,omitzero" json:"-"`
	Offset         param.Opt[string] `query:"offset,omitzero" json:"-"`
	RoadName       param.Opt[string] `query:"road_name,omitzero" json:"-"`
	SuburbLocality param.Opt[string] `query:"suburb_locality,omitzero" json:"-"`
	TownCity       param.Opt[string] `query:"town_city,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParams]'s query parameters as `url.Values`.
func (r AddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
