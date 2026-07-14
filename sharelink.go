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
// links that visitors can open without an API key or, for email-restricted links,
// after signing in with an allowed email. A share link is scoped to a single
// workspace and bundles one or more predictions and pipeline runs. The link ID is
// itself the bearer credential; treat it as a secret. Create, retrieve, and
// archive require a workspace-scoped API key with read permission on every
// referenced resource. Retrieving metadata remains available after expiry or
// archive. Viewing content and listing shared pipeline results are gated by the
// link ID and the link's access mode. Archiving a link revokes public access
// immediately; subsequent content reads return 404. The underlying predictions and
// pipelines are unaffected and remain accessible through their own authenticated
// endpoints.
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

// Retrieve metadata for a share link owned by the authenticated organization.
// Archived and expired links remain retrievable.
func (r *ShareLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ShareLinkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/share-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a share link so it no longer grants public access. Metadata remains
// retrievable and repeated calls preserve the first archive timestamp.
func (r *ShareLinkService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (res *ShareLinkArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/share-links/%s/archive", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Paginated results for one pipeline exposed by a share link. The response shape
// matches the authed pipeline-results endpoints exactly. Access is gated by the
// share-link ID and — for email-mode links — a signed compute-API JWT. Pipeline
// IDs not covered by the link return 404 indistinguishably from unknown links.
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
// matches the authed pipeline-results endpoints exactly. Access is gated by the
// share-link ID and — for email-mode links — a signed compute-API JWT. Pipeline
// IDs not covered by the link return 404 indistinguishably from unknown links.
func (r *ShareLinkService) ListPipelineResultsAutoPaging(ctx context.Context, pipelineID string, params ShareLinkListPipelineResultsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ShareLinkListPipelineResultsResponse] {
	return pagination.NewCursorPageAutoPager(r.ListPipelineResults(ctx, pipelineID, params, opts...))
}

type ShareLinkNewResponse struct {
	// Share link ID. This value is the bearer credential used to access the linked
	// resources — treat it as a secret.
	ID string `json:"id" api:"required"`
	// Access-control parameters for the share link. Discriminated by `access_mode`:
	// `public` requires no other fields; `email` requires a non-empty `allowed_emails`
	// list.
	AccessParameters ShareLinkNewResponseAccessParametersUnion `json:"access_parameters" api:"required"`
	// When the share link was archived, or null if it has never been archived.
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
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
		ID               respjson.Field
		AccessParameters respjson.Field
		ArchivedAt       respjson.Field
		CreatedAt        respjson.Field
		ExpiresAt        respjson.Field
		PipelineIDs      respjson.Field
		PredictionIDs    respjson.Field
		WorkspaceID      respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkNewResponseAccessParametersUnion contains all possible properties and
// values from
// [ShareLinkNewResponseAccessParametersPublicShareLinkAccessParameters],
// [ShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkNewResponseAccessParametersUnion struct {
	AccessMode string `json:"access_mode"`
	// This field is from variant
	// [ShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters].
	AllowedEmails []string `json:"allowed_emails"`
	JSON          struct {
		AccessMode    respjson.Field
		AllowedEmails respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkNewResponseAccessParametersUnion) AsShareLinkNewResponseAccessParametersPublicShareLinkAccessParameters() (v ShareLinkNewResponseAccessParametersPublicShareLinkAccessParameters) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkNewResponseAccessParametersUnion) AsShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters() (v ShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkNewResponseAccessParametersUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkNewResponseAccessParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public access: anyone holding the share link ID can read.
type ShareLinkNewResponseAccessParametersPublicShareLinkAccessParameters struct {
	AccessMode constant.Public `json:"access_mode" default:"public"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessMode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkNewResponseAccessParametersPublicShareLinkAccessParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkNewResponseAccessParametersPublicShareLinkAccessParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email-restricted access: only the addresses in `allowed_emails` can read the
// link.
type ShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters struct {
	AccessMode constant.Email `json:"access_mode" default:"email"`
	// Email addresses allowed to read the link. Must contain at least one address; up
	// to 100 entries.
	AllowedEmails []string `json:"allowed_emails" api:"required" format:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessMode    respjson.Field
		AllowedEmails respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkNewResponseAccessParametersEmailShareLinkAccessParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkGetResponse struct {
	// Share link ID. This value is the bearer credential used to access the linked
	// resources — treat it as a secret.
	ID string `json:"id" api:"required"`
	// Access-control parameters for the share link. Discriminated by `access_mode`:
	// `public` requires no other fields; `email` requires a non-empty `allowed_emails`
	// list.
	AccessParameters ShareLinkGetResponseAccessParametersUnion `json:"access_parameters" api:"required"`
	// When the share link was archived, or null if it has never been archived.
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
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
		ID               respjson.Field
		AccessParameters respjson.Field
		ArchivedAt       respjson.Field
		CreatedAt        respjson.Field
		ExpiresAt        respjson.Field
		PipelineIDs      respjson.Field
		PredictionIDs    respjson.Field
		WorkspaceID      respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkGetResponseAccessParametersUnion contains all possible properties and
// values from
// [ShareLinkGetResponseAccessParametersPublicShareLinkAccessParameters],
// [ShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkGetResponseAccessParametersUnion struct {
	AccessMode string `json:"access_mode"`
	// This field is from variant
	// [ShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters].
	AllowedEmails []string `json:"allowed_emails"`
	JSON          struct {
		AccessMode    respjson.Field
		AllowedEmails respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkGetResponseAccessParametersUnion) AsShareLinkGetResponseAccessParametersPublicShareLinkAccessParameters() (v ShareLinkGetResponseAccessParametersPublicShareLinkAccessParameters) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkGetResponseAccessParametersUnion) AsShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters() (v ShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkGetResponseAccessParametersUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkGetResponseAccessParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Public access: anyone holding the share link ID can read.
type ShareLinkGetResponseAccessParametersPublicShareLinkAccessParameters struct {
	AccessMode constant.Public `json:"access_mode" default:"public"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessMode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponseAccessParametersPublicShareLinkAccessParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponseAccessParametersPublicShareLinkAccessParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email-restricted access: only the addresses in `allowed_emails` can read the
// link.
type ShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters struct {
	AccessMode constant.Email `json:"access_mode" default:"email"`
	// Email addresses allowed to read the link. Must contain at least one address; up
	// to 100 entries.
	AllowedEmails []string `json:"allowed_emails" api:"required" format:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessMode    respjson.Field
		AllowedEmails respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkGetResponseAccessParametersEmailShareLinkAccessParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkArchiveResponse struct {
	ID       string `json:"id" api:"required"`
	Archived bool   `json:"archived" api:"required"`
	// When the share link was first archived.
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Archived    respjson.Field
		ArchivedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkArchiveResponse) UnmarshalJSON(data []byte) error {
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
	// Access-control parameters for the share link. Discriminated by `access_mode`:
	// `public` requires no other fields; `email` requires a non-empty `allowed_emails`
	// list.
	AccessParameters ShareLinkNewParamsAccessParametersUnion `json:"access_parameters,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ShareLinkNewParamsAccessParametersUnion struct {
	OfShareLinkNewsAccessParametersPublicShareLinkAccessParameters *ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters `json:",omitzero,inline"`
	OfShareLinkNewsAccessParametersEmailShareLinkAccessParameters  *ShareLinkNewParamsAccessParametersEmailShareLinkAccessParameters  `json:",omitzero,inline"`
	paramUnion
}

func (u ShareLinkNewParamsAccessParametersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfShareLinkNewsAccessParametersPublicShareLinkAccessParameters, u.OfShareLinkNewsAccessParametersEmailShareLinkAccessParameters)
}
func (u *ShareLinkNewParamsAccessParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func NewShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters() ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters {
	return ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters{
		AccessMode: "public",
	}
}

// Public access: anyone holding the share link ID can read.
//
// This struct has a constant value, construct it with
// [NewShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters].
type ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters struct {
	AccessMode constant.Public `json:"access_mode" default:"public"`
	paramObj
}

func (r ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters) MarshalJSON() (data []byte, err error) {
	type shadow ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShareLinkNewParamsAccessParametersPublicShareLinkAccessParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email-restricted access: only the addresses in `allowed_emails` can read the
// link.
//
// The properties AccessMode, AllowedEmails are required.
type ShareLinkNewParamsAccessParametersEmailShareLinkAccessParameters struct {
	// Email addresses allowed to read the link. Must contain at least one address; up
	// to 100 entries.
	AllowedEmails []string `json:"allowed_emails,omitzero" api:"required" format:"email"`
	// This field can be elided, and will marshal its zero value as "email".
	AccessMode constant.Email `json:"access_mode" default:"email"`
	paramObj
}

func (r ShareLinkNewParamsAccessParametersEmailShareLinkAccessParameters) MarshalJSON() (data []byte, err error) {
	type shadow ShareLinkNewParamsAccessParametersEmailShareLinkAccessParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShareLinkNewParamsAccessParametersEmailShareLinkAccessParameters) UnmarshalJSON(data []byte) error {
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
