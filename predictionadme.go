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

// Predict Tier 1 ADME summary values for a batch of small molecules specified by
// SMILES.
//
// PredictionAdmeService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPredictionAdmeService] method instead.
type PredictionAdmeService struct {
	Options []option.RequestOption
}

// NewPredictionAdmeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPredictionAdmeService(opts ...option.RequestOption) (r PredictionAdmeService) {
	r = PredictionAdmeService{}
	r.Options = opts
	return
}

// Retrieve an ADME prediction by ID, including its status and results.
func (r *PredictionAdmeService) Get(ctx context.Context, id string, query PredictionAdmeGetParams, opts ...option.RequestOption) (res *PredictionAdmeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/predictions/adme/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List ADME predictions, optionally filtered by workspace
func (r *PredictionAdmeService) List(ctx context.Context, query PredictionAdmeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[PredictionAdmeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/predictions/adme"
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

// List ADME predictions, optionally filtered by workspace
func (r *PredictionAdmeService) ListAutoPaging(ctx context.Context, query PredictionAdmeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[PredictionAdmeListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete the input, output, and result data associated with this
// prediction. The prediction record itself is retained with a `data_deleted_at`
// timestamp. This action is irreversible.
func (r *PredictionAdmeService) DeleteData(ctx context.Context, id string, opts ...option.RequestOption) (res *PredictionAdmeDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/predictions/adme/%s/delete-data", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Submit a prediction job that returns Tier 1 ADME summary values for each
// requested molecule.
func (r *PredictionAdmeService) Start(ctx context.Context, body PredictionAdmeStartParams, opts ...option.RequestOption) (res *PredictionAdmeStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/predictions/adme"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PredictionAdmeGetResponse struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error PredictionAdmeGetResponseError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input PredictionAdmeGetResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.AdmeV1 `json:"model" default:"adme-v1"`
	// Prediction output when succeeded
	Output    PredictionAdmeGetResponseOutput `json:"output" api:"required"`
	StartedAt time.Time                       `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status PredictionAdmeGetResponseStatus `json:"status" api:"required"`
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
func (r PredictionAdmeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type PredictionAdmeGetResponseError struct {
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
func (r PredictionAdmeGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type PredictionAdmeGetResponseInput struct {
	// Molecules to score. Results are returned in the same order as this list.
	Molecules []PredictionAdmeGetResponseInputMolecule `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeGetResponseInputMolecule struct {
	// SMILES string of the molecule to predict ADME properties for.
	Smiles string `json:"smiles" api:"required"`
	// Optional client-provided identifier. Returned as `external_id` in the matching
	// output item.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Smiles      respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeGetResponseInputMolecule) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeGetResponseInputMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type PredictionAdmeGetResponseOutput struct {
	// Per-molecule results in the same order as the request. Successful molecules
	// carry an `adme` summary. Failed molecules carry `status: "failed"` and a
	// non-null `error`.
	Molecules []PredictionAdmeGetResponseOutputMoleculeUnion `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeGetResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeGetResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionAdmeGetResponseOutputMoleculeUnion contains all possible properties
// and values from [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceeded],
// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailed].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionAdmeGetResponseOutputMoleculeUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme], [any]
	Adme PredictionAdmeGetResponseOutputMoleculeUnionAdme `json:"adme"`
	// This field is a union of [any],
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError]
	Error      PredictionAdmeGetResponseOutputMoleculeUnionError `json:"error"`
	Smiles     string                                            `json:"smiles"`
	Status     string                                            `json:"status"`
	ExternalID string                                            `json:"external_id"`
	JSON       struct {
		ID         respjson.Field
		Adme       respjson.Field
		Error      respjson.Field
		Smiles     respjson.Field
		Status     respjson.Field
		ExternalID respjson.Field
		raw        string
	} `json:"-"`
}

func (u PredictionAdmeGetResponseOutputMoleculeUnion) AsPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceeded() (v PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceeded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionAdmeGetResponseOutputMoleculeUnion) AsPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailed() (v PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionAdmeGetResponseOutputMoleculeUnion) RawJSON() string { return u.JSON.raw }

func (r *PredictionAdmeGetResponseOutputMoleculeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionAdmeGetResponseOutputMoleculeUnionAdme is an implicit subunion of
// [PredictionAdmeGetResponseOutputMoleculeUnion].
// PredictionAdmeGetResponseOutputMoleculeUnionAdme provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PredictionAdmeGetResponseOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedAdme]
type PredictionAdmeGetResponseOutputMoleculeUnionAdme struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedAdme any `json:",inline"`
	// This field is from variant
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme].
	Liphophilicity float64 `json:"liphophilicity"`
	// This field is from variant
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme].
	Permeability float64 `json:"permeability"`
	// This field is from variant
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme].
	Solubility PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility"`
	JSON       struct {
		OfPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedAdme respjson.Field
		Liphophilicity                                                  respjson.Field
		Permeability                                                    respjson.Field
		Solubility                                                      respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (r *PredictionAdmeGetResponseOutputMoleculeUnionAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionAdmeGetResponseOutputMoleculeUnionError is an implicit subunion of
// [PredictionAdmeGetResponseOutputMoleculeUnion].
// PredictionAdmeGetResponseOutputMoleculeUnionError provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PredictionAdmeGetResponseOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededError]
type PredictionAdmeGetResponseOutputMoleculeUnionError struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededError any `json:",inline"`
	// This field is from variant
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError].
	Code string `json:"code"`
	// This field is from variant
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError].
	Message string `json:"message"`
	// This field is from variant
	// [PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError].
	Details any `json:"details"`
	JSON    struct {
		OfPredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededError respjson.Field
		Code                                                                respjson.Field
		Message                                                             respjson.Field
		Details                                                             respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (r *PredictionAdmeGetResponseOutputMoleculeUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceeded struct {
	// Internally generated molecule identifier.
	ID string `json:"id" api:"required"`
	// Tier 1 ADME summary values for this molecule.
	Adme  PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme `json:"adme" api:"required"`
	Error any                                                              `json:"error" api:"required"`
	// Echoed SMILES from the request.
	Smiles string             `json:"smiles" api:"required"`
	Status constant.Succeeded `json:"status" default:"succeeded"`
	// Client-provided molecule identifier, if one was supplied.
	ExternalID string `json:"external_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Adme        respjson.Field
		Error       respjson.Field
		Smiles      respjson.Field
		Status      respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceeded) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceeded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tier 1 ADME summary values for this molecule.
type PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme struct {
	// Lipophilicity score from the internal LogD prediction.
	Liphophilicity float64 `json:"liphophilicity" api:"required"`
	// Permeability score for this molecule.
	Permeability float64 `json:"permeability" api:"required"`
	// Solubility judgement for this molecule.
	//
	// Any of "high-confidence", "medium-confidence", "high-risk".
	Solubility PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Liphophilicity respjson.Field
		Permeability   respjson.Field
		Solubility     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solubility judgement for this molecule.
type PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility string

const (
	PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighConfidence   PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-confidence"
	PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityMediumConfidence PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "medium-confidence"
	PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighRisk         PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-risk"
)

type PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailed struct {
	// Internally generated molecule identifier.
	ID    string                                                         `json:"id" api:"required"`
	Adme  any                                                            `json:"adme" api:"required"`
	Error PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError `json:"error" api:"required"`
	// Echoed SMILES from the request.
	Smiles string          `json:"smiles" api:"required"`
	Status constant.Failed `json:"status" default:"failed"`
	// Client-provided molecule identifier, if one was supplied.
	ExternalID string `json:"external_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Adme        respjson.Field
		Error       respjson.Field
		Smiles      respjson.Field
		Status      respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailed) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError struct {
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
func (r PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeGetResponseOutputMoleculeAdmeMoleculeFailedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeGetResponseStatus string

const (
	PredictionAdmeGetResponseStatusPending   PredictionAdmeGetResponseStatus = "pending"
	PredictionAdmeGetResponseStatusRunning   PredictionAdmeGetResponseStatus = "running"
	PredictionAdmeGetResponseStatusSucceeded PredictionAdmeGetResponseStatus = "succeeded"
	PredictionAdmeGetResponseStatusFailed    PredictionAdmeGetResponseStatus = "failed"
)

type PredictionAdmeListResponse struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error PredictionAdmeListResponseError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model     constant.AdmeV1 `json:"model" default:"adme-v1"`
	StartedAt time.Time       `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status PredictionAdmeListResponseStatus `json:"status" api:"required"`
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
func (r PredictionAdmeListResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type PredictionAdmeListResponseError struct {
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
func (r PredictionAdmeListResponseError) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeListResponseStatus string

const (
	PredictionAdmeListResponseStatusPending   PredictionAdmeListResponseStatus = "pending"
	PredictionAdmeListResponseStatusRunning   PredictionAdmeListResponseStatus = "running"
	PredictionAdmeListResponseStatusSucceeded PredictionAdmeListResponseStatus = "succeeded"
	PredictionAdmeListResponseStatusFailed    PredictionAdmeListResponseStatus = "failed"
)

type PredictionAdmeDeleteDataResponse struct {
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
func (r PredictionAdmeDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeStartResponse struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error PredictionAdmeStartResponseError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input PredictionAdmeStartResponseInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.AdmeV1 `json:"model" default:"adme-v1"`
	// Prediction output when succeeded
	Output    PredictionAdmeStartResponseOutput `json:"output" api:"required"`
	StartedAt time.Time                         `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status PredictionAdmeStartResponseStatus `json:"status" api:"required"`
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
func (r PredictionAdmeStartResponse) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type PredictionAdmeStartResponseError struct {
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
func (r PredictionAdmeStartResponseError) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeStartResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type PredictionAdmeStartResponseInput struct {
	// Molecules to score. Results are returned in the same order as this list.
	Molecules []PredictionAdmeStartResponseInputMolecule `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeStartResponseInput) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeStartResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeStartResponseInputMolecule struct {
	// SMILES string of the molecule to predict ADME properties for.
	Smiles string `json:"smiles" api:"required"`
	// Optional client-provided identifier. Returned as `external_id` in the matching
	// output item.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Smiles      respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeStartResponseInputMolecule) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeStartResponseInputMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type PredictionAdmeStartResponseOutput struct {
	// Per-molecule results in the same order as the request. Successful molecules
	// carry an `adme` summary. Failed molecules carry `status: "failed"` and a
	// non-null `error`.
	Molecules []PredictionAdmeStartResponseOutputMoleculeUnion `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeStartResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *PredictionAdmeStartResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionAdmeStartResponseOutputMoleculeUnion contains all possible properties
// and values from
// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceeded],
// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailed].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PredictionAdmeStartResponseOutputMoleculeUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme], [any]
	Adme PredictionAdmeStartResponseOutputMoleculeUnionAdme `json:"adme"`
	// This field is a union of [any],
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError]
	Error      PredictionAdmeStartResponseOutputMoleculeUnionError `json:"error"`
	Smiles     string                                              `json:"smiles"`
	Status     string                                              `json:"status"`
	ExternalID string                                              `json:"external_id"`
	JSON       struct {
		ID         respjson.Field
		Adme       respjson.Field
		Error      respjson.Field
		Smiles     respjson.Field
		Status     respjson.Field
		ExternalID respjson.Field
		raw        string
	} `json:"-"`
}

func (u PredictionAdmeStartResponseOutputMoleculeUnion) AsPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceeded() (v PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceeded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PredictionAdmeStartResponseOutputMoleculeUnion) AsPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailed() (v PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PredictionAdmeStartResponseOutputMoleculeUnion) RawJSON() string { return u.JSON.raw }

func (r *PredictionAdmeStartResponseOutputMoleculeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionAdmeStartResponseOutputMoleculeUnionAdme is an implicit subunion of
// [PredictionAdmeStartResponseOutputMoleculeUnion].
// PredictionAdmeStartResponseOutputMoleculeUnionAdme provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PredictionAdmeStartResponseOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedAdme]
type PredictionAdmeStartResponseOutputMoleculeUnionAdme struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedAdme any `json:",inline"`
	// This field is from variant
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme].
	Liphophilicity float64 `json:"liphophilicity"`
	// This field is from variant
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme].
	Permeability float64 `json:"permeability"`
	// This field is from variant
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme].
	Solubility PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility"`
	JSON       struct {
		OfPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedAdme respjson.Field
		Liphophilicity                                                    respjson.Field
		Permeability                                                      respjson.Field
		Solubility                                                        respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (r *PredictionAdmeStartResponseOutputMoleculeUnionAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PredictionAdmeStartResponseOutputMoleculeUnionError is an implicit subunion of
// [PredictionAdmeStartResponseOutputMoleculeUnion].
// PredictionAdmeStartResponseOutputMoleculeUnionError provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PredictionAdmeStartResponseOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededError]
type PredictionAdmeStartResponseOutputMoleculeUnionError struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededError any `json:",inline"`
	// This field is from variant
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError].
	Code string `json:"code"`
	// This field is from variant
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError].
	Message string `json:"message"`
	// This field is from variant
	// [PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError].
	Details any `json:"details"`
	JSON    struct {
		OfPredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededError respjson.Field
		Code                                                                  respjson.Field
		Message                                                               respjson.Field
		Details                                                               respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (r *PredictionAdmeStartResponseOutputMoleculeUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceeded struct {
	// Internally generated molecule identifier.
	ID string `json:"id" api:"required"`
	// Tier 1 ADME summary values for this molecule.
	Adme  PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme `json:"adme" api:"required"`
	Error any                                                                `json:"error" api:"required"`
	// Echoed SMILES from the request.
	Smiles string             `json:"smiles" api:"required"`
	Status constant.Succeeded `json:"status" default:"succeeded"`
	// Client-provided molecule identifier, if one was supplied.
	ExternalID string `json:"external_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Adme        respjson.Field
		Error       respjson.Field
		Smiles      respjson.Field
		Status      respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceeded) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceeded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tier 1 ADME summary values for this molecule.
type PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme struct {
	// Lipophilicity score from the internal LogD prediction.
	Liphophilicity float64 `json:"liphophilicity" api:"required"`
	// Permeability score for this molecule.
	Permeability float64 `json:"permeability" api:"required"`
	// Solubility judgement for this molecule.
	//
	// Any of "high-confidence", "medium-confidence", "high-risk".
	Solubility PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Liphophilicity respjson.Field
		Permeability   respjson.Field
		Solubility     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solubility judgement for this molecule.
type PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility string

const (
	PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighConfidence   PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-confidence"
	PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityMediumConfidence PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "medium-confidence"
	PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighRisk         PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-risk"
)

type PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailed struct {
	// Internally generated molecule identifier.
	ID    string                                                           `json:"id" api:"required"`
	Adme  any                                                              `json:"adme" api:"required"`
	Error PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError `json:"error" api:"required"`
	// Echoed SMILES from the request.
	Smiles string          `json:"smiles" api:"required"`
	Status constant.Failed `json:"status" default:"failed"`
	// Client-provided molecule identifier, if one was supplied.
	ExternalID string `json:"external_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Adme        respjson.Field
		Error       respjson.Field
		Smiles      respjson.Field
		Status      respjson.Field
		ExternalID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailed) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError struct {
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
func (r PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError) RawJSON() string {
	return r.JSON.raw
}
func (r *PredictionAdmeStartResponseOutputMoleculeAdmeMoleculeFailedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PredictionAdmeStartResponseStatus string

const (
	PredictionAdmeStartResponseStatusPending   PredictionAdmeStartResponseStatus = "pending"
	PredictionAdmeStartResponseStatusRunning   PredictionAdmeStartResponseStatus = "running"
	PredictionAdmeStartResponseStatusSucceeded PredictionAdmeStartResponseStatus = "succeeded"
	PredictionAdmeStartResponseStatusFailed    PredictionAdmeStartResponseStatus = "failed"
)

type PredictionAdmeGetParams struct {
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PredictionAdmeGetParams]'s query parameters as
// `url.Values`.
func (r PredictionAdmeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PredictionAdmeListParams struct {
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

// URLQuery serializes [PredictionAdmeListParams]'s query parameters as
// `url.Values`.
func (r PredictionAdmeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PredictionAdmeStartParams struct {
	Input PredictionAdmeStartParamsInput `json:"input,omitzero" api:"required"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Target workspace ID (admin keys only; ignored for workspace keys)
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Model to use for prediction
	//
	// This field can be elided, and will marshal its zero value as "adme-v1".
	Model constant.AdmeV1 `json:"model" default:"adme-v1"`
	paramObj
}

func (r PredictionAdmeStartParams) MarshalJSON() (data []byte, err error) {
	type shadow PredictionAdmeStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionAdmeStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Molecules is required.
type PredictionAdmeStartParamsInput struct {
	// Molecules to score. Results are returned in the same order as this list.
	Molecules []PredictionAdmeStartParamsInputMolecule `json:"molecules,omitzero" api:"required"`
	paramObj
}

func (r PredictionAdmeStartParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow PredictionAdmeStartParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionAdmeStartParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Smiles is required.
type PredictionAdmeStartParamsInputMolecule struct {
	// SMILES string of the molecule to predict ADME properties for.
	Smiles string `json:"smiles" api:"required"`
	// Optional client-provided identifier. Returned as `external_id` in the matching
	// output item.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r PredictionAdmeStartParamsInputMolecule) MarshalJSON() (data []byte, err error) {
	type shadow PredictionAdmeStartParamsInputMolecule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PredictionAdmeStartParamsInputMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
