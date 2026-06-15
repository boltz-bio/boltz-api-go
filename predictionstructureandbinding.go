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

// Predict 3D structure coordinates, per-residue confidence scores, and binding
// metrics for a molecular complex. Supports optional template-guided folding and
// per-protein MSA control.
//
// PredictionStructureAndBindingService contains methods and other services that
// help with interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPredictionStructureAndBindingService] method instead.
type PredictionStructureAndBindingService struct {
	Options []option.RequestOption
}

// NewPredictionStructureAndBindingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPredictionStructureAndBindingService(opts ...option.RequestOption) (r PredictionStructureAndBindingService) {
	r = PredictionStructureAndBindingService{}
	r.Options = opts
	return
}

// Retrieve a prediction by ID, including its status and results.
func (r *PredictionStructureAndBindingService) Get(ctx context.Context, id string, query PredictionStructureAndBindingGetParams, opts ...option.RequestOption) (res *PredictionStructureAndBindingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/predictions/structure-and-binding/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List structure and binding predictions, optionally filtered by workspace
func (r *PredictionStructureAndBindingService) List(ctx context.Context, query PredictionStructureAndBindingListParams, opts ...option.RequestOption) (res *pagination.CursorPage[PredictionStructureAndBindingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/predictions/structure-and-binding"
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

// List structure and binding predictions, optionally filtered by workspace
func (r *PredictionStructureAndBindingService) ListAutoPaging(ctx context.Context, query PredictionStructureAndBindingListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[PredictionStructureAndBindingListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// prediction. The prediction record itself is retained with a `data_deleted_at`
// timestamp. This action is irreversible.
func (r *PredictionStructureAndBindingService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *PredictionStructureAndBindingDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/predictions/structure-and-binding/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Estimate the cost of a prediction without creating any resource or consuming
// GPU.
func (r *PredictionStructureAndBindingService) EstimateCost(ctx context.Context, body PredictionStructureAndBindingEstimateCostParams, opts ...option.RequestOption) (res *PredictionStructureAndBindingEstimateCostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/predictions/structure-and-binding/estimate-cost"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submit a prediction job that produces 3D structure coordinates and confidence
// scores for the input molecular complex, with optional binding metrics. Protein
// entities can use automatic MSA generation, custom A3M/CSV MSAs, or empty MSA
// mode. Boltz-2.1 predictions can also include up to 4 CIF or PDB templates to
// guide protein-chain geometry.
func (r *PredictionStructureAndBindingService) Start(ctx context.Context, body PredictionStructureAndBindingStartParams, opts ...option.RequestOption) (res *PredictionStructureAndBindingStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/predictions/structure-and-binding"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the total token count for a prospective Boltz2 complex along with the
// maximum token count the caller's compute config will accept. Pure introspection
// — never rejects when the complex is too large; still rejects on malformed inputs
// the model could not interpret. Use the sandbox prediction endpoint if you need
// the cap enforced.
func (r *PredictionStructureAndBindingService) TokenCount(ctx context.Context, body PredictionStructureAndBindingTokenCountParams, opts ...option.RequestOption) (res *PredictionStructureAndBindingTokenCountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/predictions/structure-and-binding/token-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PredictionStructureAndBindingGetResponse struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error PredictionStructureAndBindingGetResponseError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input PredictionStructureAndBindingGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	// Prediction output when succeeded
	Output    PredictionStructureAndBindingGetResponseOutput `json:"output" api:"required"`
	StartedAt time.Time                                      `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status PredictionStructureAndBindingGetResponseStatus `json:"status" api:"required"`
	// Model version used for prediction
	Version string `json:"version" api:"required"`
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
		Error          respjson.Field
		ExpiresAt      respjson.Field
		Input          respjson.Field
		Livemode       respjson.Field
		Model          respjson.Field
		Output         respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		Version        respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type PredictionStructureAndBindingGetResponseError struct {
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
func (r PredictionStructureAndBindingGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type PredictionStructureAndBindingGetResponseInput struct {
	// Entities (proteins, RNA, DNA, ligands) forming the complex to predict. Order
	// determines chain assignment.
	Entities []PredictionStructureAndBindingGetResponseInputEntityUnion `json:"entities" api:"required"`
	Binding  PredictionStructureAndBindingGetResponseInputBindingUnion  `json:"binding"`
	// Bond constraints between atoms. Atom-level ligand references currently support
	// ligand_ccd only; ligand_smiles is unsupported.
	Bonds []PredictionStructureAndBindingGetResponseInputBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints  []PredictionStructureAndBindingGetResponseInputConstraintUnion `json:"constraints"`
	ModelOptions PredictionStructureAndBindingGetResponseInputModelOptions      `json:"model_options"`
	// Number of structure samples to generate
	NumSamples int64 `json:"num_samples"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []PredictionStructureAndBindingGetResponseInputTemplate `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities     respjson.Field
		Binding      respjson.Field
		Bonds        respjson.Field
		Constraints  respjson.Field
		ModelOptions respjson.Field
		NumSamples   respjson.Field
		Templates    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputEntityUnion contains all possible
// properties and values from
// [PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse],
// [PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse],
// [PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponse],
// [PredictionStructureAndBindingGetResponseInputEntityLigandCcdEntityResponse],
// [PredictionStructureAndBindingGetResponseInputEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification],
	// [[]PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification],
	// [[]PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification]
	Modifications PredictionStructureAndBindingGetResponseInputEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse].
	Msa  PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Msa           respjson.Field
		raw           string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputEntityUnion) AsPredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse() (v PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputEntityUnion) AsPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse() (v PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputEntityUnion) AsPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponse() (v PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputEntityUnion) AsPredictionStructureAndBindingGetResponseInputEntityLigandCcdEntityResponse() (v PredictionStructureAndBindingGetResponseInputEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputEntityUnion) AsPredictionStructureAndBindingGetResponseInputEntityLigandSmilesEntityResponse() (v PredictionStructureAndBindingGetResponseInputEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *PredictionStructureAndBindingGetResponseInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputEntityUnionModifications is an
// implicit subunion of [PredictionStructureAndBindingGetResponseInputEntityUnion].
// PredictionStructureAndBindingGetResponseInputEntityUnionModifications provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PredictionStructureAndBindingGetResponseInputEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfPredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModifications
// OfPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModifications
// OfPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModifications]
type PredictionStructureAndBindingGetResponseInputEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModifications []PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModifications []PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModifications []PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                struct {
		OfPredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModifications respjson.Field
		OfPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModifications           respjson.Field
		OfPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModifications           respjson.Field
		raw                                                                                           string
	} `json:"-"`
}

func (r *PredictionStructureAndBindingGetResponseInputEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification `json:"modifications"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Msa           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion
// contains all possible properties and values from
// [PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse],
// [PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion struct {
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Format PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Source PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source"`
	Type   string                                                                                                         `json:"type"`
	JSON   struct {
		Format respjson.Field
		Source respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) AsPredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse() (v PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) AsPredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse() (v PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format" api:"required"`
	Source PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source" api:"required"`
	Type   constant.Custom                                                                                                `json:"type" default:"custom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		Source      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat string

const (
	PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatA3m PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "a3m"
	PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatCsv PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "csv"
)

type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse struct {
	Type constant.Empty `json:"type" default:"empty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaFormat string

const (
	PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaFormatA3m PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaFormat = "a3m"
	PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaFormatCsv PredictionStructureAndBindingGetResponseInputEntityBoltz2ProteinEntityResponseMsaFormat = "csv"
)

type PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification `json:"modifications"`
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
func (r PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification `json:"modifications"`
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
func (r PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityLigandCcdEntityResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityLigandSmilesEntityResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputBindingUnion contains all possible
// properties and values from
// [PredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse],
// [PredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputBindingUnion struct {
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse].
	BinderChainID string `json:"binder_chain_id"`
	Type          string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse].
	BinderChainIDs []string `json:"binder_chain_ids"`
	JSON           struct {
		BinderChainID  respjson.Field
		Type           respjson.Field
		BinderChainIDs respjson.Field
		raw            string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputBindingUnion) AsPredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse() (v PredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputBindingUnion) AsPredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse() (v PredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputBindingUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, <50 atoms, and only
	// ligands+proteins in entities)
	BinderChainID string                        `json:"binder_chain_id" api:"required"`
	Type          constant.LigandProteinBinding `json:"type" default:"ligand_protein_binding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainID respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputBindingLigandProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse struct {
	// Chain IDs of the protein binders
	BinderChainIDs []string                       `json:"binder_chain_ids" api:"required"`
	Type           constant.ProteinProteinBinding `json:"type" default:"protein_protein_binding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainIDs respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputBindingProteinProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingGetResponseInputBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 PredictionStructureAndBindingGetResponseInputBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 PredictionStructureAndBindingGetResponseInputBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputBond) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputBondAtom1Union contains all
// possible properties and values from
// [PredictionStructureAndBindingGetResponseInputBondAtom1LigandAtomResponse],
// [PredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputBondAtom1Union) AsPredictionStructureAndBindingGetResponseInputBondAtom1LigandAtomResponse() (v PredictionStructureAndBindingGetResponseInputBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputBondAtom1Union) AsPredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse() (v PredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingGetResponseInputBondAtom1LigandAtomResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputBondAtom2Union contains all
// possible properties and values from
// [PredictionStructureAndBindingGetResponseInputBondAtom2LigandAtomResponse],
// [PredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputBondAtom2Union) AsPredictionStructureAndBindingGetResponseInputBondAtom2LigandAtomResponse() (v PredictionStructureAndBindingGetResponseInputBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputBondAtom2Union) AsPredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse() (v PredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingGetResponseInputBondAtom2LigandAtomResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputConstraintUnion contains all
// possible properties and values from
// [PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse],
// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputConstraintUnion struct {
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse].
	Token1 PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse].
	Token2 PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u PredictionStructureAndBindingGetResponseInputConstraintUnion) AsPredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse() (v PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputConstraintUnion) AsPredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse() (v PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union) AsPredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union) AsPredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union) AsPredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union) AsPredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputModelOptions struct {
	// The number of recycling steps to use for prediction. Default is 3.
	RecyclingSteps int64 `json:"recycling_steps"`
	// The number of sampling steps to use for prediction. Default is 200.
	SamplingSteps int64 `json:"sampling_steps"`
	// Diffusion step scale (temperature). Controls sampling diversity — higher values
	// produce more varied structures. Default is 1.638.
	StepScale float64 `json:"step_scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecyclingSteps respjson.Field
		SamplingSteps  respjson.Field
		StepScale      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputModelOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
type PredictionStructureAndBindingGetResponseInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains    []PredictionStructureAndBindingGetResponseInputTemplateTemplateChain   `json:"template_chains" api:"required"`
	TemplateStructure PredictionStructureAndBindingGetResponseInputTemplateTemplateStructure `json:"template_structure" api:"required"`
	// Force the template reference potential with this distance threshold in
	// angstroms. Omit to use the template without force.
	ForceThresholdAngstroms float64 `json:"force_threshold_angstroms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TemplateChains          respjson.Field
		TemplateStructure       respjson.Field
		ForceThresholdAngstroms respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputTemplate) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
type PredictionStructureAndBindingGetResponseInputTemplateTemplateChain struct {
	// Chain ID in this prediction request
	InputChainID string `json:"input_chain_id" api:"required"`
	// Corresponding chain ID in the template structure file
	TemplateChainID string `json:"template_chain_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputChainID    respjson.Field
		TemplateChainID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseInputTemplateTemplateChain) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputTemplateTemplateStructure struct {
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
func (r PredictionStructureAndBindingGetResponseInputTemplateTemplateStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputTemplateTemplateStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type PredictionStructureAndBindingGetResponseOutput struct {
	// Per-sample structure results
	AllSampleResults []PredictionStructureAndBindingGetResponseOutputAllSampleResult   `json:"all_sample_results" api:"required"`
	BestSample       PredictionStructureAndBindingGetResponseOutputBestSample          `json:"best_sample" api:"required"`
	Archive          PredictionStructureAndBindingGetResponseOutputArchive             `json:"archive"`
	BindingMetrics   PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion `json:"binding_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllSampleResults respjson.Field
		BestSample       respjson.Field
		Archive          respjson.Field
		BindingMetrics   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputAllSampleResult struct {
	Metrics         PredictionStructureAndBindingGetResponseOutputAllSampleResultMetrics         `json:"metrics" api:"required"`
	Structure       PredictionStructureAndBindingGetResponseOutputAllSampleResultStructure       `json:"structure" api:"required"`
	LigandStructure PredictionStructureAndBindingGetResponseOutputAllSampleResultLigandStructure `json:"ligand_structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics         respjson.Field
		Structure       respjson.Field
		LigandStructure respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutputAllSampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputAllSampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputAllSampleResultMetrics struct {
	// Complex interface predicted distance error. Lower is better.
	ComplexIpde float64 `json:"complex_ipde" api:"required"`
	// Complex interface pLDDT (0-1 float). Confidence at inter-chain interfaces.
	ComplexIplddt float64 `json:"complex_iplddt" api:"required"`
	// Complex predicted distance error. Lower is better.
	ComplexPde float64 `json:"complex_pde" api:"required"`
	// Complex pLDDT (0-1 float). Per-residue confidence averaged over the complex.
	ComplexPlddt float64 `json:"complex_plddt" api:"required"`
	// Interface predicted TM score (0-1). Confidence in domain interfaces.
	Iptm float64 `json:"iptm" api:"required"`
	// Ligand interface pTM (0-1). Only present when ligands are included.
	LigandIptm float64 `json:"ligand_iptm" api:"required"`
	// Protein-protein interface pTM (0-1). Only present for multi-protein complexes.
	ProteinIptm float64 `json:"protein_iptm" api:"required"`
	// Predicted TM score (0-1). Global structure quality.
	Ptm float64 `json:"ptm" api:"required"`
	// Overall structure confidence (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComplexIpde         respjson.Field
		ComplexIplddt       respjson.Field
		ComplexPde          respjson.Field
		ComplexPlddt        respjson.Field
		Iptm                respjson.Field
		LigandIptm          respjson.Field
		ProteinIptm         respjson.Field
		Ptm                 respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutputAllSampleResultMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputAllSampleResultMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputAllSampleResultStructure struct {
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
func (r PredictionStructureAndBindingGetResponseOutputAllSampleResultStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputAllSampleResultStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputAllSampleResultLigandStructure struct {
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
func (r PredictionStructureAndBindingGetResponseOutputAllSampleResultLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputAllSampleResultLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputBestSample struct {
	Metrics         PredictionStructureAndBindingGetResponseOutputBestSampleMetrics         `json:"metrics" api:"required"`
	Structure       PredictionStructureAndBindingGetResponseOutputBestSampleStructure       `json:"structure" api:"required"`
	LigandStructure PredictionStructureAndBindingGetResponseOutputBestSampleLigandStructure `json:"ligand_structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics         respjson.Field
		Structure       respjson.Field
		LigandStructure respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutputBestSample) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseOutputBestSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputBestSampleMetrics struct {
	// Complex interface predicted distance error. Lower is better.
	ComplexIpde float64 `json:"complex_ipde" api:"required"`
	// Complex interface pLDDT (0-1 float). Confidence at inter-chain interfaces.
	ComplexIplddt float64 `json:"complex_iplddt" api:"required"`
	// Complex predicted distance error. Lower is better.
	ComplexPde float64 `json:"complex_pde" api:"required"`
	// Complex pLDDT (0-1 float). Per-residue confidence averaged over the complex.
	ComplexPlddt float64 `json:"complex_plddt" api:"required"`
	// Interface predicted TM score (0-1). Confidence in domain interfaces.
	Iptm float64 `json:"iptm" api:"required"`
	// Ligand interface pTM (0-1). Only present when ligands are included.
	LigandIptm float64 `json:"ligand_iptm" api:"required"`
	// Protein-protein interface pTM (0-1). Only present for multi-protein complexes.
	ProteinIptm float64 `json:"protein_iptm" api:"required"`
	// Predicted TM score (0-1). Global structure quality.
	Ptm float64 `json:"ptm" api:"required"`
	// Overall structure confidence (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComplexIpde         respjson.Field
		ComplexIplddt       respjson.Field
		ComplexPde          respjson.Field
		ComplexPlddt        respjson.Field
		Iptm                respjson.Field
		LigandIptm          respjson.Field
		ProteinIptm         respjson.Field
		Ptm                 respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutputBestSampleMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputBestSampleMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputBestSampleStructure struct {
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
func (r PredictionStructureAndBindingGetResponseOutputBestSampleStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputBestSampleStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputBestSampleLigandStructure struct {
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
func (r PredictionStructureAndBindingGetResponseOutputBestSampleLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputBestSampleLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputArchive struct {
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
func (r PredictionStructureAndBindingGetResponseOutputArchive) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingGetResponseOutputArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion contains all
// possible properties and values from
// [PredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics],
// [PredictionStructureAndBindingGetResponseOutputBindingMetricsProteinProteinBindingMetrics].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion struct {
	BindingConfidence float64 `json:"binding_confidence"`
	// This field is from variant
	// [PredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics].
	OptimizationScore float64 `json:"optimization_score"`
	Type              string  `json:"type"`
	JSON              struct {
		BindingConfidence respjson.Field
		OptimizationScore respjson.Field
		Type              respjson.Field
		raw               string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion) AsPredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics() (v PredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion) AsPredictionStructureAndBindingGetResponseOutputBindingMetricsProteinProteinBindingMetrics() (v PredictionStructureAndBindingGetResponseOutputBindingMetricsProteinProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingGetResponseOutputBindingMetricsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics struct {
	// Confidence that binding occurs (0-1). Primary metric for hit discovery.
	BindingConfidence float64 `json:"binding_confidence" api:"required"`
	// Binding strength ranking score for lead optimization. Higher values indicate
	// stronger predicted binding.
	OptimizationScore float64                              `json:"optimization_score" api:"required"`
	Type              constant.LigandProteinBindingMetrics `json:"type" default:"ligand_protein_binding_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence respjson.Field
		OptimizationScore respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputBindingMetricsLigandProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseOutputBindingMetricsProteinProteinBindingMetrics struct {
	// Confidence that binding occurs (0-1). Primary metric for hit discovery.
	BindingConfidence float64                               `json:"binding_confidence" api:"required"`
	Type              constant.ProteinProteinBindingMetrics `json:"type" default:"protein_protein_binding_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingGetResponseOutputBindingMetricsProteinProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseOutputBindingMetricsProteinProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseStatus string

const (
	PredictionStructureAndBindingGetResponseStatusPending   PredictionStructureAndBindingGetResponseStatus = "pending"
	PredictionStructureAndBindingGetResponseStatusRunning   PredictionStructureAndBindingGetResponseStatus = "running"
	PredictionStructureAndBindingGetResponseStatusSucceeded PredictionStructureAndBindingGetResponseStatus = "succeeded"
	PredictionStructureAndBindingGetResponseStatusFailed    PredictionStructureAndBindingGetResponseStatus = "failed"
)

type PredictionStructureAndBindingListResponse struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error PredictionStructureAndBindingListResponseError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model     constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	StartedAt time.Time         `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status PredictionStructureAndBindingListResponseStatus `json:"status" api:"required"`
	// Model version used for prediction
	Version string `json:"version" api:"required"`
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
		Error          respjson.Field
		ExpiresAt      respjson.Field
		Livemode       respjson.Field
		Model          respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		Version        respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingListResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type PredictionStructureAndBindingListResponseError struct {
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
func (r PredictionStructureAndBindingListResponseError) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingListResponseStatus string

const (
	PredictionStructureAndBindingListResponseStatusPending   PredictionStructureAndBindingListResponseStatus = "pending"
	PredictionStructureAndBindingListResponseStatusRunning   PredictionStructureAndBindingListResponseStatus = "running"
	PredictionStructureAndBindingListResponseStatusSucceeded PredictionStructureAndBindingListResponseStatus = "succeeded"
	PredictionStructureAndBindingListResponseStatusFailed    PredictionStructureAndBindingListResponseStatus = "failed"
)

type PredictionStructureAndBindingDeleteDataResponse struct {
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
func (r PredictionStructureAndBindingDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimate response with monetary values encoded as decimal strings to preserve
// precision.
type PredictionStructureAndBindingEstimateCostResponse struct {
	// Cost breakdown for the billed application.
	Breakdown  PredictionStructureAndBindingEstimateCostResponseBreakdown `json:"breakdown" api:"required"`
	Disclaimer string                                                     `json:"disclaimer" api:"required"`
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
func (r PredictionStructureAndBindingEstimateCostResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingEstimateCostResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost breakdown for the billed application.
type PredictionStructureAndBindingEstimateCostResponseBreakdown struct {
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application PredictionStructureAndBindingEstimateCostResponseBreakdownApplication `json:"application" api:"required"`
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
func (r PredictionStructureAndBindingEstimateCostResponseBreakdown) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingEstimateCostResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingEstimateCostResponseBreakdownApplication string

const (
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationStructureAndBinding        PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "structure_and_binding"
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationSmallMoleculeDesign        PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "small_molecule_design"
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationSmallMoleculeLibraryScreen PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "small_molecule_library_screen"
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationProteinDesign              PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "protein_design"
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationProteinRedesign            PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "protein_redesign"
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationProteinLibraryScreen       PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "protein_library_screen"
	PredictionStructureAndBindingEstimateCostResponseBreakdownApplicationAdme                       PredictionStructureAndBindingEstimateCostResponseBreakdownApplication = "adme"
)

type PredictionStructureAndBindingStartResponse struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error PredictionStructureAndBindingStartResponseError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input PredictionStructureAndBindingStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	// Prediction output when succeeded
	Output    PredictionStructureAndBindingStartResponseOutput `json:"output" api:"required"`
	StartedAt time.Time                                        `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status PredictionStructureAndBindingStartResponseStatus `json:"status" api:"required"`
	// Model version used for prediction
	Version string `json:"version" api:"required"`
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
		Error          respjson.Field
		ExpiresAt      respjson.Field
		Input          respjson.Field
		Livemode       respjson.Field
		Model          respjson.Field
		Output         respjson.Field
		StartedAt      respjson.Field
		Status         respjson.Field
		Version        respjson.Field
		WorkspaceID    respjson.Field
		IdempotencyKey respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type PredictionStructureAndBindingStartResponseError struct {
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
func (r PredictionStructureAndBindingStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type PredictionStructureAndBindingStartResponseInput struct {
	// Entities (proteins, RNA, DNA, ligands) forming the complex to predict. Order
	// determines chain assignment.
	Entities []PredictionStructureAndBindingStartResponseInputEntityUnion `json:"entities" api:"required"`
	Binding  PredictionStructureAndBindingStartResponseInputBindingUnion  `json:"binding"`
	// Bond constraints between atoms. Atom-level ligand references currently support
	// ligand_ccd only; ligand_smiles is unsupported.
	Bonds []PredictionStructureAndBindingStartResponseInputBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints  []PredictionStructureAndBindingStartResponseInputConstraintUnion `json:"constraints"`
	ModelOptions PredictionStructureAndBindingStartResponseInputModelOptions      `json:"model_options"`
	// Number of structure samples to generate
	NumSamples int64 `json:"num_samples"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []PredictionStructureAndBindingStartResponseInputTemplate `json:"templates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities     respjson.Field
		Binding      respjson.Field
		Bonds        respjson.Field
		Constraints  respjson.Field
		ModelOptions respjson.Field
		NumSamples   respjson.Field
		Templates    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputEntityUnion contains all possible
// properties and values from
// [PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse],
// [PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse],
// [PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponse],
// [PredictionStructureAndBindingStartResponseInputEntityLigandCcdEntityResponse],
// [PredictionStructureAndBindingStartResponseInputEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification],
	// [[]PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification],
	// [[]PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification]
	Modifications PredictionStructureAndBindingStartResponseInputEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse].
	Msa  PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Msa           respjson.Field
		raw           string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputEntityUnion) AsPredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse() (v PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputEntityUnion) AsPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse() (v PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputEntityUnion) AsPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponse() (v PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputEntityUnion) AsPredictionStructureAndBindingStartResponseInputEntityLigandCcdEntityResponse() (v PredictionStructureAndBindingStartResponseInputEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputEntityUnion) AsPredictionStructureAndBindingStartResponseInputEntityLigandSmilesEntityResponse() (v PredictionStructureAndBindingStartResponseInputEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputEntityUnionModifications is an
// implicit subunion of
// [PredictionStructureAndBindingStartResponseInputEntityUnion].
// PredictionStructureAndBindingStartResponseInputEntityUnionModifications provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PredictionStructureAndBindingStartResponseInputEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfPredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModifications
// OfPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModifications
// OfPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModifications]
type PredictionStructureAndBindingStartResponseInputEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModifications []PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModifications []PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModifications []PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                  struct {
		OfPredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModifications respjson.Field
		OfPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModifications           respjson.Field
		OfPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModifications           respjson.Field
		raw                                                                                             string
	} `json:"-"`
}

func (r *PredictionStructureAndBindingStartResponseInputEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification `json:"modifications"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Msa           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion
// contains all possible properties and values from
// [PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse],
// [PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion struct {
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Format PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Source PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source"`
	Type   string                                                                                                           `json:"type"`
	JSON   struct {
		Format respjson.Field
		Source respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) AsPredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse() (v PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) AsPredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse() (v PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format" api:"required"`
	Source PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source" api:"required"`
	Type   constant.Custom                                                                                                  `json:"type" default:"custom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Format      respjson.Field
		Source      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat string

const (
	PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatA3m PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "a3m"
	PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatCsv PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "csv"
)

type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse struct {
	Type constant.Empty `json:"type" default:"empty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaFormat string

const (
	PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaFormatA3m PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaFormat = "a3m"
	PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaFormatCsv PredictionStructureAndBindingStartResponseInputEntityBoltz2ProteinEntityResponseMsaFormat = "csv"
)

type PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification `json:"modifications"`
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
func (r PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification `json:"modifications"`
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
func (r PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityLigandCcdEntityResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityLigandSmilesEntityResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputBindingUnion contains all
// possible properties and values from
// [PredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse],
// [PredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputBindingUnion struct {
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse].
	BinderChainID string `json:"binder_chain_id"`
	Type          string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse].
	BinderChainIDs []string `json:"binder_chain_ids"`
	JSON           struct {
		BinderChainID  respjson.Field
		Type           respjson.Field
		BinderChainIDs respjson.Field
		raw            string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputBindingUnion) AsPredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse() (v PredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputBindingUnion) AsPredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse() (v PredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputBindingUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, <50 atoms, and only
	// ligands+proteins in entities)
	BinderChainID string                        `json:"binder_chain_id" api:"required"`
	Type          constant.LigandProteinBinding `json:"type" default:"ligand_protein_binding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainID respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputBindingLigandProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse struct {
	// Chain IDs of the protein binders
	BinderChainIDs []string                       `json:"binder_chain_ids" api:"required"`
	Type           constant.ProteinProteinBinding `json:"type" default:"protein_protein_binding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChainIDs respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputBindingProteinProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingStartResponseInputBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 PredictionStructureAndBindingStartResponseInputBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 PredictionStructureAndBindingStartResponseInputBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputBond) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponseInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputBondAtom1Union contains all
// possible properties and values from
// [PredictionStructureAndBindingStartResponseInputBondAtom1LigandAtomResponse],
// [PredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputBondAtom1Union) AsPredictionStructureAndBindingStartResponseInputBondAtom1LigandAtomResponse() (v PredictionStructureAndBindingStartResponseInputBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputBondAtom1Union) AsPredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse() (v PredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingStartResponseInputBondAtom1LigandAtomResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputBondAtom2Union contains all
// possible properties and values from
// [PredictionStructureAndBindingStartResponseInputBondAtom2LigandAtomResponse],
// [PredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputBondAtom2Union) AsPredictionStructureAndBindingStartResponseInputBondAtom2LigandAtomResponse() (v PredictionStructureAndBindingStartResponseInputBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputBondAtom2Union) AsPredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse() (v PredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingStartResponseInputBondAtom2LigandAtomResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputConstraintUnion contains all
// possible properties and values from
// [PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse],
// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputConstraintUnion struct {
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse].
	Token1 PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse].
	Token2 PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u PredictionStructureAndBindingStartResponseInputConstraintUnion) AsPredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse() (v PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputConstraintUnion) AsPredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse() (v PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                              `json:"type" default:"contact"`
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
func (r PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union) AsPredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union) AsPredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union) AsPredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union) AsPredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputModelOptions struct {
	// The number of recycling steps to use for prediction. Default is 3.
	RecyclingSteps int64 `json:"recycling_steps"`
	// The number of sampling steps to use for prediction. Default is 200.
	SamplingSteps int64 `json:"sampling_steps"`
	// Diffusion step scale (temperature). Controls sampling diversity — higher values
	// produce more varied structures. Default is 1.638.
	StepScale float64 `json:"step_scale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecyclingSteps respjson.Field
		SamplingSteps  respjson.Field
		StepScale      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputModelOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
type PredictionStructureAndBindingStartResponseInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains    []PredictionStructureAndBindingStartResponseInputTemplateTemplateChain   `json:"template_chains" api:"required"`
	TemplateStructure PredictionStructureAndBindingStartResponseInputTemplateTemplateStructure `json:"template_structure" api:"required"`
	// Force the template reference potential with this distance threshold in
	// angstroms. Omit to use the template without force.
	ForceThresholdAngstroms float64 `json:"force_threshold_angstroms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TemplateChains          respjson.Field
		TemplateStructure       respjson.Field
		ForceThresholdAngstroms respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputTemplate) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponseInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
type PredictionStructureAndBindingStartResponseInputTemplateTemplateChain struct {
	// Chain ID in this prediction request
	InputChainID string `json:"input_chain_id" api:"required"`
	// Corresponding chain ID in the template structure file
	TemplateChainID string `json:"template_chain_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputChainID    respjson.Field
		TemplateChainID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseInputTemplateTemplateChain) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputTemplateTemplateStructure struct {
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
func (r PredictionStructureAndBindingStartResponseInputTemplateTemplateStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputTemplateTemplateStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type PredictionStructureAndBindingStartResponseOutput struct {
	// Per-sample structure results
	AllSampleResults []PredictionStructureAndBindingStartResponseOutputAllSampleResult   `json:"all_sample_results" api:"required"`
	BestSample       PredictionStructureAndBindingStartResponseOutputBestSample          `json:"best_sample" api:"required"`
	Archive          PredictionStructureAndBindingStartResponseOutputArchive             `json:"archive"`
	BindingMetrics   PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion `json:"binding_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllSampleResults respjson.Field
		BestSample       respjson.Field
		Archive          respjson.Field
		BindingMetrics   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputAllSampleResult struct {
	Metrics         PredictionStructureAndBindingStartResponseOutputAllSampleResultMetrics         `json:"metrics" api:"required"`
	Structure       PredictionStructureAndBindingStartResponseOutputAllSampleResultStructure       `json:"structure" api:"required"`
	LigandStructure PredictionStructureAndBindingStartResponseOutputAllSampleResultLigandStructure `json:"ligand_structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics         respjson.Field
		Structure       respjson.Field
		LigandStructure respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutputAllSampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputAllSampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputAllSampleResultMetrics struct {
	// Complex interface predicted distance error. Lower is better.
	ComplexIpde float64 `json:"complex_ipde" api:"required"`
	// Complex interface pLDDT (0-1 float). Confidence at inter-chain interfaces.
	ComplexIplddt float64 `json:"complex_iplddt" api:"required"`
	// Complex predicted distance error. Lower is better.
	ComplexPde float64 `json:"complex_pde" api:"required"`
	// Complex pLDDT (0-1 float). Per-residue confidence averaged over the complex.
	ComplexPlddt float64 `json:"complex_plddt" api:"required"`
	// Interface predicted TM score (0-1). Confidence in domain interfaces.
	Iptm float64 `json:"iptm" api:"required"`
	// Ligand interface pTM (0-1). Only present when ligands are included.
	LigandIptm float64 `json:"ligand_iptm" api:"required"`
	// Protein-protein interface pTM (0-1). Only present for multi-protein complexes.
	ProteinIptm float64 `json:"protein_iptm" api:"required"`
	// Predicted TM score (0-1). Global structure quality.
	Ptm float64 `json:"ptm" api:"required"`
	// Overall structure confidence (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComplexIpde         respjson.Field
		ComplexIplddt       respjson.Field
		ComplexPde          respjson.Field
		ComplexPlddt        respjson.Field
		Iptm                respjson.Field
		LigandIptm          respjson.Field
		ProteinIptm         respjson.Field
		Ptm                 respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutputAllSampleResultMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputAllSampleResultMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputAllSampleResultStructure struct {
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
func (r PredictionStructureAndBindingStartResponseOutputAllSampleResultStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputAllSampleResultStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputAllSampleResultLigandStructure struct {
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
func (r PredictionStructureAndBindingStartResponseOutputAllSampleResultLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputAllSampleResultLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputBestSample struct {
	Metrics         PredictionStructureAndBindingStartResponseOutputBestSampleMetrics         `json:"metrics" api:"required"`
	Structure       PredictionStructureAndBindingStartResponseOutputBestSampleStructure       `json:"structure" api:"required"`
	LigandStructure PredictionStructureAndBindingStartResponseOutputBestSampleLigandStructure `json:"ligand_structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics         respjson.Field
		Structure       respjson.Field
		LigandStructure respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutputBestSample) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputBestSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputBestSampleMetrics struct {
	// Complex interface predicted distance error. Lower is better.
	ComplexIpde float64 `json:"complex_ipde" api:"required"`
	// Complex interface pLDDT (0-1 float). Confidence at inter-chain interfaces.
	ComplexIplddt float64 `json:"complex_iplddt" api:"required"`
	// Complex predicted distance error. Lower is better.
	ComplexPde float64 `json:"complex_pde" api:"required"`
	// Complex pLDDT (0-1 float). Per-residue confidence averaged over the complex.
	ComplexPlddt float64 `json:"complex_plddt" api:"required"`
	// Interface predicted TM score (0-1). Confidence in domain interfaces.
	Iptm float64 `json:"iptm" api:"required"`
	// Ligand interface pTM (0-1). Only present when ligands are included.
	LigandIptm float64 `json:"ligand_iptm" api:"required"`
	// Protein-protein interface pTM (0-1). Only present for multi-protein complexes.
	ProteinIptm float64 `json:"protein_iptm" api:"required"`
	// Predicted TM score (0-1). Global structure quality.
	Ptm float64 `json:"ptm" api:"required"`
	// Overall structure confidence (0-1).
	StructureConfidence float64 `json:"structure_confidence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComplexIpde         respjson.Field
		ComplexIplddt       respjson.Field
		ComplexPde          respjson.Field
		ComplexPlddt        respjson.Field
		Iptm                respjson.Field
		LigandIptm          respjson.Field
		ProteinIptm         respjson.Field
		Ptm                 respjson.Field
		StructureConfidence respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutputBestSampleMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputBestSampleMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputBestSampleStructure struct {
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
func (r PredictionStructureAndBindingStartResponseOutputBestSampleStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputBestSampleStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputBestSampleLigandStructure struct {
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
func (r PredictionStructureAndBindingStartResponseOutputBestSampleLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputBestSampleLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputArchive struct {
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
func (r PredictionStructureAndBindingStartResponseOutputArchive) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingStartResponseOutputArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion contains all
// possible properties and values from
// [PredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics],
// [PredictionStructureAndBindingStartResponseOutputBindingMetricsProteinProteinBindingMetrics].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion struct {
	BindingConfidence float64 `json:"binding_confidence"`
	// This field is from variant
	// [PredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics].
	OptimizationScore float64 `json:"optimization_score"`
	Type              string  `json:"type"`
	JSON              struct {
		BindingConfidence respjson.Field
		OptimizationScore respjson.Field
		Type              respjson.Field
		raw               string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion) AsPredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics() (v PredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion) AsPredictionStructureAndBindingStartResponseOutputBindingMetricsProteinProteinBindingMetrics() (v PredictionStructureAndBindingStartResponseOutputBindingMetricsProteinProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *PredictionStructureAndBindingStartResponseOutputBindingMetricsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics struct {
	// Confidence that binding occurs (0-1). Primary metric for hit discovery.
	BindingConfidence float64 `json:"binding_confidence" api:"required"`
	// Binding strength ranking score for lead optimization. Higher values indicate
	// stronger predicted binding.
	OptimizationScore float64                              `json:"optimization_score" api:"required"`
	Type              constant.LigandProteinBindingMetrics `json:"type" default:"ligand_protein_binding_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence respjson.Field
		OptimizationScore respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputBindingMetricsLigandProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseOutputBindingMetricsProteinProteinBindingMetrics struct {
	// Confidence that binding occurs (0-1). Primary metric for hit discovery.
	BindingConfidence float64                               `json:"binding_confidence" api:"required"`
	Type              constant.ProteinProteinBindingMetrics `json:"type" default:"protein_protein_binding_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BindingConfidence respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingStartResponseOutputBindingMetricsProteinProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseOutputBindingMetricsProteinProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseStatus string

const (
	PredictionStructureAndBindingStartResponseStatusPending   PredictionStructureAndBindingStartResponseStatus = "pending"
	PredictionStructureAndBindingStartResponseStatusRunning   PredictionStructureAndBindingStartResponseStatus = "running"
	PredictionStructureAndBindingStartResponseStatusSucceeded PredictionStructureAndBindingStartResponseStatus = "succeeded"
	PredictionStructureAndBindingStartResponseStatusFailed    PredictionStructureAndBindingStartResponseStatus = "failed"
)

type PredictionStructureAndBindingTokenCountResponse struct {
	// Maximum token count the caller's compute config will accept for a sandbox
	// prediction.
	MaxTokenCount int64 `json:"max_token_count" api:"required"`
	// Total token count for the complex, after applying the affinity model's internal
	// crop when relevant.
	TotalTokenCount int64 `json:"total_token_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxTokenCount   respjson.Field
		TotalTokenCount respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionStructureAndBindingTokenCountResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionStructureAndBindingTokenCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PredictionStructureAndBindingGetParams]'s query parameters
// as `url.Values`.
func (r PredictionStructureAndBindingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PredictionStructureAndBindingListParams struct {
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

// URLQuery serializes [PredictionStructureAndBindingListParams]'s query parameters
// as `url.Values`.
func (r PredictionStructureAndBindingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PredictionStructureAndBindingEstimateCostParams struct {
	Input PredictionStructureAndBindingEstimateCostParamsInput `json:"input,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Model to use for prediction
	//
	// This field can be elided, and will marshal its zero value as "boltz-2.1".
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParams) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Entities is required.
type PredictionStructureAndBindingEstimateCostParamsInput struct {
	// Entities (proteins, RNA, DNA, ligands) forming the complex to predict. Order
	// determines chain assignment.
	Entities []PredictionStructureAndBindingEstimateCostParamsInputEntityUnion `json:"entities,omitzero" api:"required"`
	// Number of structure samples to generate
	NumSamples param.Opt[int64]                                                 `json:"num_samples,omitzero"`
	Binding    PredictionStructureAndBindingEstimateCostParamsInputBindingUnion `json:"binding,omitzero"`
	// Bond constraints between atoms. Atom-level ligand references currently support
	// ligand_ccd only; ligand_smiles is unsupported.
	Bonds []PredictionStructureAndBindingEstimateCostParamsInputBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints  []PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion `json:"constraints,omitzero"`
	ModelOptions PredictionStructureAndBindingEstimateCostParamsInputModelOptions      `json:"model_options,omitzero"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []PredictionStructureAndBindingEstimateCostParamsInputTemplate `json:"templates,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputEntityUnion struct {
	OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntity *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntity `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityRnaEntity           *PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntity           `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityDnaEntity           *PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntity           `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityLigandCcdEntity     *PredictionStructureAndBindingEstimateCostParamsInputEntityLigandCcdEntity     `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityLigandSmilesEntity  *PredictionStructureAndBindingEstimateCostParamsInputEntityLigandSmilesEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityRnaEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityDnaEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityLigandCcdEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityLigandSmilesEntity)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityModification `json:"modifications,omitzero"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaUnion `json:"msa,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityModification struct {
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

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaUnion struct {
	OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa  *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa, u.OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
//
// The properties Format, Source, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat `json:"format,omitzero" api:"required"`
	// How to provide a file to the API
	Source PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion `json:"source,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "custom".
	Type constant.Custom `json:"type" default:"custom"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat string

const (
	PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatA3m PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat = "a3m"
	PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatCsv PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat = "csv"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion struct {
	OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource    *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource    `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource, u.OfPredictionStructureAndBindingEstimateCostsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source struct {
	// Base64-encoded file contents
	Data string `json:"data" api:"required"`
	// MIME type (e.g., text/csv)
	MediaType string `json:"media_type" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewPredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa() PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa {
	return PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa{
		Type: "empty",
	}
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
//
// This struct has a constant value, construct it with
// [NewPredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa].
type PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa struct {
	Type constant.Empty `json:"type" default:"empty"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntityModification struct {
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

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntityModification struct {
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

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputBindingUnion struct {
	OfPredictionStructureAndBindingEstimateCostsInputBindingLigandProteinBinding  *PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputBindingProteinProteinBinding *PredictionStructureAndBindingEstimateCostParamsInputBindingProteinProteinBinding `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputBindingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputBindingLigandProteinBinding, u.OfPredictionStructureAndBindingEstimateCostsInputBindingProteinProteinBinding)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties BinderChainID, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, <50 atoms, and only
	// ligands+proteins in entities)
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ligand_protein_binding".
	Type constant.LigandProteinBinding `json:"type" default:"ligand_protein_binding"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BinderChainIDs, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputBindingProteinProteinBinding struct {
	// Chain IDs of the protein binders
	BinderChainIDs []string `json:"binder_chain_ids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "protein_protein_binding".
	Type constant.ProteinProteinBinding `json:"type" default:"protein_protein_binding"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputBindingProteinProteinBinding) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBindingProteinProteinBinding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBindingProteinProteinBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type PredictionStructureAndBindingEstimateCostParamsInputBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputBond) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union struct {
	OfPredictionStructureAndBindingEstimateCostsInputBondAtom1LigandAtom  *PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputBondAtom1PolymerAtom *PredictionStructureAndBindingEstimateCostParamsInputBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputBondAtom1LigandAtom, u.OfPredictionStructureAndBindingEstimateCostsInputBondAtom1PolymerAtom)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputBondAtom1PolymerAtom struct {
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

func (r PredictionStructureAndBindingEstimateCostParamsInputBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union struct {
	OfPredictionStructureAndBindingEstimateCostsInputBondAtom2LigandAtom  *PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputBondAtom2PolymerAtom *PredictionStructureAndBindingEstimateCostParamsInputBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputBondAtom2LigandAtom, u.OfPredictionStructureAndBindingEstimateCostsInputBondAtom2PolymerAtom)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputBondAtom2PolymerAtom struct {
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

func (r PredictionStructureAndBindingEstimateCostParamsInputBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion struct {
	OfPredictionStructureAndBindingEstimateCostsInputConstraintPocketConstraint  *PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraint *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputConstraintPocketConstraint, u.OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraint)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint struct {
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

func (r PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1Union struct {
	OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken1PolymerContactToken *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken1LigandContactToken  *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken1PolymerContactToken, u.OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken1LigandContactToken)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2Union struct {
	OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken2PolymerContactToken *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken2LigandContactToken  *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken2PolymerContactToken, u.OfPredictionStructureAndBindingEstimateCostsInputConstraintContactConstraintToken2LigandContactToken)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingEstimateCostParamsInputModelOptions struct {
	// The number of recycling steps to use for prediction. Default is 3.
	RecyclingSteps param.Opt[int64] `json:"recycling_steps,omitzero"`
	// The number of sampling steps to use for prediction. Default is 200.
	SamplingSteps param.Opt[int64] `json:"sampling_steps,omitzero"`
	// Diffusion step scale (temperature). Controls sampling diversity — higher values
	// produce more varied structures. Default is 1.638.
	StepScale param.Opt[float64] `json:"step_scale,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputModelOptions) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputModelOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
//
// The properties TemplateChains, TemplateStructure are required.
type PredictionStructureAndBindingEstimateCostParamsInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains []PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateChain `json:"template_chains,omitzero" api:"required"`
	// How to provide a template structure file. URLs must point to a CIF or PDB file;
	// base64 uploads must use chemical/x-cif or chemical/x-pdb.
	TemplateStructure PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureUnion `json:"template_structure,omitzero" api:"required"`
	// Force the template reference potential with this distance threshold in
	// angstroms. Omit to use the template without force.
	ForceThresholdAngstroms param.Opt[float64] `json:"force_threshold_angstroms,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputTemplate) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
//
// The properties InputChainID, TemplateChainID are required.
type PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateChain struct {
	// Chain ID in this prediction request
	InputChainID string `json:"input_chain_id" api:"required"`
	// Corresponding chain ID in the template structure file
	TemplateChainID string `json:"template_chain_id" api:"required"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateChain) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureUnion struct {
	OfPredictionStructureAndBindingEstimateCostsInputTemplateTemplateStructureURLSource                     *PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureURLSource                     `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputTemplateTemplateStructureTemplateStructureBase64Source *PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputTemplateTemplateStructureURLSource, u.OfPredictionStructureAndBindingEstimateCostsInputTemplateTemplateStructureTemplateStructureBase64Source)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64Source struct {
	// Base64-encoded template structure file contents
	Data string `json:"data" api:"required"`
	// Template structure MIME type
	//
	// Any of "chemical/x-cif", "chemical/x-pdb".
	MediaType PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType `json:"media_type,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure MIME type
type PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType string

const (
	PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaTypeChemicalXCif PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType = "chemical/x-cif"
	PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaTypeChemicalXPdb PredictionStructureAndBindingEstimateCostParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType = "chemical/x-pdb"
)

type PredictionStructureAndBindingStartParams struct {
	Input PredictionStructureAndBindingStartParamsInput `json:"input,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Model to use for prediction
	//
	// This field can be elided, and will marshal its zero value as "boltz-2.1".
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	paramObj
}

func (r PredictionStructureAndBindingStartParams) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Entities is required.
type PredictionStructureAndBindingStartParamsInput struct {
	// Entities (proteins, RNA, DNA, ligands) forming the complex to predict. Order
	// determines chain assignment.
	Entities []PredictionStructureAndBindingStartParamsInputEntityUnion `json:"entities,omitzero" api:"required"`
	// Number of structure samples to generate
	NumSamples param.Opt[int64]                                          `json:"num_samples,omitzero"`
	Binding    PredictionStructureAndBindingStartParamsInputBindingUnion `json:"binding,omitzero"`
	// Bond constraints between atoms. Atom-level ligand references currently support
	// ligand_ccd only; ligand_smiles is unsupported.
	Bonds []PredictionStructureAndBindingStartParamsInputBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints  []PredictionStructureAndBindingStartParamsInputConstraintUnion `json:"constraints,omitzero"`
	ModelOptions PredictionStructureAndBindingStartParamsInputModelOptions      `json:"model_options,omitzero"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []PredictionStructureAndBindingStartParamsInputTemplate `json:"templates,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputEntityUnion struct {
	OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntity *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityRnaEntity           *PredictionStructureAndBindingStartParamsInputEntityRnaEntity           `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityDnaEntity           *PredictionStructureAndBindingStartParamsInputEntityDnaEntity           `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityLigandCcdEntity     *PredictionStructureAndBindingStartParamsInputEntityLigandCcdEntity     `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityLigandSmilesEntity  *PredictionStructureAndBindingStartParamsInputEntityLigandSmilesEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityRnaEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityDnaEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityLigandCcdEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityLigandSmilesEntity)
}
func (u *PredictionStructureAndBindingStartParamsInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityModification `json:"modifications,omitzero"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaUnion `json:"msa,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityModification struct {
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

func (r PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaUnion struct {
	OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa  *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa, u.OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa)
}
func (u *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
//
// The properties Format, Source, Type are required.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat `json:"format,omitzero" api:"required"`
	// How to provide a file to the API
	Source PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion `json:"source,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "custom".
	Type constant.Custom `json:"type" default:"custom"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat string

const (
	PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatA3m PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat = "a3m"
	PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatCsv PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat = "csv"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion struct {
	OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource    *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource    `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource, u.OfPredictionStructureAndBindingStartsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source)
}
func (u *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source struct {
	// Base64-encoded file contents
	Data string `json:"data" api:"required"`
	// MIME type (e.g., text/csv)
	MediaType string `json:"media_type" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewPredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa() PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa {
	return PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa{
		Type: "empty",
	}
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
//
// This struct has a constant value, construct it with
// [NewPredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa].
type PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa struct {
	Type constant.Empty `json:"type" default:"empty"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingStartParamsInputEntityRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityRnaEntityModification struct {
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

func (r PredictionStructureAndBindingStartParamsInputEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingStartParamsInputEntityDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityDnaEntityModification struct {
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

func (r PredictionStructureAndBindingStartParamsInputEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputBindingUnion struct {
	OfPredictionStructureAndBindingStartsInputBindingLigandProteinBinding  *PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputBindingProteinProteinBinding *PredictionStructureAndBindingStartParamsInputBindingProteinProteinBinding `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputBindingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputBindingLigandProteinBinding, u.OfPredictionStructureAndBindingStartsInputBindingProteinProteinBinding)
}
func (u *PredictionStructureAndBindingStartParamsInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties BinderChainID, Type are required.
type PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, <50 atoms, and only
	// ligands+proteins in entities)
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ligand_protein_binding".
	Type constant.LigandProteinBinding `json:"type" default:"ligand_protein_binding"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BinderChainIDs, Type are required.
type PredictionStructureAndBindingStartParamsInputBindingProteinProteinBinding struct {
	// Chain IDs of the protein binders
	BinderChainIDs []string `json:"binder_chain_ids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "protein_protein_binding".
	Type constant.ProteinProteinBinding `json:"type" default:"protein_protein_binding"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputBindingProteinProteinBinding) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBindingProteinProteinBinding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBindingProteinProteinBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type PredictionStructureAndBindingStartParamsInputBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 PredictionStructureAndBindingStartParamsInputBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 PredictionStructureAndBindingStartParamsInputBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputBond) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputBondAtom1Union struct {
	OfPredictionStructureAndBindingStartsInputBondAtom1LigandAtom  *PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputBondAtom1PolymerAtom *PredictionStructureAndBindingStartParamsInputBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputBondAtom1LigandAtom, u.OfPredictionStructureAndBindingStartsInputBondAtom1PolymerAtom)
}
func (u *PredictionStructureAndBindingStartParamsInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingStartParamsInputBondAtom1PolymerAtom struct {
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

func (r PredictionStructureAndBindingStartParamsInputBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputBondAtom2Union struct {
	OfPredictionStructureAndBindingStartsInputBondAtom2LigandAtom  *PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputBondAtom2PolymerAtom *PredictionStructureAndBindingStartParamsInputBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputBondAtom2LigandAtom, u.OfPredictionStructureAndBindingStartsInputBondAtom2PolymerAtom)
}
func (u *PredictionStructureAndBindingStartParamsInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingStartParamsInputBondAtom2PolymerAtom struct {
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

func (r PredictionStructureAndBindingStartParamsInputBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputConstraintUnion struct {
	OfPredictionStructureAndBindingStartsInputConstraintPocketConstraint  *PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputConstraintContactConstraint *PredictionStructureAndBindingStartParamsInputConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputConstraintPocketConstraint, u.OfPredictionStructureAndBindingStartsInputConstraintContactConstraint)
}
func (u *PredictionStructureAndBindingStartParamsInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint struct {
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

func (r PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1Union struct {
	OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken1PolymerContactToken *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken1LigandContactToken  *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken1PolymerContactToken, u.OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken1LigandContactToken)
}
func (u *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2Union struct {
	OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken2PolymerContactToken *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken2LigandContactToken  *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken2PolymerContactToken, u.OfPredictionStructureAndBindingStartsInputConstraintContactConstraintToken2LigandContactToken)
}
func (u *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartParamsInputModelOptions struct {
	// The number of recycling steps to use for prediction. Default is 3.
	RecyclingSteps param.Opt[int64] `json:"recycling_steps,omitzero"`
	// The number of sampling steps to use for prediction. Default is 200.
	SamplingSteps param.Opt[int64] `json:"sampling_steps,omitzero"`
	// Diffusion step scale (temperature). Controls sampling diversity — higher values
	// produce more varied structures. Default is 1.638.
	StepScale param.Opt[float64] `json:"step_scale,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputModelOptions) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputModelOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
//
// The properties TemplateChains, TemplateStructure are required.
type PredictionStructureAndBindingStartParamsInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains []PredictionStructureAndBindingStartParamsInputTemplateTemplateChain `json:"template_chains,omitzero" api:"required"`
	// How to provide a template structure file. URLs must point to a CIF or PDB file;
	// base64 uploads must use chemical/x-cif or chemical/x-pdb.
	TemplateStructure PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureUnion `json:"template_structure,omitzero" api:"required"`
	// Force the template reference potential with this distance threshold in
	// angstroms. Omit to use the template without force.
	ForceThresholdAngstroms param.Opt[float64] `json:"force_threshold_angstroms,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputTemplate) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
//
// The properties InputChainID, TemplateChainID are required.
type PredictionStructureAndBindingStartParamsInputTemplateTemplateChain struct {
	// Chain ID in this prediction request
	InputChainID string `json:"input_chain_id" api:"required"`
	// Corresponding chain ID in the template structure file
	TemplateChainID string `json:"template_chain_id" api:"required"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputTemplateTemplateChain) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputTemplateTemplateChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureUnion struct {
	OfPredictionStructureAndBindingStartsInputTemplateTemplateStructureURLSource                     *PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureURLSource                     `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputTemplateTemplateStructureTemplateStructureBase64Source *PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputTemplateTemplateStructureURLSource, u.OfPredictionStructureAndBindingStartsInputTemplateTemplateStructureTemplateStructureBase64Source)
}
func (u *PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64Source struct {
	// Base64-encoded template structure file contents
	Data string `json:"data" api:"required"`
	// Template structure MIME type
	//
	// Any of "chemical/x-cif", "chemical/x-pdb".
	MediaType PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType `json:"media_type,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure MIME type
type PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType string

const (
	PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaTypeChemicalXCif PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType = "chemical/x-cif"
	PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaTypeChemicalXPdb PredictionStructureAndBindingStartParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType = "chemical/x-pdb"
)

type PredictionStructureAndBindingTokenCountParams struct {
	Input PredictionStructureAndBindingTokenCountParamsInput `json:"input,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Model to use for prediction
	//
	// This field can be elided, and will marshal its zero value as "boltz-2.1".
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParams) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Entities is required.
type PredictionStructureAndBindingTokenCountParamsInput struct {
	// Entities (proteins, RNA, DNA, ligands) forming the complex to predict. Order
	// determines chain assignment.
	Entities []PredictionStructureAndBindingTokenCountParamsInputEntityUnion `json:"entities,omitzero" api:"required"`
	// Number of structure samples to generate
	NumSamples param.Opt[int64]                                               `json:"num_samples,omitzero"`
	Binding    PredictionStructureAndBindingTokenCountParamsInputBindingUnion `json:"binding,omitzero"`
	// Bond constraints between atoms. Atom-level ligand references currently support
	// ligand_ccd only; ligand_smiles is unsupported.
	Bonds []PredictionStructureAndBindingTokenCountParamsInputBond `json:"bonds,omitzero"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints  []PredictionStructureAndBindingTokenCountParamsInputConstraintUnion `json:"constraints,omitzero"`
	ModelOptions PredictionStructureAndBindingTokenCountParamsInputModelOptions      `json:"model_options,omitzero"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []PredictionStructureAndBindingTokenCountParamsInputTemplate `json:"templates,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputEntityUnion struct {
	OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntity *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntity `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputEntityRnaEntity           *PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntity           `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputEntityDnaEntity           *PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntity           `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputEntityLigandCcdEntity     *PredictionStructureAndBindingTokenCountParamsInputEntityLigandCcdEntity     `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputEntityLigandSmilesEntity  *PredictionStructureAndBindingTokenCountParamsInputEntityLigandSmilesEntity  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntity,
		u.OfPredictionStructureAndBindingTokenCountsInputEntityRnaEntity,
		u.OfPredictionStructureAndBindingTokenCountsInputEntityDnaEntity,
		u.OfPredictionStructureAndBindingTokenCountsInputEntityLigandCcdEntity,
		u.OfPredictionStructureAndBindingTokenCountsInputEntityLigandSmilesEntity)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityModification `json:"modifications,omitzero"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaUnion `json:"msa,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityModification struct {
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

func (r PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaUnion struct {
	OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa  *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa, u.OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
//
// The properties Format, Source, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat `json:"format,omitzero" api:"required"`
	// How to provide a file to the API
	Source PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion `json:"source,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "custom".
	Type constant.Custom `json:"type" default:"custom"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat string

const (
	PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatA3m PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat = "a3m"
	PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormatCsv PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaFormat = "csv"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion struct {
	OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource    *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource    `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource, u.OfPredictionStructureAndBindingTokenCountsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source struct {
	// Base64-encoded file contents
	Data string `json:"data" api:"required"`
	// MIME type (e.g., text/csv)
	MediaType string `json:"media_type" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2CustomMsaSourceBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewPredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa() PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa {
	return PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa{
		Type: "empty",
	}
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
//
// This struct has a constant value, construct it with
// [NewPredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa].
type PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa struct {
	Type constant.Empty `json:"type" default:"empty"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityBoltz2ProteinEntityMsaBoltz2EmptyMsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "rna".
	Type constant.Rna `json:"type" default:"rna"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntityModification struct {
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

func (r PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "dna".
	Type constant.Dna `json:"type" default:"dna"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
//
// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntityModification struct {
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

func (r PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityLigandCcdEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// CCD code (e.g., ATP, ADP)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_ccd".
	Type constant.LigandCcd `json:"type" default:"ligand_ccd"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityLigandCcdEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityLigandCcdEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingTokenCountParamsInputEntityLigandSmilesEntity struct {
	// Chain IDs for this ligand
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// SMILES string representing the ligand
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_smiles".
	Type constant.LigandSmiles `json:"type" default:"ligand_smiles"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputEntityLigandSmilesEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputEntityLigandSmilesEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputBindingUnion struct {
	OfPredictionStructureAndBindingTokenCountsInputBindingLigandProteinBinding  *PredictionStructureAndBindingTokenCountParamsInputBindingLigandProteinBinding  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputBindingProteinProteinBinding *PredictionStructureAndBindingTokenCountParamsInputBindingProteinProteinBinding `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputBindingUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputBindingLigandProteinBinding, u.OfPredictionStructureAndBindingTokenCountsInputBindingProteinProteinBinding)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties BinderChainID, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputBindingLigandProteinBinding struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, <50 atoms, and only
	// ligands+proteins in entities)
	BinderChainID string `json:"binder_chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "ligand_protein_binding".
	Type constant.LigandProteinBinding `json:"type" default:"ligand_protein_binding"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputBindingLigandProteinBinding) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBindingLigandProteinBinding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBindingLigandProteinBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BinderChainIDs, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputBindingProteinProteinBinding struct {
	// Chain IDs of the protein binders
	BinderChainIDs []string `json:"binder_chain_ids,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as
	// "protein_protein_binding".
	Type constant.ProteinProteinBinding `json:"type" default:"protein_protein_binding"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputBindingProteinProteinBinding) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBindingProteinProteinBinding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBindingProteinProteinBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties Atom1, Atom2 are required.
type PredictionStructureAndBindingTokenCountParamsInputBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 PredictionStructureAndBindingTokenCountParamsInputBondAtom1Union `json:"atom1,omitzero" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 PredictionStructureAndBindingTokenCountParamsInputBondAtom2Union `json:"atom2,omitzero" api:"required"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputBond) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBond
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputBondAtom1Union struct {
	OfPredictionStructureAndBindingTokenCountsInputBondAtom1LigandAtom  *PredictionStructureAndBindingTokenCountParamsInputBondAtom1LigandAtom  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputBondAtom1PolymerAtom *PredictionStructureAndBindingTokenCountParamsInputBondAtom1PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputBondAtom1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputBondAtom1LigandAtom, u.OfPredictionStructureAndBindingTokenCountsInputBondAtom1PolymerAtom)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputBondAtom1LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputBondAtom1LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBondAtom1LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBondAtom1LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputBondAtom1PolymerAtom struct {
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

func (r PredictionStructureAndBindingTokenCountParamsInputBondAtom1PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBondAtom1PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBondAtom1PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputBondAtom2Union struct {
	OfPredictionStructureAndBindingTokenCountsInputBondAtom2LigandAtom  *PredictionStructureAndBindingTokenCountParamsInputBondAtom2LigandAtom  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputBondAtom2PolymerAtom *PredictionStructureAndBindingTokenCountParamsInputBondAtom2PolymerAtom `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputBondAtom2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputBondAtom2LigandAtom, u.OfPredictionStructureAndBindingTokenCountsInputBondAtom2PolymerAtom)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputBondAtom2LigandAtom struct {
	// Standardized atom name (verifiable in CIF file on RCSB). Atom-level references
	// to ligand_smiles entities are currently unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID containing the atom
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_atom".
	Type constant.LigandAtom `json:"type" default:"ligand_atom"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputBondAtom2LigandAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBondAtom2LigandAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBondAtom2LigandAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AtomName, ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputBondAtom2PolymerAtom struct {
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

func (r PredictionStructureAndBindingTokenCountParamsInputBondAtom2PolymerAtom) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputBondAtom2PolymerAtom
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputBondAtom2PolymerAtom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputConstraintUnion struct {
	OfPredictionStructureAndBindingTokenCountsInputConstraintPocketConstraint  *PredictionStructureAndBindingTokenCountParamsInputConstraintPocketConstraint  `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraint *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraint `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputConstraintUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputConstraintPocketConstraint, u.OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraint)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Constrains the binder to interact with specific pocket residues on the target.
//
// The properties BinderChainID, ContactResidues, MaxDistanceAngstrom, Type are
// required.
type PredictionStructureAndBindingTokenCountParamsInputConstraintPocketConstraint struct {
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

func (r PredictionStructureAndBindingTokenCountParamsInputConstraintPocketConstraint) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputConstraintPocketConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputConstraintPocketConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
//
// The properties MaxDistanceAngstrom, Token1, Token2, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraint struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1Union `json:"token1,omitzero" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2Union `json:"token2,omitzero" api:"required"`
	// Whether to force the constraint
	Force param.Opt[bool] `json:"force,omitzero"`
	// This field can be elided, and will marshal its zero value as "contact".
	Type constant.Contact `json:"type" default:"contact"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraint) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraint
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1Union struct {
	OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken1PolymerContactToken *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1PolymerContactToken `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken1LigandContactToken  *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken1PolymerContactToken, u.OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken1LigandContactToken)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken1LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2Union struct {
	OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken2PolymerContactToken *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2PolymerContactToken `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken2LigandContactToken  *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2LigandContactToken  `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken2PolymerContactToken, u.OfPredictionStructureAndBindingTokenCountsInputConstraintContactConstraintToken2LigandContactToken)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainID, ResidueIndex, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2PolymerContactToken struct {
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// 0-based residue index
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// This field can be elided, and will marshal its zero value as "polymer_contact".
	Type constant.PolymerContact `json:"type" default:"polymer_contact"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2PolymerContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2PolymerContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2PolymerContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
//
// The properties AtomName, ChainID, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2LigandContactToken struct {
	// Atom name. Atom-level references to ligand_smiles entities are currently
	// unsupported; use ligand_ccd instead.
	AtomName string `json:"atom_name" api:"required"`
	// Chain ID
	ChainID string `json:"chain_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "ligand_contact".
	Type constant.LigandContact `json:"type" default:"ligand_contact"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2LigandContactToken) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2LigandContactToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputConstraintContactConstraintToken2LigandContactToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingTokenCountParamsInputModelOptions struct {
	// The number of recycling steps to use for prediction. Default is 3.
	RecyclingSteps param.Opt[int64] `json:"recycling_steps,omitzero"`
	// The number of sampling steps to use for prediction. Default is 200.
	SamplingSteps param.Opt[int64] `json:"sampling_steps,omitzero"`
	// Diffusion step scale (temperature). Controls sampling diversity — higher values
	// produce more varied structures. Default is 1.638.
	StepScale param.Opt[float64] `json:"step_scale,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputModelOptions) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputModelOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
//
// The properties TemplateChains, TemplateStructure are required.
type PredictionStructureAndBindingTokenCountParamsInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains []PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateChain `json:"template_chains,omitzero" api:"required"`
	// How to provide a template structure file. URLs must point to a CIF or PDB file;
	// base64 uploads must use chemical/x-cif or chemical/x-pdb.
	TemplateStructure PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureUnion `json:"template_structure,omitzero" api:"required"`
	// Force the template reference potential with this distance threshold in
	// angstroms. Omit to use the template without force.
	ForceThresholdAngstroms param.Opt[float64] `json:"force_threshold_angstroms,omitzero"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputTemplate) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
//
// The properties InputChainID, TemplateChainID are required.
type PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateChain struct {
	// Chain ID in this prediction request
	InputChainID string `json:"input_chain_id" api:"required"`
	// Corresponding chain ID in the template structure file
	TemplateChainID string `json:"template_chain_id" api:"required"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateChain) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateChain
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureUnion struct {
	OfPredictionStructureAndBindingTokenCountsInputTemplateTemplateStructureURLSource                     *PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureURLSource                     `json:",omitzero,inline"`
	OfPredictionStructureAndBindingTokenCountsInputTemplateTemplateStructureTemplateStructureBase64Source *PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64Source `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingTokenCountsInputTemplateTemplateStructureURLSource, u.OfPredictionStructureAndBindingTokenCountsInputTemplateTemplateStructureTemplateStructureBase64Source)
}
func (u *PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Type, URL are required.
type PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureURLSource struct {
	URL string `json:"url" api:"required" format:"uri"`
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type" default:"url"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureURLSource) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MediaType, Type are required.
type PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64Source struct {
	// Base64-encoded template structure file contents
	Data string `json:"data" api:"required"`
	// Template structure MIME type
	//
	// Any of "chemical/x-cif", "chemical/x-pdb".
	MediaType PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType `json:"media_type,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "base64".
	Type constant.Base64 `json:"type" default:"base64"`
	paramObj
}

func (r PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure MIME type
type PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType string

const (
	PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaTypeChemicalXCif PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType = "chemical/x-cif"
	PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaTypeChemicalXPdb PredictionStructureAndBindingTokenCountParamsInputTemplateTemplateStructureTemplateStructureBase64SourceMediaType = "chemical/x-pdb"
)
