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
		OfProteinDesignRunInput: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInput{
			BinderSpecification: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationUnion{
				OfProteinDesignEstimateCostsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpec: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpec{
					ChainSelection: map[string]boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
						"B": {
							OfProteinDesignEstimateCostsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
								CropResidues: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
									OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
								},
								DesignMotifs: []boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
									OfProteinDesignEstimateCostsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
										DesignLengthRange: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
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
					Modality: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
					Structure: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecStructureUnion{
						OfProteinDesignEstimateCostsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
							URL: "https://example.com",
						},
					},
					Rules: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecRules{
						ExcludedAminoAcids:     []string{"x"},
						ExcludedSequenceMotifs: []string{"string"},
						MaxHydrophobicFraction: boltzapi.Float(0),
					},
				},
			},
			NumProteins: 10,
			Target: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetUnion{
				OfProteinDesignEstimateCostsBodyProteinDesignRunInputTargetStructureTemplateTarget: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetStructureTemplateTarget{
					ChainSelection: map[string]boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionUnion{
						"A": {
							OfProteinDesignEstimateCostsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
								CropResidues: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
									OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
								},
								EpitopeResidues:    []int64{10, 11, 12},
								FlexibleResidues:   []int64{5, 6, 7},
								NonBindingResidues: []int64{0, 1, 2},
							},
						},
					},
					Structure: boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetStructureTemplateTargetStructureUnion{
						OfProteinDesignEstimateCostsBodyProteinDesignRunInputTargetStructureTemplateTargetStructureURLSource: &boltzapi.ProteinDesignEstimateCostParamsBodyProteinDesignRunInputTargetStructureTemplateTargetStructureURLSource{
							URL: "https://example.com",
						},
					},
				},
			},
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

func TestProteinDesignListCuratedSpecifications(t *testing.T) {
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
	_, err := client.Protein.Design.ListCuratedSpecifications(context.TODO(), boltzapi.ProteinDesignListCuratedSpecificationsParams{
		Type: boltzapi.ProteinDesignListCuratedSpecificationsParamsTypeNanobody,
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
		OfProteinDesignRunInput: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInput{
			BinderSpecification: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationUnion{
				OfProteinDesignStartsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpec: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpec{
					ChainSelection: map[string]boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
						"B": {
							OfProteinDesignStartsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
								CropResidues: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
									OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
								},
								DesignMotifs: []boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
									OfProteinDesignStartsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
										DesignLengthRange: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
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
					Modality: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
					Structure: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecStructureUnion{
						OfProteinDesignStartsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
							URL: "https://example.com",
						},
					},
					Rules: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecRules{
						ExcludedAminoAcids:     []string{"x"},
						ExcludedSequenceMotifs: []string{"string"},
						MaxHydrophobicFraction: boltzapi.Float(0),
					},
				},
			},
			NumProteins: 10,
			Target: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetUnion{
				OfProteinDesignStartsBodyProteinDesignRunInputTargetStructureTemplateTarget: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetStructureTemplateTarget{
					ChainSelection: map[string]boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionUnion{
						"A": {
							OfProteinDesignStartsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
								CropResidues: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
									OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
								},
								EpitopeResidues:    []int64{10, 11, 12},
								FlexibleResidues:   []int64{5, 6, 7},
								NonBindingResidues: []int64{0, 1, 2},
							},
						},
					},
					Structure: boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetStructureTemplateTargetStructureUnion{
						OfProteinDesignStartsBodyProteinDesignRunInputTargetStructureTemplateTargetStructureURLSource: &boltzapi.ProteinDesignStartParamsBodyProteinDesignRunInputTargetStructureTemplateTargetStructureURLSource{
							URL: "https://example.com",
						},
					},
				},
			},
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
