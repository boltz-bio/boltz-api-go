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

// Screen an existing library of small molecules against a protein target. Results
// are scored by binding confidence (likelihood of binding, for hit discovery),
// optimization score (binding strength ranking, for lead optimization), and
// structure confidence.
//
// SmallMoleculeLibraryScreenService contains methods and other services that help
// with interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmallMoleculeLibraryScreenService] method instead.
type SmallMoleculeLibraryScreenService struct {
	Options []option.RequestOption
}

// NewSmallMoleculeLibraryScreenService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSmallMoleculeLibraryScreenService(opts ...option.RequestOption) (r SmallMoleculeLibraryScreenService) {
	r = SmallMoleculeLibraryScreenService{}
	r.Options = opts
	return
}

// Retrieve a library screen by ID, including progress and status
func (r *SmallMoleculeLibraryScreenService) Get(ctx context.Context, id string, query SmallMoleculeLibraryScreenGetParams, opts ...option.RequestOption) (res *SmallMoleculeLibraryScreenGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/library-screen/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List small molecule library screens, optionally filtered by workspace
func (r *SmallMoleculeLibraryScreenService) List(ctx context.Context, query SmallMoleculeLibraryScreenListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SmallMoleculeLibraryScreenListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/small-molecule/library-screen"
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

// List small molecule library screens, optionally filtered by workspace
func (r *SmallMoleculeLibraryScreenService) ListAutoPaging(ctx context.Context, query SmallMoleculeLibraryScreenListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SmallMoleculeLibraryScreenListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// library screen. The library screen record itself is retained with a
// `data_deleted_at` timestamp. This action is irreversible.
func (r *SmallMoleculeLibraryScreenService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeLibraryScreenDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/library-screen/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Estimate the cost of a small molecule library screen without creating any
// resource or consuming GPU.
func (r *SmallMoleculeLibraryScreenService) EstimateCost(ctx context.Context, body SmallMoleculeLibraryScreenEstimateCostParams, opts ...option.RequestOption) (res *SmallMoleculeLibraryScreenEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/small-molecule/library-screen/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve paginated results from a library screen
func (r *SmallMoleculeLibraryScreenService) ListResults(ctx context.Context, id string, query SmallMoleculeLibraryScreenListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[SmallMoleculeLibraryScreenListResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/library-screen/%s/results", url.PathEscape(id))
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

// Retrieve paginated results from a library screen
func (r *SmallMoleculeLibraryScreenService) ListResultsAutoPaging(ctx context.Context, id string, query SmallMoleculeLibraryScreenListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SmallMoleculeLibraryScreenListResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, id, query, opts...))
}

// Resume a stopped small molecule library screen from its last checkpoint
func (r *SmallMoleculeLibraryScreenService) Resume(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeLibraryScreenResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/library-screen/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Screen a set of small molecule candidates against a protein target
func (r *SmallMoleculeLibraryScreenService) Start(ctx context.Context, body SmallMoleculeLibraryScreenStartParams, opts ...option.RequestOption) (res *SmallMoleculeLibraryScreenStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/small-molecule/library-screen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an in-progress library screen early
func (r *SmallMoleculeLibraryScreenService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeLibraryScreenStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/library-screen/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A small molecule library screening pipeline run
type SmallMoleculeLibraryScreenGetResponse struct {
	// Unique SmScreen identifier
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
	EngineVersion constant.String1_0                         `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeLibraryScreenGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeLibraryScreenGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                            `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeLibraryScreenGetResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                     `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeLibraryScreenGetResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                   `json:"stopped_at" api:"required" format:"date-time"`
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
func (r SmallMoleculeLibraryScreenGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseError struct {
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
func (r SmallMoleculeLibraryScreenGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeLibraryScreenGetResponseInput struct {
	Molecules SmallMoleculeLibraryScreenGetResponseInputMolecules `json:"molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeLibraryScreenGetResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeLibraryScreenGetResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules       respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputMolecules struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMolecules) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponseInputMolecules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeLibraryScreenGetResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeLibraryScreenGetResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeLibraryScreenGetResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                     `json:"type" default:"glycan"`
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeLibraryScreenGetResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                               `json:"type" default:"contact"`
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. When omitted,
	// small-molecule design and library screen use 'recommended', while Explore uses
	// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
	// hit rate. 'extra': adds additional alerts beyond the recommended set for
	// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
	// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
	// only custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. When omitted,
// small-molecule design and library screen use 'recommended', while Explore uses
// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
// hit rate. 'extra': adds additional alerts beyond the recommended set for
// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
// only custom_filters will be applied.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                               `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                        `json:"latest_result_id"`
	RejectionSummary SmallMoleculeLibraryScreenGetResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed     respjson.Field
		NumMoleculesScreened   respjson.Field
		TotalMoleculesToScreen respjson.Field
		LatestResultID         respjson.Field
		RejectionSummary       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseProgressRejectionSummary struct {
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilteredCount respjson.Field
		InvalidCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenGetResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenGetResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenGetResponseStatus string

const (
	SmallMoleculeLibraryScreenGetResponseStatusPending   SmallMoleculeLibraryScreenGetResponseStatus = "pending"
	SmallMoleculeLibraryScreenGetResponseStatusRunning   SmallMoleculeLibraryScreenGetResponseStatus = "running"
	SmallMoleculeLibraryScreenGetResponseStatusSucceeded SmallMoleculeLibraryScreenGetResponseStatus = "succeeded"
	SmallMoleculeLibraryScreenGetResponseStatusFailed    SmallMoleculeLibraryScreenGetResponseStatus = "failed"
	SmallMoleculeLibraryScreenGetResponseStatusStopped   SmallMoleculeLibraryScreenGetResponseStatus = "stopped"
)

// Summary of a small molecule library screening pipeline run (excludes input)
type SmallMoleculeLibraryScreenListResponse struct {
	// Unique SmScreenSummary identifier
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
	EngineVersion constant.String1_0                          `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeLibraryScreenListResponseError `json:"error" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                             `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeLibraryScreenListResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                      `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeLibraryScreenListResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                    `json:"stopped_at" api:"required" format:"date-time"`
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
func (r SmallMoleculeLibraryScreenListResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResponseError struct {
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
func (r SmallMoleculeLibraryScreenListResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResponseProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                         `json:"latest_result_id"`
	RejectionSummary SmallMoleculeLibraryScreenListResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed     respjson.Field
		NumMoleculesScreened   respjson.Field
		TotalMoleculesToScreen respjson.Field
		LatestResultID         respjson.Field
		RejectionSummary       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenListResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResponseProgressRejectionSummary struct {
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilteredCount respjson.Field
		InvalidCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenListResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenListResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResponseStatus string

const (
	SmallMoleculeLibraryScreenListResponseStatusPending   SmallMoleculeLibraryScreenListResponseStatus = "pending"
	SmallMoleculeLibraryScreenListResponseStatusRunning   SmallMoleculeLibraryScreenListResponseStatus = "running"
	SmallMoleculeLibraryScreenListResponseStatusSucceeded SmallMoleculeLibraryScreenListResponseStatus = "succeeded"
	SmallMoleculeLibraryScreenListResponseStatusFailed    SmallMoleculeLibraryScreenListResponseStatus = "failed"
	SmallMoleculeLibraryScreenListResponseStatusStopped   SmallMoleculeLibraryScreenListResponseStatus = "stopped"
)

type SmallMoleculeLibraryScreenDeleteDataResponse struct {
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
func (r SmallMoleculeLibraryScreenDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimate response with monetary values encoded as decimal strings to preserve
// precision.
type SmallMoleculeLibraryScreenEstimateCostResponse struct {
	// Cost breakdown for the billed application.
	Breakdown  SmallMoleculeLibraryScreenEstimateCostResponseBreakdown `json:"breakdown" api:"required"`
	Disclaimer string                                                  `json:"disclaimer" api:"required"`
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
func (r SmallMoleculeLibraryScreenEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost breakdown for the billed application.
type SmallMoleculeLibraryScreenEstimateCostResponseBreakdown struct {
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
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
func (r SmallMoleculeLibraryScreenEstimateCostResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenEstimateCostResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication string

const (
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationStructureAndBinding        SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "structure_and_binding"
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationSmallMoleculeDesign        SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "small_molecule_design"
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationSmallMoleculeLibraryScreen SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "small_molecule_library_screen"
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationProteinDesign              SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "protein_design"
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationProteinRedesign            SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "protein_redesign"
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationProteinLibraryScreen       SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "protein_library_screen"
	SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplicationAdme                       SmallMoleculeLibraryScreenEstimateCostResponseBreakdownApplication = "adme"
)

// Result for a single screened small molecule
type SmallMoleculeLibraryScreenListResultsResponse struct {
	// Unique result ID
	ID        string                                                 `json:"id" api:"required"`
	Artifacts SmallMoleculeLibraryScreenListResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                              `json:"created_at" api:"required" format:"date-time"`
	// Scoring metrics for a screened small molecule
	Metrics SmallMoleculeLibraryScreenListResultsResponseMetrics `json:"metrics" api:"required"`
	// SMILES string of the screened molecule
	Smiles string `json:"smiles" api:"required"`
	// Tier 1 ADME summary values for this molecule.
	Adme SmallMoleculeLibraryScreenListResultsResponseAdme `json:"adme"`
	// Client-provided identifier for this molecule, if provided
	ExternalID string `json:"external_id"`
	// Warnings about potential quality issues with this result.
	Warnings []SmallMoleculeLibraryScreenListResultsResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Metrics     respjson.Field
		Smiles      respjson.Field
		Adme        respjson.Field
		ExternalID  respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenListResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResultsResponseArtifacts struct {
	Archive         SmallMoleculeLibraryScreenListResultsResponseArtifactsArchive         `json:"archive" api:"required"`
	Structure       SmallMoleculeLibraryScreenListResultsResponseArtifactsStructure       `json:"structure" api:"required"`
	LigandStructure SmallMoleculeLibraryScreenListResultsResponseArtifactsLigandStructure `json:"ligand_structure"`
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
func (r SmallMoleculeLibraryScreenListResultsResponseArtifacts) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResultsResponseArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResultsResponseArtifactsArchive struct {
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
func (r SmallMoleculeLibraryScreenListResultsResponseArtifactsArchive) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenListResultsResponseArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResultsResponseArtifactsStructure struct {
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
func (r SmallMoleculeLibraryScreenListResultsResponseArtifactsStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenListResultsResponseArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResultsResponseArtifactsLigandStructure struct {
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
func (r SmallMoleculeLibraryScreenListResultsResponseArtifactsLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenListResultsResponseArtifactsLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scoring metrics for a screened small molecule
type SmallMoleculeLibraryScreenListResultsResponseMetrics struct {
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
func (r SmallMoleculeLibraryScreenListResultsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResultsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tier 1 ADME summary values for this molecule.
type SmallMoleculeLibraryScreenListResultsResponseAdme struct {
	// Lipophilicity score from the internal LogD prediction.
	Lipophilicity float64 `json:"lipophilicity" api:"required"`
	// Permeability score for this molecule.
	Permeability float64 `json:"permeability" api:"required"`
	// Solubility judgement for this molecule.
	//
	// Any of "high-confidence", "medium-confidence", "high-risk".
	Solubility SmallMoleculeLibraryScreenListResultsResponseAdmeSolubility `json:"solubility" api:"required"`
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
func (r SmallMoleculeLibraryScreenListResultsResponseAdme) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResultsResponseAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solubility judgement for this molecule.
type SmallMoleculeLibraryScreenListResultsResponseAdmeSolubility string

const (
	SmallMoleculeLibraryScreenListResultsResponseAdmeSolubilityHighConfidence   SmallMoleculeLibraryScreenListResultsResponseAdmeSolubility = "high-confidence"
	SmallMoleculeLibraryScreenListResultsResponseAdmeSolubilityMediumConfidence SmallMoleculeLibraryScreenListResultsResponseAdmeSolubility = "medium-confidence"
	SmallMoleculeLibraryScreenListResultsResponseAdmeSolubilityHighRisk         SmallMoleculeLibraryScreenListResultsResponseAdmeSolubility = "high-risk"
)

// A warning about a potential quality issue with a result
type SmallMoleculeLibraryScreenListResultsResponseWarning struct {
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
func (r SmallMoleculeLibraryScreenListResultsResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenListResultsResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A small molecule library screening pipeline run
type SmallMoleculeLibraryScreenResumeResponse struct {
	// Unique SmScreen identifier
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
	EngineVersion constant.String1_0                            `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeLibraryScreenResumeResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeLibraryScreenResumeResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                               `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeLibraryScreenResumeResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                        `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeLibraryScreenResumeResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                      `json:"stopped_at" api:"required" format:"date-time"`
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
func (r SmallMoleculeLibraryScreenResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseError struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeLibraryScreenResumeResponseInput struct {
	Molecules SmallMoleculeLibraryScreenResumeResponseInputMolecules `json:"molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeLibraryScreenResumeResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeLibraryScreenResumeResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules       respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenResumeResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputMolecules struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMolecules) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponseInputMolecules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeLibraryScreenResumeResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeLibraryScreenResumeResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                        `json:"type" default:"glycan"`
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                  `json:"type" default:"contact"`
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. When omitted,
	// small-molecule design and library screen use 'recommended', while Explore uses
	// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
	// hit rate. 'extra': adds additional alerts beyond the recommended set for
	// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
	// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
	// only custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. When omitted,
// small-molecule design and library screen use 'recommended', while Explore uses
// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
// hit rate. 'extra': adds additional alerts beyond the recommended set for
// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
// only custom_filters will be applied.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                                  `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                           `json:"latest_result_id"`
	RejectionSummary SmallMoleculeLibraryScreenResumeResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed     respjson.Field
		NumMoleculesScreened   respjson.Field
		TotalMoleculesToScreen respjson.Field
		LatestResultID         respjson.Field
		RejectionSummary       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenResumeResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenResumeResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseProgressRejectionSummary struct {
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilteredCount respjson.Field
		InvalidCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenResumeResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenResumeResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenResumeResponseStatus string

const (
	SmallMoleculeLibraryScreenResumeResponseStatusPending   SmallMoleculeLibraryScreenResumeResponseStatus = "pending"
	SmallMoleculeLibraryScreenResumeResponseStatusRunning   SmallMoleculeLibraryScreenResumeResponseStatus = "running"
	SmallMoleculeLibraryScreenResumeResponseStatusSucceeded SmallMoleculeLibraryScreenResumeResponseStatus = "succeeded"
	SmallMoleculeLibraryScreenResumeResponseStatusFailed    SmallMoleculeLibraryScreenResumeResponseStatus = "failed"
	SmallMoleculeLibraryScreenResumeResponseStatusStopped   SmallMoleculeLibraryScreenResumeResponseStatus = "stopped"
)

// A small molecule library screening pipeline run
type SmallMoleculeLibraryScreenStartResponse struct {
	// Unique SmScreen identifier
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
	EngineVersion constant.String1_0                           `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeLibraryScreenStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeLibraryScreenStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                              `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeLibraryScreenStartResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                       `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeLibraryScreenStartResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                     `json:"stopped_at" api:"required" format:"date-time"`
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
func (r SmallMoleculeLibraryScreenStartResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseError struct {
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
func (r SmallMoleculeLibraryScreenStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeLibraryScreenStartResponseInput struct {
	Molecules SmallMoleculeLibraryScreenStartResponseInputMolecules `json:"molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeLibraryScreenStartResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeLibraryScreenStartResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules       respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputMolecules struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMolecules) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponseInputMolecules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeLibraryScreenStartResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeLibraryScreenStartResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeLibraryScreenStartResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                       `json:"type" default:"glycan"`
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeLibraryScreenStartResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                 `json:"type" default:"contact"`
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. When omitted,
	// small-molecule design and library screen use 'recommended', while Explore uses
	// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
	// hit rate. 'extra': adds additional alerts beyond the recommended set for
	// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
	// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
	// only custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. When omitted,
// small-molecule design and library screen use 'recommended', while Explore uses
// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
// hit rate. 'extra': adds additional alerts beyond the recommended set for
// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
// only custom_filters will be applied.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                                 `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                          `json:"latest_result_id"`
	RejectionSummary SmallMoleculeLibraryScreenStartResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed     respjson.Field
		NumMoleculesScreened   respjson.Field
		TotalMoleculesToScreen respjson.Field
		LatestResultID         respjson.Field
		RejectionSummary       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStartResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStartResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseProgressRejectionSummary struct {
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilteredCount respjson.Field
		InvalidCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStartResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStartResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStartResponseStatus string

const (
	SmallMoleculeLibraryScreenStartResponseStatusPending   SmallMoleculeLibraryScreenStartResponseStatus = "pending"
	SmallMoleculeLibraryScreenStartResponseStatusRunning   SmallMoleculeLibraryScreenStartResponseStatus = "running"
	SmallMoleculeLibraryScreenStartResponseStatusSucceeded SmallMoleculeLibraryScreenStartResponseStatus = "succeeded"
	SmallMoleculeLibraryScreenStartResponseStatusFailed    SmallMoleculeLibraryScreenStartResponseStatus = "failed"
	SmallMoleculeLibraryScreenStartResponseStatusStopped   SmallMoleculeLibraryScreenStartResponseStatus = "stopped"
)

// A small molecule library screening pipeline run
type SmallMoleculeLibraryScreenStopResponse struct {
	// Unique SmScreen identifier
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
	EngineVersion constant.String1_0                          `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeLibraryScreenStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeLibraryScreenStopResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                             `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeLibraryScreenStopResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                      `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeLibraryScreenStopResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                    `json:"stopped_at" api:"required" format:"date-time"`
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
func (r SmallMoleculeLibraryScreenStopResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseError struct {
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
func (r SmallMoleculeLibraryScreenStopResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeLibraryScreenStopResponseInput struct {
	Molecules SmallMoleculeLibraryScreenStopResponseInputMolecules `json:"molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeLibraryScreenStopResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeLibraryScreenStopResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules       respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStopResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputMolecules struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMolecules) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponseInputMolecules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeLibraryScreenStopResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeLibraryScreenStopResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeLibraryScreenStopResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion) AsSmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                      `json:"type" default:"glycan"`
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeLibraryScreenStopResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion) AsSmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                `json:"type" default:"contact"`
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. When omitted,
	// small-molecule design and library screen use 'recommended', while Explore uses
	// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
	// hit rate. 'extra': adds additional alerts beyond the recommended set for
	// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
	// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
	// only custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. When omitted,
// small-molecule design and library screen use 'recommended', while Explore uses
// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
// hit rate. 'extra': adds additional alerts beyond the recommended set for
// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
// only custom_filters will be applied.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                                `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                         `json:"latest_result_id"`
	RejectionSummary SmallMoleculeLibraryScreenStopResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed     respjson.Field
		NumMoleculesScreened   respjson.Field
		TotalMoleculesToScreen respjson.Field
		LatestResultID         respjson.Field
		RejectionSummary       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStopResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeLibraryScreenStopResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseProgressRejectionSummary struct {
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilteredCount respjson.Field
		InvalidCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeLibraryScreenStopResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeLibraryScreenStopResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenStopResponseStatus string

const (
	SmallMoleculeLibraryScreenStopResponseStatusPending   SmallMoleculeLibraryScreenStopResponseStatus = "pending"
	SmallMoleculeLibraryScreenStopResponseStatusRunning   SmallMoleculeLibraryScreenStopResponseStatus = "running"
	SmallMoleculeLibraryScreenStopResponseStatusSucceeded SmallMoleculeLibraryScreenStopResponseStatus = "succeeded"
	SmallMoleculeLibraryScreenStopResponseStatusFailed    SmallMoleculeLibraryScreenStopResponseStatus = "failed"
	SmallMoleculeLibraryScreenStopResponseStatusStopped   SmallMoleculeLibraryScreenStopResponseStatus = "stopped"
)

type SmallMoleculeLibraryScreenGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SmallMoleculeLibraryScreenGetParams]'s query parameters as
// `url.Values`.
func (r SmallMoleculeLibraryScreenGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeLibraryScreenListParams struct {
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

// URLQuery serializes [SmallMoleculeLibraryScreenListParams]'s query parameters as
// `url.Values`.
func (r SmallMoleculeLibraryScreenListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeLibraryScreenEstimateCostParams struct {
	// List of small molecules to screen.
	Molecules []SmallMoleculeLibraryScreenEstimateCostParamsMolecule `json:"molecules,omitzero" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeLibraryScreenEstimateCostParamsTarget `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters `json:"molecule_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A small molecule with SMILES and optional ID
//
// The property Smiles is required.
type SmallMoleculeLibraryScreenEstimateCostParamsMolecule struct {
	// SMILES string of the molecule
	Smiles string `json:"smiles" api:"required"`
	// Optional identifier for this molecule
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMolecule) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMolecule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
//
// The property Entities is required.
type SmallMoleculeLibraryScreenEstimateCostParamsTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeLibraryScreenEstimateCostParamsTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion `json:"constraints,omitzero"`
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeLibraryScreenEstimateCostParamsTarget](
		"type", "no_template",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityUnion struct {
	OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityProteinEntity *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntity `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityGlycanEntity  *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityProteinEntity, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityGlycanEntity)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntityModification struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBond struct {
	Atom1 SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union struct {
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1PolymerAtom *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1CcdAtom     *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1SmilesAtom  *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1LigandAtom  *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1PolymerAtom, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1CcdAtom, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1SmilesAtom, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1LigandAtom)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1PolymerAtom struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1CcdAtom struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union struct {
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2PolymerAtom *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2CcdAtom     *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2SmilesAtom  *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2LigandAtom  *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2PolymerAtom, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2CcdAtom, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2SmilesAtom, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2LigandAtom)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2PolymerAtom struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2CcdAtom struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion struct {
	OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintPocketConstraint  *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraint *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintPocketConstraint, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraint)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1Union struct {
	OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken1PolymerContactToken *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken1LigandContactToken  *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken1PolymerContactToken, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2Union struct {
	OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken2PolymerContactToken *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken2LigandContactToken  *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken2PolymerContactToken, u.OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. When omitted,
	// small-molecule design and library screen use 'recommended', while Explore uses
	// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
	// hit rate. 'extra': adds additional alerts beyond the recommended set for
	// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
	// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
	// only custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level,omitzero"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion `json:"custom_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. When omitted,
// small-molecule design and library screen use 'recommended', while Explore uses
// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
// hit rate. 'extra': adds additional alerts beyond the recommended set for
// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
// only custom_filters will be applied.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion struct {
	OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter        *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter        `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterRdkitDescriptorFilter *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterSmartsCustomFilter    *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter    `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterSmartsCatalogFilter   *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter   `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterSmilesRegexFilter     *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter     `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter,
		u.OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterRdkitDescriptorFilter,
		u.OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterSmartsCustomFilter,
		u.OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterSmartsCatalogFilter,
		u.OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterSmilesRegexFilter)
}
func (u *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
//
// The properties MaxHba, MaxHbd, MaxLogp, MaxMw, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
//
// The property Type is required.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter struct {
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 `json:"fraction_csp3,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp `json:"mol_logp,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt `json:"mol_wt,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings `json:"num_aromatic_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors `json:"num_h_acceptors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors `json:"num_h_donors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms `json:"num_heteroatoms,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings `json:"num_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds `json:"num_rotatable_bonds,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa `json:"tpsa,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "rdkit_descriptor_filter".
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
//
// The properties Patterns, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_custom_filter".
	Type constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCustomFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
//
// The properties Catalog, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter struct {
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

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmartsCatalogFilter](
		"catalog", "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL", "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica", "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH",
	)
}

// Filter molecules by regex patterns on their SMILES representation.
//
// The properties Patterns, Type are required.
type SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smiles_regex_filter".
	Type constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	paramObj
}

func (r SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterSmilesRegexFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeLibraryScreenListResultsParams struct {
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

// URLQuery serializes [SmallMoleculeLibraryScreenListResultsParams]'s query
// parameters as `url.Values`.
func (r SmallMoleculeLibraryScreenListResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeLibraryScreenStartParams struct {
	// List of small molecules to screen.
	Molecules []SmallMoleculeLibraryScreenStartParamsMolecule `json:"molecules,omitzero" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeLibraryScreenStartParamsTarget `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeLibraryScreenStartParamsMoleculeFilters `json:"molecule_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParams) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A small molecule with SMILES and optional ID
//
// The property Smiles is required.
type SmallMoleculeLibraryScreenStartParamsMolecule struct {
	// SMILES string of the molecule
	Smiles string `json:"smiles" api:"required"`
	// Optional identifier for this molecule
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMolecule) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMolecule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
//
// The property Entities is required.
type SmallMoleculeLibraryScreenStartParamsTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeLibraryScreenStartParamsTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeLibraryScreenStartParamsTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion `json:"constraints,omitzero"`
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

func (r SmallMoleculeLibraryScreenStartParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeLibraryScreenStartParamsTarget](
		"type", "no_template",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsTargetEntityUnion struct {
	OfSmallMoleculeLibraryScreenStartsTargetEntityProteinEntity *SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntity `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetEntityGlycanEntity  *SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsTargetEntityProteinEntity, u.OfSmallMoleculeLibraryScreenStartsTargetEntityGlycanEntity)
}
func (u *SmallMoleculeLibraryScreenStartParamsTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntityModification struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBond struct {
	Atom1 SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeLibraryScreenStartParamsTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union struct {
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom1PolymerAtom *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom1CcdAtom     *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom1SmilesAtom  *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom1LigandAtom  *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom1PolymerAtom, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom1CcdAtom, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom1SmilesAtom, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom1LigandAtom)
}
func (u *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom1PolymerAtom struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom1CcdAtom struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union struct {
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom2PolymerAtom *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom2CcdAtom     *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom2SmilesAtom  *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetBondAtom2LigandAtom  *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom2PolymerAtom, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom2CcdAtom, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom2SmilesAtom, u.OfSmallMoleculeLibraryScreenStartsTargetBondAtom2LigandAtom)
}
func (u *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom2PolymerAtom struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom2CcdAtom struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion struct {
	OfSmallMoleculeLibraryScreenStartsTargetConstraintPocketConstraint  *SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraint *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsTargetConstraintPocketConstraint, u.OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraint)
}
func (u *SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1Union struct {
	OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken1PolymerContactToken *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken1LigandContactToken  *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken1PolymerContactToken, u.OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1LigandContactToken struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2Union struct {
	OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken2PolymerContactToken *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken2LigandContactToken  *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken2PolymerContactToken, u.OfSmallMoleculeLibraryScreenStartsTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2LigandContactToken struct {
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

func (r SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeLibraryScreenStartParamsMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. When omitted,
	// small-molecule design and library screen use 'recommended', while Explore uses
	// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
	// hit rate. 'extra': adds additional alerts beyond the recommended set for
	// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
	// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
	// only custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level,omitzero"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion `json:"custom_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFilters) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. When omitted,
// small-molecule design and library screen use 'recommended', while Explore uses
// 'disabled'. 'recommended': applies a curated set of alerts balancing safety and
// hit rate. 'extra': adds additional alerts beyond the recommended set for
// stricter filtering. 'aggressive': applies the most comprehensive alert set — may
// reject viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely;
// only custom_filters will be applied.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion struct {
	OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterLipinskiFilter        *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter        `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterRdkitDescriptorFilter *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterSmartsCustomFilter    *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter    `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterSmartsCatalogFilter   *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter   `json:",omitzero,inline"`
	OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterSmilesRegexFilter     *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter     `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterLipinskiFilter,
		u.OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterRdkitDescriptorFilter,
		u.OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterSmartsCustomFilter,
		u.OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterSmartsCatalogFilter,
		u.OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterSmilesRegexFilter)
}
func (u *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
//
// The properties MaxHba, MaxHbd, MaxLogp, MaxMw, Type are required.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter struct {
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

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
//
// The property Type is required.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter struct {
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 `json:"fraction_csp3,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp `json:"mol_logp,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt `json:"mol_wt,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings `json:"num_aromatic_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors `json:"num_h_acceptors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors `json:"num_h_donors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms `json:"num_heteroatoms,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings `json:"num_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds `json:"num_rotatable_bonds,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa `json:"tpsa,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "rdkit_descriptor_filter".
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
//
// The properties Patterns, Type are required.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_custom_filter".
	Type constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
//
// The properties Catalog, Type are required.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter struct {
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

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter](
		"catalog", "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL", "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica", "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH",
	)
}

// Filter molecules by regex patterns on their SMILES representation.
//
// The properties Patterns, Type are required.
type SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smiles_regex_filter".
	Type constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	paramObj
}

func (r SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
