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

func TestPredictionStructureAndBindingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Get(
		context.TODO(),
		"sab_pred_2X7Ab9Cd3Ef6Gh1JkLmN",
		boltzapi.PredictionStructureAndBindingGetParams{
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

func TestPredictionStructureAndBindingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.List(context.TODO(), boltzapi.PredictionStructureAndBindingListParams{
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

func TestPredictionStructureAndBindingDeleteData(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.DeleteData(context.TODO(), "sab_pred_2X7Ab9Cd3Ef6Gh1JkLmN")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPredictionStructureAndBindingEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.EstimateCost(context.TODO(), boltzapi.PredictionStructureAndBindingEstimateCostParams{
		Input: boltzapi.PredictionStructureAndBindingEstimateCostParamsInput{
			Entities: []boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntity: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
					Msa: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaUnion{
						OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa{
							Format: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatA3m,
							Source: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion{
								OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource{
									URL: "https://example.com",
								},
							},
						},
					},
				},
			}},
			Binding: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBindingUnion{
				OfPredictionStructureAndBindingEstimateCostsInputBindingLigandProteinBinding: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBond{{
				Atom1: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom1LigandAtom: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom2LigandAtom: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzapi.PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputConstraintPocketConstraint: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzapi.Bool(true),
				},
			}},
			ModelOptions: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputModelOptions{
				RecyclingSteps: boltzapi.Int(1),
				SamplingSteps:  boltzapi.Int(1),
				StepScale:      boltzapi.Float(1.3),
			},
			NumSamples: boltzapi.Int(1),
			Templates: []boltzapi.PredictionStructureAndBindingEstimateCostParamsInputTemplate{{
				TemplateChains: []boltzapi.PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateChain{{
					InputChainID:    "input_chain_id",
					TemplateChainID: "template_chain_id",
				}},
				TemplateStructure: boltzapi.PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureUnion{
					OfPredictionStructureAndBindingEstimateCostsInputTemplateTemplateStructureURLSource: &boltzapi.PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureURLSource{
						URL: "https://example.com",
					},
				},
				ForceThresholdAngstroms: boltzapi.Float(0),
			}},
		},
		Model:          "boltz-2.1",
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

func TestPredictionStructureAndBindingStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Start(context.TODO(), boltzapi.PredictionStructureAndBindingStartParams{
		Input: boltzapi.PredictionStructureAndBindingStartParamsInput{
			Entities: []boltzapi.PredictionStructureAndBindingStartParamsInputEntityUnion{{
				OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntity: &boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
					Msa: boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaUnion{
						OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa: &boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa{
							Format: boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatA3m,
							Source: boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion{
								OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource: &boltzapi.PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource{
									URL: "https://example.com",
								},
							},
						},
					},
				},
			}},
			Binding: boltzapi.PredictionStructureAndBindingStartParamsInputBindingUnion{
				OfPredictionStructureAndBindingStartsInputBindingLigandProteinBinding: &boltzapi.PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []boltzapi.PredictionStructureAndBindingStartParamsInputBond{{
				Atom1: boltzapi.PredictionStructureAndBindingStartParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingStartsInputBondAtom1LigandAtom: &boltzapi.PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzapi.PredictionStructureAndBindingStartParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingStartsInputBondAtom2LigandAtom: &boltzapi.PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzapi.PredictionStructureAndBindingStartParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingStartsInputConstraintPocketConstraint: &boltzapi.PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzapi.Bool(true),
				},
			}},
			ModelOptions: boltzapi.PredictionStructureAndBindingStartParamsInputModelOptions{
				RecyclingSteps: boltzapi.Int(1),
				SamplingSteps:  boltzapi.Int(1),
				StepScale:      boltzapi.Float(1.3),
			},
			NumSamples: boltzapi.Int(1),
			Templates: []boltzapi.PredictionStructureAndBindingStartParamsInputTemplate{{
				TemplateChains: []boltzapi.PredictionStructureAndBindingStartParamsInputTemplateTemplateChain{{
					InputChainID:    "input_chain_id",
					TemplateChainID: "template_chain_id",
				}},
				TemplateStructure: boltzapi.PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureUnion{
					OfPredictionStructureAndBindingStartsInputTemplateTemplateStructureURLSource: &boltzapi.PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureURLSource{
						URL: "https://example.com",
					},
				},
				ForceThresholdAngstroms: boltzapi.Float(0),
			}},
		},
		Model:          "boltz-2.1",
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
