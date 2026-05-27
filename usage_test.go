// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-api-go"
	"github.com/boltz-bio/boltz-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-api-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := boltzapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Predictions.StructureAndBinding.Start(context.TODO(), boltzapi.PredictionStructureAndBindingStartParams{
		Input: boltzapi.PredictionStructureAndBindingStartParamsInput{
			Entities: []boltzapi.PredictionStructureAndBindingStartParamsInputEntityUnion{{
				OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntity: &boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
				},
			}},
		},
		Model: "boltz-2.1",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ID)
}
