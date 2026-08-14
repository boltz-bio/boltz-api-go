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

// Resume a stopped protein library screen from its last checkpoint
func (r *ProteinLibraryScreenService) Resume(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinLibraryScreenResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/library-screen/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
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

// A protein library screening pipeline run
type ProteinLibraryScreenGetResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzprot `json:"engine" default:"boltzprot"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                   `json:"engine_version" default:"1.0"`
	Error         ProteinLibraryScreenGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                      `json:"pipeline_version" default:"1.0"`
	Progress        ProteinLibraryScreenGetResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                               `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenGetResponseStatus `json:"status" api:"required"`
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
	// listed here are included in the pipeline run — any chains omitted from this
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
	// residues. Residues not listed are excluded from the pipeline run.
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
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
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
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Bonds []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Residues []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) {
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
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                       `json:"type" default:"glycan"`
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond struct {
	Atom1 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
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
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse].
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

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
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

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse],
// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse].
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

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
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

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
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

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenGetResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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

// Summary of a protein library screening pipeline run (excludes input)
type ProteinLibraryScreenListResponse struct {
	// Unique ProteinLibraryScreenSummary identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzprot `json:"engine" default:"boltzprot"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                    `json:"engine_version" default:"1.0"`
	Error         ProteinLibraryScreenListResponseError `json:"error" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                       `json:"pipeline_version" default:"1.0"`
	Progress        ProteinLibraryScreenListResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenListResponseStatus `json:"status" api:"required"`
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
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application ProteinLibraryScreenEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
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
	ProteinLibraryScreenEstimateCostResponseBreakdownApplicationProteinRedesign            ProteinLibraryScreenEstimateCostResponseBreakdownApplication = "protein_redesign"
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
// [ProteinLibraryScreenListResultsResponseEntityLigandSmilesEntity],
// [ProteinLibraryScreenListResultsResponseEntityGlycanEntity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenListResultsResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenListResultsResponseEntityProteinEntityModification],
	// [[]ProteinLibraryScreenListResultsResponseEntityRnaEntityModification],
	// [[]ProteinLibraryScreenListResultsResponseEntityDnaEntityModification]
	Modifications ProteinLibraryScreenListResultsResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinLibraryScreenListResultsResponseEntityGlycanEntity].
	Bonds []ProteinLibraryScreenListResultsResponseEntityGlycanEntityBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenListResultsResponseEntityGlycanEntity].
	Residues []ProteinLibraryScreenListResultsResponseEntityGlycanEntityResidue `json:"residues"`
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

func (u ProteinLibraryScreenListResultsResponseEntityUnion) AsProteinLibraryScreenListResultsResponseEntityGlycanEntity() (v ProteinLibraryScreenListResultsResponseEntityGlycanEntity) {
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
	// [[]ProteinLibraryScreenListResultsResponseEntityProteinEntityModification]
	// instead of an object.
	OfProteinLibraryScreenListResultsResponseEntityProteinEntityModifications []ProteinLibraryScreenListResultsResponseEntityProteinEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenListResultsResponseEntityRnaEntityModification] instead
	// of an object.
	OfProteinLibraryScreenListResultsResponseEntityRnaEntityModifications []ProteinLibraryScreenListResultsResponseEntityRnaEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenListResultsResponseEntityDnaEntityModification] instead
	// of an object.
	OfProteinLibraryScreenListResultsResponseEntityDnaEntityModifications []ProteinLibraryScreenListResultsResponseEntityDnaEntityModification `json:",inline"`
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenListResultsResponseEntityProteinEntityModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenListResultsResponseEntityProteinEntityModification struct {
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
func (r ProteinLibraryScreenListResultsResponseEntityProteinEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenListResultsResponseEntityRnaEntityModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenListResultsResponseEntityRnaEntityModification struct {
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
func (r ProteinLibraryScreenListResultsResponseEntityRnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenListResultsResponseEntityDnaEntityModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenListResultsResponseEntityDnaEntityModification struct {
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
func (r ProteinLibraryScreenListResultsResponseEntityDnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinLibraryScreenListResultsResponseEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenListResultsResponseEntityGlycanEntityBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenListResultsResponseEntityGlycanEntityResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                    `json:"type" default:"glycan"`
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
func (r ProteinLibraryScreenListResultsResponseEntityGlycanEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinLibraryScreenListResultsResponseEntityGlycanEntityBond struct {
	Atom1 ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenListResultsResponseEntityGlycanEntityBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom1 struct {
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
func (r ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom2 struct {
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
func (r ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenListResultsResponseEntityGlycanEntityResidue struct {
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
func (r ProteinLibraryScreenListResultsResponseEntityGlycanEntityResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenListResultsResponseEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
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

// A protein library screening pipeline run
type ProteinLibraryScreenResumeResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzprot `json:"engine" default:"boltzprot"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                      `json:"engine_version" default:"1.0"`
	Error         ProteinLibraryScreenResumeResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenResumeResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                         `json:"pipeline_version" default:"1.0"`
	Progress        ProteinLibraryScreenResumeResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                  `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenResumeResponseStatus `json:"status" api:"required"`
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
func (r ProteinLibraryScreenResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseError struct {
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
func (r ProteinLibraryScreenResumeResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenResumeResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ProteinLibraryScreenResumeResponseInput struct {
	Proteins ProteinLibraryScreenResumeResponseInputProteins `json:"proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ProteinLibraryScreenResumeResponseInputTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Proteins    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenResumeResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenResumeResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputProteins struct {
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
func (r ProteinLibraryScreenResumeResponseInputProteins) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenResumeResponseInputProteins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetUnion contains all possible
// properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse].
	Structure ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                                `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
	Entities []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
	Bonds []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
	Constraints []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse].
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

func (u ProteinLibraryScreenResumeResponseInputTargetUnion) AsProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse() (v ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinLibraryScreenResumeResponseInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the pipeline run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                                 `json:"type" default:"structure_template"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
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

func (u ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the pipeline run.
	CropResidues ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                                `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Bonds []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Residues []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion].
// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                        struct {
		OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                             string
	} `json:"-"`
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                          `json:"type" default:"glycan"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond struct {
	Atom1 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse].
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

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse].
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

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                    `json:"type" default:"contact"`
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenResumeResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseProgress struct {
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
func (r ProteinLibraryScreenResumeResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinLibraryScreenResumeResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenResumeResponseStatus string

const (
	ProteinLibraryScreenResumeResponseStatusPending   ProteinLibraryScreenResumeResponseStatus = "pending"
	ProteinLibraryScreenResumeResponseStatusRunning   ProteinLibraryScreenResumeResponseStatus = "running"
	ProteinLibraryScreenResumeResponseStatusSucceeded ProteinLibraryScreenResumeResponseStatus = "succeeded"
	ProteinLibraryScreenResumeResponseStatusFailed    ProteinLibraryScreenResumeResponseStatus = "failed"
	ProteinLibraryScreenResumeResponseStatusStopped   ProteinLibraryScreenResumeResponseStatus = "stopped"
)

// A protein library screening pipeline run
type ProteinLibraryScreenStartResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzprot `json:"engine" default:"boltzprot"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                     `json:"engine_version" default:"1.0"`
	Error         ProteinLibraryScreenStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                        `json:"pipeline_version" default:"1.0"`
	Progress        ProteinLibraryScreenStartResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                 `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenStartResponseStatus `json:"status" api:"required"`
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
	// listed here are included in the pipeline run — any chains omitted from this
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
	// residues. Residues not listed are excluded from the pipeline run.
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
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
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
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Bonds []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Residues []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) {
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
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                         `json:"type" default:"glycan"`
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond struct {
	Atom1 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
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
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse].
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

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
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

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse],
// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse].
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

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
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

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
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

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenStartResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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

// A protein library screening pipeline run
type ProteinLibraryScreenStopResponse struct {
	// Unique ProteinLibraryScreen identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.Boltzprot `json:"engine" default:"boltzprot"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.String1_0                    `json:"engine_version" default:"1.0"`
	Error         ProteinLibraryScreenStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinLibraryScreenStopResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                       `json:"pipeline_version" default:"1.0"`
	Progress        ProteinLibraryScreenStopResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinLibraryScreenStopResponseStatus `json:"status" api:"required"`
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
	// listed here are included in the pipeline run — any chains omitted from this
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
	// residues. Residues not listed are excluded from the pipeline run.
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
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
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
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Bonds []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Residues []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityUnion) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) {
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
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
	// Chain IDs for this ligand
	ChainIDs []string           `json:"chain_ids" api:"required"`
	Type     constant.LigandCcd `json:"type" default:"ligand_ccd"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                        `json:"type" default:"glycan"`
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond struct {
	Atom1 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
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
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse].
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

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
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

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse],
// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse].
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

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2Union) AsProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
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

// Atom reference for a specific CCD residue in a glycan graph.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
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

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ProteinLibraryScreenStopResponseInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
	OfProteinLibraryScreenEstimateCostsProteinEntityGlycanEntity       *ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntity       `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsProteinEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityRnaEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityDnaEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityLigandCcdEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityLigandSmilesEntity,
		u.OfProteinLibraryScreenEstimateCostsProteinEntityGlycanEntity)
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModification struct {
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

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModification struct {
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

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModification struct {
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

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBond struct {
	Atom1 ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsProteinEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
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
	// listed here are included in the pipeline run — any chains omitted from this
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
	// residues. Residues not listed are excluded from the pipeline run.
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
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
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
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityGlycanEntity       *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntity       `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityProteinEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityRnaEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityDnaEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityLigandCcdEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityLigandSmilesEntity,
		u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetEntityGlycanEntity)
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBond struct {
	Atom1 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
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
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1PolymerAtom *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1CcdAtom     *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1SmilesAtom  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1LigandAtom  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1PolymerAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1CcdAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1SmilesAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom1LigandAtom)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1CcdAtom struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union struct {
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2PolymerAtom *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2CcdAtom     *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2SmilesAtom  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2LigandAtom  *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2PolymerAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2CcdAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2SmilesAtom, u.OfProteinLibraryScreenEstimateCostsTargetNoTemplateTargetBondAtom2LigandAtom)
}
func (u *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2CcdAtom struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom struct {
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

func (r ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
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

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken struct {
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenEstimateCostParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken struct {
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
	OfProteinLibraryScreenStartsProteinEntityGlycanEntity       *ProteinLibraryScreenStartParamsProteinEntityGlycanEntity       `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsProteinEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsProteinEntityProteinEntity,
		u.OfProteinLibraryScreenStartsProteinEntityRnaEntity,
		u.OfProteinLibraryScreenStartsProteinEntityDnaEntity,
		u.OfProteinLibraryScreenStartsProteinEntityLigandCcdEntity,
		u.OfProteinLibraryScreenStartsProteinEntityLigandSmilesEntity,
		u.OfProteinLibraryScreenStartsProteinEntityGlycanEntity)
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartParamsProteinEntityProteinEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityProteinEntityModification struct {
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

func (r ProteinLibraryScreenStartParamsProteinEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartParamsProteinEntityRnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityRnaEntityModification struct {
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

func (r ProteinLibraryScreenStartParamsProteinEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartParamsProteinEntityDnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityDnaEntityModification struct {
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

func (r ProteinLibraryScreenStartParamsProteinEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsProteinEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type ProteinLibraryScreenStartParamsProteinEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenStartParamsProteinEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBond struct {
	Atom1 ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type ProteinLibraryScreenStartParamsProteinEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsProteinEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsProteinEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsProteinEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
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
	// listed here are included in the pipeline run — any chains omitted from this
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
	// residues. Residues not listed are excluded from the pipeline run.
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
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
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
	OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityGlycanEntity       *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntity       `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityProteinEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityRnaEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityDnaEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityLigandCcdEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityLigandSmilesEntity,
		u.OfProteinLibraryScreenStartsTargetNoTemplateTargetEntityGlycanEntity)
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
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModification struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModification struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
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
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModification `json:"modifications,omitzero"`
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

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModification struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// One CCD code (for example ATP or ADP). This field remains a string; use a glycan
	// entity for multiple connected CCD residues.
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
//
// The properties Bonds, ChainIDs, Residues, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBond `json:"bonds,omitzero" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityResidue `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "glycan".
	Type constant.Glycan `json:"type" default:"glycan"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBond struct {
	Atom1 ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1 `json:"atom1,omitzero" api:"required"`
	Atom2 ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2 `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBond) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomID, ResidueID are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2 struct {
	// Exact atom identifier from the residue CCD entry (\_chem_comp_atom.atom_id)
	AtomID string `json:"atom_id" api:"required"`
	// Request-local ID of the glycan residue containing the atom
	ResidueID string `json:"residue_id" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Ccd are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityResidue struct {
	// Request-local residue ID used by glycan bonds and external atom references
	ID string `json:"id" api:"required"`
	// CCD code for this monosaccharide residue (for example NAG, BMA, or FUC)
	Ccd string `json:"ccd" api:"required"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityResidue) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityResidue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
//
// The properties Atom1, Atom2 are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
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
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1PolymerAtom *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1PolymerAtom `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1CcdAtom     *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1CcdAtom     `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1SmilesAtom  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1SmilesAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1LigandAtom  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1PolymerAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1CcdAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1SmilesAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom1LigandAtom)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1CcdAtom struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union struct {
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2PolymerAtom *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2PolymerAtom `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2CcdAtom     *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2CcdAtom     `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2SmilesAtom  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2SmilesAtom  `json:",omitzero,inline"`
	OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2LigandAtom  *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom  `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2PolymerAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2CcdAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2SmilesAtom, u.OfProteinLibraryScreenStartsTargetNoTemplateTargetBondAtom2LigandAtom)
}
func (u *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
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

// Atom reference for a specific CCD residue in a glycan graph.
//
// The properties AtomID, ChainID, ResidueID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2CcdAtom struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2CcdAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2CcdAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2CcdAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
//
// The properties AtomMap, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2SmilesAtom struct {
	// Numeric atom-map identifier from the input SMILES (for example 7 for [C:7])
	AtomMap int64 `json:"atom_map" api:"required"`
	// Chain ID containing the SMILES ligand
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "smiles_atom".
	Type constant.SmilesAtom `json:"type" default:"smiles_atom"`
	paramObj
}

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2SmilesAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2SmilesAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2SmilesAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
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

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken1LigandContactToken struct {
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

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
//
// The properties AtomName, ChainID, Type are required.
type ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken struct {
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

func (r ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinLibraryScreenStartParamsTargetNoTemplateTargetConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
