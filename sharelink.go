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

// Share read-only access to predictions and pipeline runs by issuing time-limited
// links that any visitor can open without an API key. A share link is scoped to a
// single workspace and bundles one or more predictions and pipeline runs. The link
// ID is itself the bearer credential — treat it as a secret. Create and delete
// require a workspace-scoped API key with read permission on every referenced
// resource; read endpoints are unauthenticated and gated by the link ID. Deleting
// a link revokes the bearer credential immediately — subsequent reads return 404.
// The underlying predictions and pipelines are not affected and remain accessible
// through their own authenticated endpoints; per-resource data retention runs on
// its own clock. The action is irreversible.
//
// ShareLinkService contains methods and other services that help with interacting
// with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShareLinkService] method instead.
type ShareLinkService struct {
	Options []option.RequestOption
}

// NewShareLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewShareLinkService(opts ...option.RequestOption) (r ShareLinkService) {
	r = ShareLinkService{}
	r.Options = opts
	return
}

// Create an unauthenticated, read-only share link covering one or more predictions
// and/or pipelines that all live in the same workspace. The returned `id` is the
// bearer credential — treat it as a secret.
func (r *ShareLinkService) New(ctx context.Context, body ShareLinkNewParams, opts ...option.RequestOption) (res *ShareLinkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/share-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Read the predictions and pipelines exposed by a share link. No authentication is
// required — the share link ID itself is the access credential. Returns 404
// indistinguishably for unknown, expired, or revoked links.
func (r *ShareLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ShareLinkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/share/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke a share link so its bearer credential can no longer be used. The record
// is retained with a `data_deleted_at` timestamp and is purged by a background
// sweep. This action is irreversible. The share link ID is sent in the body
// because it is the bearer credential and must not appear in URLs.
func (r *ShareLinkService) DeleteData(ctx context.Context, body ShareLinkDeleteDataParams, opts ...option.RequestOption) (res *ShareLinkDeleteDataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/share-links/delete-data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Paginated results for one pipeline exposed by a share link. The response shape
// matches the authed pipeline-results endpoints exactly. No authentication is
// required — the share-link ID gates access. Pipeline IDs not covered by the link
// return 404 indistinguishably from unknown links.
func (r *ShareLinkService) ListPipelineResults(ctx context.Context, pipelineID string, params ShareLinkListPipelineResultsParams, opts ...option.RequestOption) (res *pagination.CursorPage[ShareLinkListPipelineResultsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if pipelineID == "" {
		err = errors.New("missing required pipelineId parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/share/%s/pipelines/%s/results", url.PathEscape(params.ID), url.PathEscape(pipelineID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Paginated results for one pipeline exposed by a share link. The response shape
// matches the authed pipeline-results endpoints exactly. No authentication is
// required — the share-link ID gates access. Pipeline IDs not covered by the link
// return 404 indistinguishably from unknown links.
func (r *ShareLinkService) ListPipelineResultsAutoPaging(ctx context.Context, pipelineID string, params ShareLinkListPipelineResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ShareLinkListPipelineResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListPipelineResults(ctx, pipelineID, params, opts...))
}

type ShareLinkNewResponse struct {
	// Share link ID. This value is the bearer credential used to access the linked
	// resources — treat it as a secret.
	ID string `json:"id" api:"required"`
	// When the share link was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the share link stops granting access.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Pipelines exposed by this share link.
	PipelineIDs []string `json:"pipeline_ids" api:"required"`
	// Predictions exposed by this share link.
	PredictionIDs []string `json:"prediction_ids" api:"required"`
	// Workspace that owns the share link and the referenced resources.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Visitable share link URL for the deployment's public app. Present when the
	// deployment has a configured app host; otherwise construct as
	// `<your-app-host>/share/{id}`. Treat as a secret — the `{id}` segment is the
	// bearer credential.
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ExpiresAt     respjson.Field
		PipelineIDs   respjson.Field
		PredictionIDs respjson.Field
		WorkspaceID   respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponse struct {
	// Share link ID.
	ID string `json:"id" api:"required"`
	// When the share link was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the share link stops granting access.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Pipelines exposed by this share link, in the order they were registered.
	Pipelines []ShareLinkGetResponsePipelineUnion `json:"pipelines" api:"required"`
	// Predictions exposed by this share link, in the order they were registered.
	Predictions []ShareLinkGetResponsePredictionUnion `json:"predictions" api:"required"`
	// Workspace that owns the share link and the referenced resources.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Pipelines   respjson.Field
		Predictions respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnion contains all possible properties and values
// from [ShareLinkGetResponsePipelineProteinDesignRun],
// [ShareLinkGetResponsePipelineProteinRedesignRun],
// [ShareLinkGetResponsePipelineProteinLibraryScreen],
// [ShareLinkGetResponsePipelineSmDesignRun],
// [ShareLinkGetResponsePipelineSmScreen].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineUnion struct {
	ID            string    `json:"id"`
	CompletedAt   time.Time `json:"completed_at"`
	CreatedAt     time.Time `json:"created_at"`
	DataDeletedAt time.Time `json:"data_deleted_at"`
	Engine        string    `json:"engine"`
	EngineVersion string    `json:"engine_version"`
	// This field is a union of [ShareLinkGetResponsePipelineProteinDesignRunError],
	// [ShareLinkGetResponsePipelineProteinRedesignRunError],
	// [ShareLinkGetResponsePipelineProteinLibraryScreenError],
	// [ShareLinkGetResponsePipelineSmDesignRunError],
	// [ShareLinkGetResponsePipelineSmScreenError]
	Error ShareLinkGetResponsePipelineUnionError `json:"error"`
	// This field is a union of [ShareLinkGetResponsePipelineProteinDesignRunInput],
	// [ShareLinkGetResponsePipelineProteinRedesignRunInput],
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInput],
	// [ShareLinkGetResponsePipelineSmDesignRunInput],
	// [ShareLinkGetResponsePipelineSmScreenInput]
	Input           ShareLinkGetResponsePipelineUnionInput `json:"input"`
	Livemode        bool                                   `json:"livemode"`
	Pipeline        string                                 `json:"pipeline"`
	PipelineVersion string                                 `json:"pipeline_version"`
	// This field is a union of [ShareLinkGetResponsePipelineProteinDesignRunProgress],
	// [ShareLinkGetResponsePipelineProteinRedesignRunProgress],
	// [ShareLinkGetResponsePipelineProteinLibraryScreenProgress],
	// [ShareLinkGetResponsePipelineSmDesignRunProgress],
	// [ShareLinkGetResponsePipelineSmScreenProgress]
	Progress       ShareLinkGetResponsePipelineUnionProgress `json:"progress"`
	StartedAt      time.Time                                 `json:"started_at"`
	Status         string                                    `json:"status"`
	StoppedAt      time.Time                                 `json:"stopped_at"`
	WorkspaceID    string                                    `json:"workspace_id"`
	IdempotencyKey string                                    `json:"idempotency_key"`
	JSON           struct {
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
		raw             string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineUnion) AsShareLinkGetResponsePipelineProteinDesignRun() (v ShareLinkGetResponsePipelineProteinDesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineUnion) AsShareLinkGetResponsePipelineProteinRedesignRun() (v ShareLinkGetResponsePipelineProteinRedesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineUnion) AsShareLinkGetResponsePipelineProteinLibraryScreen() (v ShareLinkGetResponsePipelineProteinLibraryScreen) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineUnion) AsShareLinkGetResponsePipelineSmDesignRun() (v ShareLinkGetResponsePipelineSmDesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineUnion) AsShareLinkGetResponsePipelineSmScreen() (v ShareLinkGetResponsePipelineSmScreen) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkGetResponsePipelineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionError is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion]. ShareLinkGetResponsePipelineUnionError
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details"`
	JSON    struct {
		Code    respjson.Field
		Message respjson.Field
		Details respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInput is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion]. ShareLinkGetResponsePipelineUnionInput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionInput struct {
	// This field is from variant [ShareLinkGetResponsePipelineProteinDesignRunInput].
	BinderSpecification ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion `json:"binder_specification"`
	NumProteins         int64                                                                     `json:"num_proteins"`
	// This field is a union of
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion],
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion],
	// [ShareLinkGetResponsePipelineSmDesignRunInputTarget],
	// [ShareLinkGetResponsePipelineSmScreenInputTarget]
	Target         ShareLinkGetResponsePipelineUnionInputTarget `json:"target"`
	IdempotencyKey string                                       `json:"idempotency_key"`
	WorkspaceID    string                                       `json:"workspace_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinRedesignRunInput].
	ChainsConfig ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfig `json:"chains_config"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinRedesignRunInput].
	Structure ShareLinkGetResponsePipelineProteinRedesignRunInputStructure `json:"structure"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinRedesignRunInput].
	Rules ShareLinkGetResponsePipelineProteinRedesignRunInputRules `json:"rules"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInput].
	Proteins ShareLinkGetResponsePipelineProteinLibraryScreenInputProteins `json:"proteins"`
	// This field is from variant [ShareLinkGetResponsePipelineSmDesignRunInput].
	NumMolecules int64 `json:"num_molecules"`
	// This field is from variant [ShareLinkGetResponsePipelineSmDesignRunInput].
	ChemicalSpace string `json:"chemical_space"`
	// This field is a union of
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFilters],
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFilters]
	MoleculeFilters ShareLinkGetResponsePipelineUnionInputMoleculeFilters `json:"molecule_filters"`
	// This field is from variant [ShareLinkGetResponsePipelineSmScreenInput].
	Molecules ShareLinkGetResponsePipelineSmScreenInputMolecules `json:"molecules"`
	JSON      struct {
		BinderSpecification respjson.Field
		NumProteins         respjson.Field
		Target              respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ChainsConfig        respjson.Field
		Structure           respjson.Field
		Rules               respjson.Field
		Proteins            respjson.Field
		NumMolecules        respjson.Field
		ChemicalSpace       respjson.Field
		MoleculeFilters     respjson.Field
		Molecules           respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTarget is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTarget provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionInputTarget struct {
	// This field is a union of
	// [map[string]ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion],
	// [map[string]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion]
	ChainSelection ShareLinkGetResponsePipelineUnionInputTargetChainSelection `json:"chain_selection"`
	// This field is a union of
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure],
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure]
	Structure ShareLinkGetResponsePipelineUnionInputTargetStructure `json:"structure"`
	Type      string                                                `json:"type"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion],
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion],
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity],
	// [[]ShareLinkGetResponsePipelineSmScreenInputTargetEntity]
	Entities ShareLinkGetResponsePipelineUnionInputTargetEntities `json:"entities"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond],
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond],
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputTargetBond],
	// [[]ShareLinkGetResponsePipelineSmScreenInputTargetBond]
	Bonds ShareLinkGetResponsePipelineUnionInputTargetBonds `json:"bonds"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion],
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion],
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion],
	// [[]ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion]
	Constraints         ShareLinkGetResponsePipelineUnionInputTargetConstraints `json:"constraints"`
	EpitopeLigandChains []string                                                `json:"epitope_ligand_chains"`
	EpitopeResidues     []int64                                                 `json:"epitope_residues"`
	NonBindingResidues  []int64                                                 `json:"non_binding_residues"`
	PocketResidues      []int64                                                 `json:"pocket_residues"`
	ReferenceLigands    []string                                                `json:"reference_ligands"`
	JSON                struct {
		ChainSelection      respjson.Field
		Structure           respjson.Field
		Type                respjson.Field
		Entities            respjson.Field
		Bonds               respjson.Field
		Constraints         respjson.Field
		EpitopeLigandChains respjson.Field
		EpitopeResidues     respjson.Field
		NonBindingResidues  respjson.Field
		PocketResidues      respjson.Field
		ReferenceLigands    respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTargetChainSelection is an implicit
// subunion of [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTargetChainSelection provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionInputTargetChainSelection struct {
	ChainType string `json:"chain_type"`
	// This field is a union of
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion],
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion]
	CropResidues       ShareLinkGetResponsePipelineUnionInputTargetChainSelectionCropResidues `json:"crop_residues"`
	EpitopeResidues    []int64                                                                `json:"epitope_residues"`
	FlexibleResidues   []int64                                                                `json:"flexible_residues"`
	NonBindingResidues []int64                                                                `json:"non_binding_residues"`
	JSON               struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		raw                string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputTargetChainSelection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTargetChainSelectionCropResidues is an
// implicit subunion of [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTargetChainSelectionCropResidues provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkGetResponsePipelineUnionInputTargetChainSelectionCropResidues struct {
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

func (r *ShareLinkGetResponsePipelineUnionInputTargetChainSelectionCropResidues) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTargetStructure is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTargetStructure provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionInputTargetStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputTargetStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTargetEntities is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTargetEntities provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntities
// OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntities
// OfShareLinkGetResponsePipelineSmDesignRunInputTargetEntities
// OfShareLinkGetResponsePipelineSmScreenInputTargetEntities]
type ShareLinkGetResponsePipelineUnionInputTargetEntities struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntities []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntities []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity] instead of an
	// object.
	OfShareLinkGetResponsePipelineSmDesignRunInputTargetEntities []ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmScreenInputTargetEntity] instead of an object.
	OfShareLinkGetResponsePipelineSmScreenInputTargetEntities []ShareLinkGetResponsePipelineSmScreenInputTargetEntity `json:",inline"`
	JSON                                                      struct {
		OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntities     respjson.Field
		OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntities respjson.Field
		OfShareLinkGetResponsePipelineSmDesignRunInputTargetEntities                                  respjson.Field
		OfShareLinkGetResponsePipelineSmScreenInputTargetEntities                                     respjson.Field
		raw                                                                                           string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputTargetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTargetBonds is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTargetBonds provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBonds
// OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBonds
// OfShareLinkGetResponsePipelineSmDesignRunInputTargetBonds
// OfShareLinkGetResponsePipelineSmScreenInputTargetBonds]
type ShareLinkGetResponsePipelineUnionInputTargetBonds struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBonds []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBonds []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputTargetBond] instead of an object.
	OfShareLinkGetResponsePipelineSmDesignRunInputTargetBonds []ShareLinkGetResponsePipelineSmDesignRunInputTargetBond `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmScreenInputTargetBond] instead of an object.
	OfShareLinkGetResponsePipelineSmScreenInputTargetBonds []ShareLinkGetResponsePipelineSmScreenInputTargetBond `json:",inline"`
	JSON                                                   struct {
		OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBonds     respjson.Field
		OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBonds respjson.Field
		OfShareLinkGetResponsePipelineSmDesignRunInputTargetBonds                                  respjson.Field
		OfShareLinkGetResponsePipelineSmScreenInputTargetBonds                                     respjson.Field
		raw                                                                                        string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputTargetBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputTargetConstraints is an implicit subunion
// of [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputTargetConstraints provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraints
// OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraints
// OfShareLinkGetResponsePipelineSmDesignRunInputTargetConstraints
// OfShareLinkGetResponsePipelineSmScreenInputTargetConstraints]
type ShareLinkGetResponsePipelineUnionInputTargetConstraints struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraints []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraints []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion] instead of
	// an object.
	OfShareLinkGetResponsePipelineSmDesignRunInputTargetConstraints []ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion] instead of an
	// object.
	OfShareLinkGetResponsePipelineSmScreenInputTargetConstraints []ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion `json:",inline"`
	JSON                                                         struct {
		OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraints     respjson.Field
		OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraints respjson.Field
		OfShareLinkGetResponsePipelineSmDesignRunInputTargetConstraints                                  respjson.Field
		OfShareLinkGetResponsePipelineSmScreenInputTargetConstraints                                     respjson.Field
		raw                                                                                              string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputTargetConstraints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputMoleculeFilters is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputMoleculeFilters provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionInputMoleculeFilters struct {
	BoltzSmartsCatalogFilterLevel string `json:"boltz_smarts_catalog_filter_level"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion],
	// [[]ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion]
	CustomFilters ShareLinkGetResponsePipelineUnionInputMoleculeFiltersCustomFilters `json:"custom_filters"`
	JSON          struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		raw                           string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionInputMoleculeFiltersCustomFilters is an
// implicit subunion of [ShareLinkGetResponsePipelineUnion].
// ShareLinkGetResponsePipelineUnionInputMoleculeFiltersCustomFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilters
// OfShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilters]
type ShareLinkGetResponsePipelineUnionInputMoleculeFiltersCustomFilters struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion]
	// instead of an object.
	OfShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilters []ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion]
	// instead of an object.
	OfShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilters []ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion `json:",inline"`
	JSON                                                                    struct {
		OfShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilters respjson.Field
		OfShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilters    respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionInputMoleculeFiltersCustomFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineUnionProgress is an implicit subunion of
// [ShareLinkGetResponsePipelineUnion]. ShareLinkGetResponsePipelineUnionProgress
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineUnion].
type ShareLinkGetResponsePipelineUnionProgress struct {
	NumProteinsGenerated    int64  `json:"num_proteins_generated"`
	TotalProteinsToGenerate int64  `json:"total_proteins_to_generate"`
	LatestResultID          string `json:"latest_result_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenProgress].
	NumProteinsFailed int64 `json:"num_proteins_failed"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenProgress].
	NumProteinsScreened int64 `json:"num_proteins_screened"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenProgress].
	TotalProteinsToScreen int64 `json:"total_proteins_to_screen"`
	// This field is from variant [ShareLinkGetResponsePipelineSmDesignRunProgress].
	NumMoleculesGenerated int64 `json:"num_molecules_generated"`
	// This field is from variant [ShareLinkGetResponsePipelineSmDesignRunProgress].
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate"`
	// This field is from variant [ShareLinkGetResponsePipelineSmScreenProgress].
	NumMoleculesFailed int64 `json:"num_molecules_failed"`
	// This field is from variant [ShareLinkGetResponsePipelineSmScreenProgress].
	NumMoleculesScreened int64 `json:"num_molecules_screened"`
	// This field is from variant [ShareLinkGetResponsePipelineSmScreenProgress].
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen"`
	// This field is from variant [ShareLinkGetResponsePipelineSmScreenProgress].
	RejectionSummary ShareLinkGetResponsePipelineSmScreenProgressRejectionSummary `json:"rejection_summary"`
	JSON             struct {
		NumProteinsGenerated     respjson.Field
		TotalProteinsToGenerate  respjson.Field
		LatestResultID           respjson.Field
		NumProteinsFailed        respjson.Field
		NumProteinsScreened      respjson.Field
		TotalProteinsToScreen    respjson.Field
		NumMoleculesGenerated    respjson.Field
		TotalMoleculesToGenerate respjson.Field
		NumMoleculesFailed       respjson.Field
		NumMoleculesScreened     respjson.Field
		TotalMoleculesToScreen   respjson.Field
		RejectionSummary         respjson.Field
		raw                      string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineUnionProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A protein design pipeline run that generates novel protein binders
type ShareLinkGetResponsePipelineProteinDesignRun struct {
	// Unique ProteinDesignRun identifier
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
	EngineVersion constant.String1_0                                `json:"engine_version" default:"1.0"`
	Error         ShareLinkGetResponsePipelineProteinDesignRunError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkGetResponsePipelineProteinDesignRunInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein design
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein design
	PipelineVersion constant.String1_0                                   `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkGetResponsePipelineProteinDesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                            `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkGetResponsePipelineProteinDesignRunStatus `json:"status" api:"required"`
	StoppedAt time.Time                                          `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinDesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunError struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinDesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkGetResponsePipelineProteinDesignRunInput struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, or boltz_curated for
	// Boltz-managed nanobody and antibody defaults.
	BinderSpecification ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion `json:"binder_specification" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion `json:"target" api:"required"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinDesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse].
	ChainSelection map[string]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection"`
	Modality       string                                                                                                                                `json:"modality"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse].
	Structure ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure"`
	Type      string                                                                                                           `json:"type"`
	// This field is a union of
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules],
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules],
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules]
	Rules ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules `json:"rules"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse].
	Entities []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse].
	Bonds []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse].
	Binder ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder"`
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules
// is an implicit subunion of
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion].
// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnion].
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules struct {
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

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// pipeline run.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the pipeline run — any chains omitted from
	// this mapping are ignored. Each value defines which residues to keep
	// (crop_residues). Omit design_motifs to include the chain as fixed scaffold
	// context.
	ChainSelection map[string]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality  ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality  `json:"modality" api:"required"`
	Structure ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure" api:"required"`
	Type      constant.StructureTemplate                                                                                       `json:"type" default:"structure_template"`
	// Constraints applied during sequence design
	Rules ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	CropResidues ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	DesignMotifs []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	JSON         struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	// This field is a union of
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange],
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange]
	DesignLengthRange ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange `json:"design_length_range"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	EndIndex int64 `json:"end_index"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	StartIndex int64  `json:"start_index"`
	Type       string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// is an implicit subunion of
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange struct {
	Max  int64 `json:"max"`
	Min  int64 `json:"min"`
	JSON struct {
		Max respjson.Field
		Min respjson.Field
		raw string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range" api:"required"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range" api:"required"`
	Type              constant.Insertion                                                                                                                                                                               `json:"type" default:"insertion"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality string

const (
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityPeptide       ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "peptide"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityAntibody      ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "antibody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityNanobody      ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "nanobody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityCustomProtein ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "custom_protein"
)

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality `json:"modality" api:"required"`
	Type     constant.NoTemplate                                                                                      `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// Constraints applied during sequence design
	Rules ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	Modifications ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications]
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                                        struct {
		OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications respjson.Field
		OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications    respjson.Field
		OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications        respjson.Field
		OfShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications        respjson.Field
		raw                                                                                                                                                string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protein binder entity with designed and/or fixed segments.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse struct {
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
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality string

const (
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityPeptide       ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "peptide"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityAntibody      ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "antibody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityNanobody      ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "nanobody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityCustomProtein ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "custom_protein"
)

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder" api:"required"`
	Type   constant.BoltzCurated                                                                                    `json:"type" default:"boltz_curated"`
	// Constraints applied during sequence design
	Rules ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder string

const (
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzNanobody ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_nanobody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzAntibody ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBinder string

const (
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBinderBoltzNanobody ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBinder = "boltz_nanobody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBinderBoltzAntibody ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationBinder = "boltz_antibody"
)

type ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModality string

const (
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModalityPeptide       ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModality = "peptide"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModalityAntibody      ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModality = "antibody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModalityNanobody      ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModality = "nanobody"
	ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModalityCustomProtein ShareLinkGetResponsePipelineProteinDesignRunInputBinderSpecificationModality = "custom_protein"
)

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse].
	Structure ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                                          `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	Entities []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	Bonds []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	Constraints []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the pipeline run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                                           `json:"type" default:"structure_template"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the pipeline run.
	CropResidues ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                                          `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion].
// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                  struct {
		OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                                       string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                              `json:"type" default:"contact"`
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunProgress struct {
	// Number of protein binders generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein binders requested
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
func (r ShareLinkGetResponsePipelineProteinDesignRunProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinDesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinDesignRunStatus string

const (
	ShareLinkGetResponsePipelineProteinDesignRunStatusPending   ShareLinkGetResponsePipelineProteinDesignRunStatus = "pending"
	ShareLinkGetResponsePipelineProteinDesignRunStatusRunning   ShareLinkGetResponsePipelineProteinDesignRunStatus = "running"
	ShareLinkGetResponsePipelineProteinDesignRunStatusSucceeded ShareLinkGetResponsePipelineProteinDesignRunStatus = "succeeded"
	ShareLinkGetResponsePipelineProteinDesignRunStatusFailed    ShareLinkGetResponsePipelineProteinDesignRunStatus = "failed"
	ShareLinkGetResponsePipelineProteinDesignRunStatusStopped   ShareLinkGetResponsePipelineProteinDesignRunStatus = "stopped"
)

// A protein redesign pipeline run that inverse-folds and scores redesigned binder
// candidates
type ShareLinkGetResponsePipelineProteinRedesignRun struct {
	// Unique ProteinRedesignRun identifier
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
	EngineVersion constant.V2026_03_01                                `json:"engine_version" default:"v2026-03-01"`
	Error         ShareLinkGetResponsePipelineProteinRedesignRunError `json:"error" api:"required"`
	// Input for protein redesign from a complete target/binder CIF complex.
	Input ShareLinkGetResponsePipelineProteinRedesignRunInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein redesign
	Pipeline constant.BoltzProteinRedesign `json:"pipeline" default:"boltz-protein-redesign"`
	// Pipeline version used for protein redesign
	PipelineVersion constant.V2026_03_01                                   `json:"pipeline_version" default:"v2026-03-01"`
	Progress        ShareLinkGetResponsePipelineProteinRedesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                              `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkGetResponsePipelineProteinRedesignRunStatus `json:"status" api:"required"`
	StoppedAt time.Time                                            `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkGetResponsePipelineProteinRedesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinRedesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinRedesignRunError struct {
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
func (r ShareLinkGetResponsePipelineProteinRedesignRunError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinRedesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for protein redesign from a complete target/binder CIF complex.
type ShareLinkGetResponsePipelineProteinRedesignRunInput struct {
	// Complete chain assignment for the input CIF. Every CIF chain must appear exactly
	// once, either under target_chains or binder_chains.
	ChainsConfig ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfig `json:"chains_config" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                        `json:"num_proteins" api:"required"`
	Structure   ShareLinkGetResponsePipelineProteinRedesignRunInputStructure `json:"structure" api:"required"`
	// Optional idempotency key.
	IdempotencyKey string `json:"idempotency_key"`
	// Constraints applied during sequence design
	Rules ShareLinkGetResponsePipelineProteinRedesignRunInputRules `json:"rules"`
	// Workspace to run this redesign in.
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainsConfig   respjson.Field
		NumProteins    respjson.Field
		Structure      respjson.Field
		IdempotencyKey respjson.Field
		Rules          respjson.Field
		WorkspaceID    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinRedesignRunInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete chain assignment for the input CIF. Every CIF chain must appear exactly
// once, either under target_chains or binder_chains.
type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfig struct {
	// Binder chains keyed by CIF chain ID. Protein binder chains may provide
	// designed_residues; at least 5 total unique designed residues are required across
	// all protein binder chains. Target and binder chains must be disjoint and
	// together cover every chain in the CIF.
	BinderChains map[string]ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion `json:"binder_chains" api:"required"`
	// Target chains keyed by CIF chain ID. Target and binder chains must be disjoint
	// and together cover every chain in the CIF.
	TargetChains map[string]ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChain `json:"target_chains" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChains respjson.Field
		TargetChains respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig],
// [ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig].
	DesignedResidues []int64 `json:"designed_residues"`
	JSON             struct {
		ChainType        respjson.Field
		DesignedResidues respjson.Field
		raw              string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) AsShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig() (v ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) AsShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig() (v ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a protein binder chain. Across all protein binder chains, at
// least 5 unique designed residues are required.
type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig struct {
	ChainType constant.Protein `json:"chain_type" default:"protein"`
	// 0-indexed residue indices to redesign for this protein binder chain.
	DesignedResidues []int64 `json:"designed_residues"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType        respjson.Field
		DesignedResidues respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a fixed non-protein binder chain in the input CIF.
type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig struct {
	// Any of "ligand", "rna", "dna".
	ChainType ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType `json:"chain_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType string

const (
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainTypeLigand ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType = "ligand"
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainTypeRna    ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType = "rna"
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainTypeDna    ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType = "dna"
)

// Configuration for a fixed target chain in the input CIF.
type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChain struct {
	// Molecule type for a chain in the input CIF.
	//
	// Any of "ligand", "protein", "rna", "dna".
	ChainType ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType `json:"chain_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChain) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule type for a chain in the input CIF.
type ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType string

const (
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeLigand  ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "ligand"
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeProtein ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "protein"
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeRna     ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "rna"
	ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeDna     ShareLinkGetResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "dna"
)

type ShareLinkGetResponsePipelineProteinRedesignRunInputStructure struct {
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
func (r ShareLinkGetResponsePipelineProteinRedesignRunInputStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkGetResponsePipelineProteinRedesignRunInputRules struct {
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
func (r ShareLinkGetResponsePipelineProteinRedesignRunInputRules) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinRedesignRunInputRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinRedesignRunProgress struct {
	// Number of protein binders generated so far
	NumProteinsGenerated int64 `json:"num_proteins_generated" api:"required"`
	// Total number of protein binders requested
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
func (r ShareLinkGetResponsePipelineProteinRedesignRunProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinRedesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinRedesignRunStatus string

const (
	ShareLinkGetResponsePipelineProteinRedesignRunStatusPending   ShareLinkGetResponsePipelineProteinRedesignRunStatus = "pending"
	ShareLinkGetResponsePipelineProteinRedesignRunStatusRunning   ShareLinkGetResponsePipelineProteinRedesignRunStatus = "running"
	ShareLinkGetResponsePipelineProteinRedesignRunStatusSucceeded ShareLinkGetResponsePipelineProteinRedesignRunStatus = "succeeded"
	ShareLinkGetResponsePipelineProteinRedesignRunStatusFailed    ShareLinkGetResponsePipelineProteinRedesignRunStatus = "failed"
	ShareLinkGetResponsePipelineProteinRedesignRunStatusStopped   ShareLinkGetResponsePipelineProteinRedesignRunStatus = "stopped"
)

// A protein library screening pipeline run
type ShareLinkGetResponsePipelineProteinLibraryScreen struct {
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
	EngineVersion constant.String1_0                                    `json:"engine_version" default:"1.0"`
	Error         ShareLinkGetResponsePipelineProteinLibraryScreenError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkGetResponsePipelineProteinLibraryScreenInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                                       `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkGetResponsePipelineProteinLibraryScreenProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                                `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkGetResponsePipelineProteinLibraryScreenStatus `json:"status" api:"required"`
	StoppedAt time.Time                                              `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreen) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinLibraryScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenError struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkGetResponsePipelineProteinLibraryScreenInput struct {
	Proteins ShareLinkGetResponsePipelineProteinLibraryScreenInputProteins `json:"proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Proteins    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputProteins struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputProteins) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputProteins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse].
	Structure ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                                              `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	Entities []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	Bonds []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	Constraints []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
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

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the pipeline run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                                               `json:"type" default:"structure_template"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
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

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the pipeline run.
	CropResidues ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                                              `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion].
// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                      struct {
		OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                                           string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                                  `json:"type" default:"contact"`
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenProgress struct {
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
func (r ShareLinkGetResponsePipelineProteinLibraryScreenProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineProteinLibraryScreenProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineProteinLibraryScreenStatus string

const (
	ShareLinkGetResponsePipelineProteinLibraryScreenStatusPending   ShareLinkGetResponsePipelineProteinLibraryScreenStatus = "pending"
	ShareLinkGetResponsePipelineProteinLibraryScreenStatusRunning   ShareLinkGetResponsePipelineProteinLibraryScreenStatus = "running"
	ShareLinkGetResponsePipelineProteinLibraryScreenStatusSucceeded ShareLinkGetResponsePipelineProteinLibraryScreenStatus = "succeeded"
	ShareLinkGetResponsePipelineProteinLibraryScreenStatusFailed    ShareLinkGetResponsePipelineProteinLibraryScreenStatus = "failed"
	ShareLinkGetResponsePipelineProteinLibraryScreenStatusStopped   ShareLinkGetResponsePipelineProteinLibraryScreenStatus = "stopped"
)

// A small molecule design pipeline run that generates novel molecules
type ShareLinkGetResponsePipelineSmDesignRun struct {
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
	EngineVersion constant.String1_0                           `json:"engine_version" default:"1.0"`
	Error         ShareLinkGetResponsePipelineSmDesignRunError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkGetResponsePipelineSmDesignRunInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                              `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkGetResponsePipelineSmDesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                       `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkGetResponsePipelineSmDesignRunStatus `json:"status" api:"required"`
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
func (r ShareLinkGetResponsePipelineSmDesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunError struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkGetResponsePipelineSmDesignRunInput struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target ShareLinkGetResponsePipelineSmDesignRunInputTarget `json:"target" api:"required"`
	// Chemical space to constrain generated molecules. Currently only 'enamine_real'
	// (Enamine REAL chemical space) is supported. Additional options may be added in
	// the future.
	//
	// Any of "enamine_real".
	ChemicalSpace string `json:"chemical_space"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFilters `json:"molecule_filters"`
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
func (r ShareLinkGetResponsePipelineSmDesignRunInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type ShareLinkGetResponsePipelineSmDesignRunInputTarget struct {
	// Protein entities defining the target structure. Each entity represents a protein
	// chain.
	Entities []ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ShareLinkGetResponsePipelineSmDesignRunInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion `json:"constraints"`
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTarget) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineSmDesignRunInputTargetEntityModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetEntityModification struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse].
	Token1 ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse].
	Token2 ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion) AsShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion) AsShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                                 `json:"patterns"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunProgress struct {
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
func (r ShareLinkGetResponsePipelineSmDesignRunProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmDesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmDesignRunStatus string

const (
	ShareLinkGetResponsePipelineSmDesignRunStatusPending   ShareLinkGetResponsePipelineSmDesignRunStatus = "pending"
	ShareLinkGetResponsePipelineSmDesignRunStatusRunning   ShareLinkGetResponsePipelineSmDesignRunStatus = "running"
	ShareLinkGetResponsePipelineSmDesignRunStatusSucceeded ShareLinkGetResponsePipelineSmDesignRunStatus = "succeeded"
	ShareLinkGetResponsePipelineSmDesignRunStatusFailed    ShareLinkGetResponsePipelineSmDesignRunStatus = "failed"
	ShareLinkGetResponsePipelineSmDesignRunStatusStopped   ShareLinkGetResponsePipelineSmDesignRunStatus = "stopped"
)

// A small molecule library screening pipeline run
type ShareLinkGetResponsePipelineSmScreen struct {
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
	EngineVersion constant.String1_0                        `json:"engine_version" default:"1.0"`
	Error         ShareLinkGetResponsePipelineSmScreenError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkGetResponsePipelineSmScreenInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                           `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkGetResponsePipelineSmScreenProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                    `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkGetResponsePipelineSmScreenStatus `json:"status" api:"required"`
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
func (r ShareLinkGetResponsePipelineSmScreen) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenError struct {
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
func (r ShareLinkGetResponsePipelineSmScreenError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkGetResponsePipelineSmScreenInput struct {
	Molecules ShareLinkGetResponsePipelineSmScreenInputMolecules `json:"molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target ShareLinkGetResponsePipelineSmScreenInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters ShareLinkGetResponsePipelineSmScreenInputMoleculeFilters `json:"molecule_filters"`
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
func (r ShareLinkGetResponsePipelineSmScreenInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenInputMolecules struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMolecules) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenInputMolecules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type ShareLinkGetResponsePipelineSmScreenInputTarget struct {
	// Protein entities defining the target structure. Each entity represents a protein
	// chain.
	Entities []ShareLinkGetResponsePipelineSmScreenInputTargetEntity `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Atom-level ligand
	// references currently support ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ShareLinkGetResponsePipelineSmScreenInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints []ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion `json:"constraints"`
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
func (r ShareLinkGetResponsePipelineSmScreenInputTarget) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenInputTargetEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePipelineSmScreenInputTargetEntityModification `json:"modifications"`
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetEntity) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePipelineSmScreenInputTargetEntityModification struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmScreenInputTargetBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineSmScreenInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse],
// [ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union) AsShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union) AsShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse],
// [ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union) AsShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union) AsShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse],
// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse].
	Token1 ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse].
	Token2 ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion) AsShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion) AsShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFilters) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                              `json:"patterns"`
	// This field is from variant
	// [ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                       `json:"latest_result_id"`
	RejectionSummary ShareLinkGetResponsePipelineSmScreenProgressRejectionSummary `json:"rejection_summary"`
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
func (r ShareLinkGetResponsePipelineSmScreenProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePipelineSmScreenProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenProgressRejectionSummary struct {
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
func (r ShareLinkGetResponsePipelineSmScreenProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePipelineSmScreenProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePipelineSmScreenStatus string

const (
	ShareLinkGetResponsePipelineSmScreenStatusPending   ShareLinkGetResponsePipelineSmScreenStatus = "pending"
	ShareLinkGetResponsePipelineSmScreenStatusRunning   ShareLinkGetResponsePipelineSmScreenStatus = "running"
	ShareLinkGetResponsePipelineSmScreenStatusSucceeded ShareLinkGetResponsePipelineSmScreenStatus = "succeeded"
	ShareLinkGetResponsePipelineSmScreenStatusFailed    ShareLinkGetResponsePipelineSmScreenStatus = "failed"
	ShareLinkGetResponsePipelineSmScreenStatusStopped   ShareLinkGetResponsePipelineSmScreenStatus = "stopped"
)

type ShareLinkGetResponsePipelineStatus string

const (
	ShareLinkGetResponsePipelineStatusPending   ShareLinkGetResponsePipelineStatus = "pending"
	ShareLinkGetResponsePipelineStatusRunning   ShareLinkGetResponsePipelineStatus = "running"
	ShareLinkGetResponsePipelineStatusSucceeded ShareLinkGetResponsePipelineStatus = "succeeded"
	ShareLinkGetResponsePipelineStatusFailed    ShareLinkGetResponsePipelineStatus = "failed"
	ShareLinkGetResponsePipelineStatusStopped   ShareLinkGetResponsePipelineStatus = "stopped"
)

// ShareLinkGetResponsePredictionUnion contains all possible properties and values
// from [ShareLinkGetResponsePredictionBoltz2Prediction],
// [ShareLinkGetResponsePredictionAdmePrediction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionUnion struct {
	ID            string    `json:"id"`
	CompletedAt   time.Time `json:"completed_at"`
	CreatedAt     time.Time `json:"created_at"`
	DataDeletedAt time.Time `json:"data_deleted_at"`
	// This field is a union of [ShareLinkGetResponsePredictionBoltz2PredictionError],
	// [ShareLinkGetResponsePredictionAdmePredictionError]
	Error     ShareLinkGetResponsePredictionUnionError `json:"error"`
	ExpiresAt time.Time                                `json:"expires_at"`
	// This field is a union of [ShareLinkGetResponsePredictionBoltz2PredictionInput],
	// [ShareLinkGetResponsePredictionAdmePredictionInput]
	Input    ShareLinkGetResponsePredictionUnionInput `json:"input"`
	Livemode bool                                     `json:"livemode"`
	Model    string                                   `json:"model"`
	// This field is a union of [ShareLinkGetResponsePredictionBoltz2PredictionOutput],
	// [ShareLinkGetResponsePredictionAdmePredictionOutput]
	Output         ShareLinkGetResponsePredictionUnionOutput `json:"output"`
	StartedAt      time.Time                                 `json:"started_at"`
	Status         string                                    `json:"status"`
	Version        string                                    `json:"version"`
	WorkspaceID    string                                    `json:"workspace_id"`
	IdempotencyKey string                                    `json:"idempotency_key"`
	JSON           struct {
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
		raw            string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionUnion) AsShareLinkGetResponsePredictionBoltz2Prediction() (v ShareLinkGetResponsePredictionBoltz2Prediction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionUnion) AsShareLinkGetResponsePredictionAdmePrediction() (v ShareLinkGetResponsePredictionAdmePrediction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkGetResponsePredictionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionUnionError is an implicit subunion of
// [ShareLinkGetResponsePredictionUnion]. ShareLinkGetResponsePredictionUnionError
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePredictionUnion].
type ShareLinkGetResponsePredictionUnionError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details"`
	JSON    struct {
		Code    respjson.Field
		Message respjson.Field
		Details respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePredictionUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionUnionInput is an implicit subunion of
// [ShareLinkGetResponsePredictionUnion]. ShareLinkGetResponsePredictionUnionInput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePredictionUnion].
type ShareLinkGetResponsePredictionUnionInput struct {
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	Entities []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	Binding ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion `json:"binding"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	Bonds []ShareLinkGetResponsePredictionBoltz2PredictionInputBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	Constraints []ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	ModelOptions ShareLinkGetResponsePredictionBoltz2PredictionInputModelOptions `json:"model_options"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInput].
	Templates []ShareLinkGetResponsePredictionBoltz2PredictionInputTemplate `json:"templates"`
	// This field is from variant [ShareLinkGetResponsePredictionAdmePredictionInput].
	Molecules []ShareLinkGetResponsePredictionAdmePredictionInputMolecule `json:"molecules"`
	JSON      struct {
		Entities     respjson.Field
		Binding      respjson.Field
		Bonds        respjson.Field
		Constraints  respjson.Field
		ModelOptions respjson.Field
		NumSamples   respjson.Field
		Templates    respjson.Field
		Molecules    respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePredictionUnionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionUnionOutput is an implicit subunion of
// [ShareLinkGetResponsePredictionUnion]. ShareLinkGetResponsePredictionUnionOutput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePredictionUnion].
type ShareLinkGetResponsePredictionUnionOutput struct {
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionOutput].
	AllSampleResults []ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResult `json:"all_sample_results"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionOutput].
	BestSample ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSample `json:"best_sample"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionOutput].
	Archive ShareLinkGetResponsePredictionBoltz2PredictionOutputArchive `json:"archive"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionOutput].
	BindingMetrics ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion `json:"binding_metrics"`
	// This field is from variant [ShareLinkGetResponsePredictionAdmePredictionOutput].
	Molecules []ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion `json:"molecules"`
	JSON      struct {
		AllSampleResults respjson.Field
		BestSample       respjson.Field
		Archive          respjson.Field
		BindingMetrics   respjson.Field
		Molecules        respjson.Field
		raw              string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePredictionUnionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2Prediction struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error ShareLinkGetResponsePredictionBoltz2PredictionError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input ShareLinkGetResponsePredictionBoltz2PredictionInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	// Prediction output when succeeded
	Output    ShareLinkGetResponsePredictionBoltz2PredictionOutput `json:"output" api:"required"`
	StartedAt time.Time                                            `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status ShareLinkGetResponsePredictionBoltz2PredictionStatus `json:"status" api:"required"`
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
func (r ShareLinkGetResponsePredictionBoltz2Prediction) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionBoltz2Prediction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type ShareLinkGetResponsePredictionBoltz2PredictionError struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionBoltz2PredictionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type ShareLinkGetResponsePredictionBoltz2PredictionInput struct {
	// Entities (proteins, RNA, DNA, ligands) forming the complex to predict. Order
	// determines chain assignment.
	Entities []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion `json:"entities" api:"required"`
	Binding  ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion  `json:"binding"`
	// Bond constraints between atoms. Atom-level ligand references currently support
	// ligand_ccd only; ligand_smiles is unsupported.
	Bonds []ShareLinkGetResponsePredictionBoltz2PredictionInputBond `json:"bonds"`
	// Structural constraints (pocket and contact). Atom-level ligand references
	// currently support ligand_ccd only; ligand_smiles is unsupported.
	Constraints  []ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion `json:"constraints"`
	ModelOptions ShareLinkGetResponsePredictionBoltz2PredictionInputModelOptions      `json:"model_options"`
	// Number of structure samples to generate (1-10)
	NumSamples int64 `json:"num_samples"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []ShareLinkGetResponsePredictionBoltz2PredictionInputTemplate `json:"templates"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification],
	// [[]ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification],
	// [[]ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification]
	Modifications ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse].
	Msa  ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
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

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnionModifications is
// an implicit subunion of
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion].
// ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModifications
// OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModifications
// OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModifications]
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModifications []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModifications []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModifications []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                      struct {
		OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModifications respjson.Field
		OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModifications           respjson.Field
		OfShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModifications           respjson.Field
		raw                                                                                                 string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification `json:"modifications"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion
// contains all possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Format ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Source ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source"`
	Type   string                                                                                                               `json:"type"`
	JSON   struct {
		Format respjson.Field
		Source respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format" api:"required"`
	Source ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source" api:"required"`
	Type   constant.Custom                                                                                                      `json:"type" default:"custom"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat string

const (
	ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatA3m ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "a3m"
	ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatCsv ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "csv"
)

type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse struct {
	Type constant.Empty `json:"type" default:"empty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormat string

const (
	ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormatA3m ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormat = "a3m"
	ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormatCsv ShareLinkGetResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormat = "csv"
)

type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse].
	BinderChainID string `json:"binder_chain_id"`
	Type          string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse].
	BinderChainIDs []string `json:"binder_chain_ids"`
	JSON           struct {
		BinderChainID  respjson.Field
		Type           respjson.Field
		BinderChainIDs respjson.Field
		raw            string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, at most 100 heavy
	// atoms, and only ligands+proteins in entities)
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bond between two atoms. Atom-level ligand references currently support
// ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePredictionBoltz2PredictionInputBond struct {
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom1 ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union `json:"atom1" api:"required"`
	// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Atom2 ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBond) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union contains all
// possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union contains all
// possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	Type     string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse].
	ResidueIndex int64 `json:"residue_index"`
	JSON         struct {
		AtomName     respjson.Field
		ChainID      respjson.Field
		Type         respjson.Field
		ResidueIndex respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand atom reference. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion struct {
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse].
	Token1 ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse].
	Token2 ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion) AsShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact constraint between two tokens. Atom-level ligand references currently
// support ligand_ccd entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token1 ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token. Atom-level ligand references currently support ligand_ccd
	// entities only; ligand_smiles is unsupported.
	Token2 ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) AsShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token. Atom-level ligand references currently support ligand_ccd
// entities only; ligand_smiles is unsupported.
type ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputModelOptions struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputModelOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
type ShareLinkGetResponsePredictionBoltz2PredictionInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains    []ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateChain   `json:"template_chains" api:"required"`
	TemplateStructure ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateStructure `json:"template_structure" api:"required"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputTemplate) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
type ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateChain struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateChain) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateStructure struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionInputTemplateTemplateStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type ShareLinkGetResponsePredictionBoltz2PredictionOutput struct {
	// Per-sample structure results
	AllSampleResults []ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResult   `json:"all_sample_results" api:"required"`
	BestSample       ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSample          `json:"best_sample" api:"required"`
	Archive          ShareLinkGetResponsePredictionBoltz2PredictionOutputArchive             `json:"archive"`
	BindingMetrics   ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion `json:"binding_metrics"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResult struct {
	Metrics         ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics         `json:"metrics" api:"required"`
	Structure       ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultStructure       `json:"structure" api:"required"`
	LigandStructure ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure `json:"ligand_structure"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultStructure struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSample struct {
	Metrics         ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleMetrics         `json:"metrics" api:"required"`
	Structure       ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleStructure       `json:"structure" api:"required"`
	LigandStructure ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure `json:"ligand_structure"`
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSample) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleMetrics struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleStructure struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputArchive struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputArchive) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion contains
// all possible properties and values from
// [ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics],
// [ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion struct {
	BindingConfidence float64 `json:"binding_confidence"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics].
	OptimizationScore float64 `json:"optimization_score"`
	Type              string  `json:"type"`
	JSON              struct {
		BindingConfidence respjson.Field
		OptimizationScore respjson.Field
		Type              respjson.Field
		raw               string
	} `json:"-"`
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) AsShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics() (v ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) AsShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics() (v ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics struct {
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
func (r ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionBoltz2PredictionStatus string

const (
	ShareLinkGetResponsePredictionBoltz2PredictionStatusPending   ShareLinkGetResponsePredictionBoltz2PredictionStatus = "pending"
	ShareLinkGetResponsePredictionBoltz2PredictionStatusRunning   ShareLinkGetResponsePredictionBoltz2PredictionStatus = "running"
	ShareLinkGetResponsePredictionBoltz2PredictionStatusSucceeded ShareLinkGetResponsePredictionBoltz2PredictionStatus = "succeeded"
	ShareLinkGetResponsePredictionBoltz2PredictionStatusFailed    ShareLinkGetResponsePredictionBoltz2PredictionStatus = "failed"
)

type ShareLinkGetResponsePredictionAdmePrediction struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error ShareLinkGetResponsePredictionAdmePredictionError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input ShareLinkGetResponsePredictionAdmePredictionInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.AdmeV1 `json:"model" default:"adme-v1"`
	// Prediction output when succeeded
	Output    ShareLinkGetResponsePredictionAdmePredictionOutput `json:"output" api:"required"`
	StartedAt time.Time                                          `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status ShareLinkGetResponsePredictionAdmePredictionStatus `json:"status" api:"required"`
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
func (r ShareLinkGetResponsePredictionAdmePrediction) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionAdmePrediction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type ShareLinkGetResponsePredictionAdmePredictionError struct {
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
func (r ShareLinkGetResponsePredictionAdmePredictionError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionAdmePredictionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type ShareLinkGetResponsePredictionAdmePredictionInput struct {
	// Molecules to score (1-128 per request). Results are returned in the same order
	// as this list.
	Molecules []ShareLinkGetResponsePredictionAdmePredictionInputMolecule `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePredictionAdmePredictionInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionAdmePredictionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionAdmePredictionInputMolecule struct {
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
func (r ShareLinkGetResponsePredictionAdmePredictionInputMolecule) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionAdmePredictionInputMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type ShareLinkGetResponsePredictionAdmePredictionOutput struct {
	// Per-molecule results in the same order as the request. Successful molecules
	// carry an `adme` summary. Failed molecules carry `status: "failed"` and a
	// non-null `error`.
	Molecules []ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponsePredictionAdmePredictionOutput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponsePredictionAdmePredictionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion contains all
// possible properties and values from
// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded],
// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme],
	// [any]
	Adme ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionAdme `json:"adme"`
	// This field is a union of [any],
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError]
	Error      ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionError `json:"error"`
	Smiles     string                                                               `json:"smiles"`
	Status     string                                                               `json:"status"`
	ExternalID string                                                               `json:"external_id"`
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

func (u ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion) AsShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded() (v ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion) AsShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed() (v ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionAdme is an
// implicit subunion of
// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion].
// ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionAdme provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedAdme]
type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionAdme struct {
	// This field will be present if the value is a [any] instead of an object.
	OfShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedAdme any `json:",inline"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme].
	Lipophilicity float64 `json:"lipophilicity"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme].
	Permeability float64 `json:"permeability"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme].
	Solubility ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility"`
	JSON       struct {
		OfShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedAdme respjson.Field
		Lipophilicity                                                                      respjson.Field
		Permeability                                                                       respjson.Field
		Solubility                                                                         respjson.Field
		raw                                                                                string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionError is an
// implicit subunion of
// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion].
// ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionError provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededError]
type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionError struct {
	// This field will be present if the value is a [any] instead of an object.
	OfShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededError any `json:",inline"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError].
	Code string `json:"code"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError].
	Message string `json:"message"`
	// This field is from variant
	// [ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError].
	Details any `json:"details"`
	JSON    struct {
		OfShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededError respjson.Field
		Code                                                                                   respjson.Field
		Message                                                                                respjson.Field
		Details                                                                                respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded struct {
	// Internally generated molecule identifier.
	ID string `json:"id" api:"required"`
	// Tier 1 ADME summary values for this molecule.
	Adme  ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme `json:"adme" api:"required"`
	Error any                                                                                 `json:"error" api:"required"`
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
func (r ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tier 1 ADME summary values for this molecule.
type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme struct {
	// Lipophilicity score from the internal LogD prediction.
	Lipophilicity float64 `json:"lipophilicity" api:"required"`
	// Permeability score for this molecule.
	Permeability float64 `json:"permeability" api:"required"`
	// Solubility judgement for this molecule.
	//
	// Any of "high-confidence", "medium-confidence", "high-risk".
	Solubility ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility" api:"required"`
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
func (r ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solubility judgement for this molecule.
type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility string

const (
	ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighConfidence   ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-confidence"
	ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityMediumConfidence ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "medium-confidence"
	ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighRisk         ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-risk"
)

type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed struct {
	// Internally generated molecule identifier.
	ID    string                                                                            `json:"id" api:"required"`
	Adme  any                                                                               `json:"adme" api:"required"`
	Error ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError `json:"error" api:"required"`
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
func (r ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError struct {
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
func (r ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponsePredictionAdmePredictionStatus string

const (
	ShareLinkGetResponsePredictionAdmePredictionStatusPending   ShareLinkGetResponsePredictionAdmePredictionStatus = "pending"
	ShareLinkGetResponsePredictionAdmePredictionStatusRunning   ShareLinkGetResponsePredictionAdmePredictionStatus = "running"
	ShareLinkGetResponsePredictionAdmePredictionStatusSucceeded ShareLinkGetResponsePredictionAdmePredictionStatus = "succeeded"
	ShareLinkGetResponsePredictionAdmePredictionStatusFailed    ShareLinkGetResponsePredictionAdmePredictionStatus = "failed"
)

type ShareLinkGetResponsePredictionStatus string

const (
	ShareLinkGetResponsePredictionStatusPending   ShareLinkGetResponsePredictionStatus = "pending"
	ShareLinkGetResponsePredictionStatusRunning   ShareLinkGetResponsePredictionStatus = "running"
	ShareLinkGetResponsePredictionStatusSucceeded ShareLinkGetResponsePredictionStatus = "succeeded"
	ShareLinkGetResponsePredictionStatusFailed    ShareLinkGetResponsePredictionStatus = "failed"
)

type ShareLinkDeleteDataResponse struct {
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
func (r ShareLinkDeleteDataResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkDeleteDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single generated protein design
type ShareLinkListPipelineResultsResponse struct {
	// Unique result ID
	ID        string                                        `json:"id" api:"required"`
	Artifacts ShareLinkListPipelineResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	// Entities of the designed binder complex. Includes both designed entities and
	// fixed entities from the input.
	Entities []ShareLinkListPipelineResultsResponseEntityUnion `json:"entities" api:"required"`
	// Structural and binding quality metrics for a designed protein binder
	Metrics ShareLinkListPipelineResultsResponseMetrics `json:"metrics" api:"required"`
	// Warnings about potential quality issues with this result.
	Warnings []ShareLinkListPipelineResultsResponseWarning `json:"warnings"`
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
func (r ShareLinkListPipelineResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseArtifacts struct {
	Archive   ShareLinkListPipelineResultsResponseArtifactsArchive   `json:"archive" api:"required"`
	Structure ShareLinkListPipelineResultsResponseArtifactsStructure `json:"structure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archive     respjson.Field
		Structure   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkListPipelineResultsResponseArtifacts) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseArtifacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseArtifactsArchive struct {
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
func (r ShareLinkListPipelineResultsResponseArtifactsArchive) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseArtifactsArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseArtifactsStructure struct {
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
func (r ShareLinkListPipelineResultsResponseArtifactsStructure) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseArtifactsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkListPipelineResultsResponseEntityUnion contains all possible properties
// and values from [ShareLinkListPipelineResultsResponseEntityProteinEntity],
// [ShareLinkListPipelineResultsResponseEntityRnaEntity],
// [ShareLinkListPipelineResultsResponseEntityDnaEntity],
// [ShareLinkListPipelineResultsResponseEntityLigandCcdEntity],
// [ShareLinkListPipelineResultsResponseEntityLigandSmilesEntity].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkListPipelineResultsResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkListPipelineResultsResponseEntityProteinEntityModification],
	// [[]ShareLinkListPipelineResultsResponseEntityRnaEntityModification],
	// [[]ShareLinkListPipelineResultsResponseEntityDnaEntityModification]
	Modifications ShareLinkListPipelineResultsResponseEntityUnionModifications `json:"modifications"`
	JSON          struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkListPipelineResultsResponseEntityUnion) AsShareLinkListPipelineResultsResponseEntityProteinEntity() (v ShareLinkListPipelineResultsResponseEntityProteinEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkListPipelineResultsResponseEntityUnion) AsShareLinkListPipelineResultsResponseEntityRnaEntity() (v ShareLinkListPipelineResultsResponseEntityRnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkListPipelineResultsResponseEntityUnion) AsShareLinkListPipelineResultsResponseEntityDnaEntity() (v ShareLinkListPipelineResultsResponseEntityDnaEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkListPipelineResultsResponseEntityUnion) AsShareLinkListPipelineResultsResponseEntityLigandCcdEntity() (v ShareLinkListPipelineResultsResponseEntityLigandCcdEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkListPipelineResultsResponseEntityUnion) AsShareLinkListPipelineResultsResponseEntityLigandSmilesEntity() (v ShareLinkListPipelineResultsResponseEntityLigandSmilesEntity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkListPipelineResultsResponseEntityUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkListPipelineResultsResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkListPipelineResultsResponseEntityUnionModifications is an implicit
// subunion of [ShareLinkListPipelineResultsResponseEntityUnion].
// ShareLinkListPipelineResultsResponseEntityUnionModifications provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkListPipelineResultsResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkListPipelineResultsResponseEntityProteinEntityModifications
// OfShareLinkListPipelineResultsResponseEntityRnaEntityModifications
// OfShareLinkListPipelineResultsResponseEntityDnaEntityModifications]
type ShareLinkListPipelineResultsResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkListPipelineResultsResponseEntityProteinEntityModification] instead
	// of an object.
	OfShareLinkListPipelineResultsResponseEntityProteinEntityModifications []ShareLinkListPipelineResultsResponseEntityProteinEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkListPipelineResultsResponseEntityRnaEntityModification] instead of
	// an object.
	OfShareLinkListPipelineResultsResponseEntityRnaEntityModifications []ShareLinkListPipelineResultsResponseEntityRnaEntityModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkListPipelineResultsResponseEntityDnaEntityModification] instead of
	// an object.
	OfShareLinkListPipelineResultsResponseEntityDnaEntityModifications []ShareLinkListPipelineResultsResponseEntityDnaEntityModification `json:",inline"`
	JSON                                                               struct {
		OfShareLinkListPipelineResultsResponseEntityProteinEntityModifications respjson.Field
		OfShareLinkListPipelineResultsResponseEntityRnaEntityModifications     respjson.Field
		OfShareLinkListPipelineResultsResponseEntityDnaEntityModifications     respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (r *ShareLinkListPipelineResultsResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityProteinEntity struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkListPipelineResultsResponseEntityProteinEntityModification `json:"modifications"`
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
func (r ShareLinkListPipelineResultsResponseEntityProteinEntity) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseEntityProteinEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkListPipelineResultsResponseEntityProteinEntityModification struct {
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
func (r ShareLinkListPipelineResultsResponseEntityProteinEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityProteinEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityRnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkListPipelineResultsResponseEntityRnaEntityModification `json:"modifications"`
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
func (r ShareLinkListPipelineResultsResponseEntityRnaEntity) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseEntityRnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkListPipelineResultsResponseEntityRnaEntityModification struct {
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
func (r ShareLinkListPipelineResultsResponseEntityRnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityRnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityDnaEntity struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkListPipelineResultsResponseEntityDnaEntityModification `json:"modifications"`
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
func (r ShareLinkListPipelineResultsResponseEntityDnaEntity) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseEntityDnaEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkListPipelineResultsResponseEntityDnaEntityModification struct {
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
func (r ShareLinkListPipelineResultsResponseEntityDnaEntityModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityDnaEntityModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityLigandCcdEntity struct {
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
func (r ShareLinkListPipelineResultsResponseEntityLigandCcdEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityLigandCcdEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityLigandSmilesEntity struct {
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
func (r ShareLinkListPipelineResultsResponseEntityLigandSmilesEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityLigandSmilesEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structural and binding quality metrics for a designed protein binder
type ShareLinkListPipelineResultsResponseMetrics struct {
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
func (r ShareLinkListPipelineResultsResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warning about a potential quality issue with a result
type ShareLinkListPipelineResultsResponseWarning struct {
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
func (r ShareLinkListPipelineResultsResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkNewParams struct {
	ExpiresAt string `json:"expires_at" api:"required" format:"date-time-string"`
	// Workspace ID. Only used with admin API keys. Ignored (or validated) for
	// workspace-scoped keys.
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// Pipelines to expose through the share link. Must belong to the resolved
	// workspace. Up to 100 entries.
	PipelineIDs []string `json:"pipeline_ids,omitzero"`
	// Predictions to expose through the share link. Must belong to the resolved
	// workspace. Up to 100 entries.
	PredictionIDs []string `json:"prediction_ids,omitzero"`
	paramObj
}

func (r ShareLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ShareLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShareLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkDeleteDataParams struct {
	// Share link ID to revoke. Sent in the body — never as a URL path segment —
	// because the ID is itself the bearer credential.
	ID string `json:"id" api:"required"`
	paramObj
}

func (r ShareLinkDeleteDataParams) MarshalJSON() (data []byte, err error) {
	type shadow ShareLinkDeleteDataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShareLinkDeleteDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsParams struct {
	ID string `path:"id" api:"required" json:"-"`
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
	paramObj
}

// URLQuery serializes [ShareLinkListPipelineResultsParams]'s query parameters as
// `url.Values`.
func (r ShareLinkListPipelineResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
