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
// metrics for a molecular complex.
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
// scores for the input molecular complex, with optional binding metrics.
func (r *PredictionStructureAndBindingService) Start(ctx context.Context, body PredictionStructureAndBindingStartParams, opts ...option.RequestOption) (res *PredictionStructureAndBindingStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/predictions/structure-and-binding"
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities     respjson.Field
		Binding      respjson.Field
		Bonds        respjson.Field
		Constraints  respjson.Field
		ModelOptions respjson.Field
		NumSamples   respjson.Field
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
// [PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponse],
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
	// [[]PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification],
	// [[]PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification],
	// [[]PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification]
	Modifications PredictionStructureAndBindingGetResponseInputEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u PredictionStructureAndBindingGetResponseInputEntityUnion) AsPredictionStructureAndBindingGetResponseInputEntityProteinEntityResponse() (v PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponse) {
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
// OfPredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModifications
// OfPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModifications
// OfPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModifications]
type PredictionStructureAndBindingGetResponseInputEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModifications []PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModifications []PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModifications []PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                struct {
		OfPredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModifications respjson.Field
		OfPredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModifications     respjson.Field
		OfPredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                     string
	} `json:"-"`
}

func (r *PredictionStructureAndBindingGetResponseInputEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification `json:"modifications"`
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
func (r PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification struct {
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
func (r PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingGetResponseInputEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

type PredictionStructureAndBindingGetResponseInputEntityRnaEntityResponseModification struct {
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
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

type PredictionStructureAndBindingGetResponseInputEntityDnaEntityResponseModification struct {
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
	Metrics   PredictionStructureAndBindingGetResponseOutputAllSampleResultMetrics   `json:"metrics" api:"required"`
	Structure PredictionStructureAndBindingGetResponseOutputAllSampleResultStructure `json:"structure" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type PredictionStructureAndBindingGetResponseOutputBestSample struct {
	Metrics   PredictionStructureAndBindingGetResponseOutputBestSampleMetrics   `json:"metrics" api:"required"`
	Structure PredictionStructureAndBindingGetResponseOutputBestSampleStructure `json:"structure" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities     respjson.Field
		Binding      respjson.Field
		Bonds        respjson.Field
		Constraints  respjson.Field
		ModelOptions respjson.Field
		NumSamples   respjson.Field
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
// [PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponse],
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
	// [[]PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification],
	// [[]PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification],
	// [[]PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification]
	Modifications PredictionStructureAndBindingStartResponseInputEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u PredictionStructureAndBindingStartResponseInputEntityUnion) AsPredictionStructureAndBindingStartResponseInputEntityProteinEntityResponse() (v PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponse) {
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
// OfPredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModifications
// OfPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModifications
// OfPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModifications]
type PredictionStructureAndBindingStartResponseInputEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModifications []PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModifications []PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification]
	// instead of an object.
	OfPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModifications []PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                  struct {
		OfPredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModifications respjson.Field
		OfPredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModifications     respjson.Field
		OfPredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                       string
	} `json:"-"`
}

func (r *PredictionStructureAndBindingStartResponseInputEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification `json:"modifications"`
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
func (r PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification struct {
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
func (r PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionStructureAndBindingStartResponseInputEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

type PredictionStructureAndBindingStartResponseInputEntityRnaEntityResponseModification struct {
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
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

type PredictionStructureAndBindingStartResponseInputEntityDnaEntityResponseModification struct {
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
	Metrics   PredictionStructureAndBindingStartResponseOutputAllSampleResultMetrics   `json:"metrics" api:"required"`
	Structure PredictionStructureAndBindingStartResponseOutputAllSampleResultStructure `json:"structure" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type PredictionStructureAndBindingStartResponseOutputBestSample struct {
	Metrics   PredictionStructureAndBindingStartResponseOutputBestSampleMetrics   `json:"metrics" api:"required"`
	Structure PredictionStructureAndBindingStartResponseOutputBestSampleStructure `json:"structure" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntity      *PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity      `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityRnaEntity          *PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntity          `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityDnaEntity          *PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntity          `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityLigandCcdEntity    *PredictionStructureAndBindingEstimateCostParamsInputEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfPredictionStructureAndBindingEstimateCostsInputEntityLigandSmilesEntity *PredictionStructureAndBindingEstimateCostParamsInputEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingEstimateCostParamsInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityRnaEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityDnaEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityLigandCcdEntity,
		u.OfPredictionStructureAndBindingEstimateCostsInputEntityLigandSmilesEntity)
}
func (u *PredictionStructureAndBindingEstimateCostParamsInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityRnaEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
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
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingEstimateCostParamsInputEntityDnaEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
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
	OfPredictionStructureAndBindingStartsInputEntityProteinEntity      *PredictionStructureAndBindingStartParamsInputEntityProteinEntity      `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityRnaEntity          *PredictionStructureAndBindingStartParamsInputEntityRnaEntity          `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityDnaEntity          *PredictionStructureAndBindingStartParamsInputEntityDnaEntity          `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityLigandCcdEntity    *PredictionStructureAndBindingStartParamsInputEntityLigandCcdEntity    `json:",omitzero,inline"`
	OfPredictionStructureAndBindingStartsInputEntityLigandSmilesEntity *PredictionStructureAndBindingStartParamsInputEntityLigandSmilesEntity `json:",omitzero,inline"`
	paramUnion
}

func (u PredictionStructureAndBindingStartParamsInputEntityUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredictionStructureAndBindingStartsInputEntityProteinEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityRnaEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityDnaEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityLigandCcdEntity,
		u.OfPredictionStructureAndBindingStartsInputEntityLigandSmilesEntity)
}
func (u *PredictionStructureAndBindingStartParamsInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties ChainIDs, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string `json:"chain_ids,omitzero" api:"required"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic param.Opt[bool] `json:"cyclic,omitzero"`
	// Post-translational modifications. Optional; defaults to an empty list when
	// omitted.
	Modifications []PredictionStructureAndBindingStartParamsInputEntityProteinEntityModification `json:"modifications,omitzero"`
	// This field can be elided, and will marshal its zero value as "protein".
	Type constant.Protein `json:"type" default:"protein"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityProteinEntity) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityProteinEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityProteinEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "ccd".
	Type constant.Ccd `json:"type" default:"ccd"`
	paramObj
}

func (r PredictionStructureAndBindingStartParamsInputEntityProteinEntityModification) MarshalJSON() (data []byte, err error) {
	type shadow PredictionStructureAndBindingStartParamsInputEntityProteinEntityModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionStructureAndBindingStartParamsInputEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
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
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityRnaEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
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
	// Chemical modifications. Optional; defaults to an empty list when omitted.
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

// The properties ResidueIndex, Type, Value are required.
type PredictionStructureAndBindingStartParamsInputEntityDnaEntityModification struct {
	// 0-based index of the residue to modify
	ResidueIndex int64 `json:"residue_index" api:"required"`
	// CCD code from RCSB PDB (e.g. 'MSE' for selenomethionine, 'SEP' for
	// phosphoserine)
	Value string `json:"value" api:"required"`
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
