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

// Redesign selected protein residues in one fixed CIF structure. Use the top-level
// type discriminator to choose binder redesign, with target and binder chain
// roles, or generic redesign. Every chain in the input structure must be assigned
// exactly once. Binder results include binding and structure metrics; generic
// results include structure and secondary-structure metrics.
//
// ProteinSequenceRedesignService contains methods and other services that help
// with interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProteinSequenceRedesignService] method instead.
type ProteinSequenceRedesignService struct {
	Options []option.RequestOption
}

// NewProteinSequenceRedesignService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProteinSequenceRedesignService(opts ...option.RequestOption) (r ProteinSequenceRedesignService) {
	r = ProteinSequenceRedesignService{}
	r.Options = opts
	return
}

// Retrieve a sequence redesign run by ID, including progress and status
func (r *ProteinSequenceRedesignService) Get(ctx context.Context, id string, query ProteinSequenceRedesignGetParams, opts ...option.RequestOption) (res *ProteinSequenceRedesignGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/sequence-redesign/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List protein sequence redesign runs, optionally filtered by workspace
func (r *ProteinSequenceRedesignService) List(ctx context.Context, query ProteinSequenceRedesignListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProteinSequenceRedesignListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/protein/sequence-redesign"
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

// List protein sequence redesign runs, optionally filtered by workspace
func (r *ProteinSequenceRedesignService) ListAutoPaging(ctx context.Context, query ProteinSequenceRedesignListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProteinSequenceRedesignListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// sequence redesign run. The sequence redesign run record itself is retained with
// a `data_deleted_at` timestamp. This action is irreversible.
func (r *ProteinSequenceRedesignService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinSequenceRedesignDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/sequence-redesign/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Estimate the cost of a protein sequence redesign run without creating any
// resource or consuming GPU.
func (r *ProteinSequenceRedesignService) EstimateCost(ctx context.Context, body ProteinSequenceRedesignEstimateCostParams, opts ...option.RequestOption) (res *ProteinSequenceRedesignEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/protein/sequence-redesign/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve paginated results from a protein sequence redesign run
func (r *ProteinSequenceRedesignService) ListResults(ctx context.Context, id string, query ProteinSequenceRedesignListResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[ProteinSequenceRedesignListResultsResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/sequence-redesign/%s/results", url.PathEscape(id))
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

// Retrieve paginated results from a protein sequence redesign run
func (r *ProteinSequenceRedesignService) ListResultsAutoPaging(ctx context.Context, id string, query ProteinSequenceRedesignListResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ProteinSequenceRedesignListResultsResponseUnion] {
	return pagination.NewCursorPageAutoPager(r.ListResults(ctx, id, query, opts...))
}

// Resume a stopped protein sequence redesign run from its last checkpoint
func (r *ProteinSequenceRedesignService) Resume(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinSequenceRedesignResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/sequence-redesign/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Create a protein sequence redesign run from selected residues in a fixed input
// structure
func (r *ProteinSequenceRedesignService) Start(ctx context.Context, body ProteinSequenceRedesignStartParams, opts ...option.RequestOption) (res *ProteinSequenceRedesignStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/protein/sequence-redesign"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an in-progress protein sequence redesign run early
func (r *ProteinSequenceRedesignService) Stop(ctx context.Context, id string, opts ...option.RequestOption) (res *ProteinSequenceRedesignStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/protein/sequence-redesign/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A fixed-structure protein sequence redesign run.
type ProteinSequenceRedesignGetResponse struct {
	// Unique ProteinSequenceRedesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.BoltzProteinRedesign `json:"engine" default:"boltz-protein-redesign"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.V2026_07_14                    `json:"engine_version" default:"v2026-07-14"`
	Error         ProteinSequenceRedesignGetResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinSequenceRedesignGetResponseInputUnion `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode        bool                                       `json:"livemode" api:"required"`
	Pipeline        constant.BoltzProteinRedesign              `json:"pipeline" default:"boltz-protein-redesign"`
	PipelineVersion constant.V2026_07_14                       `json:"pipeline_version" default:"v2026-07-14"`
	Progress        ProteinSequenceRedesignGetResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                  `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinSequenceRedesignGetResponseStatus `json:"status" api:"required"`
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
func (r ProteinSequenceRedesignGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseError struct {
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
func (r ProteinSequenceRedesignGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputUnion contains all possible properties
// and values from
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponse],
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignGetResponseInputUnion struct {
	// This field is a union of
	// [[]ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion],
	// [[]ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	Entities    ProteinSequenceRedesignGetResponseInputUnionEntities `json:"entities"`
	NumProteins int64                                                `json:"num_proteins"`
	// This field is a union of
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseStructure],
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseStructure]
	Structure ProteinSequenceRedesignGetResponseInputUnionStructure `json:"structure"`
	Type      string                                                `json:"type"`
	// This field is a union of
	// [[]ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion],
	// [[]ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	GlobalDesignFilters ProteinSequenceRedesignGetResponseInputUnionGlobalDesignFilters `json:"global_design_filters"`
	IdempotencyKey      string                                                          `json:"idempotency_key"`
	WorkspaceID         string                                                          `json:"workspace_id"`
	JSON                struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinSequenceRedesignGetResponseInputUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignGetResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinSequenceRedesignGetResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputUnionEntities is an implicit subunion of
// [ProteinSequenceRedesignGetResponseInputUnion].
// ProteinSequenceRedesignGetResponseInputUnionEntities provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignGetResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntities
// OfProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntities]
type ProteinSequenceRedesignGetResponseInputUnionEntities struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion]
	// instead of an object.
	OfProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	// instead of an object.
	OfProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:",inline"`
	JSON                                                                                            struct {
		OfProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntities  respjson.Field
		OfProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntities respjson.Field
		raw                                                                                             string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignGetResponseInputUnionEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputUnionStructure is an implicit subunion of
// [ProteinSequenceRedesignGetResponseInputUnion].
// ProteinSequenceRedesignGetResponseInputUnionStructure provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignGetResponseInputUnion].
type ProteinSequenceRedesignGetResponseInputUnionStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignGetResponseInputUnionStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputUnionGlobalDesignFilters is an implicit
// subunion of [ProteinSequenceRedesignGetResponseInputUnion].
// ProteinSequenceRedesignGetResponseInputUnionGlobalDesignFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignGetResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters
// OfProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters]
type ProteinSequenceRedesignGetResponseInputUnionGlobalDesignFilters struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	JSON                                                                                                       struct {
		OfProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters  respjson.Field
		OfProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters respjson.Field
		raw                                                                                                        string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignGetResponseInputUnionGlobalDesignFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                         `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Binder                                                                               `json:"type" default:"binder"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                        `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse],
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion struct {
	ChainID string `json:"chain_id"`
	Role    string `json:"role"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse].
	Type constant.FromTemplate `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
	DesignMotifs []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	JSON         struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed target chain from the input CIF.
type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Target       `json:"role" default:"target"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Binder       `json:"role" default:"binder"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                          `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Generic                                                                               `json:"type" default:"generic"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                         `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity struct {
	ChainID string                `json:"chain_id" api:"required"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignGetResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseProgress struct {
	// Number of protein designs generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein designs requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated    respjson.Field
		TotalProteinsToGenerate respjson.Field
		LatestResultID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignGetResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignGetResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignGetResponseStatus string

const (
	ProteinSequenceRedesignGetResponseStatusPending   ProteinSequenceRedesignGetResponseStatus = "pending"
	ProteinSequenceRedesignGetResponseStatusRunning   ProteinSequenceRedesignGetResponseStatus = "running"
	ProteinSequenceRedesignGetResponseStatusSucceeded ProteinSequenceRedesignGetResponseStatus = "succeeded"
	ProteinSequenceRedesignGetResponseStatusFailed    ProteinSequenceRedesignGetResponseStatus = "failed"
	ProteinSequenceRedesignGetResponseStatusStopped   ProteinSequenceRedesignGetResponseStatus = "stopped"
)

// Summary of a protein sequence redesign run.
type ProteinSequenceRedesignListResponse struct {
	// Unique ProteinSequenceRedesignRunSummary identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.BoltzProteinRedesign `json:"engine" default:"boltz-protein-redesign"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.V2026_07_14                     `json:"engine_version" default:"v2026-07-14"`
	Error         ProteinSequenceRedesignListResponseError `json:"error" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode        bool                                        `json:"livemode" api:"required"`
	Pipeline        constant.BoltzProteinRedesign               `json:"pipeline" default:"boltz-protein-redesign"`
	PipelineVersion constant.V2026_07_14                        `json:"pipeline_version" default:"v2026-07-14"`
	Progress        ProteinSequenceRedesignListResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                   `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinSequenceRedesignListResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                 `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinSequenceRedesignListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResponseError struct {
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
func (r ProteinSequenceRedesignListResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResponseProgress struct {
	// Number of protein designs generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein designs requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated    respjson.Field
		TotalProteinsToGenerate respjson.Field
		LatestResultID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignListResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResponseStatus string

const (
	ProteinSequenceRedesignListResponseStatusPending   ProteinSequenceRedesignListResponseStatus = "pending"
	ProteinSequenceRedesignListResponseStatusRunning   ProteinSequenceRedesignListResponseStatus = "running"
	ProteinSequenceRedesignListResponseStatusSucceeded ProteinSequenceRedesignListResponseStatus = "succeeded"
	ProteinSequenceRedesignListResponseStatusFailed    ProteinSequenceRedesignListResponseStatus = "failed"
	ProteinSequenceRedesignListResponseStatusStopped   ProteinSequenceRedesignListResponseStatus = "stopped"
)

type ProteinSequenceRedesignDeleteDataResponse struct {
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
func (r ProteinSequenceRedesignDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimate response with monetary values encoded as decimal strings to preserve
// precision.
type ProteinSequenceRedesignEstimateCostResponse struct {
	// Cost breakdown for the billed application.
	Breakdown  ProteinSequenceRedesignEstimateCostResponseBreakdown `json:"breakdown" api:"required"`
	Disclaimer string                                               `json:"disclaimer" api:"required"`
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
func (r ProteinSequenceRedesignEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost breakdown for the billed application.
type ProteinSequenceRedesignEstimateCostResponseBreakdown struct {
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application ProteinSequenceRedesignEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
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
func (r ProteinSequenceRedesignEstimateCostResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignEstimateCostResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignEstimateCostResponseBreakdownApplication string

const (
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationStructureAndBinding        ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "structure_and_binding"
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationSmallMoleculeDesign        ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "small_molecule_design"
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationSmallMoleculeLibraryScreen ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "small_molecule_library_screen"
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationProteinDesign              ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "protein_design"
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationProteinRedesign            ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "protein_redesign"
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationProteinLibraryScreen       ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "protein_library_screen"
	ProteinSequenceRedesignEstimateCostResponseBreakdownApplicationAdme                       ProteinSequenceRedesignEstimateCostResponseBreakdownApplication = "adme"
)

// ProteinSequenceRedesignListResultsResponseUnion contains all possible properties
// and values from
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResult],
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignListResultsResponseUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifacts],
	// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifacts]
	Artifacts ProteinSequenceRedesignListResultsResponseUnionArtifacts `json:"artifacts"`
	CreatedAt time.Time                                                `json:"created_at"`
	// This field is a union of
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion],
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion]
	Entities ProteinSequenceRedesignListResultsResponseUnionEntities `json:"entities"`
	// This field is a union of
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics],
	// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultMetrics]
	Metrics ProteinSequenceRedesignListResultsResponseUnionMetrics `json:"metrics"`
	Type    string                                                 `json:"type"`
	// This field is a union of
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning],
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning]
	Warnings ProteinSequenceRedesignListResultsResponseUnionWarnings `json:"warnings"`
	JSON     struct {
		ID        respjson.Field
		Artifacts respjson.Field
		CreatedAt respjson.Field
		Entities  respjson.Field
		Metrics   respjson.Field
		Type      respjson.Field
		Warnings  respjson.Field
		raw       string
	} `json:"-"`
}

func (u ProteinSequenceRedesignListResultsResponseUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResult() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResult() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignListResultsResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinSequenceRedesignListResultsResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseUnionArtifacts is an implicit subunion
// of [ProteinSequenceRedesignListResultsResponseUnion].
// ProteinSequenceRedesignListResultsResponseUnionArtifacts provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseUnion].
type ProteinSequenceRedesignListResultsResponseUnionArtifacts struct {
	// This field is a union of
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsArchive],
	// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsArchive]
	Archive ProteinSequenceRedesignListResultsResponseUnionArtifactsArchive `json:"archive"`
	// This field is a union of
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsStructure],
	// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsStructure]
	Structure ProteinSequenceRedesignListResultsResponseUnionArtifactsStructure `json:"structure"`
	JSON      struct {
		Archive   respjson.Field
		Structure respjson.Field
		raw       string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseUnionArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseUnionArtifactsArchive is an implicit
// subunion of [ProteinSequenceRedesignListResultsResponseUnion].
// ProteinSequenceRedesignListResultsResponseUnionArtifactsArchive provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseUnion].
type ProteinSequenceRedesignListResultsResponseUnionArtifactsArchive struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseUnionArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseUnionArtifactsStructure is an implicit
// subunion of [ProteinSequenceRedesignListResultsResponseUnion].
// ProteinSequenceRedesignListResultsResponseUnionArtifactsStructure provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseUnion].
type ProteinSequenceRedesignListResultsResponseUnionArtifactsStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseUnionArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseUnionEntities is an implicit subunion
// of [ProteinSequenceRedesignListResultsResponseUnion].
// ProteinSequenceRedesignListResultsResponseUnionEntities provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntities
// OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntities]
type ProteinSequenceRedesignListResultsResponseUnionEntities struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntities []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntities []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion `json:",inline"`
	JSON                                                                           struct {
		OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntities  respjson.Field
		OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntities respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseUnionEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseUnionMetrics is an implicit subunion
// of [ProteinSequenceRedesignListResultsResponseUnion].
// ProteinSequenceRedesignListResultsResponseUnionMetrics provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseUnion].
type ProteinSequenceRedesignListResultsResponseUnionMetrics struct {
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics].
	BindingConfidence float64 `json:"binding_confidence"`
	HelixFraction     float64 `json:"helix_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics].
	Iptm         float64 `json:"iptm"`
	LoopFraction float64 `json:"loop_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics].
	MinInteractionPae   float64 `json:"min_interaction_pae"`
	SheetFraction       float64 `json:"sheet_fraction"`
	StructureConfidence float64 `json:"structure_confidence"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics].
	IpsaeMin float64 `json:"ipsae_min"`
	JSON     struct {
		BindingConfidence   respjson.Field
		HelixFraction       respjson.Field
		Iptm                respjson.Field
		LoopFraction        respjson.Field
		MinInteractionPae   respjson.Field
		SheetFraction       respjson.Field
		StructureConfidence respjson.Field
		IpsaeMin            respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseUnionMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseUnionWarnings is an implicit subunion
// of [ProteinSequenceRedesignListResultsResponseUnion].
// ProteinSequenceRedesignListResultsResponseUnionWarnings provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarnings
// OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarnings]
type ProteinSequenceRedesignListResultsResponseUnionWarnings struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarnings []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarnings []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning `json:",inline"`
	JSON                                                                           struct {
		OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarnings  respjson.Field
		OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarnings respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseUnionWarnings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResult struct {
	// Unique result ID.
	ID        string                                                                       `json:"id" api:"required"`
	Artifacts ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                                                    `json:"created_at" api:"required" format:"date-time"`
	// Designed and fixed entities returned for this result.
	Entities []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion `json:"entities" api:"required"`
	// Structural and binding quality metrics for a designed protein binder
	Metrics ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics `json:"metrics" api:"required"`
	Type    constant.Binder                                                            `json:"type" default:"binder"`
	// Warnings about potential quality issues with this result.
	Warnings []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Entities    respjson.Field
		Metrics     respjson.Field
		Type        respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifacts struct {
	Archive   ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsArchive   `json:"archive" api:"required"`
	Structure ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsStructure `json:"structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archive     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifacts) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsArchive struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsArchive) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsStructure struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntity],
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntity],
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntity],
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandCcdEntity],
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandSmilesEntity],
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification],
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification],
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification]
	Modifications ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity].
	Bonds []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBond `json:"bonds"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity].
	Residues []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityResidue `json:"residues"`
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

func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntity() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntity() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntity() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandCcdEntity() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandCcdEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandSmilesEntity() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandSmilesEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity() (v ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnionModifications
// is an implicit subunion of
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion].
// ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModifications
// OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModifications
// OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModifications]
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModifications []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModifications []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModifications []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification `json:",inline"`
	JSON                                                                                              struct {
		OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModifications respjson.Field
		OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModifications     respjson.Field
		OfProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModifications     respjson.Field
		raw                                                                                                   string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification `json:"modifications"`
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification `json:"modifications"`
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification `json:"modifications"`
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandCcdEntity struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandCcdEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandSmilesEntity struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandSmilesEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                `json:"type" default:"glycan"`
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBond struct {
	Atom1 ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom1 struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom2 struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityResidue struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structural and binding quality metrics for a designed protein binder
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics struct {
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
	// Lower of the target-to-binder and binder-to-target ipSAE scores using a 10
	// Angstrom PAE cutoff. Higher values indicate a more confidently predicted
	// interface.
	IpsaeMin float64 `json:"ipsae_min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence   respjson.Field
		HelixFraction       respjson.Field
		Iptm                respjson.Field
		LoopFraction        respjson.Field
		MinInteractionPae   respjson.Field
		SheetFraction       respjson.Field
		StructureConfidence respjson.Field
		IpsaeMin            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warning about a potential quality issue with a result
type ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning struct {
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
func (r ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseBinderProteinDesignResultWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResult struct {
	// Unique result ID.
	ID        string                                                                        `json:"id" api:"required"`
	Artifacts ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                                                     `json:"created_at" api:"required" format:"date-time"`
	// Designed and fixed entities returned for this result.
	Entities []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion `json:"entities" api:"required"`
	// Structure and design-quality metrics for a generic protein design.
	Metrics ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultMetrics `json:"metrics" api:"required"`
	Type    constant.Generic                                                            `json:"type" default:"generic"`
	// Warnings about potential quality issues with this result.
	Warnings []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Artifacts   respjson.Field
		CreatedAt   respjson.Field
		Entities    respjson.Field
		Metrics     respjson.Field
		Type        respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifacts struct {
	Archive   ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsArchive   `json:"archive" api:"required"`
	Structure ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsStructure `json:"structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archive     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifacts) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsArchive struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsArchive) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsStructure struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntity],
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntity],
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntity],
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandCcdEntity],
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandSmilesEntity],
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification],
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification],
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification]
	Modifications ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity].
	Bonds []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBond `json:"bonds"`
	// This field is from variant
	// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity].
	Residues []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityResidue `json:"residues"`
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

func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntity() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntity() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntity() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandCcdEntity() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandCcdEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandSmilesEntity() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandSmilesEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) AsProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity() (v ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnionModifications
// is an implicit subunion of
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion].
// ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModifications
// OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModifications
// OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModifications]
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModifications []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModifications []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification]
	// instead of an object.
	OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModifications []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification `json:",inline"`
	JSON                                                                                               struct {
		OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModifications respjson.Field
		OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModifications     respjson.Field
		OfProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModifications     respjson.Field
		raw                                                                                                    string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification `json:"modifications"`
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification `json:"modifications"`
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification `json:"modifications"`
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandCcdEntity struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandCcdEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandSmilesEntity struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandSmilesEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                 `json:"type" default:"glycan"`
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBond struct {
	Atom1 ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom1 `json:"atom1" api:"required"`
	Atom2 ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom1 struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom2 struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityResidue struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structure and design-quality metrics for a generic protein design.
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultMetrics struct {
	// Fraction of the designed sequence forming alpha helices (0-1).
	HelixFraction float64 `json:"helix_fraction" api:"required"`
	// Fraction of the designed sequence in coil/loop regions (0-1).
	LoopFraction float64 `json:"loop_fraction" api:"required"`
	// Fraction of the designed sequence forming beta sheets (0-1).
	SheetFraction float64 `json:"sheet_fraction" api:"required"`
	// Confidence in the predicted 3D structure (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HelixFraction       respjson.Field
		LoopFraction        respjson.Field
		SheetFraction       respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warning about a potential quality issue with a result
type ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning struct {
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
func (r ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignListResultsResponseGenericProteinDesignResultWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed-structure protein sequence redesign run.
type ProteinSequenceRedesignResumeResponse struct {
	// Unique ProteinSequenceRedesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.BoltzProteinRedesign `json:"engine" default:"boltz-protein-redesign"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.V2026_07_14                       `json:"engine_version" default:"v2026-07-14"`
	Error         ProteinSequenceRedesignResumeResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinSequenceRedesignResumeResponseInputUnion `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode        bool                                          `json:"livemode" api:"required"`
	Pipeline        constant.BoltzProteinRedesign                 `json:"pipeline" default:"boltz-protein-redesign"`
	PipelineVersion constant.V2026_07_14                          `json:"pipeline_version" default:"v2026-07-14"`
	Progress        ProteinSequenceRedesignResumeResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                     `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinSequenceRedesignResumeResponseStatus `json:"status" api:"required"`
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
func (r ProteinSequenceRedesignResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseError struct {
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
func (r ProteinSequenceRedesignResumeResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignResumeResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputUnion contains all possible properties
// and values from
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponse],
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignResumeResponseInputUnion struct {
	// This field is a union of
	// [[]ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion],
	// [[]ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	Entities    ProteinSequenceRedesignResumeResponseInputUnionEntities `json:"entities"`
	NumProteins int64                                                   `json:"num_proteins"`
	// This field is a union of
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseStructure],
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseStructure]
	Structure ProteinSequenceRedesignResumeResponseInputUnionStructure `json:"structure"`
	Type      string                                                   `json:"type"`
	// This field is a union of
	// [[]ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion],
	// [[]ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	GlobalDesignFilters ProteinSequenceRedesignResumeResponseInputUnionGlobalDesignFilters `json:"global_design_filters"`
	IdempotencyKey      string                                                             `json:"idempotency_key"`
	WorkspaceID         string                                                             `json:"workspace_id"`
	JSON                struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinSequenceRedesignResumeResponseInputUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignResumeResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinSequenceRedesignResumeResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputUnionEntities is an implicit subunion
// of [ProteinSequenceRedesignResumeResponseInputUnion].
// ProteinSequenceRedesignResumeResponseInputUnionEntities provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignResumeResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntities
// OfProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntities]
type ProteinSequenceRedesignResumeResponseInputUnionEntities struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion]
	// instead of an object.
	OfProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	// instead of an object.
	OfProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:",inline"`
	JSON                                                                                               struct {
		OfProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntities  respjson.Field
		OfProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntities respjson.Field
		raw                                                                                                string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignResumeResponseInputUnionEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputUnionStructure is an implicit subunion
// of [ProteinSequenceRedesignResumeResponseInputUnion].
// ProteinSequenceRedesignResumeResponseInputUnionStructure provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignResumeResponseInputUnion].
type ProteinSequenceRedesignResumeResponseInputUnionStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignResumeResponseInputUnionStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputUnionGlobalDesignFilters is an
// implicit subunion of [ProteinSequenceRedesignResumeResponseInputUnion].
// ProteinSequenceRedesignResumeResponseInputUnionGlobalDesignFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignResumeResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters
// OfProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters]
type ProteinSequenceRedesignResumeResponseInputUnionGlobalDesignFilters struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	JSON                                                                                                          struct {
		OfProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters  respjson.Field
		OfProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters respjson.Field
		raw                                                                                                           string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignResumeResponseInputUnionGlobalDesignFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                            `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Binder                                                                                  `json:"type" default:"binder"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                           `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse],
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion struct {
	ChainID string `json:"chain_id"`
	Role    string `json:"role"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse].
	Type constant.FromTemplate `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
	DesignMotifs []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	JSON         struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed target chain from the input CIF.
type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Target       `json:"role" default:"target"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Binder       `json:"role" default:"binder"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                             `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Generic                                                                                  `json:"type" default:"generic"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                            `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity struct {
	ChainID string                `json:"chain_id" api:"required"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignResumeResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseProgress struct {
	// Number of protein designs generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein designs requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated    respjson.Field
		TotalProteinsToGenerate respjson.Field
		LatestResultID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignResumeResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignResumeResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignResumeResponseStatus string

const (
	ProteinSequenceRedesignResumeResponseStatusPending   ProteinSequenceRedesignResumeResponseStatus = "pending"
	ProteinSequenceRedesignResumeResponseStatusRunning   ProteinSequenceRedesignResumeResponseStatus = "running"
	ProteinSequenceRedesignResumeResponseStatusSucceeded ProteinSequenceRedesignResumeResponseStatus = "succeeded"
	ProteinSequenceRedesignResumeResponseStatusFailed    ProteinSequenceRedesignResumeResponseStatus = "failed"
	ProteinSequenceRedesignResumeResponseStatusStopped   ProteinSequenceRedesignResumeResponseStatus = "stopped"
)

// A fixed-structure protein sequence redesign run.
type ProteinSequenceRedesignStartResponse struct {
	// Unique ProteinSequenceRedesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.BoltzProteinRedesign `json:"engine" default:"boltz-protein-redesign"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.V2026_07_14                      `json:"engine_version" default:"v2026-07-14"`
	Error         ProteinSequenceRedesignStartResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinSequenceRedesignStartResponseInputUnion `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode        bool                                         `json:"livemode" api:"required"`
	Pipeline        constant.BoltzProteinRedesign                `json:"pipeline" default:"boltz-protein-redesign"`
	PipelineVersion constant.V2026_07_14                         `json:"pipeline_version" default:"v2026-07-14"`
	Progress        ProteinSequenceRedesignStartResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                    `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinSequenceRedesignStartResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                  `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinSequenceRedesignStartResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseError struct {
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
func (r ProteinSequenceRedesignStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputUnion contains all possible properties
// and values from
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponse],
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStartResponseInputUnion struct {
	// This field is a union of
	// [[]ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion],
	// [[]ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	Entities    ProteinSequenceRedesignStartResponseInputUnionEntities `json:"entities"`
	NumProteins int64                                                  `json:"num_proteins"`
	// This field is a union of
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseStructure],
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseStructure]
	Structure ProteinSequenceRedesignStartResponseInputUnionStructure `json:"structure"`
	Type      string                                                  `json:"type"`
	// This field is a union of
	// [[]ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion],
	// [[]ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	GlobalDesignFilters ProteinSequenceRedesignStartResponseInputUnionGlobalDesignFilters `json:"global_design_filters"`
	IdempotencyKey      string                                                            `json:"idempotency_key"`
	WorkspaceID         string                                                            `json:"workspace_id"`
	JSON                struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStartResponseInputUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStartResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinSequenceRedesignStartResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputUnionEntities is an implicit subunion
// of [ProteinSequenceRedesignStartResponseInputUnion].
// ProteinSequenceRedesignStartResponseInputUnionEntities provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignStartResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntities
// OfProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntities]
type ProteinSequenceRedesignStartResponseInputUnionEntities struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion]
	// instead of an object.
	OfProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	// instead of an object.
	OfProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:",inline"`
	JSON                                                                                              struct {
		OfProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntities  respjson.Field
		OfProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntities respjson.Field
		raw                                                                                               string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignStartResponseInputUnionEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputUnionStructure is an implicit subunion
// of [ProteinSequenceRedesignStartResponseInputUnion].
// ProteinSequenceRedesignStartResponseInputUnionStructure provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignStartResponseInputUnion].
type ProteinSequenceRedesignStartResponseInputUnionStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignStartResponseInputUnionStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputUnionGlobalDesignFilters is an implicit
// subunion of [ProteinSequenceRedesignStartResponseInputUnion].
// ProteinSequenceRedesignStartResponseInputUnionGlobalDesignFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignStartResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters
// OfProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters]
type ProteinSequenceRedesignStartResponseInputUnionGlobalDesignFilters struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	JSON                                                                                                         struct {
		OfProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters  respjson.Field
		OfProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters respjson.Field
		raw                                                                                                          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignStartResponseInputUnionGlobalDesignFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                           `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Binder                                                                                 `json:"type" default:"binder"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                          `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse],
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion struct {
	ChainID string `json:"chain_id"`
	Role    string `json:"role"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse].
	Type constant.FromTemplate `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
	DesignMotifs []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	JSON         struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed target chain from the input CIF.
type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Target       `json:"role" default:"target"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Binder       `json:"role" default:"binder"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                            `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Generic                                                                                 `json:"type" default:"generic"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                           `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity struct {
	ChainID string                `json:"chain_id" api:"required"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStartResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseProgress struct {
	// Number of protein designs generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein designs requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated    respjson.Field
		TotalProteinsToGenerate respjson.Field
		LatestResultID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStartResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignStartResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStartResponseStatus string

const (
	ProteinSequenceRedesignStartResponseStatusPending   ProteinSequenceRedesignStartResponseStatus = "pending"
	ProteinSequenceRedesignStartResponseStatusRunning   ProteinSequenceRedesignStartResponseStatus = "running"
	ProteinSequenceRedesignStartResponseStatusSucceeded ProteinSequenceRedesignStartResponseStatus = "succeeded"
	ProteinSequenceRedesignStartResponseStatusFailed    ProteinSequenceRedesignStartResponseStatus = "failed"
	ProteinSequenceRedesignStartResponseStatusStopped   ProteinSequenceRedesignStartResponseStatus = "stopped"
)

// A fixed-structure protein sequence redesign run.
type ProteinSequenceRedesignStopResponse struct {
	// Unique ProteinSequenceRedesignRun identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input, output, and result data was permanently deleted. Null if data
	// has not been deleted.
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Deprecated. Use pipeline instead.
	//
	// Deprecated: Use pipeline instead.
	Engine constant.BoltzProteinRedesign `json:"engine" default:"boltz-protein-redesign"`
	// Deprecated. Use pipeline_version instead.
	//
	// Deprecated: Use pipeline_version instead.
	EngineVersion constant.V2026_07_14                     `json:"engine_version" default:"v2026-07-14"`
	Error         ProteinSequenceRedesignStopResponseError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ProteinSequenceRedesignStopResponseInputUnion `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode        bool                                        `json:"livemode" api:"required"`
	Pipeline        constant.BoltzProteinRedesign               `json:"pipeline" default:"boltz-protein-redesign"`
	PipelineVersion constant.V2026_07_14                        `json:"pipeline_version" default:"v2026-07-14"`
	Progress        ProteinSequenceRedesignStopResponseProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                   `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ProteinSequenceRedesignStopResponseStatus `json:"status" api:"required"`
	StoppedAt time.Time                                 `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ProteinSequenceRedesignStopResponse) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseError struct {
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
func (r ProteinSequenceRedesignStopResponseError) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignStopResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputUnion contains all possible properties
// and values from
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponse],
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStopResponseInputUnion struct {
	// This field is a union of
	// [[]ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion],
	// [[]ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	Entities    ProteinSequenceRedesignStopResponseInputUnionEntities `json:"entities"`
	NumProteins int64                                                 `json:"num_proteins"`
	// This field is a union of
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseStructure],
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseStructure]
	Structure ProteinSequenceRedesignStopResponseInputUnionStructure `json:"structure"`
	Type      string                                                 `json:"type"`
	// This field is a union of
	// [[]ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion],
	// [[]ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	GlobalDesignFilters ProteinSequenceRedesignStopResponseInputUnionGlobalDesignFilters `json:"global_design_filters"`
	IdempotencyKey      string                                                           `json:"idempotency_key"`
	WorkspaceID         string                                                           `json:"workspace_id"`
	JSON                struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		raw                 string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStopResponseInputUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStopResponseInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ProteinSequenceRedesignStopResponseInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputUnionEntities is an implicit subunion of
// [ProteinSequenceRedesignStopResponseInputUnion].
// ProteinSequenceRedesignStopResponseInputUnionEntities provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignStopResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntities
// OfProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntities]
type ProteinSequenceRedesignStopResponseInputUnionEntities struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion]
	// instead of an object.
	OfProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity]
	// instead of an object.
	OfProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntities []ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:",inline"`
	JSON                                                                                             struct {
		OfProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntities  respjson.Field
		OfProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntities respjson.Field
		raw                                                                                              string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignStopResponseInputUnionEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputUnionStructure is an implicit subunion
// of [ProteinSequenceRedesignStopResponseInputUnion].
// ProteinSequenceRedesignStopResponseInputUnionStructure provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignStopResponseInputUnion].
type ProteinSequenceRedesignStopResponseInputUnionStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignStopResponseInputUnionStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputUnionGlobalDesignFilters is an implicit
// subunion of [ProteinSequenceRedesignStopResponseInputUnion].
// ProteinSequenceRedesignStopResponseInputUnionGlobalDesignFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ProteinSequenceRedesignStopResponseInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters
// OfProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters]
type ProteinSequenceRedesignStopResponseInputUnionGlobalDesignFilters struct {
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	JSON                                                                                                        struct {
		OfProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters  respjson.Field
		OfProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters respjson.Field
		raw                                                                                                         string
	} `json:"-"`
}

func (r *ProteinSequenceRedesignStopResponseInputUnionGlobalDesignFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                          `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Binder                                                                                `json:"type" default:"binder"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                         `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse],
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion struct {
	ChainID string `json:"chain_id"`
	Role    string `json:"role"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse].
	Type constant.FromTemplate `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
	DesignMotifs []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	JSON         struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed target chain from the input CIF.
type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Target       `json:"role" default:"target"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Binder       `json:"role" default:"binder"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                           `json:"num_proteins" api:"required"`
	Structure   ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Generic                                                                                `json:"type" default:"generic"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                          `json:"idempotency_key"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities            respjson.Field
		NumProteins         respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity struct {
	ChainID string                `json:"chain_id" api:"required"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif `json:"design_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID      respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion `json:"filters" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64           `json:"residues" api:"required"`
	Type     constant.Residues `json:"type" default:"residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		Residues    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string                    `json:"amino_acids" api:"required"`
	Type       constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
	MaxFraction float64                         `json:"max_fraction" api:"required"`
	Type        constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFraction respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string                        `json:"motifs" api:"required"`
	Type   constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Motifs      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ProteinSequenceRedesignStopResponseInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseProgress struct {
	// Number of protein designs generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein designs requested
	TotalProteinsToGenerate int64 `json:"total_proteins_to_generate" api:"required"`
	// ID of the most recently generated result
	LatestResultID string `json:"latest_result_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumProteinsGenerated    respjson.Field
		TotalProteinsToGenerate respjson.Field
		LatestResultID          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProteinSequenceRedesignStopResponseProgress) RawJSON() string { return r.JSON.raw }
func (r *ProteinSequenceRedesignStopResponseProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignStopResponseStatus string

const (
	ProteinSequenceRedesignStopResponseStatusPending   ProteinSequenceRedesignStopResponseStatus = "pending"
	ProteinSequenceRedesignStopResponseStatusRunning   ProteinSequenceRedesignStopResponseStatus = "running"
	ProteinSequenceRedesignStopResponseStatusSucceeded ProteinSequenceRedesignStopResponseStatus = "succeeded"
	ProteinSequenceRedesignStopResponseStatusFailed    ProteinSequenceRedesignStopResponseStatus = "failed"
	ProteinSequenceRedesignStopResponseStatusStopped   ProteinSequenceRedesignStopResponseStatus = "stopped"
)

type ProteinSequenceRedesignGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProteinSequenceRedesignGetParams]'s query parameters as
// `url.Values`.
func (r ProteinSequenceRedesignGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinSequenceRedesignListParams struct {
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

// URLQuery serializes [ProteinSequenceRedesignListParams]'s query parameters as
// `url.Values`.
func (r ProteinSequenceRedesignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinSequenceRedesignEstimateCostParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfBinderProteinSequenceRedesignRunInput *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfGenericProteinSequenceRedesignRunInput *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInput `json:",inline"`

	paramObj
}

func (u ProteinSequenceRedesignEstimateCostParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBinderProteinSequenceRedesignRunInput, u.OfGenericProteinSequenceRedesignRunInput)
}
func (r *ProteinSequenceRedesignEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Entities, NumProteins, Structure, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInput struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion `json:"entities,omitzero" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure      ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion `json:"structure,omitzero" api:"required"`
	IdempotencyKey param.Opt[string]                                                                                `json:"idempotency_key,omitzero"`
	// Workspace to run this redesign in.
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion `json:"global_design_filters,omitzero"`
	// This field can be elided, and will marshal its zero value as "binder".
	Type constant.Binder `json:"type" default:"binder"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInput) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// A fixed target chain from the input CIF.
//
// The properties ChainID, Role, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity struct {
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "target".
	Role constant.Target `json:"role" default:"target"`
	// This field can be elided, and will marshal its zero value as "from_template".
	Type constant.FromTemplate `json:"type" default:"from_template"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainID, Role, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity struct {
	ChainID string `json:"chain_id" api:"required"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif `json:"design_motifs,omitzero"`
	// This field can be elided, and will marshal its zero value as "binder".
	Role constant.Binder `json:"role" default:"binder"`
	// This field can be elided, and will marshal its zero value as "from_template".
	Type constant.FromTemplate `json:"type" default:"from_template"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filters, Residues, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion `json:"filters,omitzero" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64 `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "residues".
	Type constant.Residues `json:"type" default:"residues"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputStructureURLSource       *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource       `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputStructureURLSource, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source struct {
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

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Entities, NumProteins, Structure, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInput struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntity `json:"entities,omitzero" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure      ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion `json:"structure,omitzero" api:"required"`
	IdempotencyKey param.Opt[string]                                                                                 `json:"idempotency_key,omitzero"`
	// Workspace to run this redesign in.
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion `json:"global_design_filters,omitzero"`
	// This field can be elided, and will marshal its zero value as "generic".
	Type constant.Generic `json:"type" default:"generic"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInput) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainID, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntity struct {
	ChainID string `json:"chain_id" api:"required"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif `json:"design_motifs,omitzero"`
	// This field can be elided, and will marshal its zero value as "from_template".
	Type constant.FromTemplate `json:"type" default:"from_template"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filters, Residues, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion `json:"filters,omitzero" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64 `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "residues".
	Type constant.Residues `json:"type" default:"residues"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputStructureURLSource       *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource       `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputStructureURLSource, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source struct {
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

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion struct {
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignEstimateCostsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignEstimateCostParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProteinSequenceRedesignListResultsParams struct {
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

// URLQuery serializes [ProteinSequenceRedesignListResultsParams]'s query
// parameters as `url.Values`.
func (r ProteinSequenceRedesignListResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProteinSequenceRedesignStartParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfBinderProteinSequenceRedesignRunInput *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInput `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfGenericProteinSequenceRedesignRunInput *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInput `json:",inline"`

	paramObj
}

func (u ProteinSequenceRedesignStartParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBinderProteinSequenceRedesignRunInput, u.OfGenericProteinSequenceRedesignRunInput)
}
func (r *ProteinSequenceRedesignStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Entities, NumProteins, Structure, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInput struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion `json:"entities,omitzero" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure      ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion `json:"structure,omitzero" api:"required"`
	IdempotencyKey param.Opt[string]                                                                         `json:"idempotency_key,omitzero"`
	// Workspace to run this redesign in.
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion `json:"global_design_filters,omitzero"`
	// This field can be elided, and will marshal its zero value as "binder".
	Type constant.Binder `json:"type" default:"binder"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInput) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion struct {
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity)
}
func (u *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// A fixed target chain from the input CIF.
//
// The properties ChainID, Role, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity struct {
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "target".
	Role constant.Target `json:"role" default:"target"`
	// This field can be elided, and will marshal its zero value as "from_template".
	Type constant.FromTemplate `json:"type" default:"from_template"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignTargetEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainID, Role, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity struct {
	ChainID string `json:"chain_id" api:"required"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif `json:"design_motifs,omitzero"`
	// This field can be elided, and will marshal its zero value as "binder".
	Role constant.Binder `json:"role" default:"binder"`
	// This field can be elided, and will marshal its zero value as "from_template".
	Type constant.FromTemplate `json:"type" default:"from_template"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filters, Residues, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion `json:"filters,omitzero" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64 `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "residues".
	Type constant.Residues `json:"type" default:"residues"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion struct {
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputEntityBinderSequenceRedesignBinderEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion struct {
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputStructureURLSource       *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource       `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputStructureURLSource, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source)
}
func (u *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source struct {
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

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion struct {
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignStartsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyBinderProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Entities, NumProteins, Structure, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInput struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntity `json:"entities,omitzero" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// How to provide a CIF structure file. URLs are auto-detected; base64 uploads must
	// use chemical/x-cif media type.
	Structure      ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion `json:"structure,omitzero" api:"required"`
	IdempotencyKey param.Opt[string]                                                                          `json:"idempotency_key,omitzero"`
	// Workspace to run this redesign in.
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion `json:"global_design_filters,omitzero"`
	// This field can be elided, and will marshal its zero value as "generic".
	Type constant.Generic `json:"type" default:"generic"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInput) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainID, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntity struct {
	ChainID string `json:"chain_id" api:"required"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif `json:"design_motifs,omitzero"`
	// This field can be elided, and will marshal its zero value as "from_template".
	Type constant.FromTemplate `json:"type" default:"from_template"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntity) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filters, Residues, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion `json:"filters,omitzero" api:"required"`
	// 0-indexed residues to redesign on this chain.
	Residues []int64 `json:"residues,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "residues".
	Type constant.Residues `json:"type" default:"residues"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion struct {
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputEntityDesignMotifFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion struct {
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputStructureURLSource       *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource       `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputStructureURLSource, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source)
}
func (u *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source struct {
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

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputStructureCifBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion struct {
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter     `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter `json:",omitzero,inline"`
	OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter `json:",omitzero,inline"`
	paramUnion
}

func (u ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter, u.OfProteinSequenceRedesignStartsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter)
}
func (u *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties AminoAcids, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter struct {
	// Single-letter amino-acid codes that must not occur in the filtered designed
	// region.
	AminoAcids []string `json:"amino_acids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_amino_acids".
	Type constant.ExcludedAminoAcids `json:"type" default:"excluded_amino_acids"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedAminoAcidsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxFraction, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter struct {
	MaxFraction float64 `json:"max_fraction" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "max_hydrophobic_fraction".
	Type constant.MaxHydrophobicFraction `json:"type" default:"max_hydrophobic_fraction"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterMaxHydrophobicFractionDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Motifs, Type are required.
type ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter struct {
	// Sequence motifs that must not occur. X matches any single residue.
	Motifs []string `json:"motifs,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "excluded_sequence_motifs".
	Type constant.ExcludedSequenceMotifs `json:"type" default:"excluded_sequence_motifs"`
	paramObj
}

func (r ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) MarshalJSON() (data []byte, err error) {
	type shadow ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProteinSequenceRedesignStartParamsBodyGenericProteinSequenceRedesignRunInputGlobalDesignFilterExcludedSequenceMotifsDesignFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
