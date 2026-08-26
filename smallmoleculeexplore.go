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

// Explore a large library of small molecules against a protein target without
// screening all of it. Submit the whole library and a budget; molecules are chosen
// to score as results arrive, so each choice is informed by everything scored so
// far. Results use the same scores as a library screen, and progress reports the
// library size alongside the budget.
//
// SmallMoleculeExploreService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmallMoleculeExploreService] method instead.
type SmallMoleculeExploreService struct {
	Options []option.RequestOption
}

// NewSmallMoleculeExploreService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSmallMoleculeExploreService(opts ...option.RequestOption) (r SmallMoleculeExploreService) {
	r = SmallMoleculeExploreService{}
	r.Options = opts
	return
}

// Retrieve an exploration by ID. Once library preparation completes, progress
// reports the accepted library size alongside the budget.
func (r *SmallMoleculeExploreService) Get(ctx context.Context, id string, query SmallMoleculeExploreGetParams, opts ...option.RequestOption) (res *SmallMoleculeExploreGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/explore/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve paginated results from an exploration. Results appear as molecules are
// scored, and remain retrievable if the run fails partway.
func (r *SmallMoleculeExploreService) ListResults(ctx context.Context, id string, query SmallMoleculeExploreListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[SmallMoleculeExploreListResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/explore/%s/results", url.PathEscape(id))
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

// Retrieve paginated results from an exploration. Results appear as molecules are
// scored, and remain retrievable if the run fails partway.
func (r *SmallMoleculeExploreService) ListResultsAutoPaging(ctx context.Context, id string, query SmallMoleculeExploreListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SmallMoleculeExploreListResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, id, query, opts...))
}

// Resume a stopped exploration. Selection continues from the surrogate as it
// stood, so molecules already scored still inform what is chosen next.
func (r *SmallMoleculeExploreService) Resume(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeExploreResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/explore/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Explore a large library against a protein target without screening all of it.
// Submit the whole library and a budget; molecules are chosen to score as results
// arrive, so each choice is informed by everything scored before it.
func (r *SmallMoleculeExploreService) Start(ctx context.Context, body SmallMoleculeExploreStartParams, opts ...option.RequestOption) (res *SmallMoleculeExploreStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/small-molecule/explore"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an in-progress exploration early. Molecules already scored are kept and
// remain retrievable; no further molecules are selected.
func (r *SmallMoleculeExploreService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *SmallMoleculeExploreStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/small-molecule/explore/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A small molecule library exploration pipeline run
type SmallMoleculeExploreGetResponse struct {
	// Unique SmExplore identifier
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
	Error         SmallMoleculeExploreGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeExploreGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library exploration
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule exploration
	PipelineVersion constant.String1_0                      `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeExploreGetResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                               `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeExploreGetResponseStatus `json:"status" api:"required"`
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
func (r SmallMoleculeExploreGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseError struct {
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
func (r SmallMoleculeExploreGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeExploreGetResponseInput struct {
	// How many molecules to score. Must not exceed the accepted library size or
	// 5,000,000.
	Budget  int64                                       `json:"budget" api:"required"`
	Library SmallMoleculeExploreGetResponseInputLibrary `json:"library" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeExploreGetResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeExploreGetResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Budget          respjson.Field
		Library         respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputLibrary struct {
	// Delimited-text format of the molecule library.
	//
	// Any of "csv", "tsv".
	Format SmallMoleculeExploreGetResponseInputLibraryFormat `json:"format" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	SmilesColumn string `json:"smiles_column" api:"required"`
	// The original submitted URL for URL sources, or a temporary download URL for an
	// uploaded base64 source.
	Source SmallMoleculeExploreGetResponseInputLibrarySourceUnion `json:"source" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	IDColumn string `json:"id_column"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format       respjson.Field
		SmilesColumn respjson.Field
		Source       respjson.Field
		IDColumn     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseInputLibrary) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseInputLibrary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delimited-text format of the molecule library.
type SmallMoleculeExploreGetResponseInputLibraryFormat string

const (
	SmallMoleculeExploreGetResponseInputLibraryFormatCsv SmallMoleculeExploreGetResponseInputLibraryFormat = "csv"
	SmallMoleculeExploreGetResponseInputLibraryFormatTsv SmallMoleculeExploreGetResponseInputLibraryFormat = "tsv"
)

// SmallMoleculeExploreGetResponseInputLibrarySourceUnion contains all possible
// properties and values from
// [SmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse],
// [SmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputLibrarySourceUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse].
	Type constant.URL `json:"type"`
	URL  string       `json:"url"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse].
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		Type         respjson.Field
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreGetResponseInputLibrarySourceUnion) AsSmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse() (v SmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputLibrarySourceUnion) AsSmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse() (v SmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputLibrarySourceUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreGetResponseInputLibrarySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse struct {
	Type constant.URL `json:"type" default:"url"`
	URL  string       `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputLibrarySourceURLSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputLibrarySourceFileOutputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeExploreGetResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeExploreGetResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeExploreGetResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeExploreGetResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeExploreGetResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreGetResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeExploreGetResponseInputTargetEntityUnion) AsSmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetEntityUnion) AsSmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreGetResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
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
func (r SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeExploreGetResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeExploreGetResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeExploreGetResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreGetResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeExploreGetResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputTargetBondAtom1Union) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreGetResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreGetResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeExploreGetResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreGetResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeExploreGetResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputTargetBondAtom2Union) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreGetResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreGetResponseInputTargetConstraintUnion contains all possible
// properties and values from
// [SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeExploreGetResponseInputTargetConstraintUnion) AsSmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetConstraintUnion) AsSmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreGetResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeExploreGetResponseInputMoleculeFilters struct {
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
	BoltzSmartsCatalogFilterLevel SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
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
type SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeExploreGetResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                         `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreGetResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseProgress struct {
	// Molecules that reached terminal failure. These do not consume budget — each is
	// replaced by another selection.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Molecules that produced a usable result. The run completes when this reaches the
	// budget.
	NumMoleculesScored int64 `json:"num_molecules_scored" api:"required"`
	// The requested budget: how many of the library will be scored.
	TotalMoleculesToScore int64 `json:"total_molecules_to_score" api:"required"`
	// ID of the most recently scored result
	LatestResultID string `json:"latest_result_id"`
	// Distinct molecules accepted after validation, filtering and de-duplication.
	// Omitted while the submitted library is being prepared.
	LibrarySize      int64                                                   `json:"library_size"`
	RejectionSummary SmallMoleculeExploreGetResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed    respjson.Field
		NumMoleculesScored    respjson.Field
		TotalMoleculesToScore respjson.Field
		LatestResultID        respjson.Field
		LibrarySize           respjson.Field
		RejectionSummary      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseProgressRejectionSummary struct {
	// Number of submitted rows that collapsed onto a molecule already in the library.
	// Exploration works over distinct molecules, so duplicates are merged rather than
	// scored twice.
	DuplicateCount int64 `json:"duplicate_count" api:"required"`
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		FilteredCount  respjson.Field
		InvalidCount   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreGetResponseProgressRejectionSummary) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreGetResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreGetResponseStatus string

const (
	SmallMoleculeExploreGetResponseStatusPending   SmallMoleculeExploreGetResponseStatus = "pending"
	SmallMoleculeExploreGetResponseStatusRunning   SmallMoleculeExploreGetResponseStatus = "running"
	SmallMoleculeExploreGetResponseStatusSucceeded SmallMoleculeExploreGetResponseStatus = "succeeded"
	SmallMoleculeExploreGetResponseStatusFailed    SmallMoleculeExploreGetResponseStatus = "failed"
	SmallMoleculeExploreGetResponseStatusStopped   SmallMoleculeExploreGetResponseStatus = "stopped"
)

// Result for a single scored small molecule
type SmallMoleculeExploreListResultsResponse struct {
	// Unique result ID
	ID        string                                           `json:"id" api:"required"`
	Artifacts SmallMoleculeExploreListResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                        `json:"created_at" api:"required" format:"date-time"`
	// Scoring metrics for a screened small molecule
	Metrics SmallMoleculeExploreListResultsResponseMetrics `json:"metrics" api:"required"`
	// SMILES string of the scored molecule
	Smiles string `json:"smiles" api:"required"`
	// Client-provided identifier for this molecule, if provided
	ExternalID string `json:"external_id"`
	// Warnings about potential quality issues with this result.
	Warnings []SmallMoleculeExploreListResultsResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Metrics     respjson.Field
		Smiles      respjson.Field
		ExternalID  respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreListResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreListResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreListResultsResponseArtifacts struct {
	Archive         SmallMoleculeExploreListResultsResponseArtifactsArchive         `json:"archive" api:"required"`
	Structure       SmallMoleculeExploreListResultsResponseArtifactsStructure       `json:"structure" api:"required"`
	LigandStructure SmallMoleculeExploreListResultsResponseArtifactsLigandStructure `json:"ligand_structure"`
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
func (r SmallMoleculeExploreListResultsResponseArtifacts) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreListResultsResponseArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreListResultsResponseArtifactsArchive struct {
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
func (r SmallMoleculeExploreListResultsResponseArtifactsArchive) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreListResultsResponseArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreListResultsResponseArtifactsStructure struct {
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
func (r SmallMoleculeExploreListResultsResponseArtifactsStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreListResultsResponseArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreListResultsResponseArtifactsLigandStructure struct {
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
func (r SmallMoleculeExploreListResultsResponseArtifactsLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreListResultsResponseArtifactsLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scoring metrics for a screened small molecule
type SmallMoleculeExploreListResultsResponseMetrics struct {
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
func (r SmallMoleculeExploreListResultsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreListResultsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warning about a potential quality issue with a result
type SmallMoleculeExploreListResultsResponseWarning struct {
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
func (r SmallMoleculeExploreListResultsResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreListResultsResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A small molecule library exploration pipeline run
type SmallMoleculeExploreResumeResponse struct {
	// Unique SmExplore identifier
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
	EngineVersion constant.String1_0                      `json:"engine_version" default:"1.0"`
	Error         SmallMoleculeExploreResumeResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeExploreResumeResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library exploration
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule exploration
	PipelineVersion constant.String1_0                         `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeExploreResumeResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                  `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeExploreResumeResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                `json:"stopped_at" api:"required" format:"date-time"`
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
func (r SmallMoleculeExploreResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseError struct {
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
func (r SmallMoleculeExploreResumeResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeExploreResumeResponseInput struct {
	// How many molecules to score. Must not exceed the accepted library size or
	// 5,000,000.
	Budget  int64                                          `json:"budget" api:"required"`
	Library SmallMoleculeExploreResumeResponseInputLibrary `json:"library" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeExploreResumeResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeExploreResumeResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Budget          respjson.Field
		Library         respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputLibrary struct {
	// Delimited-text format of the molecule library.
	//
	// Any of "csv", "tsv".
	Format SmallMoleculeExploreResumeResponseInputLibraryFormat `json:"format" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	SmilesColumn string `json:"smiles_column" api:"required"`
	// The original submitted URL for URL sources, or a temporary download URL for an
	// uploaded base64 source.
	Source SmallMoleculeExploreResumeResponseInputLibrarySourceUnion `json:"source" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	IDColumn string `json:"id_column"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format       respjson.Field
		SmilesColumn respjson.Field
		Source       respjson.Field
		IDColumn     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseInputLibrary) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseInputLibrary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delimited-text format of the molecule library.
type SmallMoleculeExploreResumeResponseInputLibraryFormat string

const (
	SmallMoleculeExploreResumeResponseInputLibraryFormatCsv SmallMoleculeExploreResumeResponseInputLibraryFormat = "csv"
	SmallMoleculeExploreResumeResponseInputLibraryFormatTsv SmallMoleculeExploreResumeResponseInputLibraryFormat = "tsv"
)

// SmallMoleculeExploreResumeResponseInputLibrarySourceUnion contains all possible
// properties and values from
// [SmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse],
// [SmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputLibrarySourceUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse].
	Type constant.URL `json:"type"`
	URL  string       `json:"url"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse].
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		Type         respjson.Field
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreResumeResponseInputLibrarySourceUnion) AsSmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse() (v SmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputLibrarySourceUnion) AsSmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse() (v SmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputLibrarySourceUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputLibrarySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse struct {
	Type constant.URL `json:"type" default:"url"`
	URL  string       `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputLibrarySourceURLSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputLibrarySourceFileOutputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeExploreResumeResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeExploreResumeResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeExploreResumeResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeExploreResumeResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeExploreResumeResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreResumeResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeExploreResumeResponseInputTargetEntityUnion) AsSmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetEntityUnion) AsSmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreResumeResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                  `json:"type" default:"glycan"`
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
func (r SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeExploreResumeResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union contains all
// possible properties and values from
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union contains all
// possible properties and values from
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreResumeResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeExploreResumeResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreResumeResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreResumeResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeExploreResumeResponseInputTargetConstraintUnion) AsSmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetConstraintUnion) AsSmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                            `json:"type" default:"contact"`
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
func (r SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeExploreResumeResponseInputMoleculeFilters struct {
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
	BoltzSmartsCatalogFilterLevel SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
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
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeExploreResumeResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                            `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseProgress struct {
	// Molecules that reached terminal failure. These do not consume budget — each is
	// replaced by another selection.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Molecules that produced a usable result. The run completes when this reaches the
	// budget.
	NumMoleculesScored int64 `json:"num_molecules_scored" api:"required"`
	// The requested budget: how many of the library will be scored.
	TotalMoleculesToScore int64 `json:"total_molecules_to_score" api:"required"`
	// ID of the most recently scored result
	LatestResultID string `json:"latest_result_id"`
	// Distinct molecules accepted after validation, filtering and de-duplication.
	// Omitted while the submitted library is being prepared.
	LibrarySize      int64                                                      `json:"library_size"`
	RejectionSummary SmallMoleculeExploreResumeResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed    respjson.Field
		NumMoleculesScored    respjson.Field
		TotalMoleculesToScore respjson.Field
		LatestResultID        respjson.Field
		LibrarySize           respjson.Field
		RejectionSummary      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreResumeResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseProgressRejectionSummary struct {
	// Number of submitted rows that collapsed onto a molecule already in the library.
	// Exploration works over distinct molecules, so duplicates are merged rather than
	// scored twice.
	DuplicateCount int64 `json:"duplicate_count" api:"required"`
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		FilteredCount  respjson.Field
		InvalidCount   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreResumeResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreResumeResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreResumeResponseStatus string

const (
	SmallMoleculeExploreResumeResponseStatusPending   SmallMoleculeExploreResumeResponseStatus = "pending"
	SmallMoleculeExploreResumeResponseStatusRunning   SmallMoleculeExploreResumeResponseStatus = "running"
	SmallMoleculeExploreResumeResponseStatusSucceeded SmallMoleculeExploreResumeResponseStatus = "succeeded"
	SmallMoleculeExploreResumeResponseStatusFailed    SmallMoleculeExploreResumeResponseStatus = "failed"
	SmallMoleculeExploreResumeResponseStatusStopped   SmallMoleculeExploreResumeResponseStatus = "stopped"
)

// A small molecule library exploration pipeline run
type SmallMoleculeExploreStartResponse struct {
	// Unique SmExplore identifier
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
	Error         SmallMoleculeExploreStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeExploreStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library exploration
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule exploration
	PipelineVersion constant.String1_0                        `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeExploreStartResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                 `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeExploreStartResponseStatus `json:"status" api:"required"`
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
func (r SmallMoleculeExploreStartResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseError struct {
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
func (r SmallMoleculeExploreStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeExploreStartResponseInput struct {
	// How many molecules to score. Must not exceed the accepted library size or
	// 5,000,000.
	Budget  int64                                         `json:"budget" api:"required"`
	Library SmallMoleculeExploreStartResponseInputLibrary `json:"library" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeExploreStartResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeExploreStartResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Budget          respjson.Field
		Library         respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputLibrary struct {
	// Delimited-text format of the molecule library.
	//
	// Any of "csv", "tsv".
	Format SmallMoleculeExploreStartResponseInputLibraryFormat `json:"format" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	SmilesColumn string `json:"smiles_column" api:"required"`
	// The original submitted URL for URL sources, or a temporary download URL for an
	// uploaded base64 source.
	Source SmallMoleculeExploreStartResponseInputLibrarySourceUnion `json:"source" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	IDColumn string `json:"id_column"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format       respjson.Field
		SmilesColumn respjson.Field
		Source       respjson.Field
		IDColumn     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseInputLibrary) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseInputLibrary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delimited-text format of the molecule library.
type SmallMoleculeExploreStartResponseInputLibraryFormat string

const (
	SmallMoleculeExploreStartResponseInputLibraryFormatCsv SmallMoleculeExploreStartResponseInputLibraryFormat = "csv"
	SmallMoleculeExploreStartResponseInputLibraryFormatTsv SmallMoleculeExploreStartResponseInputLibraryFormat = "tsv"
)

// SmallMoleculeExploreStartResponseInputLibrarySourceUnion contains all possible
// properties and values from
// [SmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse],
// [SmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputLibrarySourceUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse].
	Type constant.URL `json:"type"`
	URL  string       `json:"url"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse].
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		Type         respjson.Field
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreStartResponseInputLibrarySourceUnion) AsSmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse() (v SmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputLibrarySourceUnion) AsSmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse() (v SmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputLibrarySourceUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreStartResponseInputLibrarySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse struct {
	Type constant.URL `json:"type" default:"url"`
	URL  string       `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputLibrarySourceURLSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputLibrarySourceFileOutputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeExploreStartResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeExploreStartResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeExploreStartResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeExploreStartResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeExploreStartResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStartResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeExploreStartResponseInputTargetEntityUnion) AsSmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetEntityUnion) AsSmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreStartResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
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
func (r SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeExploreStartResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeExploreStartResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeExploreStartResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStartResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeExploreStartResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreStartResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStartResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeExploreStartResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStartResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeExploreStartResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreStartResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStartResponseInputTargetConstraintUnion contains all
// possible properties and values from
// [SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeExploreStartResponseInputTargetConstraintUnion) AsSmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetConstraintUnion) AsSmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStartResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeExploreStartResponseInputMoleculeFilters struct {
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
	BoltzSmartsCatalogFilterLevel SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
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
type SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeExploreStartResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                           `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseProgress struct {
	// Molecules that reached terminal failure. These do not consume budget — each is
	// replaced by another selection.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Molecules that produced a usable result. The run completes when this reaches the
	// budget.
	NumMoleculesScored int64 `json:"num_molecules_scored" api:"required"`
	// The requested budget: how many of the library will be scored.
	TotalMoleculesToScore int64 `json:"total_molecules_to_score" api:"required"`
	// ID of the most recently scored result
	LatestResultID string `json:"latest_result_id"`
	// Distinct molecules accepted after validation, filtering and de-duplication.
	// Omitted while the submitted library is being prepared.
	LibrarySize      int64                                                     `json:"library_size"`
	RejectionSummary SmallMoleculeExploreStartResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed    respjson.Field
		NumMoleculesScored    respjson.Field
		TotalMoleculesToScore respjson.Field
		LatestResultID        respjson.Field
		LibrarySize           respjson.Field
		RejectionSummary      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStartResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseProgressRejectionSummary struct {
	// Number of submitted rows that collapsed onto a molecule already in the library.
	// Exploration works over distinct molecules, so duplicates are merged rather than
	// scored twice.
	DuplicateCount int64 `json:"duplicate_count" api:"required"`
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		FilteredCount  respjson.Field
		InvalidCount   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStartResponseProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStartResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStartResponseStatus string

const (
	SmallMoleculeExploreStartResponseStatusPending   SmallMoleculeExploreStartResponseStatus = "pending"
	SmallMoleculeExploreStartResponseStatusRunning   SmallMoleculeExploreStartResponseStatus = "running"
	SmallMoleculeExploreStartResponseStatusSucceeded SmallMoleculeExploreStartResponseStatus = "succeeded"
	SmallMoleculeExploreStartResponseStatusFailed    SmallMoleculeExploreStartResponseStatus = "failed"
	SmallMoleculeExploreStartResponseStatusStopped   SmallMoleculeExploreStartResponseStatus = "stopped"
)

// A small molecule library exploration pipeline run
type SmallMoleculeExploreStopResponse struct {
	// Unique SmExplore identifier
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
	Error         SmallMoleculeExploreStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input SmallMoleculeExploreStopResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library exploration
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule exploration
	PipelineVersion constant.String1_0                       `json:"pipeline_version" default:"1.0"`
	Progress        SmallMoleculeExploreStopResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    SmallMoleculeExploreStopResponseStatus `json:"status" api:"required"`
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
func (r SmallMoleculeExploreStopResponse) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseError struct {
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
func (r SmallMoleculeExploreStopResponseError) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type SmallMoleculeExploreStopResponseInput struct {
	// How many molecules to score. Must not exceed the accepted library size or
	// 5,000,000.
	Budget  int64                                        `json:"budget" api:"required"`
	Library SmallMoleculeExploreStopResponseInputLibrary `json:"library" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeExploreStopResponseInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeExploreStopResponseInputMoleculeFilters `json:"molecule_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Budget          respjson.Field
		Library         respjson.Field
		Target          respjson.Field
		MoleculeFilters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseInput) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputLibrary struct {
	// Delimited-text format of the molecule library.
	//
	// Any of "csv", "tsv".
	Format SmallMoleculeExploreStopResponseInputLibraryFormat `json:"format" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	SmilesColumn string `json:"smiles_column" api:"required"`
	// The original submitted URL for URL sources, or a temporary download URL for an
	// uploaded base64 source.
	Source SmallMoleculeExploreStopResponseInputLibrarySourceUnion `json:"source" api:"required"`
	// An exact, case-sensitive column name without leading or trailing whitespace.
	IDColumn string `json:"id_column"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format       respjson.Field
		SmilesColumn respjson.Field
		Source       respjson.Field
		IDColumn     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseInputLibrary) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseInputLibrary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delimited-text format of the molecule library.
type SmallMoleculeExploreStopResponseInputLibraryFormat string

const (
	SmallMoleculeExploreStopResponseInputLibraryFormatCsv SmallMoleculeExploreStopResponseInputLibraryFormat = "csv"
	SmallMoleculeExploreStopResponseInputLibraryFormatTsv SmallMoleculeExploreStopResponseInputLibraryFormat = "tsv"
)

// SmallMoleculeExploreStopResponseInputLibrarySourceUnion contains all possible
// properties and values from
// [SmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse],
// [SmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputLibrarySourceUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse].
	Type constant.URL `json:"type"`
	URL  string       `json:"url"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse].
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		Type         respjson.Field
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreStopResponseInputLibrarySourceUnion) AsSmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse() (v SmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputLibrarySourceUnion) AsSmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse() (v SmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputLibrarySourceUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreStopResponseInputLibrarySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse struct {
	Type constant.URL `json:"type" default:"url"`
	URL  string       `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputLibrarySourceURLSourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputLibrarySourceFileOutputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type SmallMoleculeExploreStopResponseInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeExploreStopResponseInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeExploreStopResponseInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeExploreStopResponseInputTargetConstraintUnion `json:"constraints"`
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
func (r SmallMoleculeExploreStopResponseInputTarget) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStopResponseInputTargetEntityUnion contains all possible
// properties and values from
// [SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse],
// [SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse].
	Modifications []SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse].
	Bonds []SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse].
	Residues []SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u SmallMoleculeExploreStopResponseInputTargetEntityUnion) AsSmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse() (v SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetEntityUnion) AsSmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse() (v SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputTargetEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *SmallMoleculeExploreStopResponseInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponseModification struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
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
func (r SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type SmallMoleculeExploreStopResponseInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeExploreStopResponseInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeExploreStopResponseInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStopResponseInputTargetBondAtom1Union contains all possible
// properties and values from
// [SmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse],
// [SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse],
// [SmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse],
// [SmallMoleculeExploreStopResponseInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse].
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

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom1Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom1LigandAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreStopResponseInputTargetBondAtom1LigandAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStopResponseInputTargetBondAtom2Union contains all possible
// properties and values from
// [SmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse],
// [SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse],
// [SmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse],
// [SmallMoleculeExploreStopResponseInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse].
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

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetBondAtom2Union) AsSmallMoleculeExploreStopResponseInputTargetBondAtom2LigandAtomResponse() (v SmallMoleculeExploreStopResponseInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type SmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type SmallMoleculeExploreStopResponseInputTargetBondAtom2LigandAtomResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStopResponseInputTargetConstraintUnion contains all possible
// properties and values from
// [SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse],
// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputTargetConstraintUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse].
	Token1 SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse].
	Token2 SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u SmallMoleculeExploreStopResponseInputTargetConstraintUnion) AsSmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse() (v SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetConstraintUnion) AsSmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse() (v SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStopResponseInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union) AsSmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union) AsSmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeExploreStopResponseInputMoleculeFilters struct {
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
	BoltzSmartsCatalogFilterLevel SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseInputMoleculeFilters) UnmarshalJSON(data []byte) error {
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
type SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeExploreStopResponseInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion contains
// all possible properties and values from
// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                          `json:"patterns"`
	// This field is from variant
	// [SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) AsSmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *SmallMoleculeExploreStopResponseInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseProgress struct {
	// Molecules that reached terminal failure. These do not consume budget — each is
	// replaced by another selection.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Molecules that produced a usable result. The run completes when this reaches the
	// budget.
	NumMoleculesScored int64 `json:"num_molecules_scored" api:"required"`
	// The requested budget: how many of the library will be scored.
	TotalMoleculesToScore int64 `json:"total_molecules_to_score" api:"required"`
	// ID of the most recently scored result
	LatestResultID string `json:"latest_result_id"`
	// Distinct molecules accepted after validation, filtering and de-duplication.
	// Omitted while the submitted library is being prepared.
	LibrarySize      int64                                                    `json:"library_size"`
	RejectionSummary SmallMoleculeExploreStopResponseProgressRejectionSummary `json:"rejection_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumMoleculesFailed    respjson.Field
		NumMoleculesScored    respjson.Field
		TotalMoleculesToScore respjson.Field
		LatestResultID        respjson.Field
		LibrarySize           respjson.Field
		RejectionSummary      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseProgressRejectionSummary struct {
	// Number of submitted rows that collapsed onto a molecule already in the library.
	// Exploration works over distinct molecules, so duplicates are merged rather than
	// scored twice.
	DuplicateCount int64 `json:"duplicate_count" api:"required"`
	// Number of submitted molecules removed by server-side filtering rules.
	FilteredCount int64 `json:"filtered_count" api:"required"`
	// Number of submitted molecules rejected as invalid input.
	InvalidCount int64 `json:"invalid_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateCount respjson.Field
		FilteredCount  respjson.Field
		InvalidCount   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SmallMoleculeExploreStopResponseProgressRejectionSummary) RawJSON() string { return r.JSON.raw }
func (r *SmallMoleculeExploreStopResponseProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SmallMoleculeExploreStopResponseStatus string

const (
	SmallMoleculeExploreStopResponseStatusPending   SmallMoleculeExploreStopResponseStatus = "pending"
	SmallMoleculeExploreStopResponseStatusRunning   SmallMoleculeExploreStopResponseStatus = "running"
	SmallMoleculeExploreStopResponseStatusSucceeded SmallMoleculeExploreStopResponseStatus = "succeeded"
	SmallMoleculeExploreStopResponseStatusFailed    SmallMoleculeExploreStopResponseStatus = "failed"
	SmallMoleculeExploreStopResponseStatusStopped   SmallMoleculeExploreStopResponseStatus = "stopped"
)

type SmallMoleculeExploreGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SmallMoleculeExploreGetParams]'s query parameters as
// `url.Values`.
func (r SmallMoleculeExploreGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeExploreListResultsParams struct {
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

// URLQuery serializes [SmallMoleculeExploreListResultsParams]'s query parameters
// as `url.Values`.
func (r SmallMoleculeExploreListResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SmallMoleculeExploreStartParams struct {
	// How many molecules to score. Must not exceed the accepted library size or
	// 5,000,000.
	Budget int64 `json:"budget" api:"required"`
	// CSV or TSV molecule library, limited to 375 MiB and 5,000,000 data records. URL
	// sources can use the full file limit. Base64 sources are also subject to the
	// API's 50 MiB JSON request-body limit, so use a URL source for larger files. The
	// file must be UTF-8 and may contain only the selected SMILES and ID columns;
	// column order does not matter. Candidate IDs are limited to 1,024 UTF-8 bytes;
	// missing or blank IDs default to the zero-based data-record index.
	Library SmallMoleculeExploreStartParamsLibrary `json:"library,omitzero" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target SmallMoleculeExploreStartParamsTarget `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters SmallMoleculeExploreStartParamsMoleculeFilters `json:"molecule_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParams) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CSV or TSV molecule library, limited to 375 MiB and 5,000,000 data records. URL
// sources can use the full file limit. Base64 sources are also subject to the
// API's 50 MiB JSON request-body limit, so use a URL source for larger files. The
// file must be UTF-8 and may contain only the selected SMILES and ID columns;
// column order does not matter. Candidate IDs are limited to 1,024 UTF-8 bytes;
// missing or blank IDs default to the zero-based data-record index.
//
// The properties Format, Source are required.
type SmallMoleculeExploreStartParamsLibrary struct {
	// Delimited-text format of the molecule library.
	//
	// Any of "csv", "tsv".
	Format SmallMoleculeExploreStartParamsLibraryFormat `json:"format,omitzero" api:"required"`
	// How to provide a file to the API
	Source SmallMoleculeExploreStartParamsLibrarySourceUnion `json:"source,omitzero" api:"required"`
	// Column containing candidate IDs. When omitted, a distinct `id` column is used if
	// present; otherwise IDs are generated from zero-based data-record indexes.
	IDColumn param.Opt[string] `json:"id_column,omitzero"`
	// Column containing SMILES strings. Defaults to `smiles` when omitted.
	SmilesColumn param.Opt[string] `json:"smiles_column,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsLibrary) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsLibrary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsLibrary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Delimited-text format of the molecule library.
type SmallMoleculeExploreStartParamsLibraryFormat string

const (
	SmallMoleculeExploreStartParamsLibraryFormatCsv SmallMoleculeExploreStartParamsLibraryFormat = "csv"
	SmallMoleculeExploreStartParamsLibraryFormatTsv SmallMoleculeExploreStartParamsLibraryFormat = "tsv"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsLibrarySourceUnion struct {
	OfSmallMoleculeExploreStartsLibrarySourceURLSource    *SmallMoleculeExploreStartParamsLibrarySourceURLSource    `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsLibrarySourceBase64Source *SmallMoleculeExploreStartParamsLibrarySourceBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsLibrarySourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsLibrarySourceURLSource, u.OfSmallMoleculeExploreStartsLibrarySourceBase64Source)
}
func (u *SmallMoleculeExploreStartParamsLibrarySourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type SmallMoleculeExploreStartParamsLibrarySourceURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsLibrarySourceURLSource) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsLibrarySourceURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsLibrarySourceURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type SmallMoleculeExploreStartParamsLibrarySourceBase64Source struct {
	// Base64-encoded file contents
	Data string `json:"data" api:"required"`
	// MIME type (e.g., text/csv)
	MediaType string `json:"media_type" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsLibrarySourceBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsLibrarySourceBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsLibrarySourceBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
//
// The property Entities is required.
type SmallMoleculeExploreStartParamsTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []SmallMoleculeExploreStartParamsTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []SmallMoleculeExploreStartParamsTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []SmallMoleculeExploreStartParamsTargetConstraintUnion `json:"constraints,omitzero"`
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

func (r SmallMoleculeExploreStartParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeExploreStartParamsTarget](
		"type", "no_template",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsTargetEntityUnion struct {
	OfSmallMoleculeExploreStartsTargetEntityProteinEntity *SmallMoleculeExploreStartParamsTargetEntityProteinEntity `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetEntityGlycanEntity  *SmallMoleculeExploreStartParamsTargetEntityGlycanEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsTargetEntityProteinEntity, u.OfSmallMoleculeExploreStartsTargetEntityGlycanEntity)
}
func (u *SmallMoleculeExploreStartParamsTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type SmallMoleculeExploreStartParamsTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []SmallMoleculeExploreStartParamsTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type SmallMoleculeExploreStartParamsTargetEntityProteinEntityModification struct {
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

func (r SmallMoleculeExploreStartParamsTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type SmallMoleculeExploreStartParamsTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []SmallMoleculeExploreStartParamsTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBond struct {
	Atom1 SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type SmallMoleculeExploreStartParamsTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type SmallMoleculeExploreStartParamsTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 SmallMoleculeExploreStartParamsTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 SmallMoleculeExploreStartParamsTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsTargetBondAtom1Union struct {
	OfSmallMoleculeExploreStartsTargetBondAtom1PolymerAtom *SmallMoleculeExploreStartParamsTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetBondAtom1CcdAtom     *SmallMoleculeExploreStartParamsTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetBondAtom1SmilesAtom  *SmallMoleculeExploreStartParamsTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetBondAtom1LigandAtom  *SmallMoleculeExploreStartParamsTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsTargetBondAtom1PolymerAtom, u.OfSmallMoleculeExploreStartsTargetBondAtom1CcdAtom, u.OfSmallMoleculeExploreStartsTargetBondAtom1SmilesAtom, u.OfSmallMoleculeExploreStartsTargetBondAtom1LigandAtom)
}
func (u *SmallMoleculeExploreStartParamsTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom1PolymerAtom struct {
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

func (r SmallMoleculeExploreStartParamsTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom1CcdAtom struct {
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

func (r SmallMoleculeExploreStartParamsTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom1LigandAtom struct {
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

func (r SmallMoleculeExploreStartParamsTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsTargetBondAtom2Union struct {
	OfSmallMoleculeExploreStartsTargetBondAtom2PolymerAtom *SmallMoleculeExploreStartParamsTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetBondAtom2CcdAtom     *SmallMoleculeExploreStartParamsTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetBondAtom2SmilesAtom  *SmallMoleculeExploreStartParamsTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetBondAtom2LigandAtom  *SmallMoleculeExploreStartParamsTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsTargetBondAtom2PolymerAtom, u.OfSmallMoleculeExploreStartsTargetBondAtom2CcdAtom, u.OfSmallMoleculeExploreStartsTargetBondAtom2SmilesAtom, u.OfSmallMoleculeExploreStartsTargetBondAtom2LigandAtom)
}
func (u *SmallMoleculeExploreStartParamsTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom2PolymerAtom struct {
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

func (r SmallMoleculeExploreStartParamsTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom2CcdAtom struct {
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

func (r SmallMoleculeExploreStartParamsTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeExploreStartParamsTargetBondAtom2LigandAtom struct {
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

func (r SmallMoleculeExploreStartParamsTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsTargetConstraintUnion struct {
	OfSmallMoleculeExploreStartsTargetConstraintPocketConstraint  *SmallMoleculeExploreStartParamsTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetConstraintContactConstraint *SmallMoleculeExploreStartParamsTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsTargetConstraintPocketConstraint, u.OfSmallMoleculeExploreStartsTargetConstraintContactConstraint)
}
func (u *SmallMoleculeExploreStartParamsTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type SmallMoleculeExploreStartParamsTargetConstraintPocketConstraint struct {
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

func (r SmallMoleculeExploreStartParamsTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1Union struct {
	OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken1PolymerContactToken *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken1LigandContactToken  *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken1PolymerContactToken, u.OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1LigandContactToken struct {
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

func (r SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2Union struct {
	OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken2PolymerContactToken *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken2LigandContactToken  *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken2PolymerContactToken, u.OfSmallMoleculeExploreStartsTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2LigandContactToken struct {
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

func (r SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type SmallMoleculeExploreStartParamsMoleculeFilters struct {
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
	BoltzSmartsCatalogFilterLevel SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level,omitzero"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterUnion `json:"custom_filters,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFilters) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFilters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFilters) UnmarshalJSON(data []byte) error {
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
type SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    SmallMoleculeExploreStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterUnion struct {
	OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterLipinskiFilter        *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterLipinskiFilter        `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterRdkitDescriptorFilter *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterSmartsCustomFilter    *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter    `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterSmartsCatalogFilter   *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter   `json:",omitzero,inline"`
	OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterSmilesRegexFilter     *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter     `json:",omitzero,inline"`
	paramUnion
}

func (u SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterLipinskiFilter,
		u.OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterRdkitDescriptorFilter,
		u.OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterSmartsCustomFilter,
		u.OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterSmartsCatalogFilter,
		u.OfSmallMoleculeExploreStartsMoleculeFiltersCustomFilterSmilesRegexFilter)
}
func (u *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
//
// The properties MaxHba, MaxHbd, MaxLogp, MaxMw, Type are required.
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterLipinskiFilter struct {
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

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterLipinskiFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterLipinskiFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterLipinskiFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
//
// The property Type is required.
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter struct {
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 `json:"fraction_csp3,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp `json:"mol_logp,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt `json:"mol_wt,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings `json:"num_aromatic_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors `json:"num_h_acceptors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors `json:"num_h_donors,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms `json:"num_heteroatoms,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings `json:"num_rings,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds `json:"num_rotatable_bonds,omitzero"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa `json:"tpsa,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "rdkit_descriptor_filter".
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3 struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa struct {
	// Maximum allowed value (inclusive)
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum allowed value (inclusive)
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterRdkitDescriptorFilterTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
//
// The properties Patterns, Type are required.
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter struct {
	// SMARTS patterns. Molecules matching any pattern are rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smarts_custom_filter".
	Type constant.SmartsCustomFilter `json:"type" default:"smarts_custom_filter"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCustomFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
//
// The properties Catalog, Type are required.
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter struct {
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

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmartsCatalogFilter](
		"catalog", "PAINS", "PAINS_A", "PAINS_B", "PAINS_C", "BRENK", "CHEMBL", "CHEMBL_BMS", "CHEMBL_Dundee", "CHEMBL_Glaxo", "CHEMBL_Inpharmatica", "CHEMBL_LINT", "CHEMBL_MLSMR", "CHEMBL_SureChEMBL", "NIH",
	)
}

// Filter molecules by regex patterns on their SMILES representation.
//
// The properties Patterns, Type are required.
type SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter struct {
	// Regex patterns applied to SMILES strings. Molecules matching any pattern are
	// rejected.
	Patterns []string `json:"patterns,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "smiles_regex_filter".
	Type constant.SmilesRegexFilter `json:"type" default:"smiles_regex_filter"`
	paramObj
}

func (r SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter) MarshalJSON() (data []byte, err error) {
	type shadow SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SmallMoleculeExploreStartParamsMoleculeFiltersCustomFilterSmilesRegexFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
