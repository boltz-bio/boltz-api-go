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

func TestSmallMoleculeExploreGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Explore.Get(
		context.TODO(),
		"id",
		boltzapi.SmallMoleculeExploreGetParams{
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

func TestSmallMoleculeExploreListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Explore.ListResults(
		context.TODO(),
		"id",
		boltzapi.SmallMoleculeExploreListResultsParams{
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

func TestSmallMoleculeExploreResume(t *testing.T) {
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
	_, err := client.SmallMolecule.Explore.Resume(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeExploreStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Explore.Start(context.TODO(), boltzapi.SmallMoleculeExploreStartParams{
		Budget: 1,
		Library: boltzapi.SmallMoleculeExploreStartParamsLibrary{
			Format: boltzapi.SmallMoleculeExploreStartParamsLibraryFormatCsv,
			Source: boltzapi.SmallMoleculeExploreStartParamsLibrarySourceUnion{
				OfSmallMoleculeExploreStartsLibrarySourceURLSource: &boltzapi.SmallMoleculeExploreStartParamsLibrarySourceURLSource{
					URL: "https://example.com",
				},
			},
			IDColumn:     boltzapi.String("id_column"),
			SmilesColumn: boltzapi.String("smiles_column"),
		},
		Target: boltzapi.SmallMoleculeExploreStartParamsTarget{
			Entities: []boltzapi.SmallMoleculeExploreStartParamsTargetEntityUnion{{
				OfSmallMoleculeExploreStartsTargetEntityProteinEntity: &boltzapi.SmallMoleculeExploreStartParamsTargetEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzapi.Bool(true),
					Modifications: []boltzapi.SmallMoleculeExploreStartParamsTargetEntityProteinEntityModification{{
						ResidueIndex: 0,
						Value:        "value",
					}},
				},
			}},
			Bonds: []boltzapi.SmallMoleculeExploreStartParamsTargetBond{{
				Atom1: boltzapi.SmallMoleculeExploreStartParamsTargetBondAtom1Union{
					OfSmallMoleculeExploreStartsTargetBondAtom1PolymerAtom: &boltzapi.SmallMoleculeExploreStartParamsTargetBondAtom1PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
				Atom2: boltzapi.SmallMoleculeExploreStartParamsTargetBondAtom2Union{
					OfSmallMoleculeExploreStartsTargetBondAtom2PolymerAtom: &boltzapi.SmallMoleculeExploreStartParamsTargetBondAtom2PolymerAtom{
						AtomName:     "atom_name",
						ChainID:      "chain_id",
						ResidueIndex: 0,
					},
				},
			}},
			Constraints: []boltzapi.SmallMoleculeExploreStartParamsTargetConstraintUnion{{
				OfSmallMoleculeExploreStartsTargetConstraintPocketConstraint: &boltzapi.SmallMoleculeExploreStartParamsTargetConstraintPocketConstraint{
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
		MoleculeFilters: boltzapi.SmallMoleculeExploreStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzapi.SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzapi.SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzapi.SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeExploreStop(t *testing.T) {
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
	_, err := client.SmallMolecule.Explore.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
