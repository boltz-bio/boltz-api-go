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

// Screen an existing library of proteins against a target structure. Results are
// scored by binding confidence (likelihood of protein-protein interaction) and
// structure confidence.
//
// ProteinLibraryScreenService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProteinLibraryScreenService] method instead.
type ProteinLibraryScreenService struct {
	Options []option.RequestOption
}

// NewProteinLibraryScreenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProteinLibraryScreenService(opts ...option.RequestOption) (r ProteinLibraryScreenService) {
	r = ProteinLibraryScreenService{}
	r.Options = opts
	return
}

// Retrieve a library screen by ID, including progress and status
func (r *ProteinLibraryScreenService) Get(ctx context.Context, id string, query ProteinLibraryScreenGetParams, opts ...option.RequestOption) (res *ProteinLibraryScreenGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/library-screen/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List protein library screens, optionally filtered by workspace
func (r *ProteinLibraryScreenService) List(ctx context.Context, query ProteinLibraryScreenListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProteinLibraryScreenListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/protein/library-screen"
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

// List protein library screens, optionally filtered by workspace
func (r *ProteinLibraryScreenService) ListAutoPaging(ctx context.Context, query ProteinLibraryScreenListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProteinLibraryScreenListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// library screen. The library screen record itself is retained with a
// `data_deleted_at` timestamp. This action is irreversible.
func (r *ProteinLibraryScreenService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinLibraryScreenDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/library-screen/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Estimate the cost of a protein library screen without creating any resource or
// consuming GPU.
func (r *ProteinLibraryScreenService) EstimateCost(ctx context.Context, body ProteinLibraryScreenEstimateCostParams, opts ...option.RequestOption) (res *ProteinLibraryScreenEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/protein/library-screen/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve paginated results from a protein library screen
func (r *ProteinLibraryScreenService) ListResults(ctx context.Context, id string, query ProteinLibraryScreenListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProteinLibraryScreenListResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/library-screen/%s/results", url.PathEscape(id))
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

// Retrieve paginated results from a protein library screen
func (r *ProteinLibraryScreenService) ListResultsAutoPaging(ctx context.Context, id string, query ProteinLibraryScreenListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProteinLibraryScreenListResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, id, query, opts...))
}

// Screen a set of protein candidates against a target
func (r *ProteinLibraryScreenService) Start(ctx context.Context, body ProteinLibraryScreenStartParams, opts ...option.RequestOption) (res *ProteinLibraryScreenStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/protein/library-screen"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an in-progress protein library screen early
func (r *ProteinLibraryScreenService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinLibraryScreenStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/library-screen/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A protein library screening engine run
type ProteinLibraryScreenGetResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein library screen
	Engine constant.BoltzProteinScreen `json:"engine" default:"boltz-protein-screen"`
	// Engine version used for protein library screen
	EngineVersion string                               `json:"engine_version" api:"required"`
	Error         ProteinLibraryScreenGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                                    `json:"livemode" api:"required"`
	Progress  ProteinLibraryScreenGetResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                               `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenGetResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                             `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DataDeletedAt  respjson.Field
		Engine         respjson.Field
		EngineVersion  respjson.Field
		Error          respjson.Field
		Input          respjson.Field
		Livemode       respjson.Field
		Progress       respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		StoppedAt      respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseError struct {
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
func (r ProteinLibraryScreenGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinLibraryScreenGetResponseInput struct {
	Proteins ProteinLibraryScreenGetResponseInputProteins `json:"proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinLibraryScreenGetResponseInputTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Proteins    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputProteins struct {
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
func (r ProteinLibraryScreenGetResponseInputProteins) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenGetResponseInputProteins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetUnion contains all possible properties
// and values from
// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                             `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse].
	NonBindingResidues map[string][]int64 `json:"non_binding_residues"`
	JSON               struct {
		ChainSelection      respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		Entities            respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetUnion) AsProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse() (v ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinLibraryScreenGetResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                              `json:"type" default:"structure_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainSelection respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	NonBindingResidues []int64 `json:"non_binding_residues"`
	JSON               struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		raw                string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// 0-indexed residue indices where binder contact is desired (the epitope). All
	// indices must be present in crop_residues and must not overlap
	// non_binding_residues.
	EpitopeResidues []int64 `json:"epitope_residues"`
	// 0-indexed residue indices allowed to move during design (e.g. flexible loop
	// regions). All indices must be present in crop_residues.
	FlexibleResidues []int64 `json:"flexible_residues"`
	// 0-indexed residue indices where binder contact should be discouraged. All
	// indices must be present in crop_residues and must not overlap epitope_residues.
	NonBindingResidues []int64 `json:"non_binding_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	// This field will be present if the value is a [[]int64] instead of an object.
	OfIntArray []int64 `json:",inline"`
	// This field will be present if the value is a [constant.All] instead of an
	// object.
	OfAll constant.All `json:",inline"`
	JSON  struct {
		OfIntArray respjson.Field
		OfAll      respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                             `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// Chain IDs of ligand entities that are part of the binding epitope. Ligands are
	// marked as epitope in full (no residue-level selection).
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// Polymer chain residues where binder contact is desired (the epitope). Each key
	// is a chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap non_binding_residues on the same chain.
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// Polymer chain residues where binder contact should be discouraged. Each key is a
	// chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap epitope_residues on the same chain.
	NonBindingResidues map[string][]int64 `json:"non_binding_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		Type                respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion],
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion],
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion]
	Modifications ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion `json:",inline"`
	JSON                                                                                                     struct {
		OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                          string
	} `json:"-"`
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string              `json:"chain_ids" api:"required"`
	Type     constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                 `json:"type" default:"contact"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseProgress struct {
	// Number of accepted proteins that reached terminal failure during screening.
	NumProteinsFailed int64 `json:"num_proteins_failed" api:"required"`
	// Number of accepted proteins that produced usable screening results.
	NumProteinsScreened int64 `json:"num_proteins_screened" api:"required"`
	// Total number of proteins accepted into the screening run.
	TotalProteinsToScreen int64 `json:"total_proteins_to_screen" api:"required"`
	// ID of the latest result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsFailed     respjson.Field
		NumProteinsScreened   respjson.Field
		TotalProteinsToScreen respjson.Field
		LatestResultID        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseStatus string

const (
	ProteinLibraryScreenGetResponseStatusPending   ProteinLibraryScreenGetResponseStatus = "pending"
	ProteinLibraryScreenGetResponseStatusRunning   ProteinLibraryScreenGetResponseStatus = "running"
	ProteinLibraryScreenGetResponseStatusSucceeded ProteinLibraryScreenGetResponseStatus = "succeeded"
	ProteinLibraryScreenGetResponseStatusFailed    ProteinLibraryScreenGetResponseStatus = "failed"
	ProteinLibraryScreenGetResponseStatusStopped   ProteinLibraryScreenGetResponseStatus = "stopped"
)

// Summary of a protein library screening engine run (excludes input)
type ProteinLibraryScreenListResponse struct {
	// Unique ProteinLibraryScreenSummary identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein library screen
	Engine constant.BoltzProteinScreen `json:"engine" default:"boltz-protein-screen"`
	// Engine version used for protein library screen
	EngineVersion string                                `json:"engine_version" api:"required"`
	Error         ProteinLibraryScreenListResponseError `json:"error" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                                     `json:"livemode" api:"required"`
	Progress  ProteinLibraryScreenListResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenListResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                              `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DataDeletedAt  respjson.Field
		Engine         respjson.Field
		EngineVersion  respjson.Field
		Error          respjson.Field
		Livemode       respjson.Field
		Progress       respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		StoppedAt      respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResponseError struct {
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
func (r ProteinLibraryScreenListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResponseProgress struct {
	// Number of accepted proteins that reached terminal failure during screening.
	NumProteinsFailed int64 `json:"num_proteins_failed" api:"required"`
	// Number of accepted proteins that produced usable screening results.
	NumProteinsScreened int64 `json:"num_proteins_screened" api:"required"`
	// Total number of proteins accepted into the screening run.
	TotalProteinsToScreen int64 `json:"total_proteins_to_screen" api:"required"`
	// ID of the latest result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsFailed     respjson.Field
		NumProteinsScreened   respjson.Field
		TotalProteinsToScreen respjson.Field
		LatestResultID        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResponseStatus string

const (
	ProteinLibraryScreenListResponseStatusPending   ProteinLibraryScreenListResponseStatus = "pending"
	ProteinLibraryScreenListResponseStatusRunning   ProteinLibraryScreenListResponseStatus = "running"
	ProteinLibraryScreenListResponseStatusSucceeded ProteinLibraryScreenListResponseStatus = "succeeded"
	ProteinLibraryScreenListResponseStatusFailed    ProteinLibraryScreenListResponseStatus = "failed"
	ProteinLibraryScreenListResponseStatusStopped   ProteinLibraryScreenListResponseStatus = "stopped"
)

type ProteinLibraryScreenDeleteDataResponse struct {
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
func (r ProteinLibraryScreenDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimate response with monetary values encoded as decimal strings to preserve
// precision.
type ProteinLibraryScreenEstimateCostResponse struct {
	// Cost breakdown for the billed application.
	Breakdown  ProteinLibraryScreenEstimateCostResponseBreakdown `json:"breakdown" api:"required"`
	Disclaimer string                                            `json:"disclaimer" api:"required"`
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
func (r ProteinLibraryScreenEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost breakdown for the billed application.
type ProteinLibraryScreenEstimateCostResponseBreakdown struct {
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_library_screen",
	// "adme".
	Application ProteinLibraryScreenEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
	// Estimated cost per displayed unit as a decimal string, rounded up to 4 decimal
	// places. This may include token-size multipliers or generation overhead;
	// estimated_cost_usd is the authoritative total.
	CostPerUnitUsd string `json:"cost_per_unit_usd" api:"required"`
	// Number of units shown for the estimate. For structure-and-binding, this is the
	// requested number of samples. For protein and small-molecule design/screen
	// endpoints, this is the requested number of proteins or molecules.
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
func (r ProteinLibraryScreenEstimateCostResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenEstimateCostResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenEstimateCostResponseBreakdownApplication string

const (
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationStructureAndBinding        ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "structure_and_binding"
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationSmallMoleculeDesign        ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "small_molecule_design"
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationSmallMoleculeLibraryScreen ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "small_molecule_library_screen"
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationProteinDesign              ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "protein_design"
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationProteinLibraryScreen       ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "protein_library_screen"
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationAdme                       ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "adme"
)

// Result for a single screened protein
type ProteinLibraryScreenListResultsResponse struct {
	// Unique result ID
	ID        string                                           `json:"id" api:"required"`
	Artifacts ProteinLibraryScreenListResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                        `json:"created_at" api:"required" format:"date-time"`
	// Entities of the screened complex. Includes both screened and fixed entities from
	// the input.
	Entities []ProteinLibraryScreenListResultsResponseEntityUnion `json:"entities" api:"required"`
	// Structural and binding quality metrics for a screened protein
	Metrics ProteinLibraryScreenListResultsResponseMetrics `json:"metrics" api:"required"`
	// Client-provided identifier for this protein, if provided
	ExternalID string `json:"external_id"`
	// Warnings about potential quality issues with this result.
	Warnings []ProteinLibraryScreenListResultsResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Entities    respjson.Field
		Metrics     respjson.Field
		ExternalID  respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseArtifacts struct {
	Archive   ProteinLibraryScreenListResultsResponseArtifactsArchive   `json:"archive" api:"required"`
	Structure ProteinLibraryScreenListResultsResponseArtifactsStructure `json:"structure" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archive     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResultsResponseArtifacts) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponseArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseArtifactsArchive struct {
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
func (r ProteinLibraryScreenListResultsResponseArtifactsArchive) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponseArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseArtifactsStructure struct {
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
func (r ProteinLibraryScreenListResultsResponseArtifactsStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenListResultsResponseEntityUnion contains all possible
// properties and values from
// [ProteinLibraryScreenListResultsResponseEntityProteinEntity],
// [ProteinLibraryScreenListResultsResponseEntityRnaEntity],
// [ProteinLibraryScreenListResultsResponseEntityDnaEntity],
// [ProteinLibraryScreenListResultsResponseEntityLigandCcdEntity],
// [ProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenListResultsResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion],
	// [[]ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion],
	// [[]ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion]
	Modifications ProteinLibraryScreenListResultsResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinLibraryScreenListResultsResponseEntityUnion) AsProteinLibraryScreenListResultsResponseEntityProteinEntity() (v ProteinLibraryScreenListResultsResponseEntityProteinEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityUnion) AsProteinLibraryScreenListResultsResponseEntityRnaEntity() (v ProteinLibraryScreenListResultsResponseEntityRnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityUnion) AsProteinLibraryScreenListResultsResponseEntityDnaEntity() (v ProteinLibraryScreenListResultsResponseEntityDnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityUnion) AsProteinLibraryScreenListResultsResponseEntityLigandCcdEntity() (v ProteinLibraryScreenListResultsResponseEntityLigandCcdEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityUnion) AsProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity() (v ProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenListResultsResponseEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinLibraryScreenListResultsResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenListResultsResponseEntityUnionModifications is an implicit
// subunion of [ProteinLibraryScreenListResultsResponseEntityUnion].
// ProteinLibraryScreenListResultsResponseEntityUnionModifications provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinLibraryScreenListResultsResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinLibraryScreenListResultsResponseEntityProteinEntityModifications
// OfProteinLibraryScreenListResultsResponseEntityRnaEntityModifications
// OfProteinLibraryScreenListResultsResponseEntityDnaEntityModifications]
type ProteinLibraryScreenListResultsResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenListResultsResponseEntityProteinEntityModifications []ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenListResultsResponseEntityRnaEntityModifications []ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenListResultsResponseEntityDnaEntityModifications []ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion `json:",inline"`
	JSON                                                                  struct {
		OfProteinLibraryScreenListResultsResponseEntityProteinEntityModifications respjson.Field
		OfProteinLibraryScreenListResultsResponseEntityRnaEntityModifications     respjson.Field
		OfProteinLibraryScreenListResultsResponseEntityDnaEntityModifications     respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (r *ProteinLibraryScreenListResultsResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenListResultsResponseEntityProteinEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationCcdModification],
// [ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationSmilesModification].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion) AsProteinLibraryScreenListResultsResponseEntityProteinEntityModificationCcdModification() (v ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationCcdModification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion) AsProteinLibraryScreenListResultsResponseEntityProteinEntityModificationSmilesModification() (v ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationSmilesModification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationCcdModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationSmilesModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityProteinEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenListResultsResponseEntityRnaEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponseEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion contains
// all possible properties and values from
// [ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationCcdModification],
// [ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationSmilesModification].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion) AsProteinLibraryScreenListResultsResponseEntityRnaEntityModificationCcdModification() (v ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationCcdModification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion) AsProteinLibraryScreenListResultsResponseEntityRnaEntityModificationSmilesModification() (v ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationSmilesModification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationCcdModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationSmilesModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityRnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenListResultsResponseEntityDnaEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponseEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion contains
// all possible properties and values from
// [ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationCcdModification],
// [ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationSmilesModification].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion) AsProteinLibraryScreenListResultsResponseEntityDnaEntityModificationCcdModification() (v ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationCcdModification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion) AsProteinLibraryScreenListResultsResponseEntityDnaEntityModificationSmilesModification() (v ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationSmilesModification) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationCcdModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationSmilesModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityDnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResultsResponseEntityLigandCcdEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string              `json:"chain_ids" api:"required"`
	Type     constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structural and binding quality metrics for a screened protein
type ProteinLibraryScreenListResultsResponseMetrics struct {
	// Confidence that the protein binds the target (0-1). Primary metric for hit
	// discovery.
	BindingConfidence float64 `json:"binding_confidence" api:"required"`
	// Fraction of the sequence forming alpha helices (0-1).
	HelixFraction float64 `json:"helix_fraction" api:"required"`
	// Interface predicted TM score (0-1). Confidence in the protein-protein interface.
	Iptm float64 `json:"iptm" api:"required"`
	// Fraction of the sequence in coil/loop regions (0-1).
	LoopFraction float64 `json:"loop_fraction" api:"required"`
	// Minimum predicted aligned error at the interface (Angstroms). Lower values
	// indicate higher confidence.
	MinInteractionPae float64 `json:"min_interaction_pae" api:"required"`
	// Fraction of the sequence forming beta sheets (0-1).
	SheetFraction float64 `json:"sheet_fraction" api:"required"`
	// Confidence in the predicted 3D structure (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence   respjson.Field
		HelixFraction       respjson.Field
		Iptm                respjson.Field
		LoopFraction        respjson.Field
		MinInteractionPae   respjson.Field
		SheetFraction       respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResultsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warning about a potential quality issue with a result
type ProteinLibraryScreenListResultsResponseWarning struct {
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
func (r ProteinLibraryScreenListResultsResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenListResultsResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A protein library screening engine run
type ProteinLibraryScreenStartResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein library screen
	Engine constant.BoltzProteinScreen `json:"engine" default:"boltz-protein-screen"`
	// Engine version used for protein library screen
	EngineVersion string                                 `json:"engine_version" api:"required"`
	Error         ProteinLibraryScreenStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                                      `json:"livemode" api:"required"`
	Progress  ProteinLibraryScreenStartResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                                 `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenStartResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                               `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DataDeletedAt  respjson.Field
		Engine         respjson.Field
		EngineVersion  respjson.Field
		Error          respjson.Field
		Input          respjson.Field
		Livemode       respjson.Field
		Progress       respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		StoppedAt      respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseError struct {
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
func (r ProteinLibraryScreenStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinLibraryScreenStartResponseInput struct {
	Proteins ProteinLibraryScreenStartResponseInputProteins `json:"proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinLibraryScreenStartResponseInputTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Proteins    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputProteins struct {
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
func (r ProteinLibraryScreenStartResponseInputProteins) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStartResponseInputProteins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetUnion contains all possible
// properties and values from
// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                               `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse].
	NonBindingResidues map[string][]int64 `json:"non_binding_residues"`
	JSON               struct {
		ChainSelection      respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		Entities            respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetUnion) AsProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse() (v ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinLibraryScreenStartResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                                `json:"type" default:"structure_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainSelection respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	NonBindingResidues []int64 `json:"non_binding_residues"`
	JSON               struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		raw                string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// 0-indexed residue indices where binder contact is desired (the epitope). All
	// indices must be present in crop_residues and must not overlap
	// non_binding_residues.
	EpitopeResidues []int64 `json:"epitope_residues"`
	// 0-indexed residue indices allowed to move during design (e.g. flexible loop
	// regions). All indices must be present in crop_residues.
	FlexibleResidues []int64 `json:"flexible_residues"`
	// 0-indexed residue indices where binder contact should be discouraged. All
	// indices must be present in crop_residues and must not overlap epitope_residues.
	NonBindingResidues []int64 `json:"non_binding_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	// This field will be present if the value is a [[]int64] instead of an object.
	OfIntArray []int64 `json:",inline"`
	// This field will be present if the value is a [constant.All] instead of an
	// object.
	OfAll constant.All `json:",inline"`
	JSON  struct {
		OfIntArray respjson.Field
		OfAll      respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                               `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// Chain IDs of ligand entities that are part of the binding epitope. Ligands are
	// marked as epitope in full (no residue-level selection).
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// Polymer chain residues where binder contact is desired (the epitope). Each key
	// is a chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap non_binding_residues on the same chain.
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// Polymer chain residues where binder contact should be discouraged. Each key is a
	// chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap epitope_residues on the same chain.
	NonBindingResidues map[string][]int64 `json:"non_binding_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		Type                respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion],
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion],
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion]
	Modifications ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion `json:",inline"`
	JSON                                                                                                       struct {
		OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                            string
	} `json:"-"`
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string              `json:"chain_ids" api:"required"`
	Type     constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                   `json:"type" default:"contact"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseProgress struct {
	// Number of accepted proteins that reached terminal failure during screening.
	NumProteinsFailed int64 `json:"num_proteins_failed" api:"required"`
	// Number of accepted proteins that produced usable screening results.
	NumProteinsScreened int64 `json:"num_proteins_screened" api:"required"`
	// Total number of proteins accepted into the screening run.
	TotalProteinsToScreen int64 `json:"total_proteins_to_screen" api:"required"`
	// ID of the latest result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsFailed     respjson.Field
		NumProteinsScreened   respjson.Field
		TotalProteinsToScreen respjson.Field
		LatestResultID        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStartResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseStatus string

const (
	ProteinLibraryScreenStartResponseStatusPending   ProteinLibraryScreenStartResponseStatus = "pending"
	ProteinLibraryScreenStartResponseStatusRunning   ProteinLibraryScreenStartResponseStatus = "running"
	ProteinLibraryScreenStartResponseStatusSucceeded ProteinLibraryScreenStartResponseStatus = "succeeded"
	ProteinLibraryScreenStartResponseStatusFailed    ProteinLibraryScreenStartResponseStatus = "failed"
	ProteinLibraryScreenStartResponseStatusStopped   ProteinLibraryScreenStartResponseStatus = "stopped"
)

// A protein library screening engine run
type ProteinLibraryScreenStopResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein library screen
	Engine constant.BoltzProteinScreen `json:"engine" default:"boltz-protein-screen"`
	// Engine version used for protein library screen
	EngineVersion string                                `json:"engine_version" api:"required"`
	Error         ProteinLibraryScreenStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenStopResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                                     `json:"livemode" api:"required"`
	Progress  ProteinLibraryScreenStopResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenStopResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                              `json:"stopped_at" api:"required" format:"date-time"`
	// Workspace ID
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Client-provided idempotency key
	IdempotencyKey string `json:"idempotency_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CompletedAt    respjson.Field
		CreatedAt      respjson.Field
		DataDeletedAt  respjson.Field
		Engine         respjson.Field
		EngineVersion  respjson.Field
		Error          respjson.Field
		Input          respjson.Field
		Livemode       respjson.Field
		Progress       respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		StoppedAt      respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseError struct {
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
func (r ProteinLibraryScreenStopResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStopResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinLibraryScreenStopResponseInput struct {
	Proteins ProteinLibraryScreenStopResponseInputProteins `json:"proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinLibraryScreenStopResponseInputTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Proteins    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStopResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputProteins struct {
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
func (r ProteinLibraryScreenStopResponseInputProteins) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStopResponseInputProteins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetUnion contains all possible
// properties and values from
// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                              `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse].
	NonBindingResidues map[string][]int64 `json:"non_binding_residues"`
	JSON               struct {
		ChainSelection      respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		Entities            respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetUnion) AsProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse() (v ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinLibraryScreenStopResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                               `json:"type" default:"structure_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainSelection respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	NonBindingResidues []int64 `json:"non_binding_residues"`
	JSON               struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		raw                string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// 0-indexed residue indices where binder contact is desired (the epitope). All
	// indices must be present in crop_residues and must not overlap
	// non_binding_residues.
	EpitopeResidues []int64 `json:"epitope_residues"`
	// 0-indexed residue indices allowed to move during design (e.g. flexible loop
	// regions). All indices must be present in crop_residues.
	FlexibleResidues []int64 `json:"flexible_residues"`
	// 0-indexed residue indices where binder contact should be discouraged. All
	// indices must be present in crop_residues and must not overlap epitope_residues.
	NonBindingResidues []int64 `json:"non_binding_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	// This field will be present if the value is a [[]int64] instead of an object.
	OfIntArray []int64 `json:",inline"`
	// This field will be present if the value is a [constant.All] instead of an
	// object.
	OfAll constant.All `json:",inline"`
	JSON  struct {
		OfIntArray respjson.Field
		OfAll      respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                              `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// Chain IDs of ligand entities that are part of the binding epitope. Ligands are
	// marked as epitope in full (no residue-level selection).
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// Polymer chain residues where binder contact is desired (the epitope). Each key
	// is a chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap non_binding_residues on the same chain.
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// Polymer chain residues where binder contact should be discouraged. Each key is a
	// chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap epitope_residues on the same chain.
	NonBindingResidues map[string][]int64 `json:"non_binding_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		Type                respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion],
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion],
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion]
	Modifications ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion]
	// instead of an object.
	OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion `json:",inline"`
	JSON                                                                                                      struct {
		OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                           string
	} `json:"-"`
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion `json:"modifications"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion struct {
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	Value        string `json:"value"`
	JSON         struct {
		ResidueIndex respjson.Field
		Type         respjson.Field
		Value        respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64        `json:"residue_index" api:"required"`
	Type         constant.Ccd `json:"type" default:"ccd"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationCcdModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse struct {
	// 0-based index of the residue to modify
	ResidueIndex int64           `json:"residue_index" api:"required"`
	Type         constant.Smiles `json:"type" default:"smiles"`
	// SMILES string for the modification
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModificationSmilesModificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string              `json:"chain_ids" api:"required"`
	Type     constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                  `json:"type" default:"contact"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseProgress struct {
	// Number of accepted proteins that reached terminal failure during screening.
	NumProteinsFailed int64 `json:"num_proteins_failed" api:"required"`
	// Number of accepted proteins that produced usable screening results.
	NumProteinsScreened int64 `json:"num_proteins_screened" api:"required"`
	// Total number of proteins accepted into the screening run.
	TotalProteinsToScreen int64 `json:"total_proteins_to_screen" api:"required"`
	// ID of the latest result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsFailed     respjson.Field
		NumProteinsScreened   respjson.Field
		TotalProteinsToScreen respjson.Field
		LatestResultID        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenStopResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseStatus string

const (
	ProteinLibraryScreenStopResponseStatusPending   ProteinLibraryScreenStopResponseStatus = "pending"
	ProteinLibraryScreenStopResponseStatusRunning   ProteinLibraryScreenStopResponseStatus = "running"
	ProteinLibraryScreenStopResponseStatusSucceeded ProteinLibraryScreenStopResponseStatus = "succeeded"
	ProteinLibraryScreenStopResponseStatusFailed    ProteinLibraryScreenStopResponseStatus = "failed"
	ProteinLibraryScreenStopResponseStatusStopped   ProteinLibraryScreenStopResponseStatus = "stopped"
)

type ProteinLibraryScreenGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProteinLibraryScreenGetParams]'s query parameters as
// `url.Values`.
func (r ProteinLibraryScreenGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinLibraryScreenListParams struct {
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

// URLQuery serializes [ProteinLibraryScreenListParams]'s query parameters as
// `url.Values`.
func (r ProteinLibraryScreenListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinLibraryScreenEstimateCostParams struct {
	// List of protein entries to screen.
	Proteins []ProteinLibraryScreenEstimateCostParamsProtein `json:"proteins,omitzero" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinLibraryScreenEstimateCostParamsTargetUnion `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A protein screen entry with entities and optional ID
//
// The property Entities is required.
type ProteinLibraryScreenEstimateCostParamsProtein struct {
	// Entities that make up this protein complex
	Entities []ProteinLibraryScreenEstimateCostParamsProteinEntityUnion `json:"entities,omitzero" api:"required"`
	// Optional client-provided identifier for this entry
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProtein) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProtein
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProtein) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsProteinEntityUnion struct {
	OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity      *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity      `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntity          *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntity          *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityLigandCcdEntity    *ProteinLibraryScreenEstimateCostParamsProteinEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityLigandSmilesEntity *ProteinLibraryScreenEstimateCostParamsProteinEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsProteinEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityLigandCcdEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityLigandSmilesEntity)
}
func (u *ProteinLibraryScreenEstimateCostParamsProteinEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion struct {
	OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationCcdModification    *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationSmilesModification *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationCcdModification, u.OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationUnion struct {
	OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntityModificationCcdModification    *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntityModificationSmilesModification *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntityModificationCcdModification, u.OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationUnion struct {
	OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntityModificationCcdModification    *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntityModificationSmilesModification *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntityModificationCcdModification, u.OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTarget *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTarget        *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTarget        `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTarget, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTarget)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
//
// The properties ChainSelection, Structure, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion `json:"chain_selection,omitzero" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion `json:"structure,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "structure_template".
	Type constant.StructureTemplate `json:"type" default:"structure_template"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec, u.OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
//
// The properties ChainType, CropResidues are required.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues,omitzero" api:"required"`
	// 0-indexed residue indices where binder contact is desired (the epitope). All
	// indices must be present in crop_residues and must not overlap
	// non_binding_residues.
	EpitopeResidues []int64 `json:"epitope_residues,omitzero"`
	// 0-indexed residue indices allowed to move during design (e.g. flexible loop
	// regions). All indices must be present in crop_residues.
	FlexibleResidues []int64 `json:"flexible_residues,omitzero"`
	// 0-indexed residue indices where binder contact should be discouraged. All
	// indices must be present in crop_residues and must not overlap epitope_residues.
	NonBindingResidues []int64 `json:"non_binding_residues,omitzero"`
	// This field can be elided, and will marshal its zero value as "polymer".
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	OfIntArray []int64 `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIntArray, u.OfAll)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func NewProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec() ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec {
	return ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec{
		ChainType: "ligand",
	}
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
//
// This struct has a constant value, construct it with
// [NewProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec].
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureURLSource       *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource       `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureCifBase64Source *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureURLSource, u.OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureCifBase64Source)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source struct {
	// Base64-encoded CIF file contents
	Data string `json:"data" api:"required"`
	// Must be chemical/x-cif for CIF files
	//
	// This field can be elided, and will marshal its zero value as "chemical/x-cif".
	MediaType constant.ChemicalXCif `json:"media_type" default:"chemical/x-cif"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
//
// The properties Entities, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTarget struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintUnion `json:"constraints,omitzero"`
	// Chain IDs of ligand entities that are part of the binding epitope. Ligands are
	// marked as epitope in full (no residue-level selection).
	EpitopeLigandChains []string `json:"epitope_ligand_chains,omitzero"`
	// Polymer chain residues where binder contact is desired (the epitope). Each key
	// is a chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap non_binding_residues on the same chain.
	EpitopeResidues map[string][]int64 `json:"epitope_residues,omitzero"`
	// Polymer chain residues where binder contact should be discouraged. Each key is a
	// chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap epitope_residues on the same chain.
	NonBindingResidues map[string][]int64 `json:"non_binding_residues,omitzero"`
	// This field can be elided, and will marshal its zero value as "no_template".
	Type constant.NoTemplate `json:"type" default:"no_template"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntity      *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity      `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntity          *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntity          *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityLigandCcdEntity    *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityLigandSmilesEntity *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityLigandCcdEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityLigandSmilesEntity)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification    *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification    *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification    *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1LigandAtom  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1PolymerAtom *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1LigandAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1PolymerAtom)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2LigandAtom  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2PolymerAtom *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2LigandAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2PolymerAtom)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintUnion struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintPocketConstraint  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraint *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintPocketConstraint, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraint)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsParams struct {
	// Return results after this ID
	AfterID param.Opt[string] `query:"after_id,omitzero" json:"-"`
	// Return results before this ID
	BeforeID param.Opt[string] `query:"before_id,omitzero" json:"-"`
	// Max results to return. Defaults to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProteinLibraryScreenListResultsParams]'s query parameters
// as `url.Values`.
func (r ProteinLibraryScreenListResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinLibraryScreenStartParams struct {
	// List of protein entries to screen.
	Proteins []ProteinLibraryScreenStartParamsProtein `json:"proteins,omitzero" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinLibraryScreenStartParamsTargetUnion `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	paramObj
}

func (r ProteinLibraryScreenStartParams) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A protein screen entry with entities and optional ID
//
// The property Entities is required.
type ProteinLibraryScreenStartParamsProtein struct {
	// Entities that make up this protein complex
	Entities []ProteinLibraryScreenStartParamsProteinEntityUnion `json:"entities,omitzero" api:"required"`
	// Optional client-provided identifier for this entry
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProtein) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProtein
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProtein) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsProteinEntityUnion struct {
	OfProteinLibraryScreenStartsProteinEntityProteinEntity      *ProteinLibraryScreenStartParamsProteinEntityProteinEntity      `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityRnaEntity          *ProteinLibraryScreenStartParamsProteinEntityRnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityDnaEntity          *ProteinLibraryScreenStartParamsProteinEntityDnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityLigandCcdEntity    *ProteinLibraryScreenStartParamsProteinEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityLigandSmilesEntity *ProteinLibraryScreenStartParamsProteinEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsProteinEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsProteinEntityProteinEntity,
		u.OfProteinLibraryScreenStartsProteinEntityRnaEntity,
		u.OfProteinLibraryScreenStartsProteinEntityDnaEntity,
		u.OfProteinLibraryScreenStartsProteinEntityLigandCcdEntity,
		u.OfProteinLibraryScreenStartsProteinEntityLigandSmilesEntity)
}
func (u *ProteinLibraryScreenStartParamsProteinEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion struct {
	OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationCcdModification    *ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationSmilesModification *ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationCcdModification, u.OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationUnion struct {
	OfProteinLibraryScreenStartsProteinEntityRnaEntityModificationCcdModification    *ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityRnaEntityModificationSmilesModification *ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsProteinEntityRnaEntityModificationCcdModification, u.OfProteinLibraryScreenStartsProteinEntityRnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityRnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationUnion struct {
	OfProteinLibraryScreenStartsProteinEntityDnaEntityModificationCcdModification    *ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsProteinEntityDnaEntityModificationSmilesModification *ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsProteinEntityDnaEntityModificationCcdModification, u.OfProteinLibraryScreenStartsProteinEntityDnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityDnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetUnion struct {
	OfProteinLibraryScreenStartsTargetStructureTemplateTarget *ProteinLibraryScreenStartParamsTargetStructureTemplateTarget `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTarget        *ProteinLibraryScreenStartParamsTargetNoTemplateTarget        `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetStructureTemplateTarget, u.OfProteinLibraryScreenStartsTargetNoTemplateTarget)
}
func (u *ProteinLibraryScreenStartParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
//
// The properties ChainSelection, Structure, Type are required.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTarget struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion `json:"chain_selection,omitzero" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion `json:"structure,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "structure_template".
	Type constant.StructureTemplate `json:"type" default:"structure_template"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetStructureTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetStructureTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetStructureTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion struct {
	OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec, u.OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec)
}
func (u *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
//
// The properties ChainType, CropResidues are required.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues,omitzero" api:"required"`
	// 0-indexed residue indices where binder contact is desired (the epitope). All
	// indices must be present in crop_residues and must not overlap
	// non_binding_residues.
	EpitopeResidues []int64 `json:"epitope_residues,omitzero"`
	// 0-indexed residue indices allowed to move during design (e.g. flexible loop
	// regions). All indices must be present in crop_residues.
	FlexibleResidues []int64 `json:"flexible_residues,omitzero"`
	// 0-indexed residue indices where binder contact should be discouraged. All
	// indices must be present in crop_residues and must not overlap epitope_residues.
	NonBindingResidues []int64 `json:"non_binding_residues,omitzero"`
	// This field can be elided, and will marshal its zero value as "polymer".
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	OfIntArray []int64 `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIntArray, u.OfAll)
}
func (u *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func NewProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec() ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec {
	return ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec{
		ChainType: "ligand",
	}
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
//
// This struct has a constant value, construct it with
// [NewProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec].
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion struct {
	OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureURLSource       *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource       `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureCifBase64Source *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureURLSource, u.OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureCifBase64Source)
}
func (u *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureCifBase64Source struct {
	// Base64-encoded CIF file contents
	Data string `json:"data" api:"required"`
	// Must be chemical/x-cif for CIF files
	//
	// This field can be elided, and will marshal its zero value as "chemical/x-cif".
	MediaType constant.ChemicalXCif `json:"media_type" default:"chemical/x-cif"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
//
// The properties Entities, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTarget struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintUnion `json:"constraints,omitzero"`
	// Chain IDs of ligand entities that are part of the binding epitope. Ligands are
	// marked as epitope in full (no residue-level selection).
	EpitopeLigandChains []string `json:"epitope_ligand_chains,omitzero"`
	// Polymer chain residues where binder contact is desired (the epitope). Each key
	// is a chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap non_binding_residues on the same chain.
	EpitopeResidues map[string][]int64 `json:"epitope_residues,omitzero"`
	// Polymer chain residues where binder contact should be discouraged. Each key is a
	// chain ID of a polymer entity, each value is an array of 0-indexed residue
	// indices. Residues must not overlap epitope_residues on the same chain.
	NonBindingResidues map[string][]int64 `json:"non_binding_residues,omitzero"`
	// This field can be elided, and will marshal its zero value as "no_template".
	Type constant.NoTemplate `json:"type" default:"no_template"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityUnion struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntity      *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntity      `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntity          *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntity          *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntity          `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityLigandCcdEntity    *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityLigandSmilesEntity *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityLigandCcdEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityLigandSmilesEntity)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification    *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification    *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
	Modifications []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification    *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification    `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationCcdModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// SMILES string for the modification
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles".
	Type constant.Smiles `json:"type" default:"smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModificationSmilesModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1LigandAtom  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1PolymerAtom *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1LigandAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1PolymerAtom)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2LigandAtom  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2PolymerAtom *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2LigandAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2PolymerAtom)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintUnion struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintPocketConstraint  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraint *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintPocketConstraint, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraint)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintPocketConstraint struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
