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

func TestProteinDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Get(
		context.TODO(),
		"id",
		boltzapi.ProteinDesignGetParams{
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

func TestProteinDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.List(context.TODO(), boltzapi.ProteinDesignListParams{
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

func TestProteinDesignDeleteData(t *testing.T) {
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
	_, err := client.Protein.Design.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinDesignEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.EstimateCost(context.TODO(), boltzapi.ProteinDesignEstimateCostParams{
		BinderSpecification: boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationUnion{
			OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpec: &boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
										Max: 8,
										Min: 4,
									},
									EndIndex:   5,
									StartIndex: 0,
								},
							}},
						},
					},
				},
				Modality: boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: boltzapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: boltzapi.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: boltzapi.ProteinDesignEstimateCostParamsTargetUnion{
			OfProteinDesignEstimateCostsTargetStructureTemplateTarget: &boltzapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:    []int64{10, 11, 12},
							FlexibleResidues:   []int64{5, 6, 7},
							NonBindingResidues: []int64{0, 1, 2},
						},
					},
				},
				Structure: boltzapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureURLSource: &boltzapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
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

func TestProteinDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.ListResults(
		context.TODO(),
		"id",
		boltzapi.ProteinDesignListResultsParams{
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

func TestProteinDesignResume(t *testing.T) {
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
	_, err := client.Protein.Design.Resume(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Start(context.TODO(), boltzapi.ProteinDesignStartParams{
		BinderSpecification: boltzapi.ProteinDesignStartParamsBinderSpecificationUnion{
			OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpec: &boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
										Max: 8,
										Min: 4,
									},
									EndIndex:   5,
									StartIndex: 0,
								},
							}},
						},
					},
				},
				Modality: boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: boltzapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: boltzapi.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: boltzapi.ProteinDesignStartParamsTargetUnion{
			OfProteinDesignStartsTargetStructureTemplateTarget: &boltzapi.ProteinDesignStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzapi.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzapi.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzapi.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:    []int64{10, 11, 12},
							FlexibleResidues:   []int64{5, 6, 7},
							NonBindingResidues: []int64{0, 1, 2},
						},
					},
				},
				Structure: boltzapi.ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignStartsTargetStructureTemplateTargetStructureURLSource: &boltzapi.ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource{
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

func TestProteinDesignStop(t *testing.T) {
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
	_, err := client.Protein.Design.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
