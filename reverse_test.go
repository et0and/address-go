// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/et0and/address-go"
	"github.com/et0and/address-go/internal/testutil"
	"github.com/et0and/address-go/option"
)

func TestReverseGeocodeWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := address.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Reverse.Geocode(context.TODO(), address.ReverseGeocodeParams{
		Point:  "point",
		Format: address.String("format"),
		Limit:  address.String("limit"),
	})
	if err != nil {
		var apierr *address.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
