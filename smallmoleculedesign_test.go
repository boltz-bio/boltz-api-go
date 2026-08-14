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

func TestSmallMoleculeDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Get(
		context.TODO(),
		"id",
		boltzapi.SmallMoleculeDesignGetParams{
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

func TestSmallMoleculeDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.List(context.TODO(), boltzapi.SmallMoleculeDesignListParams{
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

func TestSmallMoleculeDesignDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeDesignEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.EstimateCost(context.TODO(), boltzapi.SmallMoleculeDesignEstimateCostParams{
		NumMolecules: 10,
		Target: boltzapi.SmallMoleculeDesignEstimateCostParamsTarget{
			Entities: []boltzapi.SmallMoleculeDesignEstimateCostParamsTargetEntityUnion{{
				OfSmallMoleculeDesignEstimateCostsTargetEntityProteinEntity: &boltzapi.SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			Bonds: []boltzapi.SmallMoleculeDesignEstimateCostParamsTargetBond{{
				Atom1: boltzapi.SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union{
					OfSmallMoleculeDesignEstimateCostsTargetBondAtom1PolymerAtom: &boltzapi.SmallMoleculeDesignEstimateCostParamsTargetBondAtom1PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
				Atom2: boltzapi.SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union{
					OfSmallMoleculeDesignEstimateCostsTargetBondAtom2PolymerAtom: &boltzapi.SmallMoleculeDesignEstimateCostParamsTargetBondAtom2PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
			}},
			Constraints: []boltzapi.SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion{{
				OfSmallMoleculeDesignEstimateCostsTargetConstraintPocketConstraint: &boltzapi.SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint{
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
		ChemicalSpace:  boltzapi.SmallMoleculeDesignEstimateCostParamsChemicalSpaceEnamineReal,
		IdempotencyKey: boltzapi.String("idempotency_key"),
		MoleculeFilters: boltzapi.SmallMoleculeDesignEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzapi.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzapi.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &boltzapi.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.ListResults(
		context.TODO(),
		"id",
		boltzapi.SmallMoleculeDesignListResultsParams{
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

func TestSmallMoleculeDesignResume(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Resume(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Start(context.TODO(), boltzapi.SmallMoleculeDesignStartParams{
		NumMolecules: 10,
		Target: boltzapi.SmallMoleculeDesignStartParamsTarget{
			Entities: []boltzapi.SmallMoleculeDesignStartParamsTargetEntityUnion{{
				OfSmallMoleculeDesignStartsTargetEntityProteinEntity: &boltzapi.SmallMoleculeDesignStartParamsTargetEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.SmallMoleculeDesignStartParamsTargetEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			Bonds: []boltzapi.SmallMoleculeDesignStartParamsTargetBond{{
				Atom1: boltzapi.SmallMoleculeDesignStartParamsTargetBondAtom1Union{
					OfSmallMoleculeDesignStartsTargetBondAtom1PolymerAtom: &boltzapi.SmallMoleculeDesignStartParamsTargetBondAtom1PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
				Atom2: boltzapi.SmallMoleculeDesignStartParamsTargetBondAtom2Union{
					OfSmallMoleculeDesignStartsTargetBondAtom2PolymerAtom: &boltzapi.SmallMoleculeDesignStartParamsTargetBondAtom2PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
			}},
			Constraints: []boltzapi.SmallMoleculeDesignStartParamsTargetConstraintUnion{{
				OfSmallMoleculeDesignStartsTargetConstraintPocketConstraint: &boltzapi.SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint{
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
		ChemicalSpace:  boltzapi.SmallMoleculeDesignStartParamsChemicalSpaceEnamineReal,
		IdempotencyKey: boltzapi.String("idempotency_key"),
		MoleculeFilters: boltzapi.SmallMoleculeDesignStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzapi.SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzapi.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzapi.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeDesignStop(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
