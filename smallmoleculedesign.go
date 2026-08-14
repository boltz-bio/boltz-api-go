// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/boltz-bio/boltz-api-go/internal/apijson"
	"github.com/boltz-bio/boltz-api-go/internal/apiquery"
	"github.com/boltz-bio/boltz-api-go/internal/requestconfig"
	"github.com/boltz-bio/boltz-api-go/option"
	"github.com/boltz-bio/boltz-api-go/packages/pagination"
	"github.com/boltz-bio/boltz-api-go/packages/param"
	"github.com/boltz-bio/boltz-api-go/packages/respjson"
	"github.com/boltz-bio/boltz-api-go/shared/constant"
)

// Generate novel small molecules optimized for binding to a protein target.
// Results are scored by binding confidence (likelihood of binding, for hit
// discovery), optimization score (binding strength ranking, for lead
// optimization), and structure confidence.
//
// SmallMoleculeDesignService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmallMoleculeDesignService] method instead.
type SmallMoleculeDesignService struct {
	Options []option.RequestOption
}

// NewSmallMoleculeDesignService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSmallMoleculeDesignService(opts ...option.RequestOption) (r SmallMoleculeDesignService) {
	r = SmallMoleculeDesignService{}
	r.Options = opts
	return
}

// Retrieve a design run by ID, including progress and status
func (r *SmallMoleculeDesignService) Get(ctx context.Context, id string, query SmallMoleculeDesignGetParams, opts ...option.RequestOption) (res *SmallMoleculeDesignGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/design/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List small molecule design runs, optionally filtered by workspace
func (r *SmallMoleculeDesignService) List(ctx context.Context, query SmallMoleculeDesignListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SmallMoleculeDesignListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/small-molecule/design"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List small molecule design runs, optionally filtered by workspace
func (r *SmallMoleculeDesignService) ListAutoPaging(ctx context.Context, query SmallMoleculeDesignListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SmallMoleculeDesignListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// design run. The design run record itself is retained with a `data_deleted_at`
// timestamp. This action is irreversible.
func (r *SmallMoleculeDesignService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeDesignDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/design/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Estimate the billed cost of a small molecule design run without creating any
// resource or consuming GPU. Includes generation charges implied by the scheduler
// iteration cap plus structure-scoring charges for each requested molecule.
func (r *SmallMoleculeDesignService) EstimateCost(ctx context.Context, body SmallMoleculeDesignEstimateCostParams, opts ...option.RequestOption) (res *SmallMoleculeDesignEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/small-molecule/design/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve paginated results from a design run
func (r *SmallMoleculeDesignService) ListResults(ctx context.Context, id string, query SmallMoleculeDesignListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[SmallMoleculeDesignListResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/design/%s/results", url.PathEscape(id))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve paginated results from a design run
func (r *SmallMoleculeDesignService) ListResultsAutoPaging(ctx context.Context, id string, query SmallMoleculeDesignListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SmallMoleculeDesignListResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, id, query, opts...))
}

// Resume a stopped small molecule design run from its last checkpoint
func (r *SmallMoleculeDesignService) Resume(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeDesignResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/design/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Create a new design run that generates novel small molecule candidates for a
// protein target
func (r *SmallMoleculeDesignService) Start(ctx context.Context, body SmallMoleculeDesignStartParams, opts ...option.RequestOption) (res *SmallMoleculeDesignStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/small-molecule/design"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an in-progress design run early
func (r *SmallMoleculeDesignService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeDesignStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/design/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A small molecule design pipeline run that generates novel molecules
type SmallMoleculeDesignGetResponse struct {
	// Unique SmDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzmol `json:"engine" default:"boltzmol"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                  `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeDesignGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeDesignGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                     `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeDesignGetResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                              `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeDesignGetResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                            `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		DataDeletedAt   respjson.Field
		Engine          respjson.Field
		EngineVersion   respjson.Field
		Error           respjson.Field
		Input           respjson.Field
		Livemode        respjson.Field
		Pipeline        respjson.Field
		PipelineVersion respjson.Field
		Progress        respjson.Field
		StartedAt       respjson.Field
		Status          respjson.Field
		StoppedAt       respjson.Field
		WorkspaceID     respjson.Field
		IdempotencyKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseError struct {
	// Machine-readable error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Additional field-level error details keyed by input path, when available.
	Details any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeDesignGetResponseInput struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeDesignGetResponseInputTarget `json:"target" api:"required"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace SmallMoleculeDesignGetResponseInputChemicalSpace `json:"chemical_space"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeDesignGetResponseInputMoleculeFilters `json:"molecule_filters"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMolecules    respjson.Field
		Target          respjson.Field
		ChemicalSpace   respjson.Field
		IdempotencyKey  respjson.Field
		MoleculeFilters respjson.Field
		WorkspaceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeDesignGetResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeDesignGetResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeDesignGetResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeDesignGetResponseInputTargetConstraintUnion `json:"constraints"`
	// Binding pocket residues, keyed by chain ID. Each key is a chain ID (e.g. "A")
	// and the value is an array of 0-indexed residue indices that define the binding
	// pocket on that chain. When provided, these residues guide pocket extraction and
	// add a derived pocket constraint during affinity predictions. That derived
	// constraint remains separate from any explicit pocket constraints in
	// target.constraints. When omitted, the model auto-detects the pocket.
	PocketResidues map[string][]int64 `json:"pocket_residues"`
	// Reference ligands as SMILES strings that help the model identify the binding
	// pocket. When omitted, a set of drug-like default ligands is used for pocket
	// detection.
	ReferenceLigands []string `json:"reference_ligands"`
	// Target is defined directly by protein sequences rather than a structure
	// template.
	//
	// Any of "no_template".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities         respjson.Field
		Bonds            respjson.Field
		Constraints      respjson.Field
		PocketResidues   respjson.Field
		ReferenceLigands respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignGetResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
	JSON     struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Bonds         respjson.Field
		Residues      respjson.Field
		raw           string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputTargetEntityUnion) AsSmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetEntityUnion) AsSmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignGetResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponseModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// Modification format. Only CCD polymer modifications are supported.
	Type constant.Ccd `json:"type" default:"ccd"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                              `json:"type" default:"glycan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bonds       respjson.Field
		ChainIDs    respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Ccd         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeDesignGetResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeDesignGetResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeDesignGetResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignGetResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeDesignGetResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputTargetBondAtom1Union) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignGetResponseInputTargetBondAtom1LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignGetResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeDesignGetResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignGetResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeDesignGetResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputTargetBondAtom2Union) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignGetResponseInputTargetBondAtom2LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignGetResponseInputTargetConstraintUnion contains all possible
// properties and values from
// [SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
	JSON   struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputTargetConstraintUnion) AsSmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetConstraintUnion) AsSmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputTargetConstraintUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignGetResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse struct {
	// Chain ID of the binder molecule
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// Binding pocket residues keyed by chain ID. Each key is a chain ID (e.g. "A") and
	// the value is an array of 0-indexed residue indices that define the pocket on
	// that chain.
	ContactResidues map[string][]int64 `json:"contact_residues" api:"required"`
	// Maximum allowed distance in Angstroms between binder and pocket residues.
	// Typical range: 4-8 A.
	MaxDistanceAngstrom float64         `json:"max_distance_angstrom" api:"required"`
	Type                constant.Pocket `json:"type" default:"pocket"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                        `json:"type" default:"contact"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxDistanceAngstrom respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type SmallMoleculeDesignGetResponseInputChemicalSpace string

const (
	SmallMoleculeDesignGetResponseInputChemicalSpaceEnamineReal SmallMoleculeDesignGetResponseInputChemicalSpace = "enamine_real"
	SmallMoleculeDesignGetResponseInputChemicalSpaceNone        SmallMoleculeDesignGetResponseInputChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeDesignGetResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeDesignGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion contains all
// possible properties and values from
// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                        `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
	Catalog string `json:"catalog"`
	JSON    struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		FractionCsp3         respjson.Field
		MolLogp              respjson.Field
		MolWt                respjson.Field
		NumAromaticRings     respjson.Field
		NumHAcceptors        respjson.Field
		NumHDonors           respjson.Field
		NumHeteroatoms       respjson.Field
		NumRings             respjson.Field
		NumRotatableBonds    respjson.Field
		Tpsa                 respjson.Field
		Patterns             respjson.Field
		Catalog              respjson.Field
		raw                  string
	} `json:"-"`
}

func (u SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
	// Maximum number of hydrogen bond acceptors. Lipinski threshold: 10
	MaxHba float64 `json:"max_hba" api:"required"`
	// Maximum number of hydrogen bond donors. Lipinski threshold: 5
	MaxHbd float64 `json:"max_hbd" api:"required"`
	// Maximum LogP. Lipinski threshold: 5
	MaxLogp float64 `json:"max_logp" api:"required"`
	// Maximum molecular weight (Da). Lipinski threshold: 500
	MaxMw float64                 `json:"max_mw" api:"required"`
	Type  constant.LipinskiFilter `json:"type" default:"lipinski_filter"`
	// If true, one rule violation is allowed (classic Rule of Five). Defaults to false
	// (all rules must pass).
	AllowSingleViolation bool `json:"allow_single_violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		FractionCsp3      respjson.Field
		MolLogp           respjson.Field
		MolWt             respjson.Field
		NumAromaticRings  respjson.Field
		NumHAcceptors     respjson.Field
		NumHDonors        respjson.Field
		NumHeteroatoms    respjson.Field
		NumRings          respjson.Field
		NumRotatableBonds respjson.Field
		Tpsa              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string                    `json:"patterns" api:"required"`
	Type     constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
	// Predefined SMARTS catalog to apply. PAINS, BRENK, ChEMBL, and NIH catalogs
	// reject known problematic substructures.
	//
	// Any of "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL",
	// "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica",
	// "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH".
	Catalog string                       `json:"catalog" api:"required"`
	Type    constant.SmartsCatalogFilter `json:"type" default:"smarts_catalog_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Catalog     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string                   `json:"patterns" api:"required"`
	Type     constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseProgress struct {
	// Number of molecules generated so far
	NumMoleculesGenerated int64 `json:"num_molecules_generated" api:"required"`
	// Total number of molecules requested
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesGenerated    respjson.Field
		TotalMoleculesToGenerate respjson.Field
		LatestResultID           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignGetResponseStatus string

const (
	SmallMoleculeDesignGetResponseStatusPending   SmallMoleculeDesignGetResponseStatus = "pending"
	SmallMoleculeDesignGetResponseStatusRunning   SmallMoleculeDesignGetResponseStatus = "running"
	SmallMoleculeDesignGetResponseStatusSucceeded SmallMoleculeDesignGetResponseStatus = "succeeded"
	SmallMoleculeDesignGetResponseStatusFailed    SmallMoleculeDesignGetResponseStatus = "failed"
	SmallMoleculeDesignGetResponseStatusStopped   SmallMoleculeDesignGetResponseStatus = "stopped"
)

// Summary of a small molecule design pipeline run (excludes input)
type SmallMoleculeDesignListResponse struct {
	// Unique SmDesignRunSummary identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzmol `json:"engine" default:"boltzmol"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                   `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeDesignListResponseError `json:"error" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                      `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeDesignListResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                               `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeDesignListResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                             `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		DataDeletedAt   respjson.Field
		Engine          respjson.Field
		EngineVersion   respjson.Field
		Error           respjson.Field
		Livemode        respjson.Field
		Pipeline        respjson.Field
		PipelineVersion respjson.Field
		Progress        respjson.Field
		StartedAt       respjson.Field
		Status          respjson.Field
		StoppedAt       respjson.Field
		WorkspaceID     respjson.Field
		IdempotencyKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResponseError struct {
	// Machine-readable error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Additional field-level error details keyed by input path, when available.
	Details any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResponseProgress struct {
	// Number of molecules generated so far
	NumMoleculesGenerated int64 `json:"num_molecules_generated" api:"required"`
	// Total number of molecules requested
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesGenerated    respjson.Field
		TotalMoleculesToGenerate respjson.Field
		LatestResultID           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResponseStatus string

const (
	SmallMoleculeDesignListResponseStatusPending   SmallMoleculeDesignListResponseStatus = "pending"
	SmallMoleculeDesignListResponseStatusRunning   SmallMoleculeDesignListResponseStatus = "running"
	SmallMoleculeDesignListResponseStatusSucceeded SmallMoleculeDesignListResponseStatus = "succeeded"
	SmallMoleculeDesignListResponseStatusFailed    SmallMoleculeDesignListResponseStatus = "failed"
	SmallMoleculeDesignListResponseStatusStopped   SmallMoleculeDesignListResponseStatus = "stopped"
)

type SmallMoleculeDesignDeleteDataResponse struct {
	// ID of the resource whose data was deleted
	ID          string `json:"id" api:"required"`
	DataDeleted bool   `json:"data_deleted" api:"required"`
	// When the data was deleted
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DataDeleted   respjson.Field
		DataDeletedAt respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimate response with monetary values encoded as decimal strings to preserve
// precision.
type SmallMoleculeDesignEstimateCostResponse struct {
	// Cost breakdown for the billed application.
	Breakdown  SmallMoleculeDesignEstimateCostResponseBreakdown `json:"breakdown" api:"required"`
	Disclaimer string                                           `json:"disclaimer" api:"required"`
	// Estimated total cost as a decimal string
	EstimatedCostUsd string `json:"estimated_cost_usd" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakdown        respjson.Field
		Disclaimer       respjson.Field
		EstimatedCostUsd respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost breakdown for the billed application.
type SmallMoleculeDesignEstimateCostResponseBreakdown struct {
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application SmallMoleculeDesignEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
	// Estimated cost per displayed unit as a decimal string, rounded up to 4 decimal
	// places. This may include token-size multipliers or generation overhead;
	// estimated_cost_usd is the authoritative total.
	CostPerUnitUsd string `json:"cost_per_unit_usd" api:"required"`
	// Number of billable units in the estimate. The unit depends on the endpoint:
	// samples for structure-and-binding, molecules for ADME, and requested proteins or
	// molecules for design/screen endpoints.
	NumUnits int64 `json:"num_units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Application    respjson.Field
		CostPerUnitUsd respjson.Field
		NumUnits       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignEstimateCostResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignEstimateCostResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignEstimateCostResponseBreakdownApplication string

const (
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationStructureAndBinding        SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "structure_and_binding"
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationSmallMoleculeDesign        SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "small_molecule_design"
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationSmallMoleculeLibraryScreen SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "small_molecule_library_screen"
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationProteinDesign              SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "protein_design"
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationProteinRedesign            SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "protein_redesign"
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationProteinLibraryScreen       SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "protein_library_screen"
	SmallMoleculeDesignEstimateCostResponseBreakdownApplicationAdme                       SmallMoleculeDesignEstimateCostResponseBreakdownApplication = "adme"
)

// A single designed small molecule result
type SmallMoleculeDesignListResultsResponse struct {
	// Unique result ID
	ID        string                                          `json:"id" api:"required"`
	Artifacts SmallMoleculeDesignListResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	// Scoring metrics for a designed small molecule
	Metrics SmallMoleculeDesignListResultsResponseMetrics `json:"metrics" api:"required"`
	// SMILES string of the designed molecule
	Smiles string `json:"smiles" api:"required"`
	// Tier 1 ADME summary values for this molecule.
	Adme SmallMoleculeDesignListResultsResponseAdme `json:"adme"`
	// Warnings about potential quality issues with this result.
	Warnings []SmallMoleculeDesignListResultsResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Metrics     respjson.Field
		Smiles      respjson.Field
		Adme        respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResultsResponseArtifacts struct {
	Archive         SmallMoleculeDesignListResultsResponseArtifactsArchive         `json:"archive" api:"required"`
	Structure       SmallMoleculeDesignListResultsResponseArtifactsStructure       `json:"structure" api:"required"`
	LigandStructure SmallMoleculeDesignListResultsResponseArtifactsLigandStructure `json:"ligand_structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archive         respjson.Field
		Structure       respjson.Field
		LigandStructure respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseArtifacts) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponseArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResultsResponseArtifactsArchive struct {
	// URL to download the file
	URL string `json:"url" api:"required" format:"uri"`
	// When the presigned URL expires
	URLExpiresAt time.Time `json:"url_expires_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseArtifactsArchive) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponseArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResultsResponseArtifactsStructure struct {
	// URL to download the file
	URL string `json:"url" api:"required" format:"uri"`
	// When the presigned URL expires
	URLExpiresAt time.Time `json:"url_expires_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseArtifactsStructure) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponseArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResultsResponseArtifactsLigandStructure struct {
	// URL to download the file
	URL string `json:"url" api:"required" format:"uri"`
	// When the presigned URL expires
	URLExpiresAt time.Time `json:"url_expires_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseArtifactsLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignListResultsResponseArtifactsLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scoring metrics for a designed small molecule
type SmallMoleculeDesignListResultsResponseMetrics struct {
	// Confidence that the molecule binds the target (0-1). Primary metric for hit
	// discovery.
	BindingConfidence float64 `json:"binding_confidence" api:"required"`
	// Interface pLDDT for the complex (0-1 float). Confidence at the binding
	// interface.
	ComplexIplddt float64 `json:"complex_iplddt" api:"required"`
	// pLDDT for the full complex (0-1 float).
	ComplexPlddt float64 `json:"complex_plddt" api:"required"`
	// Interface predicted TM score (0-1). Confidence in relative positioning of ligand
	// and protein.
	Iptm float64 `json:"iptm" api:"required"`
	// Binding strength ranking score for lead optimization. Higher values indicate
	// stronger predicted binding.
	OptimizationScore float64 `json:"optimization_score" api:"required"`
	// Predicted TM score (0-1). Global structure quality metric.
	Ptm float64 `json:"ptm" api:"required"`
	// Confidence in the predicted 3D structure (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence   respjson.Field
		ComplexIplddt       respjson.Field
		ComplexPlddt        respjson.Field
		Iptm                respjson.Field
		OptimizationScore   respjson.Field
		Ptm                 respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tier 1 ADME summary values for this molecule.
type SmallMoleculeDesignListResultsResponseAdme struct {
	// Lipophilicity score from the internal LogD prediction.
	Lipophilicity float64 `json:"lipophilicity" api:"required"`
	// Permeability score for this molecule.
	Permeability float64 `json:"permeability" api:"required"`
	// Solubility judgement for this molecule.
	//
	// Any of "high-confidence", "medium-confidence", "high-risk".
	Solubility SmallMoleculeDesignListResultsResponseAdmeSolubility `json:"solubility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lipophilicity respjson.Field
		Permeability  respjson.Field
		Solubility    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseAdme) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponseAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solubility judgement for this molecule.
type SmallMoleculeDesignListResultsResponseAdmeSolubility string

const (
	SmallMoleculeDesignListResultsResponseAdmeSolubilityHighConfidence   SmallMoleculeDesignListResultsResponseAdmeSolubility = "high-confidence"
	SmallMoleculeDesignListResultsResponseAdmeSolubilityMediumConfidence SmallMoleculeDesignListResultsResponseAdmeSolubility = "medium-confidence"
	SmallMoleculeDesignListResultsResponseAdmeSolubilityHighRisk         SmallMoleculeDesignListResultsResponseAdmeSolubility = "high-risk"
)

// A warning about a potential quality issue with a result
type SmallMoleculeDesignListResultsResponseWarning struct {
	// Machine-readable warning code (e.g. "low_confidence", "unusual_geometry")
	Code string `json:"code" api:"required"`
	// Human-readable description of the warning
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignListResultsResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignListResultsResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A small molecule design pipeline run that generates novel molecules
type SmallMoleculeDesignResumeResponse struct {
	// Unique SmDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzmol `json:"engine" default:"boltzmol"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                     `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeDesignResumeResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeDesignResumeResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                        `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeDesignResumeResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                 `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeDesignResumeResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                               `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		DataDeletedAt   respjson.Field
		Engine          respjson.Field
		EngineVersion   respjson.Field
		Error           respjson.Field
		Input           respjson.Field
		Livemode        respjson.Field
		Pipeline        respjson.Field
		PipelineVersion respjson.Field
		Progress        respjson.Field
		StartedAt       respjson.Field
		Status          respjson.Field
		StoppedAt       respjson.Field
		WorkspaceID     respjson.Field
		IdempotencyKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseError struct {
	// Machine-readable error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Additional field-level error details keyed by input path, when available.
	Details any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeDesignResumeResponseInput struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeDesignResumeResponseInputTarget `json:"target" api:"required"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace SmallMoleculeDesignResumeResponseInputChemicalSpace `json:"chemical_space"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeDesignResumeResponseInputMoleculeFilters `json:"molecule_filters"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMolecules    respjson.Field
		Target          respjson.Field
		ChemicalSpace   respjson.Field
		IdempotencyKey  respjson.Field
		MoleculeFilters respjson.Field
		WorkspaceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeDesignResumeResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeDesignResumeResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeDesignResumeResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeDesignResumeResponseInputTargetConstraintUnion `json:"constraints"`
	// Binding pocket residues, keyed by chain ID. Each key is a chain ID (e.g. "A")
	// and the value is an array of 0-indexed residue indices that define the binding
	// pocket on that chain. When provided, these residues guide pocket extraction and
	// add a derived pocket constraint during affinity predictions. That derived
	// constraint remains separate from any explicit pocket constraints in
	// target.constraints. When omitted, the model auto-detects the pocket.
	PocketResidues map[string][]int64 `json:"pocket_residues"`
	// Reference ligands as SMILES strings that help the model identify the binding
	// pocket. When omitted, a set of drug-like default ligands is used for pocket
	// detection.
	ReferenceLigands []string `json:"reference_ligands"`
	// Target is defined directly by protein sequences rather than a structure
	// template.
	//
	// Any of "no_template".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities         respjson.Field
		Bonds            respjson.Field
		Constraints      respjson.Field
		PocketResidues   respjson.Field
		ReferenceLigands respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignResumeResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
	JSON     struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Bonds         respjson.Field
		Residues      respjson.Field
		raw           string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputTargetEntityUnion) AsSmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetEntityUnion) AsSmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignResumeResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponseModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// Modification format. Only CCD polymer modifications are supported.
	Type constant.Ccd `json:"type" default:"ccd"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                 `json:"type" default:"glycan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bonds       respjson.Field
		ChainIDs    respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Ccd         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeDesignResumeResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom1LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignResumeResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeDesignResumeResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignResumeResponseInputTargetBondAtom2LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignResumeResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
	JSON   struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputTargetConstraintUnion) AsSmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetConstraintUnion) AsSmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse struct {
	// Chain ID of the binder molecule
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// Binding pocket residues keyed by chain ID. Each key is a chain ID (e.g. "A") and
	// the value is an array of 0-indexed residue indices that define the pocket on
	// that chain.
	ContactResidues map[string][]int64 `json:"contact_residues" api:"required"`
	// Maximum allowed distance in Angstroms between binder and pocket residues.
	// Typical range: 4-8 A.
	MaxDistanceAngstrom float64         `json:"max_distance_angstrom" api:"required"`
	Type                constant.Pocket `json:"type" default:"pocket"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                           `json:"type" default:"contact"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxDistanceAngstrom respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type SmallMoleculeDesignResumeResponseInputChemicalSpace string

const (
	SmallMoleculeDesignResumeResponseInputChemicalSpaceEnamineReal SmallMoleculeDesignResumeResponseInputChemicalSpace = "enamine_real"
	SmallMoleculeDesignResumeResponseInputChemicalSpaceNone        SmallMoleculeDesignResumeResponseInputChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeDesignResumeResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeDesignResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                           `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
	Catalog string `json:"catalog"`
	JSON    struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		FractionCsp3         respjson.Field
		MolLogp              respjson.Field
		MolWt                respjson.Field
		NumAromaticRings     respjson.Field
		NumHAcceptors        respjson.Field
		NumHDonors           respjson.Field
		NumHeteroatoms       respjson.Field
		NumRings             respjson.Field
		NumRotatableBonds    respjson.Field
		Tpsa                 respjson.Field
		Patterns             respjson.Field
		Catalog              respjson.Field
		raw                  string
	} `json:"-"`
}

func (u SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
	// Maximum number of hydrogen bond acceptors. Lipinski threshold: 10
	MaxHba float64 `json:"max_hba" api:"required"`
	// Maximum number of hydrogen bond donors. Lipinski threshold: 5
	MaxHbd float64 `json:"max_hbd" api:"required"`
	// Maximum LogP. Lipinski threshold: 5
	MaxLogp float64 `json:"max_logp" api:"required"`
	// Maximum molecular weight (Da). Lipinski threshold: 500
	MaxMw float64                 `json:"max_mw" api:"required"`
	Type  constant.LipinskiFilter `json:"type" default:"lipinski_filter"`
	// If true, one rule violation is allowed (classic Rule of Five). Defaults to false
	// (all rules must pass).
	AllowSingleViolation bool `json:"allow_single_violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		FractionCsp3      respjson.Field
		MolLogp           respjson.Field
		MolWt             respjson.Field
		NumAromaticRings  respjson.Field
		NumHAcceptors     respjson.Field
		NumHDonors        respjson.Field
		NumHeteroatoms    respjson.Field
		NumRings          respjson.Field
		NumRotatableBonds respjson.Field
		Tpsa              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string                    `json:"patterns" api:"required"`
	Type     constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
	// Predefined SMARTS catalog to apply. PAINS, BRENK, ChEMBL, and NIH catalogs
	// reject known problematic substructures.
	//
	// Any of "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL",
	// "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica",
	// "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH".
	Catalog string                       `json:"catalog" api:"required"`
	Type    constant.SmartsCatalogFilter `json:"type" default:"smarts_catalog_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Catalog     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string                   `json:"patterns" api:"required"`
	Type     constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseProgress struct {
	// Number of molecules generated so far
	NumMoleculesGenerated int64 `json:"num_molecules_generated" api:"required"`
	// Total number of molecules requested
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesGenerated    respjson.Field
		TotalMoleculesToGenerate respjson.Field
		LatestResultID           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignResumeResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignResumeResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignResumeResponseStatus string

const (
	SmallMoleculeDesignResumeResponseStatusPending   SmallMoleculeDesignResumeResponseStatus = "pending"
	SmallMoleculeDesignResumeResponseStatusRunning   SmallMoleculeDesignResumeResponseStatus = "running"
	SmallMoleculeDesignResumeResponseStatusSucceeded SmallMoleculeDesignResumeResponseStatus = "succeeded"
	SmallMoleculeDesignResumeResponseStatusFailed    SmallMoleculeDesignResumeResponseStatus = "failed"
	SmallMoleculeDesignResumeResponseStatusStopped   SmallMoleculeDesignResumeResponseStatus = "stopped"
)

// A small molecule design pipeline run that generates novel molecules
type SmallMoleculeDesignStartResponse struct {
	// Unique SmDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzmol `json:"engine" default:"boltzmol"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                    `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeDesignStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeDesignStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                       `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeDesignStartResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeDesignStartResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                              `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		DataDeletedAt   respjson.Field
		Engine          respjson.Field
		EngineVersion   respjson.Field
		Error           respjson.Field
		Input           respjson.Field
		Livemode        respjson.Field
		Pipeline        respjson.Field
		PipelineVersion respjson.Field
		Progress        respjson.Field
		StartedAt       respjson.Field
		Status          respjson.Field
		StoppedAt       respjson.Field
		WorkspaceID     respjson.Field
		IdempotencyKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseError struct {
	// Machine-readable error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Additional field-level error details keyed by input path, when available.
	Details any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeDesignStartResponseInput struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeDesignStartResponseInputTarget `json:"target" api:"required"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace SmallMoleculeDesignStartResponseInputChemicalSpace `json:"chemical_space"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeDesignStartResponseInputMoleculeFilters `json:"molecule_filters"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMolecules    respjson.Field
		Target          respjson.Field
		ChemicalSpace   respjson.Field
		IdempotencyKey  respjson.Field
		MoleculeFilters respjson.Field
		WorkspaceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeDesignStartResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeDesignStartResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeDesignStartResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeDesignStartResponseInputTargetConstraintUnion `json:"constraints"`
	// Binding pocket residues, keyed by chain ID. Each key is a chain ID (e.g. "A")
	// and the value is an array of 0-indexed residue indices that define the binding
	// pocket on that chain. When provided, these residues guide pocket extraction and
	// add a derived pocket constraint during affinity predictions. That derived
	// constraint remains separate from any explicit pocket constraints in
	// target.constraints. When omitted, the model auto-detects the pocket.
	PocketResidues map[string][]int64 `json:"pocket_residues"`
	// Reference ligands as SMILES strings that help the model identify the binding
	// pocket. When omitted, a set of drug-like default ligands is used for pocket
	// detection.
	ReferenceLigands []string `json:"reference_ligands"`
	// Target is defined directly by protein sequences rather than a structure
	// template.
	//
	// Any of "no_template".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities         respjson.Field
		Bonds            respjson.Field
		Constraints      respjson.Field
		PocketResidues   respjson.Field
		ReferenceLigands respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStartResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
	JSON     struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Bonds         respjson.Field
		Residues      respjson.Field
		raw           string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputTargetEntityUnion) AsSmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetEntityUnion) AsSmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignStartResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponseModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// Modification format. Only CCD polymer modifications are supported.
	Type constant.Ccd `json:"type" default:"ccd"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                `json:"type" default:"glycan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bonds       respjson.Field
		ChainIDs    respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Ccd         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeDesignStartResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeDesignStartResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeDesignStartResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStartResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeDesignStartResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignStartResponseInputTargetBondAtom1LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStartResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeDesignStartResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStartResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeDesignStartResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignStartResponseInputTargetBondAtom2LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStartResponseInputTargetConstraintUnion contains all possible
// properties and values from
// [SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
	JSON   struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputTargetConstraintUnion) AsSmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetConstraintUnion) AsSmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStartResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse struct {
	// Chain ID of the binder molecule
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// Binding pocket residues keyed by chain ID. Each key is a chain ID (e.g. "A") and
	// the value is an array of 0-indexed residue indices that define the pocket on
	// that chain.
	ContactResidues map[string][]int64 `json:"contact_residues" api:"required"`
	// Maximum allowed distance in Angstroms between binder and pocket residues.
	// Typical range: 4-8 A.
	MaxDistanceAngstrom float64         `json:"max_distance_angstrom" api:"required"`
	Type                constant.Pocket `json:"type" default:"pocket"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                          `json:"type" default:"contact"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxDistanceAngstrom respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type SmallMoleculeDesignStartResponseInputChemicalSpace string

const (
	SmallMoleculeDesignStartResponseInputChemicalSpaceEnamineReal SmallMoleculeDesignStartResponseInputChemicalSpace = "enamine_real"
	SmallMoleculeDesignStartResponseInputChemicalSpaceNone        SmallMoleculeDesignStartResponseInputChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeDesignStartResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeDesignStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                          `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
	Catalog string `json:"catalog"`
	JSON    struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		FractionCsp3         respjson.Field
		MolLogp              respjson.Field
		MolWt                respjson.Field
		NumAromaticRings     respjson.Field
		NumHAcceptors        respjson.Field
		NumHDonors           respjson.Field
		NumHeteroatoms       respjson.Field
		NumRings             respjson.Field
		NumRotatableBonds    respjson.Field
		Tpsa                 respjson.Field
		Patterns             respjson.Field
		Catalog              respjson.Field
		raw                  string
	} `json:"-"`
}

func (u SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
	// Maximum number of hydrogen bond acceptors. Lipinski threshold: 10
	MaxHba float64 `json:"max_hba" api:"required"`
	// Maximum number of hydrogen bond donors. Lipinski threshold: 5
	MaxHbd float64 `json:"max_hbd" api:"required"`
	// Maximum LogP. Lipinski threshold: 5
	MaxLogp float64 `json:"max_logp" api:"required"`
	// Maximum molecular weight (Da). Lipinski threshold: 500
	MaxMw float64                 `json:"max_mw" api:"required"`
	Type  constant.LipinskiFilter `json:"type" default:"lipinski_filter"`
	// If true, one rule violation is allowed (classic Rule of Five). Defaults to false
	// (all rules must pass).
	AllowSingleViolation bool `json:"allow_single_violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		FractionCsp3      respjson.Field
		MolLogp           respjson.Field
		MolWt             respjson.Field
		NumAromaticRings  respjson.Field
		NumHAcceptors     respjson.Field
		NumHDonors        respjson.Field
		NumHeteroatoms    respjson.Field
		NumRings          respjson.Field
		NumRotatableBonds respjson.Field
		Tpsa              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string                    `json:"patterns" api:"required"`
	Type     constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
	// Predefined SMARTS catalog to apply. PAINS, BRENK, ChEMBL, and NIH catalogs
	// reject known problematic substructures.
	//
	// Any of "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL",
	// "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica",
	// "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH".
	Catalog string                       `json:"catalog" api:"required"`
	Type    constant.SmartsCatalogFilter `json:"type" default:"smarts_catalog_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Catalog     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string                   `json:"patterns" api:"required"`
	Type     constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseProgress struct {
	// Number of molecules generated so far
	NumMoleculesGenerated int64 `json:"num_molecules_generated" api:"required"`
	// Total number of molecules requested
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesGenerated    respjson.Field
		TotalMoleculesToGenerate respjson.Field
		LatestResultID           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStartResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStartResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStartResponseStatus string

const (
	SmallMoleculeDesignStartResponseStatusPending   SmallMoleculeDesignStartResponseStatus = "pending"
	SmallMoleculeDesignStartResponseStatusRunning   SmallMoleculeDesignStartResponseStatus = "running"
	SmallMoleculeDesignStartResponseStatusSucceeded SmallMoleculeDesignStartResponseStatus = "succeeded"
	SmallMoleculeDesignStartResponseStatusFailed    SmallMoleculeDesignStartResponseStatus = "failed"
	SmallMoleculeDesignStartResponseStatusStopped   SmallMoleculeDesignStartResponseStatus = "stopped"
)

// A small molecule design pipeline run that generates novel molecules
type SmallMoleculeDesignStopResponse struct {
	// Unique SmDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzmol `json:"engine" default:"boltzmol"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                   `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeDesignStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeDesignStopResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                      `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeDesignStopResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                               `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeDesignStopResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                             `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		DataDeletedAt   respjson.Field
		Engine          respjson.Field
		EngineVersion   respjson.Field
		Error           respjson.Field
		Input           respjson.Field
		Livemode        respjson.Field
		Pipeline        respjson.Field
		PipelineVersion respjson.Field
		Progress        respjson.Field
		StartedAt       respjson.Field
		Status          respjson.Field
		StoppedAt       respjson.Field
		WorkspaceID     respjson.Field
		IdempotencyKey  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseError struct {
	// Machine-readable error code
	Code string `json:"code" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Additional field-level error details keyed by input path, when available.
	Details any `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeDesignStopResponseInput struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeDesignStopResponseInputTarget `json:"target" api:"required"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace SmallMoleculeDesignStopResponseInputChemicalSpace `json:"chemical_space"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeDesignStopResponseInputMoleculeFilters `json:"molecule_filters"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMolecules    respjson.Field
		Target          respjson.Field
		ChemicalSpace   respjson.Field
		IdempotencyKey  respjson.Field
		MoleculeFilters respjson.Field
		WorkspaceID     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeDesignStopResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeDesignStopResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeDesignStopResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeDesignStopResponseInputTargetConstraintUnion `json:"constraints"`
	// Binding pocket residues, keyed by chain ID. Each key is a chain ID (e.g. "A")
	// and the value is an array of 0-indexed residue indices that define the binding
	// pocket on that chain. When provided, these residues guide pocket extraction and
	// add a derived pocket constraint during affinity predictions. That derived
	// constraint remains separate from any explicit pocket constraints in
	// target.constraints. When omitted, the model auto-detects the pocket.
	PocketResidues map[string][]int64 `json:"pocket_residues"`
	// Reference ligands as SMILES strings that help the model identify the binding
	// pocket. When omitted, a set of drug-like default ligands is used for pocket
	// detection.
	ReferenceLigands []string `json:"reference_ligands"`
	// Target is defined directly by protein sequences rather than a structure
	// template.
	//
	// Any of "no_template".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities         respjson.Field
		Bonds            respjson.Field
		Constraints      respjson.Field
		PocketResidues   respjson.Field
		ReferenceLigands respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStopResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
	JSON     struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Bonds         respjson.Field
		Residues      respjson.Field
		raw           string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputTargetEntityUnion) AsSmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetEntityUnion) AsSmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignStopResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponseModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// Modification format. Only CCD polymer modifications are supported.
	Type constant.Ccd `json:"type" default:"ccd"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                               `json:"type" default:"glycan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bonds       respjson.Field
		ChainIDs    respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ResidueID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Ccd         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeDesignStopResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeDesignStopResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeDesignStopResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStopResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeDesignStopResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom1Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputTargetBondAtom1Union) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignStopResponseInputTargetBondAtom1LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStopResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeDesignStopResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse].
	AtomMap int64 `json:"atom_map"`
	JSON    struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomID       respjson.Field
		ResidueID    respjson.Field
		AtomMap      respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetBondAtom2Union) AsSmallMoleculeDesignStopResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeDesignStopResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputTargetBondAtom2Union) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                `json:"residue_index" api:"required"`
	Type         constant.PolymerAtom `json:"type" default:"polymer_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string           `json:"residue_id" api:"required"`
	Type      constant.CcdAtom `json:"type" default:"ccd_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomID      respjson.Field
		ChainID     respjson.Field
		ResidueID   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.SmilesAtom `json:"type" default:"smiles_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomMap     respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeDesignStopResponseInputTargetBondAtom2LigandAtomResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string              `json:"chain_id" api:"required"`
	Type    constant.LigandAtom `json:"type" default:"ligand_atom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStopResponseInputTargetConstraintUnion contains all possible
// properties and values from
// [SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
	JSON   struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		raw                 string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputTargetConstraintUnion) AsSmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetConstraintUnion) AsSmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStopResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse struct {
	// Chain ID of the binder molecule
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// Binding pocket residues keyed by chain ID. Each key is a chain ID (e.g. "A") and
	// the value is an array of 0-indexed residue indices that define the pocket on
	// that chain.
	ContactResidues map[string][]int64 `json:"contact_residues" api:"required"`
	// Maximum allowed distance in Angstroms between binder and pocket residues.
	// Typical range: 4-8 A.
	MaxDistanceAngstrom float64         `json:"max_distance_angstrom" api:"required"`
	Type                constant.Pocket `json:"type" default:"pocket"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainID       respjson.Field
		ContactResidues     respjson.Field
		MaxDistanceAngstrom respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                         `json:"type" default:"contact"`
	// Whether to force the constraint
	Force bool `json:"force"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxDistanceAngstrom respjson.Field
		Token1              respjson.Field
		Token2              respjson.Field
		Type                respjson.Field
		Force               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64                   `json:"residue_index" api:"required"`
	Type         constant.PolymerContact `json:"type" default:"polymer_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string                 `json:"chain_id" api:"required"`
	Type    constant.LigandContact `json:"type" default:"ligand_contact"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AtomName    respjson.Field
		ChainID     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type SmallMoleculeDesignStopResponseInputChemicalSpace string

const (
	SmallMoleculeDesignStopResponseInputChemicalSpaceEnamineReal SmallMoleculeDesignStopResponseInputChemicalSpace = "enamine_real"
	SmallMoleculeDesignStopResponseInputChemicalSpaceNone        SmallMoleculeDesignStopResponseInputChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeDesignStopResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeDesignStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                         `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
	Catalog string `json:"catalog"`
	JSON    struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		FractionCsp3         respjson.Field
		MolLogp              respjson.Field
		MolWt                respjson.Field
		NumAromaticRings     respjson.Field
		NumHAcceptors        respjson.Field
		NumHDonors           respjson.Field
		NumHeteroatoms       respjson.Field
		NumRings             respjson.Field
		NumRotatableBonds    respjson.Field
		Tpsa                 respjson.Field
		Patterns             respjson.Field
		Catalog              respjson.Field
		raw                  string
	} `json:"-"`
}

func (u SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
	// Maximum number of hydrogen bond acceptors. Lipinski threshold: 10
	MaxHba float64 `json:"max_hba" api:"required"`
	// Maximum number of hydrogen bond donors. Lipinski threshold: 5
	MaxHbd float64 `json:"max_hbd" api:"required"`
	// Maximum LogP. Lipinski threshold: 5
	MaxLogp float64 `json:"max_logp" api:"required"`
	// Maximum molecular weight (Da). Lipinski threshold: 500
	MaxMw float64                 `json:"max_mw" api:"required"`
	Type  constant.LipinskiFilter `json:"type" default:"lipinski_filter"`
	// If true, one rule violation is allowed (classic Rule of Five). Defaults to false
	// (all rules must pass).
	AllowSingleViolation bool `json:"allow_single_violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxHba               respjson.Field
		MaxHbd               respjson.Field
		MaxLogp              respjson.Field
		MaxMw                respjson.Field
		Type                 respjson.Field
		AllowSingleViolation respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type              respjson.Field
		FractionCsp3      respjson.Field
		MolLogp           respjson.Field
		MolWt             respjson.Field
		NumAromaticRings  respjson.Field
		NumHAcceptors     respjson.Field
		NumHDonors        respjson.Field
		NumHeteroatoms    respjson.Field
		NumRings          respjson.Field
		NumRotatableBonds respjson.Field
		Tpsa              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
	// Maximum allowed value (inclusive)
	Max float64 `json:"max"`
	// Minimum allowed value (inclusive)
	Min float64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string                    `json:"patterns" api:"required"`
	Type     constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
	// Predefined SMARTS catalog to apply. PAINS, BRENK, ChEMBL, and NIH catalogs
	// reject known problematic substructures.
	//
	// Any of "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL",
	// "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica",
	// "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH".
	Catalog string                       `json:"catalog" api:"required"`
	Type    constant.SmartsCatalogFilter `json:"type" default:"smarts_catalog_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Catalog     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string                   `json:"patterns" api:"required"`
	Type     constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeDesignStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseProgress struct {
	// Number of molecules generated so far
	NumMoleculesGenerated int64 `json:"num_molecules_generated" api:"required"`
	// Total number of molecules requested
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesGenerated    respjson.Field
		TotalMoleculesToGenerate respjson.Field
		LatestResultID           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeDesignStopResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeDesignStopResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignStopResponseStatus string

const (
	SmallMoleculeDesignStopResponseStatusPending   SmallMoleculeDesignStopResponseStatus = "pending"
	SmallMoleculeDesignStopResponseStatusRunning   SmallMoleculeDesignStopResponseStatus = "running"
	SmallMoleculeDesignStopResponseStatusSucceeded SmallMoleculeDesignStopResponseStatus = "succeeded"
	SmallMoleculeDesignStopResponseStatusFailed    SmallMoleculeDesignStopResponseStatus = "failed"
	SmallMoleculeDesignStopResponseStatusStopped   SmallMoleculeDesignStopResponseStatus = "stopped"
)

type SmallMoleculeDesignGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SmallMoleculeDesignGetParams]'s query parameters as
// `url.Values`.
func (r SmallMoleculeDesignGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeDesignListParams struct {
	// Return results after this ID
	AfterID param.Opt[string] `query:"after_id,omitzero" json:"-"`
	// Return results before this ID
	BeforeID param.Opt[string] `query:"before_id,omitzero" json:"-"`
	// Max items to return. Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by workspace ID. Only used with admin API keys. If not provided, defaults
	// to the workspace associated with the API key, or the default workspace for admin
	// keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SmallMoleculeDesignListParams]'s query parameters as
// `url.Values`.
func (r SmallMoleculeDesignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeDesignEstimateCostParams struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeDesignEstimateCostParamsTarget `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace SmallMoleculeDesignEstimateCostParamsChemicalSpace `json:"chemical_space,omitzero"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeDesignEstimateCostParamsMoleculeFilters `json:"molecule_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
//
// The property Entities is required.
type SmallMoleculeDesignEstimateCostParamsTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeDesignEstimateCostParamsTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeDesignEstimateCostParamsTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion `json:"constraints,omitzero"`
	// Binding pocket residues, keyed by chain ID. Each key is a chain ID (e.g. "A")
	// and the value is an array of 0-indexed residue indices that define the binding
	// pocket on that chain. When provided, these residues guide pocket extraction and
	// add a derived pocket constraint during affinity predictions. That derived
	// constraint remains separate from any explicit pocket constraints in
	// target.constraints. When omitted, the model auto-detects the pocket.
	PocketResidues map[string][]int64 `json:"pocket_residues,omitzero"`
	// Reference ligands as SMILES strings that help the model identify the binding
	// pocket. When omitted, a set of drug-like default ligands is used for pocket
	// detection.
	ReferenceLigands []string `json:"reference_ligands,omitzero"`
	// Target is defined directly by protein sequences rather than a structure
	// template.
	//
	// Any of "no_template".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeDesignEstimateCostParamsTarget](
		"type", "no_template",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsTargetEntityUnion struct {
	OfSmallMoleculeDesignEstimateCostsTargetEntityProteinEntity *SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntity `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetEntityGlycanEntity  *SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsTargetEntityProteinEntity, u.OfSmallMoleculeDesignEstimateCostsTargetEntityGlycanEntity)
}
func (u *SmallMoleculeDesignEstimateCostParamsTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// Modification format. Only CCD polymer modifications are supported.
	//
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBond struct {
	Atom1 SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeDesignEstimateCostParamsTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union struct {
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom1PolymerAtom *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom1CcdAtom     *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom1SmilesAtom  *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom1LigandAtom  *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom1PolymerAtom, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom1CcdAtom, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom1SmilesAtom, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom1LigandAtom)
}
func (u *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom1PolymerAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_atom".
	Type constant.PolymerAtom `json:"type" default:"polymer_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom1CcdAtom struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string `json:"residue_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd_atom".
	Type constant.CcdAtom `json:"type" default:"ccd_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union struct {
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom2PolymerAtom *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom2CcdAtom     *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom2SmilesAtom  *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetBondAtom2LigandAtom  *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom2PolymerAtom, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom2CcdAtom, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom2SmilesAtom, u.OfSmallMoleculeDesignEstimateCostsTargetBondAtom2LigandAtom)
}
func (u *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom2PolymerAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_atom".
	Type constant.PolymerAtom `json:"type" default:"polymer_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom2CcdAtom struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string `json:"residue_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd_atom".
	Type constant.CcdAtom `json:"type" default:"ccd_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion struct {
	OfSmallMoleculeDesignEstimateCostsTargetConstraintPocketConstraint  *SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraint *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsTargetConstraintPocketConstraint, u.OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraint)
}
func (u *SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint struct {
	// Chain ID of the binder molecule
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// Binding pocket residues keyed by chain ID. Each key is a chain ID (e.g. "A") and
	// the value is an array of 0-indexed residue indices that define the pocket on
	// that chain.
	ContactResidues map[string][]int64 `json:"contact_residues,omitzero" api:"required"`
	// Maximum allowed distance in Angstroms between binder and pocket residues.
	// Typical range: 4-8 A.
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "pocket".
	Type constant.Pocket `json:"type" default:"pocket"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1Union struct {
	OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken1PolymerContactToken *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken1LigandContactToken  *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken1PolymerContactToken, u.OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2Union struct {
	OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken2PolymerContactToken *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken2LigandContactToken  *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken2PolymerContactToken, u.OfSmallMoleculeDesignEstimateCostsTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type SmallMoleculeDesignEstimateCostParamsChemicalSpace string

const (
	SmallMoleculeDesignEstimateCostParamsChemicalSpaceEnamineReal SmallMoleculeDesignEstimateCostParamsChemicalSpace = "enamine_real"
	SmallMoleculeDesignEstimateCostParamsChemicalSpaceNone        SmallMoleculeDesignEstimateCostParamsChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeDesignEstimateCostParamsMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level,omitzero"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion `json:"custom_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFilters) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion struct {
	OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter        *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter        `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterRdkitDescriptorFilter *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterSmartsCustomFilter    *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter    `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterSmartsCatalogFilter   *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter   `json:",omitzero,inline"`
	OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterSmilesRegexFilter     *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter     `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter,
		u.OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterRdkitDescriptorFilter,
		u.OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterSmartsCustomFilter,
		u.OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterSmartsCatalogFilter,
		u.OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterSmilesRegexFilter)
}
func (u *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
//
// The properties MaxHba, MaxHbd, MaxLogp, MaxMw, Type are required.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter struct {
	// Maximum number of hydrogen bond acceptors. Lipinski threshold: 10
	MaxHba float64 `json:"max_hba" api:"required"`
	// Maximum number of hydrogen bond donors. Lipinski threshold: 5
	MaxHbd float64 `json:"max_hbd" api:"required"`
	// Maximum LogP. Lipinski threshold: 5
	MaxLogp float64 `json:"max_logp" api:"required"`
	// Maximum molecular weight (Da). Lipinski threshold: 500
	MaxMw float64 `json:"max_mw" api:"required"`
	// If true, one rule violation is allowed (classic Rule of Five). Defaults to false
	// (all rules must pass).
	AllowSingleViolation param.Opt[bool] `json:"allow_single_violation,omitzero"`
	// This field can be elided, and will marshal its zero value as "lipinski_filter".
	Type constant.LipinskiFilter `json:"type" default:"lipinski_filter"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
//
// The property Type is required.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter struct {
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 `json:"fraction_csp3,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp `json:"mol_logp,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt `json:"mol_wt,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings `json:"num_aromatic_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors `json:"num_h_acceptors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors `json:"num_h_donors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms `json:"num_heteroatoms,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings `json:"num_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds `json:"num_rotatable_bonds,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa `json:"tpsa,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "rdkit_descriptor_filter".
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
//
// The properties Patterns, Type are required.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_custom_filter".
	Type constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
//
// The properties Catalog, Type are required.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter struct {
	// Predefined SMARTS catalog to apply. PAINS, BRENK, ChEMBL, and NIH catalogs
	// reject known problematic substructures.
	//
	// Any of "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL",
	// "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica",
	// "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH".
	Catalog string `json:"catalog,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_catalog_filter".
	Type constant.SmartsCatalogFilter `json:"type" default:"smarts_catalog_filter"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter](
		"catalog", "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL", "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica", "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH",
	)
}

// Filter molecules by regex patterns on their SMILES representation.
//
// The properties Patterns, Type are required.
type SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smiles_regex_filter".
	Type constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	paramObj
}

func (r SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeDesignListResultsParams struct {
	// Return results after this ID
	AfterID param.Opt[string] `query:"after_id,omitzero" json:"-"`
	// Return results before this ID
	BeforeID param.Opt[string] `query:"before_id,omitzero" json:"-"`
	// Comma-separated list of result IDs to filter by (max 200). Only results whose ID
	// matches one of these is returned; missing IDs are silently skipped. Composes
	// with `limit`, `after_id`, and `before_id` — the filter is applied before
	// pagination.
	IDs param.Opt[string] `query:"ids,omitzero" json:"-"`
	// Max results to return. Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SmallMoleculeDesignListResultsParams]'s query parameters as
// `url.Values`.
func (r SmallMoleculeDesignListResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeDesignStartParams struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeDesignStartParamsTarget `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace SmallMoleculeDesignStartParamsChemicalSpace `json:"chemical_space,omitzero"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeDesignStartParamsMoleculeFilters `json:"molecule_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParams) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
//
// The property Entities is required.
type SmallMoleculeDesignStartParamsTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeDesignStartParamsTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeDesignStartParamsTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeDesignStartParamsTargetConstraintUnion `json:"constraints,omitzero"`
	// Binding pocket residues, keyed by chain ID. Each key is a chain ID (e.g. "A")
	// and the value is an array of 0-indexed residue indices that define the binding
	// pocket on that chain. When provided, these residues guide pocket extraction and
	// add a derived pocket constraint during affinity predictions. That derived
	// constraint remains separate from any explicit pocket constraints in
	// target.constraints. When omitted, the model auto-detects the pocket.
	PocketResidues map[string][]int64 `json:"pocket_residues,omitzero"`
	// Reference ligands as SMILES strings that help the model identify the binding
	// pocket. When omitted, a set of drug-like default ligands is used for pocket
	// detection.
	ReferenceLigands []string `json:"reference_ligands,omitzero"`
	// Target is defined directly by protein sequences rather than a structure
	// template.
	//
	// Any of "no_template".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeDesignStartParamsTarget](
		"type", "no_template",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsTargetEntityUnion struct {
	OfSmallMoleculeDesignStartsTargetEntityProteinEntity *SmallMoleculeDesignStartParamsTargetEntityProteinEntity `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetEntityGlycanEntity  *SmallMoleculeDesignStartParamsTargetEntityGlycanEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsTargetEntityProteinEntity, u.OfSmallMoleculeDesignStartsTargetEntityGlycanEntity)
}
func (u *SmallMoleculeDesignStartParamsTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type SmallMoleculeDesignStartParamsTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeDesignStartParamsTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type SmallMoleculeDesignStartParamsTargetEntityProteinEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// Modification format. Only CCD polymer modifications are supported.
	//
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type SmallMoleculeDesignStartParamsTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeDesignStartParamsTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBond struct {
	Atom1 SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type SmallMoleculeDesignStartParamsTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeDesignStartParamsTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeDesignStartParamsTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeDesignStartParamsTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsTargetBondAtom1Union struct {
	OfSmallMoleculeDesignStartsTargetBondAtom1PolymerAtom *SmallMoleculeDesignStartParamsTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetBondAtom1CcdAtom     *SmallMoleculeDesignStartParamsTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetBondAtom1SmilesAtom  *SmallMoleculeDesignStartParamsTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetBondAtom1LigandAtom  *SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsTargetBondAtom1PolymerAtom, u.OfSmallMoleculeDesignStartsTargetBondAtom1CcdAtom, u.OfSmallMoleculeDesignStartsTargetBondAtom1SmilesAtom, u.OfSmallMoleculeDesignStartsTargetBondAtom1LigandAtom)
}
func (u *SmallMoleculeDesignStartParamsTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom1PolymerAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_atom".
	Type constant.PolymerAtom `json:"type" default:"polymer_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom1CcdAtom struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string `json:"residue_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd_atom".
	Type constant.CcdAtom `json:"type" default:"ccd_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsTargetBondAtom2Union struct {
	OfSmallMoleculeDesignStartsTargetBondAtom2PolymerAtom *SmallMoleculeDesignStartParamsTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetBondAtom2CcdAtom     *SmallMoleculeDesignStartParamsTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetBondAtom2SmilesAtom  *SmallMoleculeDesignStartParamsTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetBondAtom2LigandAtom  *SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsTargetBondAtom2PolymerAtom, u.OfSmallMoleculeDesignStartsTargetBondAtom2CcdAtom, u.OfSmallMoleculeDesignStartsTargetBondAtom2SmilesAtom, u.OfSmallMoleculeDesignStartsTargetBondAtom2LigandAtom)
}
func (u *SmallMoleculeDesignStartParamsTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom2PolymerAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB)
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_atom".
	Type constant.PolymerAtom `json:"type" default:"polymer_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom2CcdAtom struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Chain ID containing the CCD residue
	ChainID string `json:"chain_id" api:"required"`
	// Request-local residue ID declared by the graph entity
	ResidueID string `json:"residue_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd_atom".
	Type constant.CcdAtom `json:"type" default:"ccd_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsTargetConstraintUnion struct {
	OfSmallMoleculeDesignStartsTargetConstraintPocketConstraint  *SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetConstraintContactConstraint *SmallMoleculeDesignStartParamsTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsTargetConstraintPocketConstraint, u.OfSmallMoleculeDesignStartsTargetConstraintContactConstraint)
}
func (u *SmallMoleculeDesignStartParamsTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint struct {
	// Chain ID of the binder molecule
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// Binding pocket residues keyed by chain ID. Each key is a chain ID (e.g. "A") and
	// the value is an array of 0-indexed residue indices that define the pocket on
	// that chain.
	ContactResidues map[string][]int64 `json:"contact_residues,omitzero" api:"required"`
	// Maximum allowed distance in Angstroms between binder and pocket residues.
	// Typical range: 4-8 A.
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "pocket".
	Type constant.Pocket `json:"type" default:"pocket"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1Union struct {
	OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken1PolymerContactToken *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken1LigandContactToken  *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken1PolymerContactToken, u.OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2Union struct {
	OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken2PolymerContactToken *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken2LigandContactToken  *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken2PolymerContactToken, u.OfSmallMoleculeDesignStartsTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. For ligand_ccd, use the standardized CIF atom name. For
	// ligand_smiles, explicitly label the atom with numeric atom-map notation: [C:1]
	// is referenced as C1 and [O:2] as O2. The resulting name must be unique within
	// the molecule and at most four characters.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type SmallMoleculeDesignStartParamsChemicalSpace string

const (
	SmallMoleculeDesignStartParamsChemicalSpaceEnamineReal SmallMoleculeDesignStartParamsChemicalSpace = "enamine_real"
	SmallMoleculeDesignStartParamsChemicalSpaceNone        SmallMoleculeDesignStartParamsChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeDesignStartParamsMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level,omitzero"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion `json:"custom_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFilters) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion struct {
	OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterLipinskiFilter        *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter        `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterRdkitDescriptorFilter *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterSmartsCustomFilter    *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter    `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterSmartsCatalogFilter   *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter   `json:",omitzero,inline"`
	OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterSmilesRegexFilter     *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter     `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterLipinskiFilter,
		u.OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterRdkitDescriptorFilter,
		u.OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterSmartsCustomFilter,
		u.OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterSmartsCatalogFilter,
		u.OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterSmilesRegexFilter)
}
func (u *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
//
// The properties MaxHba, MaxHbd, MaxLogp, MaxMw, Type are required.
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter struct {
	// Maximum number of hydrogen bond acceptors. Lipinski threshold: 10
	MaxHba float64 `json:"max_hba" api:"required"`
	// Maximum number of hydrogen bond donors. Lipinski threshold: 5
	MaxHbd float64 `json:"max_hbd" api:"required"`
	// Maximum LogP. Lipinski threshold: 5
	MaxLogp float64 `json:"max_logp" api:"required"`
	// Maximum molecular weight (Da). Lipinski threshold: 500
	MaxMw float64 `json:"max_mw" api:"required"`
	// If true, one rule violation is allowed (classic Rule of Five). Defaults to false
	// (all rules must pass).
	AllowSingleViolation param.Opt[bool] `json:"allow_single_violation,omitzero"`
	// This field can be elided, and will marshal its zero value as "lipinski_filter".
	Type constant.LipinskiFilter `json:"type" default:"lipinski_filter"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
//
// The property Type is required.
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter struct {
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 `json:"fraction_csp3,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp `json:"mol_logp,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt `json:"mol_wt,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings `json:"num_aromatic_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors `json:"num_h_acceptors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors `json:"num_h_donors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms `json:"num_heteroatoms,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings `json:"num_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds `json:"num_rotatable_bonds,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa `json:"tpsa,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "rdkit_descriptor_filter".
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
//
// The properties Patterns, Type are required.
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_custom_filter".
	Type constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
//
// The properties Catalog, Type are required.
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter struct {
	// Predefined SMARTS catalog to apply. PAINS, BRENK, ChEMBL, and NIH catalogs
	// reject known problematic substructures.
	//
	// Any of "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL",
	// "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica",
	// "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH".
	Catalog string `json:"catalog,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_catalog_filter".
	Type constant.SmartsCatalogFilter `json:"type" default:"smarts_catalog_filter"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter](
		"catalog", "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL", "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica", "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH",
	)
}

// Filter molecules by regex patterns on their SMILES representation.
//
// The properties Patterns, Type are required.
type SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smiles_regex_filter".
	Type constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	paramObj
}

func (r SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
