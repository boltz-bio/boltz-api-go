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

func TestSmallMoleculeLibraryScreenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Get(
		context.TODO(),
		"id",
		boltzapi.SmallMoleculeLibraryScreenGetParams{
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

func TestSmallMoleculeLibraryScreenListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.List(context.TODO(), boltzapi.SmallMoleculeLibraryScreenListParams{
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

func TestSmallMoleculeLibraryScreenDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeLibraryScreenEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.EstimateCost(context.TODO(), boltzapi.SmallMoleculeLibraryScreenEstimateCostParams{
		Molecules: []boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsMolecule{{
			Smiles: "smiles",
			ID:     boltzapi.String("id"),
		}},
		Target: boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTarget{
			Entities: []boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityProteinEntity: &boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			Bonds: []boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetBond{{
				Atom1: boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1PolymerAtom: &boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
				Atom2: boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2PolymerAtom: &boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
			}},
			Constraints: []boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintPocketConstraint: &boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzapi.Bool(true),
				},
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
			Type:             "no_template",
		},
		IdempotencyKey: boltzapi.String("idempotency_key"),
		MoleculeFilters: boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &boltzapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: boltzapi.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeLibraryScreenListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.ListResults(
		context.TODO(),
		"id",
		boltzapi.SmallMoleculeLibraryScreenListResultsParams{
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

func TestSmallMoleculeLibraryScreenResume(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Resume(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeLibraryScreenStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Start(context.TODO(), boltzapi.SmallMoleculeLibraryScreenStartParams{
		Molecules: []boltzapi.SmallMoleculeLibraryScreenStartParamsMolecule{{
			Smiles: "smiles",
			ID:     boltzapi.String("id"),
		}},
		Target: boltzapi.SmallMoleculeLibraryScreenStartParamsTarget{
			Entities: []boltzapi.SmallMoleculeLibraryScreenStartParamsTargetEntityUnion{{
				OfSmallMoleculeLibraryScreenStartsTargetEntityProteinEntity: &boltzapi.SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			Bonds: []boltzapi.SmallMoleculeLibraryScreenStartParamsTargetBond{{
				Atom1: boltzapi.SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union{
					OfSmallMoleculeLibraryScreenStartsTargetBondAtom1PolymerAtom: &boltzapi.SmallMoleculeLibraryScreenStartParamsTargetBondAtom1PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
				Atom2: boltzapi.SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union{
					OfSmallMoleculeLibraryScreenStartsTargetBondAtom2PolymerAtom: &boltzapi.SmallMoleculeLibraryScreenStartParamsTargetBondAtom2PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
			}},
			Constraints: []boltzapi.SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion{{
				OfSmallMoleculeLibraryScreenStartsTargetConstraintPocketConstraint: &boltzapi.SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzapi.Bool(true),
				},
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
			Type:             "no_template",
		},
		IdempotencyKey: boltzapi.String("idempotency_key"),
		MoleculeFilters: boltzapi.SmallMoleculeLibraryScreenStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzapi.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzapi.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzapi.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: boltzapi.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeLibraryScreenStop(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
