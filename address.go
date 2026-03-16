// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address

import (
	"github.com/et0and/address-go/internal/apijson"
	"github.com/et0and/address-go/packages/respjson"
)

type GetAPIInfoResponse struct {
	Endpoints []string `json:"endpoints" api:"required"`
	Name      string   `json:"name" api:"required"`
	Version   string   `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoints   respjson.Field
		Name        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetAPIInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *GetAPIInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
