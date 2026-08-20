// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-api-go"
	"github.com/boltz-bio/boltz-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-api-go/option"
)

func TestProteinSequenceRedesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.Get(
		context.TODO(),
		"id",
		boltzapi.ProteinSequenceRedesignGetParams{
			WorkspaceID: boltzapi.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.List(context.TODO(), boltzapi.ProteinSequenceRedesignListParams{
		AfterID:     boltzapi.String("after_id"),
		BeforeID:    boltzapi.String("before_id"),
		Limit:       boltzapi.Int(1),
		WorkspaceID: boltzapi.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignDeleteData(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.EstimateCost(context.TODO(), boltzapi.ProteinSequenceRedesignEstimateCostParams{
		OfBinderProteinSequenceRedesignRunInput: &boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInput{
			Entities: []boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion{{
				OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity: &boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity{
					ChainID: "x",
				},
			}, {
				OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity: &boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity{
					ChainID: "x",
				},
			}},
			NumProteins: 1,
			Structure: boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion{
				OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputStructureURLSource: &boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource{
					URL: "https://example.com",
				},
			},
			GlobalDesignFilters: []boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion{{
				OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter: &boltzapi.ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter{
					AminoAcids: []string{"I"},
				},
			}},
			IdempotencyKey: boltzapi.String("idempotency_key"),
			WorkspaceID:    boltzapi.String("workspace_id"),
		},
	})
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.ListResults(
		context.TODO(),
		"id",
		boltzapi.ProteinSequenceRedesignListResultsParams{
			AfterID:     boltzapi.String("after_id"),
			BeforeID:    boltzapi.String("before_id"),
			IDs:         boltzapi.String("ids"),
			Limit:       boltzapi.Int(1),
			WorkspaceID: boltzapi.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignResume(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.Resume(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.Start(context.TODO(), boltzapi.ProteinSequenceRedesignStartParams{
		OfBinderProteinSequenceRedesignRunInput: &boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInput{
			Entities: []boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion{{
				OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity: &boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity{
					ChainID: "x",
				},
			}, {
				OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity: &boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity{
					ChainID: "x",
				},
			}},
			NumProteins: 1,
			Structure: boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion{
				OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputStructureURLSource: &boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource{
					URL: "https://example.com",
				},
			},
			GlobalDesignFilters: []boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion{{
				OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter: &boltzapi.ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter{
					AminoAcids: []string{"I"},
				},
			}},
			IdempotencyKey: boltzapi.String("idempotency_key"),
			WorkspaceID:    boltzapi.String("workspace_id"),
		},
	})
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinSequenceRedesignStop(t *testing.T) {
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
	_, err := client.Protein.SequenceRedesign.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
