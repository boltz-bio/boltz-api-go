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

// Generate novel protein binders optimized for binding to a target structure.
// Binder specifications can be provided directly, uploaded as structure templates,
// or selected from Boltz-managed curated nanobody and antibody defaults. Results
// are scored by binding confidence (likelihood of protein-protein interaction) and
// structure confidence.
//
// ProteinDesignService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProteinDesignService] method instead.
type ProteinDesignService struct {
	Options []option.RequestOption
}

// NewProteinDesignService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProteinDesignService(opts ...option.RequestOption) (r ProteinDesignService) {
	r = ProteinDesignService{}
	r.Options = opts
	return
}

// Retrieve a design run by ID, including progress and status
func (r *ProteinDesignService) Get(ctx context.Context, id string, query ProteinDesignGetParams, opts ...option.RequestOption) (res *ProteinDesignGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/design/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List protein design runs, optionally filtered by workspace
func (r *ProteinDesignService) List(ctx context.Context, query ProteinDesignListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProteinDesignListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/protein/design"
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

// List protein design runs, optionally filtered by workspace
func (r *ProteinDesignService) ListAutoPaging(ctx context.Context, query ProteinDesignListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProteinDesignListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// design run. The design run record itself is retained with a `data_deleted_at`
// timestamp. This action is irreversible.
func (r *ProteinDesignService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinDesignDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/design/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Estimate the cost of a protein design run without creating any resource or
// consuming GPU.
func (r *ProteinDesignService) EstimateCost(ctx context.Context, body ProteinDesignEstimateCostParams, opts ...option.RequestOption) (res *ProteinDesignEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/protein/design/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve paginated results from a protein design run
func (r *ProteinDesignService) ListResults(ctx context.Context, id string, query ProteinDesignListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProteinDesignListResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/design/%s/results", url.PathEscape(id))
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

// Retrieve paginated results from a protein design run
func (r *ProteinDesignService) ListResultsAutoPaging(ctx context.Context, id string, query ProteinDesignListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProteinDesignListResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, id, query, opts...))
}

// Create a new design run that generates novel protein binder candidates
func (r *ProteinDesignService) Start(ctx context.Context, body ProteinDesignStartParams, opts ...option.RequestOption) (res *ProteinDesignStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/protein/design"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an in-progress protein design run early
func (r *ProteinDesignService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinDesignStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/design/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A protein design engine run that generates novel protein binders
type ProteinDesignGetResponse struct {
	// Unique ProteinDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein design
	Engine constant.BoltzProteinDesign `json:"engine" default:"boltz-protein-design"`
	// Engine version used for protein design
	EngineVersion string                        `json:"engine_version" api:"required"`
	Error         ProteinDesignGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinDesignGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                             `json:"livemode" api:"required"`
	Progress  ProteinDesignGetResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                        `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinDesignGetResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                      `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinDesignGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseError struct {
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
func (r ProteinDesignGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinDesignGetResponseInput struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, or boltz_curated for
	// Boltz-managed nanobody and antibody defaults.
	BinderSpecification ProteinDesignGetResponseInputBinderSpecificationUnion `json:"binder_specification" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinDesignGetResponseInputTargetUnion `json:"target" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderSpecification respjson.Field
		NumProteins         respjson.Field
		Target              respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationUnion contains all possible
// properties and values from
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse],
// [ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputBinderSpecificationUnion struct {
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse].
	ChainSelection map[string]ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection"`
	Modality       string                                                                                                            `json:"modality"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse].
	Structure ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure"`
	Type      string                                                                                       `json:"type"`
	// This field is a union of
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules],
	// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules],
	// [ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules]
	Rules ProteinDesignGetResponseInputBinderSpecificationUnionRules `json:"rules"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse].
	Entities []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse].
	Bonds []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse].
	Binder ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder"`
	JSON   struct {
		ChainSelection respjson.Field
		Modality       respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		Rules          respjson.Field
		Entities       respjson.Field
		Bonds          respjson.Field
		Binder         respjson.Field
		raw            string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputBinderSpecificationUnion) AsProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse() (v ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationUnion) AsProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse() (v ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignGetResponseInputBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationUnionRules is an implicit
// subunion of [ProteinDesignGetResponseInputBinderSpecificationUnion].
// ProteinDesignGetResponseInputBinderSpecificationUnionRules provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignGetResponseInputBinderSpecificationUnion].
type ProteinDesignGetResponseInputBinderSpecificationUnionRules struct {
	ExcludedAminoAcids     []string `json:"excluded_amino_acids"`
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	MaxHydrophobicFraction float64  `json:"max_hydrophobic_fraction"`
	JSON                   struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *ProteinDesignGetResponseInputBinderSpecificationUnionRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// engine run.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep (crop_residues).
	// Omit design_motifs to include the chain as fixed scaffold context.
	ChainSelection map[string]ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality  ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality  `json:"modality" api:"required"`
	Structure ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure" api:"required"`
	Type      constant.StructureTemplate                                                                   `json:"type" default:"structure_template"`
	// Constraints applied during sequence design
	Rules ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainSelection respjson.Field
		Modality       respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		Rules          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec],
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	CropResidues ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	DesignMotifs []ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	JSON         struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec() (v ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec() (v ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion
// contains all possible properties and values from
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif],
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	// This field is a union of
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange],
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange]
	DesignLengthRange ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange `json:"design_length_range"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	EndIndex int64 `json:"end_index"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	StartIndex int64  `json:"start_index"`
	Type       string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
	AfterResidueIndex int64 `json:"after_residue_index"`
	JSON              struct {
		DesignLengthRange respjson.Field
		EndIndex          respjson.Field
		StartIndex        respjson.Field
		Type              respjson.Field
		AfterResidueIndex respjson.Field
		raw               string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif() (v ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif() (v ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// is an implicit subunion of
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
// ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange struct {
	Max  int64 `json:"max"`
	Min  int64 `json:"min"`
	JSON struct {
		Max respjson.Field
		Min respjson.Field
		raw string
	} `json:"-"`
}

func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range" api:"required"`
	// 0-indexed end residue (inclusive)
	EndIndex int64 `json:"end_index" api:"required"`
	// 0-indexed start residue (inclusive)
	StartIndex int64                `json:"start_index" api:"required"`
	Type       constant.Replacement `json:"type" default:"replacement"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DesignLengthRange respjson.Field
		EndIndex          respjson.Field
		StartIndex        respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range" api:"required"`
	Type              constant.Insertion                                                                                                                                                           `json:"type" default:"insertion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterResidueIndex respjson.Field
		DesignLengthRange respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality string

const (
	ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityPeptide       ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "peptide"
	ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityAntibody      ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "antibody"
	ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityNanobody      ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "nanobody"
	ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityCustomProtein ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "custom_protein"
)

type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality `json:"modality" api:"required"`
	Type     constant.NoTemplate                                                                  `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// Constraints applied during sequence design
	Rules ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Modality    respjson.Field
		Type        respjson.Field
		Bonds       respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion
// contains all possible properties and values from
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification],
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification],
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification],
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	Modifications ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
// ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications
// OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications
// OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications
// OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications]
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                    struct {
		OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications respjson.Field
		OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications    respjson.Field
		OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications        respjson.Field
		OfProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications        respjson.Field
		raw                                                                                                                            string
	} `json:"-"`
}

func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protein binder entity with designed and/or fixed segments.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string                 `json:"chain_ids" api:"required"`
	Type     constant.DesignedProtein `json:"type" default:"designed_protein"`
	// Binder sequence specification. Fixed amino acids are written as literal
	// single-letter codes. Designed regions are written as a length (fixed) or a
	// length range (min..max). Example: "MKTAYI5..10VKSHFSRQ" means fixed MKTAYI, then
	// 5-10 designed residues, then fixed VKSHFSRQ. "20" means 20 fully designed
	// residues. "ACDE8GHI" means fixed ACDE, then 8 designed residues, then fixed GHI.
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse struct {
	// Chain IDs to assign to this entity
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code from RCSB PDB (e.g. 'ATP', 'ADP')
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality string

const (
	ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityPeptide       ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "peptide"
	ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityAntibody      ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "antibody"
	ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityNanobody      ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "nanobody"
	ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityCustomProtein ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "custom_protein"
)

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse],
// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse() (v ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
type ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder" api:"required"`
	Type   constant.BoltzCurated                                                                `json:"type" default:"boltz_curated"`
	// Constraints applied during sequence design
	Rules ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binder      respjson.Field
		Type        respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder string

const (
	ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzNanobody ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_nanobody"
	ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzAntibody ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignGetResponseInputBinderSpecificationBinder string

const (
	ProteinDesignGetResponseInputBinderSpecificationBinderBoltzNanobody ProteinDesignGetResponseInputBinderSpecificationBinder = "boltz_nanobody"
	ProteinDesignGetResponseInputBinderSpecificationBinderBoltzAntibody ProteinDesignGetResponseInputBinderSpecificationBinder = "boltz_antibody"
)

type ProteinDesignGetResponseInputBinderSpecificationModality string

const (
	ProteinDesignGetResponseInputBinderSpecificationModalityPeptide       ProteinDesignGetResponseInputBinderSpecificationModality = "peptide"
	ProteinDesignGetResponseInputBinderSpecificationModalityAntibody      ProteinDesignGetResponseInputBinderSpecificationModality = "antibody"
	ProteinDesignGetResponseInputBinderSpecificationModalityNanobody      ProteinDesignGetResponseInputBinderSpecificationModality = "nanobody"
	ProteinDesignGetResponseInputBinderSpecificationModalityCustomProtein ProteinDesignGetResponseInputBinderSpecificationModality = "custom_protein"
)

// ProteinDesignGetResponseInputTargetUnion contains all possible properties and
// values from
// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                      `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponse].
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

func (u ProteinDesignGetResponseInputTargetUnion) AsProteinDesignGetResponseInputTargetStructureTemplateTargetResponse() (v ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignGetResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                       `json:"type" default:"structure_template"`
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
func (r ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
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

func (u ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                      `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion contains
// all possible properties and values from
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                              struct {
		OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                   string
	} `json:"-"`
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                          `json:"type" default:"contact"`
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseProgress struct {
	// Number of protein binders generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein binders requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// Cumulative number of designs produced by the model before filtering and
	// deduplication. Useful for monitoring generation efficiency.
	TotalProteinsGeneratedPreFilter int64 `json:"total_proteins_generated_pre_filter"`
	// Cumulative number of designs that survived filtering and deduplication and
	// became scoreable samples.
	TotalProteinsInserted int64 `json:"total_proteins_inserted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated            respjson.Field
		TotalProteinsToGenerate         respjson.Field
		LatestResultID                  respjson.Field
		TotalProteinsGeneratedPreFilter respjson.Field
		TotalProteinsInserted           respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignGetResponseStatus string

const (
	ProteinDesignGetResponseStatusPending   ProteinDesignGetResponseStatus = "pending"
	ProteinDesignGetResponseStatusRunning   ProteinDesignGetResponseStatus = "running"
	ProteinDesignGetResponseStatusSucceeded ProteinDesignGetResponseStatus = "succeeded"
	ProteinDesignGetResponseStatusFailed    ProteinDesignGetResponseStatus = "failed"
	ProteinDesignGetResponseStatusStopped   ProteinDesignGetResponseStatus = "stopped"
)

// Summary of a protein design engine run (excludes input)
type ProteinDesignListResponse struct {
	// Unique ProteinDesignRunSummary identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein design
	Engine constant.BoltzProteinDesign `json:"engine" default:"boltz-protein-design"`
	// Engine version used for protein design
	EngineVersion string                         `json:"engine_version" api:"required"`
	Error         ProteinDesignListResponseError `json:"error" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                              `json:"livemode" api:"required"`
	Progress  ProteinDesignListResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                         `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinDesignListResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                       `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinDesignListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResponseError struct {
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
func (r ProteinDesignListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResponseProgress struct {
	// Number of protein binders generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein binders requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// Cumulative number of designs produced by the model before filtering and
	// deduplication. Useful for monitoring generation efficiency.
	TotalProteinsGeneratedPreFilter int64 `json:"total_proteins_generated_pre_filter"`
	// Cumulative number of designs that survived filtering and deduplication and
	// became scoreable samples.
	TotalProteinsInserted int64 `json:"total_proteins_inserted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated            respjson.Field
		TotalProteinsToGenerate         respjson.Field
		LatestResultID                  respjson.Field
		TotalProteinsGeneratedPreFilter respjson.Field
		TotalProteinsInserted           respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignListResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResponseStatus string

const (
	ProteinDesignListResponseStatusPending   ProteinDesignListResponseStatus = "pending"
	ProteinDesignListResponseStatusRunning   ProteinDesignListResponseStatus = "running"
	ProteinDesignListResponseStatusSucceeded ProteinDesignListResponseStatus = "succeeded"
	ProteinDesignListResponseStatusFailed    ProteinDesignListResponseStatus = "failed"
	ProteinDesignListResponseStatusStopped   ProteinDesignListResponseStatus = "stopped"
)

type ProteinDesignDeleteDataResponse struct {
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
func (r ProteinDesignDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimate response with monetary values encoded as decimal strings to preserve
// precision.
type ProteinDesignEstimateCostResponse struct {
	// Cost breakdown for the billed application.
	Breakdown  ProteinDesignEstimateCostResponseBreakdown `json:"breakdown" api:"required"`
	Disclaimer string                                     `json:"disclaimer" api:"required"`
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
func (r ProteinDesignEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost breakdown for the billed application.
type ProteinDesignEstimateCostResponseBreakdown struct {
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application ProteinDesignEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
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
func (r ProteinDesignEstimateCostResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignEstimateCostResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignEstimateCostResponseBreakdownApplication string

const (
	ProteinDesignEstimateCostResponseBreakdownApplicationStructureAndBinding        ProteinDesignEstimateCostResponseBreakdownApplication = "structure_and_binding"
	ProteinDesignEstimateCostResponseBreakdownApplicationSmallMoleculeDesign        ProteinDesignEstimateCostResponseBreakdownApplication = "small_molecule_design"
	ProteinDesignEstimateCostResponseBreakdownApplicationSmallMoleculeLibraryScreen ProteinDesignEstimateCostResponseBreakdownApplication = "small_molecule_library_screen"
	ProteinDesignEstimateCostResponseBreakdownApplicationProteinDesign              ProteinDesignEstimateCostResponseBreakdownApplication = "protein_design"
	ProteinDesignEstimateCostResponseBreakdownApplicationProteinRedesign            ProteinDesignEstimateCostResponseBreakdownApplication = "protein_redesign"
	ProteinDesignEstimateCostResponseBreakdownApplicationProteinLibraryScreen       ProteinDesignEstimateCostResponseBreakdownApplication = "protein_library_screen"
	ProteinDesignEstimateCostResponseBreakdownApplicationAdme                       ProteinDesignEstimateCostResponseBreakdownApplication = "adme"
)

// A single generated protein design
type ProteinDesignListResultsResponse struct {
	// Unique result ID
	ID        string                                    `json:"id" api:"required"`
	Artifacts ProteinDesignListResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	// Entities of the designed binder complex. Includes both designed entities and
	// fixed entities from the input.
	Entities []ProteinDesignListResultsResponseEntityUnion `json:"entities" api:"required"`
	// Structural and binding quality metrics for a designed protein binder
	Metrics ProteinDesignListResultsResponseMetrics `json:"metrics" api:"required"`
	// Warnings about potential quality issues with this result.
	Warnings []ProteinDesignListResultsResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Entities    respjson.Field
		Metrics     respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignListResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseArtifacts struct {
	Archive   ProteinDesignListResultsResponseArtifactsArchive   `json:"archive" api:"required"`
	Structure ProteinDesignListResultsResponseArtifactsStructure `json:"structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archive     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignListResultsResponseArtifacts) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseArtifactsArchive struct {
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
func (r ProteinDesignListResultsResponseArtifactsArchive) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseArtifactsStructure struct {
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
func (r ProteinDesignListResultsResponseArtifactsStructure) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignListResultsResponseEntityUnion contains all possible properties and
// values from [ProteinDesignListResultsResponseEntityProteinEntity],
// [ProteinDesignListResultsResponseEntityRnaEntity],
// [ProteinDesignListResultsResponseEntityDnaEntity],
// [ProteinDesignListResultsResponseEntityLigandCcdEntity],
// [ProteinDesignListResultsResponseEntityLigandSmilesEntity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignListResultsResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignListResultsResponseEntityProteinEntityModification],
	// [[]ProteinDesignListResultsResponseEntityRnaEntityModification],
	// [[]ProteinDesignListResultsResponseEntityDnaEntityModification]
	Modifications ProteinDesignListResultsResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignListResultsResponseEntityUnion) AsProteinDesignListResultsResponseEntityProteinEntity() (v ProteinDesignListResultsResponseEntityProteinEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignListResultsResponseEntityUnion) AsProteinDesignListResultsResponseEntityRnaEntity() (v ProteinDesignListResultsResponseEntityRnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignListResultsResponseEntityUnion) AsProteinDesignListResultsResponseEntityDnaEntity() (v ProteinDesignListResultsResponseEntityDnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignListResultsResponseEntityUnion) AsProteinDesignListResultsResponseEntityLigandCcdEntity() (v ProteinDesignListResultsResponseEntityLigandCcdEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignListResultsResponseEntityUnion) AsProteinDesignListResultsResponseEntityLigandSmilesEntity() (v ProteinDesignListResultsResponseEntityLigandSmilesEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignListResultsResponseEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignListResultsResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignListResultsResponseEntityUnionModifications is an implicit subunion
// of [ProteinDesignListResultsResponseEntityUnion].
// ProteinDesignListResultsResponseEntityUnionModifications provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignListResultsResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignListResultsResponseEntityProteinEntityModifications
// OfProteinDesignListResultsResponseEntityRnaEntityModifications
// OfProteinDesignListResultsResponseEntityDnaEntityModifications]
type ProteinDesignListResultsResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignListResultsResponseEntityProteinEntityModification] instead of
	// an object.
	OfProteinDesignListResultsResponseEntityProteinEntityModifications []ProteinDesignListResultsResponseEntityProteinEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignListResultsResponseEntityRnaEntityModification] instead of an
	// object.
	OfProteinDesignListResultsResponseEntityRnaEntityModifications []ProteinDesignListResultsResponseEntityRnaEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignListResultsResponseEntityDnaEntityModification] instead of an
	// object.
	OfProteinDesignListResultsResponseEntityDnaEntityModifications []ProteinDesignListResultsResponseEntityDnaEntityModification `json:",inline"`
	JSON                                                           struct {
		OfProteinDesignListResultsResponseEntityProteinEntityModifications respjson.Field
		OfProteinDesignListResultsResponseEntityRnaEntityModifications     respjson.Field
		OfProteinDesignListResultsResponseEntityDnaEntityModifications     respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (r *ProteinDesignListResultsResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinDesignListResultsResponseEntityProteinEntityModification `json:"modifications"`
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
func (r ProteinDesignListResultsResponseEntityProteinEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignListResultsResponseEntityProteinEntityModification struct {
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
func (r ProteinDesignListResultsResponseEntityProteinEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignListResultsResponseEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignListResultsResponseEntityRnaEntityModification `json:"modifications"`
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
func (r ProteinDesignListResultsResponseEntityRnaEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignListResultsResponseEntityRnaEntityModification struct {
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
func (r ProteinDesignListResultsResponseEntityRnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignListResultsResponseEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignListResultsResponseEntityDnaEntityModification `json:"modifications"`
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
func (r ProteinDesignListResultsResponseEntityDnaEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignListResultsResponseEntityDnaEntityModification struct {
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
func (r ProteinDesignListResultsResponseEntityDnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignListResultsResponseEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseEntityLigandCcdEntity struct {
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
func (r ProteinDesignListResultsResponseEntityLigandCcdEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsResponseEntityLigandSmilesEntity struct {
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
func (r ProteinDesignListResultsResponseEntityLigandSmilesEntity) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structural and binding quality metrics for a designed protein binder
type ProteinDesignListResultsResponseMetrics struct {
	// Confidence that the designed binder binds the target (0-1). Primary metric for
	// hit discovery.
	BindingConfidence float64 `json:"binding_confidence" api:"required"`
	// Fraction of the designed sequence forming alpha helices (0-1).
	HelixFraction float64 `json:"helix_fraction" api:"required"`
	// Interface predicted TM score (0-1). Confidence in the protein-protein interface.
	Iptm float64 `json:"iptm" api:"required"`
	// Fraction of the designed sequence in coil/loop regions (0-1).
	LoopFraction float64 `json:"loop_fraction" api:"required"`
	// Minimum predicted aligned error at the interface (Angstroms). Lower values
	// indicate higher confidence.
	MinInteractionPae float64 `json:"min_interaction_pae" api:"required"`
	// Fraction of the designed sequence forming beta sheets (0-1).
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
func (r ProteinDesignListResultsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warning about a potential quality issue with a result
type ProteinDesignListResultsResponseWarning struct {
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
func (r ProteinDesignListResultsResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignListResultsResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A protein design engine run that generates novel protein binders
type ProteinDesignStartResponse struct {
	// Unique ProteinDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein design
	Engine constant.BoltzProteinDesign `json:"engine" default:"boltz-protein-design"`
	// Engine version used for protein design
	EngineVersion string                          `json:"engine_version" api:"required"`
	Error         ProteinDesignStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinDesignStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                               `json:"livemode" api:"required"`
	Progress  ProteinDesignStartResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                          `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinDesignStartResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                        `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinDesignStartResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseError struct {
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
func (r ProteinDesignStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinDesignStartResponseInput struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, or boltz_curated for
	// Boltz-managed nanobody and antibody defaults.
	BinderSpecification ProteinDesignStartResponseInputBinderSpecificationUnion `json:"binder_specification" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinDesignStartResponseInputTargetUnion `json:"target" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderSpecification respjson.Field
		NumProteins         respjson.Field
		Target              respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationUnion contains all possible
// properties and values from
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse],
// [ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputBinderSpecificationUnion struct {
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse].
	ChainSelection map[string]ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection"`
	Modality       string                                                                                                              `json:"modality"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse].
	Structure ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure"`
	Type      string                                                                                         `json:"type"`
	// This field is a union of
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules],
	// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules],
	// [ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules]
	Rules ProteinDesignStartResponseInputBinderSpecificationUnionRules `json:"rules"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse].
	Entities []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse].
	Bonds []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse].
	Binder ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder"`
	JSON   struct {
		ChainSelection respjson.Field
		Modality       respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		Rules          respjson.Field
		Entities       respjson.Field
		Bonds          respjson.Field
		Binder         respjson.Field
		raw            string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputBinderSpecificationUnion) AsProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse() (v ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationUnion) AsProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse() (v ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignStartResponseInputBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationUnionRules is an implicit
// subunion of [ProteinDesignStartResponseInputBinderSpecificationUnion].
// ProteinDesignStartResponseInputBinderSpecificationUnionRules provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStartResponseInputBinderSpecificationUnion].
type ProteinDesignStartResponseInputBinderSpecificationUnionRules struct {
	ExcludedAminoAcids     []string `json:"excluded_amino_acids"`
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	MaxHydrophobicFraction float64  `json:"max_hydrophobic_fraction"`
	JSON                   struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *ProteinDesignStartResponseInputBinderSpecificationUnionRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// engine run.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep (crop_residues).
	// Omit design_motifs to include the chain as fixed scaffold context.
	ChainSelection map[string]ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality  ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality  `json:"modality" api:"required"`
	Structure ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure" api:"required"`
	Type      constant.StructureTemplate                                                                     `json:"type" default:"structure_template"`
	// Constraints applied during sequence design
	Rules ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainSelection respjson.Field
		Modality       respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		Rules          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec],
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	CropResidues ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	DesignMotifs []ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	JSON         struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec() (v ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec() (v ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion
// contains all possible properties and values from
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif],
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	// This field is a union of
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange],
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange]
	DesignLengthRange ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange `json:"design_length_range"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	EndIndex int64 `json:"end_index"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	StartIndex int64  `json:"start_index"`
	Type       string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
	AfterResidueIndex int64 `json:"after_residue_index"`
	JSON              struct {
		DesignLengthRange respjson.Field
		EndIndex          respjson.Field
		StartIndex        respjson.Field
		Type              respjson.Field
		AfterResidueIndex respjson.Field
		raw               string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif() (v ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif() (v ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// is an implicit subunion of
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
// ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange struct {
	Max  int64 `json:"max"`
	Min  int64 `json:"min"`
	JSON struct {
		Max respjson.Field
		Min respjson.Field
		raw string
	} `json:"-"`
}

func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range" api:"required"`
	// 0-indexed end residue (inclusive)
	EndIndex int64 `json:"end_index" api:"required"`
	// 0-indexed start residue (inclusive)
	StartIndex int64                `json:"start_index" api:"required"`
	Type       constant.Replacement `json:"type" default:"replacement"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DesignLengthRange respjson.Field
		EndIndex          respjson.Field
		StartIndex        respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range" api:"required"`
	Type              constant.Insertion                                                                                                                                                             `json:"type" default:"insertion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterResidueIndex respjson.Field
		DesignLengthRange respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality string

const (
	ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityPeptide       ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "peptide"
	ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityAntibody      ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "antibody"
	ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityNanobody      ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "nanobody"
	ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityCustomProtein ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "custom_protein"
)

type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality `json:"modality" api:"required"`
	Type     constant.NoTemplate                                                                    `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// Constraints applied during sequence design
	Rules ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Modality    respjson.Field
		Type        respjson.Field
		Bonds       respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion
// contains all possible properties and values from
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification],
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification],
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification],
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	Modifications ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
// ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications
// OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications
// OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications
// OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications]
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                      struct {
		OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications respjson.Field
		OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications    respjson.Field
		OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications        respjson.Field
		OfProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications        respjson.Field
		raw                                                                                                                              string
	} `json:"-"`
}

func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protein binder entity with designed and/or fixed segments.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string                 `json:"chain_ids" api:"required"`
	Type     constant.DesignedProtein `json:"type" default:"designed_protein"`
	// Binder sequence specification. Fixed amino acids are written as literal
	// single-letter codes. Designed regions are written as a length (fixed) or a
	// length range (min..max). Example: "MKTAYI5..10VKSHFSRQ" means fixed MKTAYI, then
	// 5-10 designed residues, then fixed VKSHFSRQ. "20" means 20 fully designed
	// residues. "ACDE8GHI" means fixed ACDE, then 8 designed residues, then fixed GHI.
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse struct {
	// Chain IDs to assign to this entity
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code from RCSB PDB (e.g. 'ATP', 'ADP')
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality string

const (
	ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityPeptide       ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "peptide"
	ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityAntibody      ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "antibody"
	ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityNanobody      ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "nanobody"
	ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityCustomProtein ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "custom_protein"
)

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse],
// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse() (v ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
type ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder" api:"required"`
	Type   constant.BoltzCurated                                                                  `json:"type" default:"boltz_curated"`
	// Constraints applied during sequence design
	Rules ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binder      respjson.Field
		Type        respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder string

const (
	ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzNanobody ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_nanobody"
	ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzAntibody ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignStartResponseInputBinderSpecificationBinder string

const (
	ProteinDesignStartResponseInputBinderSpecificationBinderBoltzNanobody ProteinDesignStartResponseInputBinderSpecificationBinder = "boltz_nanobody"
	ProteinDesignStartResponseInputBinderSpecificationBinderBoltzAntibody ProteinDesignStartResponseInputBinderSpecificationBinder = "boltz_antibody"
)

type ProteinDesignStartResponseInputBinderSpecificationModality string

const (
	ProteinDesignStartResponseInputBinderSpecificationModalityPeptide       ProteinDesignStartResponseInputBinderSpecificationModality = "peptide"
	ProteinDesignStartResponseInputBinderSpecificationModalityAntibody      ProteinDesignStartResponseInputBinderSpecificationModality = "antibody"
	ProteinDesignStartResponseInputBinderSpecificationModalityNanobody      ProteinDesignStartResponseInputBinderSpecificationModality = "nanobody"
	ProteinDesignStartResponseInputBinderSpecificationModalityCustomProtein ProteinDesignStartResponseInputBinderSpecificationModality = "custom_protein"
)

// ProteinDesignStartResponseInputTargetUnion contains all possible properties and
// values from
// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                        `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponse].
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

func (u ProteinDesignStartResponseInputTargetUnion) AsProteinDesignStartResponseInputTargetStructureTemplateTargetResponse() (v ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignStartResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                         `json:"type" default:"structure_template"`
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
func (r ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
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

func (u ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                        `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                struct {
		OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                     string
	} `json:"-"`
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                            `json:"type" default:"contact"`
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseProgress struct {
	// Number of protein binders generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein binders requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// Cumulative number of designs produced by the model before filtering and
	// deduplication. Useful for monitoring generation efficiency.
	TotalProteinsGeneratedPreFilter int64 `json:"total_proteins_generated_pre_filter"`
	// Cumulative number of designs that survived filtering and deduplication and
	// became scoreable samples.
	TotalProteinsInserted int64 `json:"total_proteins_inserted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated            respjson.Field
		TotalProteinsToGenerate         respjson.Field
		LatestResultID                  respjson.Field
		TotalProteinsGeneratedPreFilter respjson.Field
		TotalProteinsInserted           respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStartResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStartResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartResponseStatus string

const (
	ProteinDesignStartResponseStatusPending   ProteinDesignStartResponseStatus = "pending"
	ProteinDesignStartResponseStatusRunning   ProteinDesignStartResponseStatus = "running"
	ProteinDesignStartResponseStatusSucceeded ProteinDesignStartResponseStatus = "succeeded"
	ProteinDesignStartResponseStatusFailed    ProteinDesignStartResponseStatus = "failed"
	ProteinDesignStartResponseStatusStopped   ProteinDesignStartResponseStatus = "stopped"
)

// A protein design engine run that generates novel protein binders
type ProteinDesignStopResponse struct {
	// Unique ProteinDesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Engine used for protein design
	Engine constant.BoltzProteinDesign `json:"engine" default:"boltz-protein-design"`
	// Engine version used for protein design
	EngineVersion string                         `json:"engine_version" api:"required"`
	Error         ProteinDesignStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinDesignStopResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode  bool                              `json:"livemode" api:"required"`
	Progress  ProteinDesignStopResponseProgress `json:"progress" api:"required"`
	StartedAt time.Time                         `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinDesignStopResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                       `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinDesignStopResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseError struct {
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
func (r ProteinDesignStopResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStopResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinDesignStopResponseInput struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, or boltz_curated for
	// Boltz-managed nanobody and antibody defaults.
	BinderSpecification ProteinDesignStopResponseInputBinderSpecificationUnion `json:"binder_specification" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinDesignStopResponseInputTargetUnion `json:"target" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderSpecification respjson.Field
		NumProteins         respjson.Field
		Target              respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStopResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationUnion contains all possible
// properties and values from
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse],
// [ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputBinderSpecificationUnion struct {
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse].
	ChainSelection map[string]ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection"`
	Modality       string                                                                                                             `json:"modality"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse].
	Structure ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure"`
	Type      string                                                                                        `json:"type"`
	// This field is a union of
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules],
	// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules],
	// [ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules]
	Rules ProteinDesignStopResponseInputBinderSpecificationUnionRules `json:"rules"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse].
	Entities []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse].
	Bonds []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse].
	Binder ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder"`
	JSON   struct {
		ChainSelection respjson.Field
		Modality       respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		Rules          respjson.Field
		Entities       respjson.Field
		Bonds          respjson.Field
		Binder         respjson.Field
		raw            string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputBinderSpecificationUnion) AsProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse() (v ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationUnion) AsProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse() (v ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignStopResponseInputBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationUnionRules is an implicit
// subunion of [ProteinDesignStopResponseInputBinderSpecificationUnion].
// ProteinDesignStopResponseInputBinderSpecificationUnionRules provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStopResponseInputBinderSpecificationUnion].
type ProteinDesignStopResponseInputBinderSpecificationUnionRules struct {
	ExcludedAminoAcids     []string `json:"excluded_amino_acids"`
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	MaxHydrophobicFraction float64  `json:"max_hydrophobic_fraction"`
	JSON                   struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *ProteinDesignStopResponseInputBinderSpecificationUnionRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// engine run.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep (crop_residues).
	// Omit design_motifs to include the chain as fixed scaffold context.
	ChainSelection map[string]ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality  ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality  `json:"modality" api:"required"`
	Structure ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure" api:"required"`
	Type      constant.StructureTemplate                                                                    `json:"type" default:"structure_template"`
	// Constraints applied during sequence design
	Rules ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainSelection respjson.Field
		Modality       respjson.Field
		Structure      respjson.Field
		Type           respjson.Field
		Rules          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec],
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	CropResidues ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	DesignMotifs []ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	JSON         struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec() (v ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec() (v ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion
// contains all possible properties and values from
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif],
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	// This field is a union of
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange],
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange]
	DesignLengthRange ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange `json:"design_length_range"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	EndIndex int64 `json:"end_index"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	StartIndex int64  `json:"start_index"`
	Type       string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
	AfterResidueIndex int64 `json:"after_residue_index"`
	JSON              struct {
		DesignLengthRange respjson.Field
		EndIndex          respjson.Field
		StartIndex        respjson.Field
		Type              respjson.Field
		AfterResidueIndex respjson.Field
		raw               string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif() (v ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif() (v ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// is an implicit subunion of
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
// ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange struct {
	Max  int64 `json:"max"`
	Min  int64 `json:"min"`
	JSON struct {
		Max respjson.Field
		Min respjson.Field
		raw string
	} `json:"-"`
}

func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range" api:"required"`
	// 0-indexed end residue (inclusive)
	EndIndex int64 `json:"end_index" api:"required"`
	// 0-indexed start residue (inclusive)
	StartIndex int64                `json:"start_index" api:"required"`
	Type       constant.Replacement `json:"type" default:"replacement"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DesignLengthRange respjson.Field
		EndIndex          respjson.Field
		StartIndex        respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range" api:"required"`
	Type              constant.Insertion                                                                                                                                                            `json:"type" default:"insertion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterResidueIndex respjson.Field
		DesignLengthRange respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality string

const (
	ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityPeptide       ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "peptide"
	ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityAntibody      ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "antibody"
	ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityNanobody      ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "nanobody"
	ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModalityCustomProtein ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "custom_protein"
)

type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationStructureTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality `json:"modality" api:"required"`
	Type     constant.NoTemplate                                                                   `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// Constraints applied during sequence design
	Rules ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Modality    respjson.Field
		Type        respjson.Field
		Bonds       respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion
// contains all possible properties and values from
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification],
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification],
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification],
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	Modifications ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
// ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications
// OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications
// OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications
// OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications]
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                     struct {
		OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications respjson.Field
		OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications    respjson.Field
		OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications        respjson.Field
		OfProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications        respjson.Field
		raw                                                                                                                             string
	} `json:"-"`
}

func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protein binder entity with designed and/or fixed segments.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string                 `json:"chain_ids" api:"required"`
	Type     constant.DesignedProtein `json:"type" default:"designed_protein"`
	// Binder sequence specification. Fixed amino acids are written as literal
	// single-letter codes. Designed regions are written as a length (fixed) or a
	// length range (min..max). Example: "MKTAYI5..10VKSHFSRQ" means fixed MKTAYI, then
	// 5-10 designed residues, then fixed VKSHFSRQ. "20" means 20 fully designed
	// residues. "ACDE8GHI" means fixed ACDE, then 8 designed residues, then fixed GHI.
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse struct {
	// Chain IDs to assign to this entity
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// CCD code from RCSB PDB (e.g. 'ATP', 'ADP')
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality string

const (
	ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityPeptide       ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "peptide"
	ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityAntibody      ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "antibody"
	ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityNanobody      ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "nanobody"
	ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModalityCustomProtein ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseModality = "custom_protein"
)

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse],
// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse() (v ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationNoTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
type ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder" api:"required"`
	Type   constant.BoltzCurated                                                                 `json:"type" default:"boltz_curated"`
	// Constraints applied during sequence design
	Rules ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules `json:"rules"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Binder      respjson.Field
		Type        respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder string

const (
	ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzNanobody ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_nanobody"
	ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzAntibody ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules struct {
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs"`
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction float64 `json:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludedAminoAcids     respjson.Field
		ExcludedSequenceMotifs respjson.Field
		MaxHydrophobicFraction respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignStopResponseInputBinderSpecificationBinder string

const (
	ProteinDesignStopResponseInputBinderSpecificationBinderBoltzNanobody ProteinDesignStopResponseInputBinderSpecificationBinder = "boltz_nanobody"
	ProteinDesignStopResponseInputBinderSpecificationBinderBoltzAntibody ProteinDesignStopResponseInputBinderSpecificationBinder = "boltz_antibody"
)

type ProteinDesignStopResponseInputBinderSpecificationModality string

const (
	ProteinDesignStopResponseInputBinderSpecificationModalityPeptide       ProteinDesignStopResponseInputBinderSpecificationModality = "peptide"
	ProteinDesignStopResponseInputBinderSpecificationModalityAntibody      ProteinDesignStopResponseInputBinderSpecificationModality = "antibody"
	ProteinDesignStopResponseInputBinderSpecificationModalityNanobody      ProteinDesignStopResponseInputBinderSpecificationModality = "nanobody"
	ProteinDesignStopResponseInputBinderSpecificationModalityCustomProtein ProteinDesignStopResponseInputBinderSpecificationModality = "custom_protein"
)

// ProteinDesignStopResponseInputTargetUnion contains all possible properties and
// values from
// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                       `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponse].
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

func (u ProteinDesignStopResponseInputTargetUnion) AsProteinDesignStopResponseInputTargetStructureTemplateTargetResponse() (v ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinDesignStopResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                        `json:"type" default:"structure_template"`
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
func (r ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
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

func (u ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                       `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion contains
// all possible properties and values from
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                               struct {
		OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                    string
	} `json:"-"`
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                           `json:"type" default:"contact"`
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinDesignStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseProgress struct {
	// Number of protein binders generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein binders requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// Cumulative number of designs produced by the model before filtering and
	// deduplication. Useful for monitoring generation efficiency.
	TotalProteinsGeneratedPreFilter int64 `json:"total_proteins_generated_pre_filter"`
	// Cumulative number of designs that survived filtering and deduplication and
	// became scoreable samples.
	TotalProteinsInserted int64 `json:"total_proteins_inserted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated            respjson.Field
		TotalProteinsToGenerate         respjson.Field
		LatestResultID                  respjson.Field
		TotalProteinsGeneratedPreFilter respjson.Field
		TotalProteinsInserted           respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinDesignStopResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinDesignStopResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStopResponseStatus string

const (
	ProteinDesignStopResponseStatusPending   ProteinDesignStopResponseStatus = "pending"
	ProteinDesignStopResponseStatusRunning   ProteinDesignStopResponseStatus = "running"
	ProteinDesignStopResponseStatusSucceeded ProteinDesignStopResponseStatus = "succeeded"
	ProteinDesignStopResponseStatusFailed    ProteinDesignStopResponseStatus = "failed"
	ProteinDesignStopResponseStatusStopped   ProteinDesignStopResponseStatus = "stopped"
)

type ProteinDesignGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProteinDesignGetParams]'s query parameters as `url.Values`.
func (r ProteinDesignGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinDesignListParams struct {
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

// URLQuery serializes [ProteinDesignListParams]'s query parameters as
// `url.Values`.
func (r ProteinDesignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinDesignEstimateCostParams struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, or boltz_curated for
	// Boltz-managed nanobody and antibody defaults.
	BinderSpecification ProteinDesignEstimateCostParamsBinderSpecificationUnion `json:"binder_specification,omitzero" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinDesignEstimateCostParamsTargetUnion `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	paramObj
}

func (r ProteinDesignEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationUnion struct {
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpec *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpec        *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpec        `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationBoltzCuratedBinderSpec      *ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpec      `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpec, u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpec, u.OfProteinDesignEstimateCostsBinderSpecificationBoltzCuratedBinderSpec)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// engine run.
//
// The properties ChainSelection, Modality, Structure, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep (crop_residues).
	// Omit design_motifs to include the chain as fixed scaffold context.
	ChainSelection map[string]ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion `json:"chain_selection,omitzero" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModality `json:"modality,omitzero" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion `json:"structure,omitzero" api:"required"`
	// Constraints applied during sequence design
	Rules ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules `json:"rules,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "structure_template".
	Type constant.StructureTemplate `json:"type" default:"structure_template"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion struct {
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec  *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
//
// The properties ChainType, CropResidues are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec struct {
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues,omitzero" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs,omitzero"`
	// This field can be elided, and will marshal its zero value as "polymer".
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
	OfIntArray []int64 `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIntArray, u.OfAll)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif   *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif   `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
//
// The properties DesignLengthRange, EndIndex, StartIndex, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range,omitzero" api:"required"`
	// 0-indexed end residue (inclusive)
	EndIndex int64 `json:"end_index" api:"required"`
	// 0-indexed start residue (inclusive)
	StartIndex int64 `json:"start_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "replacement".
	Type constant.Replacement `json:"type" default:"replacement"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
//
// The properties Max, Min are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
//
// The properties AfterResidueIndex, DesignLengthRange, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "insertion".
	Type constant.Insertion `json:"type" default:"insertion"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
//
// The properties Max, Min are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec() ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec {
	return ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec{
		ChainType: "ligand",
	}
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
//
// This struct has a constant value, construct it with
// [NewProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec].
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModality string

const (
	ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide       ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModality = "peptide"
	ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityAntibody      ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModality = "antibody"
	ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityNanobody      ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModality = "nanobody"
	ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityCustomProtein ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModality = "custom_protein"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion struct {
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureURLSource       *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource       `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureURLSource, u.OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules struct {
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction param.Opt[float64] `json:"max_hydrophobic_fraction,omitzero"`
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids,omitzero"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs,omitzero"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
//
// The properties Entities, Modality, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpec struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityUnion `json:"entities,omitzero" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModality `json:"modality,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBond `json:"bonds,omitzero"`
	// Constraints applied during sequence design
	Rules ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecRules `json:"rules,omitzero"`
	// This field can be elided, and will marshal its zero value as "no_template".
	Type constant.NoTemplate `json:"type" default:"no_template"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityUnion struct {
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity   *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity   `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity      *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity      `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity          *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity          `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity          *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity          `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity    *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity    `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity,
		u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity,
		u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity,
		u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity,
		u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity,
		u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Protein binder entity with designed and/or fixed segments.
//
// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Binder sequence specification. Fixed amino acids are written as literal
	// single-letter codes. Designed regions are written as a length (fixed) or a
	// length range (min..max). Example: "MKTAYI5..10VKSHFSRQ" means fixed MKTAYI, then
	// 5-10 designed residues, then fixed VKSHFSRQ. "20" means 20 fully designed
	// residues. "ACDE8GHI" means fixed ACDE, then 8 designed residues, then fixed GHI.
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "designed_protein".
	Type constant.DesignedProtein `json:"type" default:"designed_protein"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
//
// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code from RCSB PDB (e.g. 'ATP', 'ADP')
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModality string

const (
	ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModalityPeptide       ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModality = "peptide"
	ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModalityAntibody      ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModality = "antibody"
	ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModalityNanobody      ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModality = "nanobody"
	ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModalityCustomProtein ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecModality = "custom_protein"
)

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union struct {
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom  *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom, u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union struct {
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom  *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom, u.OfProteinDesignEstimateCostsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom)
}
func (u *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom struct {
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

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecRules struct {
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction param.Opt[float64] `json:"max_hydrophobic_fraction,omitzero"`
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids,omitzero"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs,omitzero"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecRules) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationNoTemplateBinderSpecRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
//
// The properties Binder, Type are required.
type ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpec struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecBinder `json:"binder,omitzero" api:"required"`
	// Constraints applied during sequence design
	Rules ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecRules `json:"rules,omitzero"`
	// This field can be elided, and will marshal its zero value as "boltz_curated".
	Type constant.BoltzCurated `json:"type" default:"boltz_curated"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecBinder string

const (
	ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecBinderBoltzNanobody ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecBinder = "boltz_nanobody"
	ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecBinderBoltzAntibody ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecRules struct {
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction param.Opt[float64] `json:"max_hydrophobic_fraction,omitzero"`
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids,omitzero"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs,omitzero"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecRules) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsBinderSpecificationBoltzCuratedBinderSpecRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetUnion struct {
	OfProteinDesignEstimateCostsTargetStructureTemplateTarget *ProteinDesignEstimateCostParamsTargetStructureTemplateTarget `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTarget        *ProteinDesignEstimateCostParamsTargetNoTemplateTarget        `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetStructureTemplateTarget, u.OfProteinDesignEstimateCostsTargetNoTemplateTarget)
}
func (u *ProteinDesignEstimateCostParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
//
// The properties ChainSelection, Structure, Type are required.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTarget struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion `json:"chain_selection,omitzero" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion `json:"structure,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "structure_template".
	Type constant.StructureTemplate `json:"type" default:"structure_template"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetStructureTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetStructureTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetStructureTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion struct {
	OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec, u.OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec)
}
func (u *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
//
// The properties ChainType, CropResidues are required.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues,omitzero" api:"required"`
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

func (r ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	OfIntArray []int64 `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIntArray, u.OfAll)
}
func (u *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func NewProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec() ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec {
	return ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec{
		ChainType: "ligand",
	}
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
//
// This struct has a constant value, construct it with
// [NewProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec].
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion struct {
	OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureURLSource       *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource       `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureCifBase64Source *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureURLSource, u.OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureCifBase64Source)
}
func (u *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source struct {
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

func (r ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
//
// The properties Entities, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTarget struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinDesignEstimateCostParamsTargetNoTemplateTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintUnion `json:"constraints,omitzero"`
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityUnion struct {
	OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityProteinEntity      *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity      `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityRnaEntity          *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity          `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityDnaEntity          *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity          `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityLigandCcdEntity    *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityLigandSmilesEntity *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityProteinEntity,
		u.OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityRnaEntity,
		u.OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityDnaEntity,
		u.OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityLigandCcdEntity,
		u.OfProteinDesignEstimateCostsTargetNoTemplateTargetEntityLigandSmilesEntity)
}
func (u *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification struct {
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1Union struct {
	OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom1LigandAtom  *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom1PolymerAtom *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom1LigandAtom, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom1PolymerAtom)
}
func (u *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom struct {
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2Union struct {
	OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom2LigandAtom  *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom2PolymerAtom *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom2LigandAtom, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetBondAtom2PolymerAtom)
}
func (u *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom struct {
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintUnion struct {
	OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintPocketConstraint  *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraint *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintPocketConstraint, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraint)
}
func (u *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint struct {
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

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union struct {
	OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union struct {
	OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken, u.OfProteinDesignEstimateCostsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignListResultsParams struct {
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

// URLQuery serializes [ProteinDesignListResultsParams]'s query parameters as
// `url.Values`.
func (r ProteinDesignListResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinDesignStartParams struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, or boltz_curated for
	// Boltz-managed nanobody and antibody defaults.
	BinderSpecification ProteinDesignStartParamsBinderSpecificationUnion `json:"binder_specification,omitzero" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinDesignStartParamsTargetUnion `json:"target,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	paramObj
}

func (r ProteinDesignStartParams) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationUnion struct {
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpec *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpec        *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpec        `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationBoltzCuratedBinderSpec      *ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpec      `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpec, u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpec, u.OfProteinDesignStartsBinderSpecificationBoltzCuratedBinderSpec)
}
func (u *ProteinDesignStartParamsBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// engine run.
//
// The properties ChainSelection, Modality, Structure, Type are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep (crop_residues).
	// Omit design_motifs to include the chain as fixed scaffold context.
	ChainSelection map[string]ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion `json:"chain_selection,omitzero" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModality `json:"modality,omitzero" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion `json:"structure,omitzero" api:"required"`
	// Constraints applied during sequence design
	Rules ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules `json:"rules,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "structure_template".
	Type constant.StructureTemplate `json:"type" default:"structure_template"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion struct {
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec  *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec)
}
func (u *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
//
// The properties ChainType, CropResidues are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec struct {
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues,omitzero" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs,omitzero"`
	// This field can be elided, and will marshal its zero value as "polymer".
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
	OfIntArray []int64 `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIntArray, u.OfAll)
}
func (u *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif   *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif   `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif)
}
func (u *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
//
// The properties DesignLengthRange, EndIndex, StartIndex, Type are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range,omitzero" api:"required"`
	// 0-indexed end residue (inclusive)
	EndIndex int64 `json:"end_index" api:"required"`
	// 0-indexed start residue (inclusive)
	StartIndex int64 `json:"start_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "replacement".
	Type constant.Replacement `json:"type" default:"replacement"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
//
// The properties Max, Min are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
//
// The properties AfterResidueIndex, DesignLengthRange, Type are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "insertion".
	Type constant.Insertion `json:"type" default:"insertion"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
//
// The properties Max, Min are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
	// Maximum sequence length in residues. Must be >= min.
	Max int64 `json:"max" api:"required"`
	// Minimum sequence length in residues
	Min int64 `json:"min" api:"required"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec() ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec {
	return ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec{
		ChainType: "ligand",
	}
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
//
// This struct has a constant value, construct it with
// [NewProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec].
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModality string

const (
	ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide       ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModality = "peptide"
	ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityAntibody      ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModality = "antibody"
	ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityNanobody      ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModality = "nanobody"
	ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityCustomProtein ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModality = "custom_protein"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion struct {
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureURLSource       *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource       `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureURLSource, u.OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source)
}
func (u *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source struct {
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

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules struct {
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction param.Opt[float64] `json:"max_hydrophobic_fraction,omitzero"`
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids,omitzero"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs,omitzero"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
//
// The properties Entities, Modality, Type are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpec struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityUnion `json:"entities,omitzero" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModality `json:"modality,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBond `json:"bonds,omitzero"`
	// Constraints applied during sequence design
	Rules ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecRules `json:"rules,omitzero"`
	// This field can be elided, and will marshal its zero value as "no_template".
	Type constant.NoTemplate `json:"type" default:"no_template"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityUnion struct {
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity   *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity   `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity      *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity      `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity          *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity          `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity          *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity          `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity    *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity    `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity,
		u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity,
		u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity,
		u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity,
		u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity,
		u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity)
}
func (u *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Protein binder entity with designed and/or fixed segments.
//
// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Binder sequence specification. Fixed amino acids are written as literal
	// single-letter codes. Designed regions are written as a length (fixed) or a
	// length range (min..max). Example: "MKTAYI5..10VKSHFSRQ" means fixed MKTAYI, then
	// 5-10 designed residues, then fixed VKSHFSRQ. "20" means 20 fully designed
	// residues. "ACDE8GHI" means fixed ACDE, then 8 designed residues, then fixed GHI.
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "designed_protein".
	Type constant.DesignedProtein `json:"type" default:"designed_protein"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification struct {
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

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityDesignedProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
//
// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification struct {
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

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification struct {
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

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification struct {
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

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity struct {
	// Chain IDs to assign to this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code from RCSB PDB (e.g. 'ATP', 'ADP')
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecEntityFixedLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModality string

const (
	ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModalityPeptide       ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModality = "peptide"
	ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModalityAntibody      ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModality = "antibody"
	ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModalityNanobody      ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModality = "nanobody"
	ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModalityCustomProtein ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecModality = "custom_protein"
)

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union struct {
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom  *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom, u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom)
}
func (u *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom struct {
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

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union struct {
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom  *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom, u.OfProteinDesignStartsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom)
}
func (u *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom struct {
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

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecRules struct {
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction param.Opt[float64] `json:"max_hydrophobic_fraction,omitzero"`
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids,omitzero"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs,omitzero"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecRules) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationNoTemplateBinderSpecRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
//
// The properties Binder, Type are required.
type ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpec struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecBinder `json:"binder,omitzero" api:"required"`
	// Constraints applied during sequence design
	Rules ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecRules `json:"rules,omitzero"`
	// This field can be elided, and will marshal its zero value as "boltz_curated".
	Type constant.BoltzCurated `json:"type" default:"boltz_curated"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecBinder string

const (
	ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecBinderBoltzNanobody ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecBinder = "boltz_nanobody"
	ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecBinderBoltzAntibody ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecRules struct {
	// Maximum allowed fraction of hydrophobic residues (I, L, V, M, F, W, Y) in
	// designed regions. Designs exceeding this threshold are filtered out before
	// scoring. Leave empty to disable.
	MaxHydrophobicFraction param.Opt[float64] `json:"max_hydrophobic_fraction,omitzero"`
	// Single-letter amino acid codes to exclude from design (e.g. ['C', 'P'] to
	// exclude cysteine and proline)
	ExcludedAminoAcids []string `json:"excluded_amino_acids,omitzero"`
	// Sequence motifs to exclude from designed regions. Designs containing any of
	// these motifs are filtered out before scoring. Use X as a single-residue wildcard
	// (e.g. "NGS", "NXS").
	ExcludedSequenceMotifs []string `json:"excluded_sequence_motifs,omitzero"`
	paramObj
}

func (r ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecRules) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsBinderSpecificationBoltzCuratedBinderSpecRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetUnion struct {
	OfProteinDesignStartsTargetStructureTemplateTarget *ProteinDesignStartParamsTargetStructureTemplateTarget `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTarget        *ProteinDesignStartParamsTargetNoTemplateTarget        `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetStructureTemplateTarget, u.OfProteinDesignStartsTargetNoTemplateTarget)
}
func (u *ProteinDesignStartParamsTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
//
// The properties ChainSelection, Structure, Type are required.
type ProteinDesignStartParamsTargetStructureTemplateTarget struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the engine run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion `json:"chain_selection,omitzero" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion `json:"structure,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "structure_template".
	Type constant.StructureTemplate `json:"type" default:"structure_template"`
	paramObj
}

func (r ProteinDesignStartParamsTargetStructureTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetStructureTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetStructureTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion struct {
	OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec *ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec `json:",omitzero,inline"`
	OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  *ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec, u.OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec)
}
func (u *ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
//
// The properties ChainType, CropResidues are required.
type ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the engine run.
	CropResidues ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues,omitzero" api:"required"`
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

func (r ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
	OfIntArray []int64 `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIntArray, u.OfAll)
}
func (u *ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func NewProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec() ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec {
	return ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec{
		ChainType: "ligand",
	}
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
//
// This struct has a constant value, construct it with
// [NewProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec].
type ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	paramObj
}

func (r ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion struct {
	OfProteinDesignStartsTargetStructureTemplateTargetStructureURLSource       *ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource       `json:",omitzero,inline"`
	OfProteinDesignStartsTargetStructureTemplateTargetStructureCifBase64Source *ProteinDesignStartParamsTargetStructureTemplateTargetStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetStructureTemplateTargetStructureURLSource, u.OfProteinDesignStartsTargetStructureTemplateTargetStructureCifBase64Source)
}
func (u *ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinDesignStartParamsTargetStructureTemplateTargetStructureCifBase64Source struct {
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

func (r ProteinDesignStartParamsTargetStructureTemplateTargetStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetStructureTemplateTargetStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetStructureTemplateTargetStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
//
// The properties Entities, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTarget struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinDesignStartParamsTargetNoTemplateTargetEntityUnion `json:"entities,omitzero" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ProteinDesignStartParamsTargetNoTemplateTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ProteinDesignStartParamsTargetNoTemplateTargetConstraintUnion `json:"constraints,omitzero"`
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

func (r ProteinDesignStartParamsTargetNoTemplateTarget) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityUnion struct {
	OfProteinDesignStartsTargetNoTemplateTargetEntityProteinEntity      *ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntity      `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetEntityRnaEntity          *ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntity          `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetEntityDnaEntity          *ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntity          `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetEntityLigandCcdEntity    *ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetEntityLigandSmilesEntity *ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetNoTemplateTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetNoTemplateTargetEntityProteinEntity,
		u.OfProteinDesignStartsTargetNoTemplateTargetEntityRnaEntity,
		u.OfProteinDesignStartsTargetNoTemplateTargetEntityDnaEntity,
		u.OfProteinDesignStartsTargetNoTemplateTargetEntityLigandCcdEntity,
		u.OfProteinDesignStartsTargetNoTemplateTargetEntityLigandSmilesEntity)
}
func (u *ProteinDesignStartParamsTargetNoTemplateTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntityModification struct {
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

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntityModification struct {
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

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntityModification struct {
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

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type ProteinDesignStartParamsTargetNoTemplateTargetBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1Union struct {
	OfProteinDesignStartsTargetNoTemplateTargetBondAtom1LigandAtom  *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetBondAtom1PolymerAtom *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetNoTemplateTargetBondAtom1LigandAtom, u.OfProteinDesignStartsTargetNoTemplateTargetBondAtom1PolymerAtom)
}
func (u *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom struct {
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

func (r ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2Union struct {
	OfProteinDesignStartsTargetNoTemplateTargetBondAtom2LigandAtom  *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetBondAtom2PolymerAtom *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetNoTemplateTargetBondAtom2LigandAtom, u.OfProteinDesignStartsTargetNoTemplateTargetBondAtom2PolymerAtom)
}
func (u *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom struct {
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

func (r ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintUnion struct {
	OfProteinDesignStartsTargetNoTemplateTargetConstraintPocketConstraint  *ProteinDesignStartParamsTargetNoTemplateTargetConstraintPocketConstraint  `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraint *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetNoTemplateTargetConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetNoTemplateTargetConstraintPocketConstraint, u.OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraint)
}
func (u *ProteinDesignStartParamsTargetNoTemplateTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintPocketConstraint struct {
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

func (r ProteinDesignStartParamsTargetNoTemplateTargetConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union struct {
	OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken, u.OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken)
}
func (u *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union struct {
	OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken, u.OfProteinDesignStartsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken)
}
func (u *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinDesignStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
