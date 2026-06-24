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

func TestProteinLibraryScreenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Get(
		context.TODO(),
		"id",
		boltzapi.ProteinLibraryScreenGetParams{
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

func TestProteinLibraryScreenListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.List(context.TODO(), boltzapi.ProteinLibraryScreenListParams{
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

func TestProteinLibraryScreenDeleteData(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.EstimateCost(context.TODO(), boltzapi.ProteinLibraryScreenEstimateCostParams{
		Proteins: []boltzapi.ProteinLibraryScreenEstimateCostParamsProtein{{
			Entities: []boltzapi.ProteinLibraryScreenEstimateCostParamsProteinEntityUnion{{
				OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity: &boltzapi.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			ID: boltzapi.String("id"),
		}},
		Target: boltzapi.ProteinLibraryScreenEstimateCostParamsTargetUnion{
			OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTarget: &boltzapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:    []int64{10, 11, 12},
							FlexibleResidues:   []int64{5, 6, 7},
							NonBindingResidues: []int64{0, 1, 2},
						},
					},
				},
				Structure: boltzapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureURLSource: &boltzapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: boltzapi.String("idempotency_key"),
		WorkspaceID:    boltzapi.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.ListResults(
		context.TODO(),
		"id",
		boltzapi.ProteinLibraryScreenListResultsParams{
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

func TestProteinLibraryScreenResume(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Resume(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Start(context.TODO(), boltzapi.ProteinLibraryScreenStartParams{
		Proteins: []boltzapi.ProteinLibraryScreenStartParamsProtein{{
			Entities: []boltzapi.ProteinLibraryScreenStartParamsProteinEntityUnion{{
				OfProteinLibraryScreenStartsProteinEntityProteinEntity: &boltzapi.ProteinLibraryScreenStartParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			ID: boltzapi.String("id"),
		}},
		Target: boltzapi.ProteinLibraryScreenStartParamsTargetUnion{
			OfProteinLibraryScreenStartsTargetStructureTemplateTarget: &boltzapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:    []int64{10, 11, 12},
							FlexibleResidues:   []int64{5, 6, 7},
							NonBindingResidues: []int64{0, 1, 2},
						},
					},
				},
				Structure: boltzapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureURLSource: &boltzapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: boltzapi.String("idempotency_key"),
		WorkspaceID:    boltzapi.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenStop(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
