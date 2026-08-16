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
// archive require an API key or supported OAuth bearer token with read permission
// on every referenced resource. Retrieving metadata remains available after expiry
// or archive. Viewing content and listing shared pipeline results are gated by the
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

// Create a read-only share link covering one or more predictions and/or pipelines
// that all live in the same workspace. Public links require only the returned
// bearer ID; email-restricted links also require a signed-in viewer whose email is
// allowed. Treat the returned `id` as a secret.
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

// Read the predictions and pipelines exposed by a share link. Public links require
// no authentication — the share link ID itself is the access credential.
// Email-mode links additionally require a signed compute-API JWT (minted for the
// browser session by Lab, or presented directly by CLI/SDK callers). Returns 404
// indistinguishably for unknown, expired, or archived links.
func (r *ShareLinkService) Read(ctx context.Context, id string, opts ...option.RequestOption) (res *ShareLinkReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/share/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	// Unique result ID.
	ID        string                                        `json:"id" api:"required"`
	Artifacts ShareLinkListPipelineResultsResponseArtifacts `json:"artifacts" api:"required"`
	CreatedAt time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	// Entities in the designed complex, including designed and fixed input entities.
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
// [ShareLinkListPipelineResultsResponseEntityLigandSmilesEntity],
// [ShareLinkListPipelineResultsResponseEntityGlycanEntity].
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
	// This field is from variant
	// [ShareLinkListPipelineResultsResponseEntityGlycanEntity].
	Bonds []ShareLinkListPipelineResultsResponseEntityGlycanEntityBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkListPipelineResultsResponseEntityGlycanEntity].
	Residues []ShareLinkListPipelineResultsResponseEntityGlycanEntityResidue `json:"residues"`
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

func (u ShareLinkListPipelineResultsResponseEntityUnion) AsShareLinkListPipelineResultsResponseEntityGlycanEntity() (v ShareLinkListPipelineResultsResponseEntityGlycanEntity) {
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

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkListPipelineResultsResponseEntityGlycanEntity struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkListPipelineResultsResponseEntityGlycanEntityBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkListPipelineResultsResponseEntityGlycanEntityResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                 `json:"type" default:"glycan"`
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
func (r ShareLinkListPipelineResultsResponseEntityGlycanEntity) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListPipelineResultsResponseEntityGlycanEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkListPipelineResultsResponseEntityGlycanEntityBond struct {
	Atom1 ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkListPipelineResultsResponseEntityGlycanEntityBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityGlycanEntityBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom1 struct {
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
func (r ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom2 struct {
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
func (r ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityGlycanEntityBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListPipelineResultsResponseEntityGlycanEntityResidue struct {
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
func (r ShareLinkListPipelineResultsResponseEntityGlycanEntityResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkListPipelineResultsResponseEntityGlycanEntityResidue) UnmarshalJSON(data []byte) error {
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

type ShareLinkReadResponse struct {
	// Share link ID.
	ID string `json:"id" api:"required"`
	// When the share link was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the share link stops granting access.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Pipelines exposed by this share link, in the order they were registered.
	Pipelines []ShareLinkReadResponsePipelineUnion `json:"pipelines" api:"required"`
	// Predictions exposed by this share link, in the order they were registered.
	Predictions []ShareLinkReadResponsePredictionUnion `json:"predictions" api:"required"`
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
func (r ShareLinkReadResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnion contains all possible properties and values
// from [ShareLinkReadResponsePipelineProteinDesignRun],
// [ShareLinkReadResponsePipelineProteinRedesignRun],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRun],
// [ShareLinkReadResponsePipelineProteinLibraryScreen],
// [ShareLinkReadResponsePipelineSmDesignRun],
// [ShareLinkReadResponsePipelineSmScreen].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineUnion struct {
	ID            string    `json:"id"`
	CompletedAt   time.Time `json:"completed_at"`
	CreatedAt     time.Time `json:"created_at"`
	DataDeletedAt time.Time `json:"data_deleted_at"`
	Engine        string    `json:"engine"`
	EngineVersion string    `json:"engine_version"`
	// This field is a union of [ShareLinkReadResponsePipelineProteinDesignRunError],
	// [ShareLinkReadResponsePipelineProteinRedesignRunError],
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunError],
	// [ShareLinkReadResponsePipelineProteinLibraryScreenError],
	// [ShareLinkReadResponsePipelineSmDesignRunError],
	// [ShareLinkReadResponsePipelineSmScreenError]
	Error ShareLinkReadResponsePipelineUnionError `json:"error"`
	// This field is a union of [ShareLinkReadResponsePipelineProteinDesignRunInput],
	// [ShareLinkReadResponsePipelineProteinRedesignRunInput],
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion],
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInput],
	// [ShareLinkReadResponsePipelineSmDesignRunInput],
	// [ShareLinkReadResponsePipelineSmScreenInput]
	Input           ShareLinkReadResponsePipelineUnionInput `json:"input"`
	Livemode        bool                                    `json:"livemode"`
	Pipeline        string                                  `json:"pipeline"`
	PipelineVersion string                                  `json:"pipeline_version"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunProgress],
	// [ShareLinkReadResponsePipelineProteinRedesignRunProgress],
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunProgress],
	// [ShareLinkReadResponsePipelineProteinLibraryScreenProgress],
	// [ShareLinkReadResponsePipelineSmDesignRunProgress],
	// [ShareLinkReadResponsePipelineSmScreenProgress]
	Progress       ShareLinkReadResponsePipelineUnionProgress `json:"progress"`
	StartedAt      time.Time                                  `json:"started_at"`
	Status         string                                     `json:"status"`
	StoppedAt      time.Time                                  `json:"stopped_at"`
	WorkspaceID    string                                     `json:"workspace_id"`
	IdempotencyKey string                                     `json:"idempotency_key"`
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

func (u ShareLinkReadResponsePipelineUnion) AsShareLinkReadResponsePipelineProteinDesignRun() (v ShareLinkReadResponsePipelineProteinDesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineUnion) AsShareLinkReadResponsePipelineProteinRedesignRun() (v ShareLinkReadResponsePipelineProteinRedesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRun() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineUnion) AsShareLinkReadResponsePipelineProteinLibraryScreen() (v ShareLinkReadResponsePipelineProteinLibraryScreen) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineUnion) AsShareLinkReadResponsePipelineSmDesignRun() (v ShareLinkReadResponsePipelineSmDesignRun) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineUnion) AsShareLinkReadResponsePipelineSmScreen() (v ShareLinkReadResponsePipelineSmScreen) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkReadResponsePipelineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionError is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion]. ShareLinkReadResponsePipelineUnionError
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionError struct {
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

func (r *ShareLinkReadResponsePipelineUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInput is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion]. ShareLinkReadResponsePipelineUnionInput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionInput struct {
	// This field is from variant [ShareLinkReadResponsePipelineProteinDesignRunInput].
	BinderSpecification ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion `json:"binder_specification"`
	NumProteins         int64                                                                      `json:"num_proteins"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion],
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion],
	// [ShareLinkReadResponsePipelineSmDesignRunInputTarget],
	// [ShareLinkReadResponsePipelineSmScreenInputTarget]
	Target         ShareLinkReadResponsePipelineUnionInputTarget `json:"target"`
	IdempotencyKey string                                        `json:"idempotency_key"`
	WorkspaceID    string                                        `json:"workspace_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinRedesignRunInput].
	ChainsConfig ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfig `json:"chains_config"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinRedesignRunInputStructure],
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseStructure],
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseStructure]
	Structure ShareLinkReadResponsePipelineUnionInputStructure `json:"structure"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinRedesignRunInput].
	Rules ShareLinkReadResponsePipelineProteinRedesignRunInputRules `json:"rules"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion],
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity]
	Entities ShareLinkReadResponsePipelineUnionInputEntities `json:"entities"`
	Type     string                                          `json:"type"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion],
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	GlobalDesignFilters ShareLinkReadResponsePipelineUnionInputGlobalDesignFilters `json:"global_design_filters"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInput].
	Proteins ShareLinkReadResponsePipelineProteinLibraryScreenInputProteins `json:"proteins"`
	// This field is from variant [ShareLinkReadResponsePipelineSmDesignRunInput].
	NumMolecules int64 `json:"num_molecules"`
	// This field is from variant [ShareLinkReadResponsePipelineSmDesignRunInput].
	ChemicalSpace ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpace `json:"chemical_space"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFilters],
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFilters]
	MoleculeFilters ShareLinkReadResponsePipelineUnionInputMoleculeFilters `json:"molecule_filters"`
	// This field is from variant [ShareLinkReadResponsePipelineSmScreenInput].
	Molecules ShareLinkReadResponsePipelineSmScreenInputMolecules `json:"molecules"`
	JSON      struct {
		BinderSpecification respjson.Field
		NumProteins         respjson.Field
		Target              respjson.Field
		IdempotencyKey      respjson.Field
		WorkspaceID         respjson.Field
		ChainsConfig        respjson.Field
		Structure           respjson.Field
		Rules               respjson.Field
		Entities            respjson.Field
		Type                respjson.Field
		GlobalDesignFilters respjson.Field
		Proteins            respjson.Field
		NumMolecules        respjson.Field
		ChemicalSpace       respjson.Field
		MoleculeFilters     respjson.Field
		Molecules           respjson.Field
		raw                 string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTarget is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTarget provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionInputTarget struct {
	// This field is a union of
	// [map[string]ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion],
	// [map[string]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion]
	ChainSelection ShareLinkReadResponsePipelineUnionInputTargetChainSelection `json:"chain_selection"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure],
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure]
	Structure ShareLinkReadResponsePipelineUnionInputTargetStructure `json:"structure"`
	Type      string                                                 `json:"type"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion],
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion],
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion],
	// [[]ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion]
	Entities ShareLinkReadResponsePipelineUnionInputTargetEntities `json:"entities"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond],
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond],
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputTargetBond],
	// [[]ShareLinkReadResponsePipelineSmScreenInputTargetBond]
	Bonds ShareLinkReadResponsePipelineUnionInputTargetBonds `json:"bonds"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion],
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion],
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion],
	// [[]ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion]
	Constraints         ShareLinkReadResponsePipelineUnionInputTargetConstraints `json:"constraints"`
	EpitopeLigandChains []string                                                 `json:"epitope_ligand_chains"`
	EpitopeResidues     []int64                                                  `json:"epitope_residues"`
	NonBindingResidues  []int64                                                  `json:"non_binding_residues"`
	PocketResidues      []int64                                                  `json:"pocket_residues"`
	ReferenceLigands    []string                                                 `json:"reference_ligands"`
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

func (r *ShareLinkReadResponsePipelineUnionInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTargetChainSelection is an implicit
// subunion of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTargetChainSelection provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionInputTargetChainSelection struct {
	ChainType string `json:"chain_type"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion],
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion]
	CropResidues       ShareLinkReadResponsePipelineUnionInputTargetChainSelectionCropResidues `json:"crop_residues"`
	EpitopeResidues    []int64                                                                 `json:"epitope_residues"`
	FlexibleResidues   []int64                                                                 `json:"flexible_residues"`
	NonBindingResidues []int64                                                                 `json:"non_binding_residues"`
	Ccd                string                                                                  `json:"ccd"`
	Smiles             string                                                                  `json:"smiles"`
	JSON               struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		Ccd                respjson.Field
		Smiles             respjson.Field
		raw                string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputTargetChainSelection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTargetChainSelectionCropResidues is an
// implicit subunion of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTargetChainSelectionCropResidues provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkReadResponsePipelineUnionInputTargetChainSelectionCropResidues struct {
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

func (r *ShareLinkReadResponsePipelineUnionInputTargetChainSelectionCropResidues) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTargetStructure is an implicit subunion
// of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTargetStructure provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionInputTargetStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputTargetStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTargetEntities is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTargetEntities provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntities
// OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntities
// OfShareLinkReadResponsePipelineSmDesignRunInputTargetEntities
// OfShareLinkReadResponsePipelineSmScreenInputTargetEntities]
type ShareLinkReadResponsePipelineUnionInputTargetEntities struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntities []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntities []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion] instead of an
	// object.
	OfShareLinkReadResponsePipelineSmDesignRunInputTargetEntities []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion] instead of an
	// object.
	OfShareLinkReadResponsePipelineSmScreenInputTargetEntities []ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion `json:",inline"`
	JSON                                                       struct {
		OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntities     respjson.Field
		OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntities respjson.Field
		OfShareLinkReadResponsePipelineSmDesignRunInputTargetEntities                                  respjson.Field
		OfShareLinkReadResponsePipelineSmScreenInputTargetEntities                                     respjson.Field
		raw                                                                                            string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputTargetEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTargetBonds is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTargetBonds provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBonds
// OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBonds
// OfShareLinkReadResponsePipelineSmDesignRunInputTargetBonds
// OfShareLinkReadResponsePipelineSmScreenInputTargetBonds]
type ShareLinkReadResponsePipelineUnionInputTargetBonds struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBonds []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBonds []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputTargetBond] instead of an
	// object.
	OfShareLinkReadResponsePipelineSmDesignRunInputTargetBonds []ShareLinkReadResponsePipelineSmDesignRunInputTargetBond `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmScreenInputTargetBond] instead of an object.
	OfShareLinkReadResponsePipelineSmScreenInputTargetBonds []ShareLinkReadResponsePipelineSmScreenInputTargetBond `json:",inline"`
	JSON                                                    struct {
		OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBonds     respjson.Field
		OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBonds respjson.Field
		OfShareLinkReadResponsePipelineSmDesignRunInputTargetBonds                                  respjson.Field
		OfShareLinkReadResponsePipelineSmScreenInputTargetBonds                                     respjson.Field
		raw                                                                                         string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputTargetBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputTargetConstraints is an implicit subunion
// of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputTargetConstraints provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraints
// OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraints
// OfShareLinkReadResponsePipelineSmDesignRunInputTargetConstraints
// OfShareLinkReadResponsePipelineSmScreenInputTargetConstraints]
type ShareLinkReadResponsePipelineUnionInputTargetConstraints struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraints []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraints []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion] instead
	// of an object.
	OfShareLinkReadResponsePipelineSmDesignRunInputTargetConstraints []ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion] instead of
	// an object.
	OfShareLinkReadResponsePipelineSmScreenInputTargetConstraints []ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion `json:",inline"`
	JSON                                                          struct {
		OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraints     respjson.Field
		OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraints respjson.Field
		OfShareLinkReadResponsePipelineSmDesignRunInputTargetConstraints                                  respjson.Field
		OfShareLinkReadResponsePipelineSmScreenInputTargetConstraints                                     respjson.Field
		raw                                                                                               string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputTargetConstraints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputStructure is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputStructure provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionInputStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputEntities is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputEntities provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntities
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntities]
type ShareLinkReadResponsePipelineUnionInputEntities struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntities []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntities []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity `json:",inline"`
	JSON                                                                                                                 struct {
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntities  respjson.Field
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntities respjson.Field
		raw                                                                                                                  string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputGlobalDesignFilters is an implicit
// subunion of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputGlobalDesignFilters provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters]
type ShareLinkReadResponsePipelineUnionInputGlobalDesignFilters struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	JSON                                                                                                                            struct {
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters  respjson.Field
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters respjson.Field
		raw                                                                                                                             string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputGlobalDesignFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputMoleculeFilters is an implicit subunion
// of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputMoleculeFilters provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionInputMoleculeFilters struct {
	BoltzSmartsCatalogFilterLevel string `json:"boltz_smarts_catalog_filter_level"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion],
	// [[]ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion]
	CustomFilters ShareLinkReadResponsePipelineUnionInputMoleculeFiltersCustomFilters `json:"custom_filters"`
	JSON          struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		raw                           string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionInputMoleculeFiltersCustomFilters is an
// implicit subunion of [ShareLinkReadResponsePipelineUnion].
// ShareLinkReadResponsePipelineUnionInputMoleculeFiltersCustomFilters provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilters
// OfShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilters]
type ShareLinkReadResponsePipelineUnionInputMoleculeFiltersCustomFilters struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilters []ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilters []ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion `json:",inline"`
	JSON                                                                     struct {
		OfShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilters respjson.Field
		OfShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilters    respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineUnionInputMoleculeFiltersCustomFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineUnionProgress is an implicit subunion of
// [ShareLinkReadResponsePipelineUnion]. ShareLinkReadResponsePipelineUnionProgress
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineUnion].
type ShareLinkReadResponsePipelineUnionProgress struct {
	NumProteinsGenerated    int64  `json:"num_proteins_generated"`
	TotalProteinsToGenerate int64  `json:"total_proteins_to_generate"`
	LatestResultID          string `json:"latest_result_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenProgress].
	NumProteinsFailed int64 `json:"num_proteins_failed"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenProgress].
	NumProteinsScreened int64 `json:"num_proteins_screened"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenProgress].
	TotalProteinsToScreen int64 `json:"total_proteins_to_screen"`
	// This field is from variant [ShareLinkReadResponsePipelineSmDesignRunProgress].
	NumMoleculesGenerated int64 `json:"num_molecules_generated"`
	// This field is from variant [ShareLinkReadResponsePipelineSmDesignRunProgress].
	TotalMoleculesToGenerate int64 `json:"total_molecules_to_generate"`
	// This field is from variant [ShareLinkReadResponsePipelineSmScreenProgress].
	NumMoleculesFailed int64 `json:"num_molecules_failed"`
	// This field is from variant [ShareLinkReadResponsePipelineSmScreenProgress].
	NumMoleculesScreened int64 `json:"num_molecules_screened"`
	// This field is from variant [ShareLinkReadResponsePipelineSmScreenProgress].
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen"`
	// This field is from variant [ShareLinkReadResponsePipelineSmScreenProgress].
	RejectionSummary ShareLinkReadResponsePipelineSmScreenProgressRejectionSummary `json:"rejection_summary"`
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

func (r *ShareLinkReadResponsePipelineUnionProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A protein design pipeline run that generates novel protein binders
type ShareLinkReadResponsePipelineProteinDesignRun struct {
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
	EngineVersion constant.String1_0                                 `json:"engine_version" default:"1.0"`
	Error         ShareLinkReadResponsePipelineProteinDesignRunError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkReadResponsePipelineProteinDesignRunInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein design
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein design
	PipelineVersion constant.String1_0                                    `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkReadResponsePipelineProteinDesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                             `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkReadResponsePipelineProteinDesignRunStatus `json:"status" api:"required"`
	StoppedAt time.Time                                           `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinDesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunError struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinDesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkReadResponsePipelineProteinDesignRunInput struct {
	// Binder specification for protein design. Use no_template for sequence-defined
	// binders, structure_template for uploaded binder structures, boltz_curated for
	// Boltz-managed nanobody and antibody defaults, or
	// uniformly_sampled_specifications to sample uniformly across multiple binder
	// specifications.
	BinderSpecification ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion `json:"binder_specification" api:"required"`
	// Number of protein designs to generate. Must be between 10 and 1,000,000.
	NumProteins int64 `json:"num_proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion `json:"target" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinDesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse].
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection"`
	Modality       string                                                                                                                                 `json:"modality"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse].
	Structure ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure"`
	Type      string                                                                                                            `json:"type"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules],
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules],
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules]
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules `json:"rules"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse].
	Entities []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse].
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse].
	Binder ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse].
	BinderSpecifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion `json:"binder_specifications"`
	JSON                 struct {
		ChainSelection       respjson.Field
		Modality             respjson.Field
		Structure            respjson.Field
		Type                 respjson.Field
		Rules                respjson.Field
		Entities             respjson.Field
		Bonds                respjson.Field
		Binder               respjson.Field
		BinderSpecifications respjson.Field
		raw                  string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnion].
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules struct {
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

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUnionRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// pipeline run.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the pipeline run — any chains omitted from
	// this mapping are ignored. Each value defines which residues to keep
	// (crop_residues). Omit design_motifs to include the chain as fixed scaffold
	// context.
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality  ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality  `json:"modality" api:"required"`
	Structure ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure" api:"required"`
	Type      constant.StructureTemplate                                                                                        `json:"type" default:"structure_template"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	CropResidues ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	DesignMotifs []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	JSON         struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange],
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange]
	DesignLengthRange ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange `json:"design_length_range"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	EndIndex int64 `json:"end_index"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	StartIndex int64  `json:"start_index"`
	Type       string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange struct {
	Max  int64 `json:"max"`
	Min  int64 `json:"min"`
	JSON struct {
		Max respjson.Field
		Min respjson.Field
		raw string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range" api:"required"`
	Type              constant.Insertion                                                                                                                                                                                `json:"type" default:"insertion"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityPeptide       ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "peptide"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityAntibody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "antibody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityNanobody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModalityCustomProtein ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseModality = "custom_protein"
)

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationStructureTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality `json:"modality" api:"required"`
	Type     constant.NoTemplate                                                                                       `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	Modifications ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications]
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                                         struct {
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications    respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications        respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications        respjson.Field
		raw                                                                                                                                                 string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protein binder entity with designed and/or fixed segments.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse struct {
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
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                                                      `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityPeptide       ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "peptide"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityAntibody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "antibody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityNanobody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModalityCustomProtein ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseModality = "custom_protein"
)

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationNoTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder" api:"required"`
	Type   constant.BoltzCurated                                                                                     `json:"type" default:"boltz_curated"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzNanobody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzAntibody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBoltzCuratedBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of binder specifications sampled uniformly during protein design.
// This lets one run explore multiple binder definitions while keeping each
// generation request shape unchanged.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse struct {
	// Binder specifications to sample uniformly when generating designs. Each
	// generation samples one specification from this list; over larger runs this gives
	// roughly equal representation.
	BinderSpecifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion `json:"binder_specifications" api:"required"`
	Type                 constant.UniformlySampledSpecifications                                                                                           `json:"type" default:"uniformly_sampled_specifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderSpecifications respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse].
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection"`
	Modality       string                                                                                                                                                                                      `json:"modality"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse].
	Structure ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure"`
	Type      string                                                                                                                                                                 `json:"type"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseRules],
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseRules],
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseRules]
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnionRules `json:"rules"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse].
	Entities []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse].
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse].
	Binder ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder"`
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnionRules
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnionRules
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnion].
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnionRules struct {
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

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationUnionRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification starting from an existing 3D structure. Upload a CIF/PDB
// file and select which chains to include, which residues to keep, and which
// regions to redesign. Only chains included in chain_selection are part of the
// pipeline run.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse struct {
	// Chains selected from the uploaded binder structure, keyed by chain ID. Only
	// chains listed here are included in the pipeline run — any chains omitted from
	// this mapping are ignored. Each value defines which residues to keep
	// (crop_residues). Omit design_motifs to include the chain as fixed scaffold
	// context.
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality  ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModality  `json:"modality" api:"required"`
	Structure ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseStructure `json:"structure" api:"required"`
	Type      constant.StructureTemplate                                                                                                                                             `json:"type" default:"structure_template"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	CropResidues ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec].
	DesignMotifs []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
	JSON         struct {
		ChainType    respjson.Field
		CropResidues respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain crop and design specification for a polymer chain in
// structure_template mode.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are removed before design.
	CropResidues ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
	// Optional motifs (replacement or insertion) defining which regions to redesign on
	// this chain. Omit this field to include the chain as fixed scaffold context.
	DesignMotifs []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion `json:"design_motifs"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion struct {
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange],
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange]
	DesignLengthRange ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange `json:"design_length_range"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	EndIndex int64 `json:"end_index"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif].
	StartIndex int64  `json:"start_index"`
	Type       string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion].
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange struct {
	Max  int64 `json:"max"`
	Min  int64 `json:"min"`
	JSON struct {
		Max respjson.Field
		Min respjson.Field
		raw string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnionDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace a contiguous region of the sequence with a designed segment. Residues
// from start_index to end_index (inclusive) are replaced with a new sequence of
// the specified length.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif struct {
	// Allowed sequence length range for designed regions
	DesignLengthRange ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange `json:"design_length_range" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Insert a designed segment at a specific position in the sequence.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif struct {
	// 0-indexed position after which to insert. Use -1 to insert before the first
	// residue.
	AfterResidueIndex int64 `json:"after_residue_index" api:"required"`
	// Allowed sequence length range for designed regions
	DesignLengthRange ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange `json:"design_length_range" api:"required"`
	Type              constant.Insertion                                                                                                                                                                                                                                     `json:"type" default:"insertion"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allowed sequence length range for designed regions
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplatePolymerChainSpecDesignMotifInsertionMotifDesignLengthRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in structure_template mode. The full
// ligand is always included.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseChainSelectionStructureTemplateLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModality string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModalityPeptide       ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModality = "peptide"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModalityAntibody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModality = "antibody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModalityNanobody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModality = "nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModalityCustomProtein ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseModality = "custom_protein"
)

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseRules struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationStructureTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Binder specification without a structural template. Define the binder from
// sequence components (fixed and designed segments) without providing a starting
// 3D structure.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse struct {
	// Binder entities composing the design. At least one must be a designed_protein
	// entity. Additional fixed entities (RNA, DNA, ligands) can be included as part of
	// the complex.
	Entities []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion `json:"entities" api:"required"`
	// Any of "peptide", "antibody", "nanobody", "custom_protein".
	Modality ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModality `json:"modality" api:"required"`
	Type     constant.NoTemplate                                                                                                                                            `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the binder complex. If defining bonds
	// where an atom is part of a designed protein chain, assume residue indices count
	// designed regions as the minimum length. Example: designed protein "1..3C1..2",
	// "C" is residue 1 (0-indexed) of the designed protein.
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBond `json:"bonds"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	Modifications ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications]
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                                                                                              struct {
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModifications respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModifications    respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModifications        respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModifications        respjson.Field
		raw                                                                                                                                                                                                      string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Protein binder entity with designed and/or fixed segments.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse struct {
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
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityDesignedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed protein entity whose sequence is not redesigned.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse struct {
	// Chain IDs to assign to this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// Optional CCD polymer modifications. Defaults to [] when omitted. SMILES
	// modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityFixedLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                                                                                                           `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModality string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModalityPeptide       ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModality = "peptide"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModalityAntibody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModality = "antibody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModalityNanobody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModality = "nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModalityCustomProtein ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseModality = "custom_protein"
)

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseRules struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationNoTemplateBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder specification. Choose a curated nanobody or
// antibody family and Boltz will select from maintained template lists during
// design. The curated lists are managed by Boltz and may be updated over time to
// improve quality and coverage.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse struct {
	// Boltz-managed curated binder family. Boltz maintains and may update the
	// underlying template lists on behalf of customers.
	//
	// Any of "boltz_nanobody", "boltz_antibody".
	Binder ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinder `json:"binder" api:"required"`
	Type   constant.BoltzCurated                                                                                                                                          `json:"type" default:"boltz_curated"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinder string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzNanobody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinderBoltzAntibody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseBinder = "boltz_antibody"
)

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseRules struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBoltzCuratedBinderSpecResponseRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBinder string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBinderBoltzNanobody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBinder = "boltz_nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBinderBoltzAntibody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationBinder = "boltz_antibody"
)

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModality string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModalityPeptide       ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModality = "peptide"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModalityAntibody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModality = "antibody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModalityNanobody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModality = "nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModalityCustomProtein ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationUniformlySampledBinderSpecResponseBinderSpecificationModality = "custom_protein"
)

// Boltz-managed curated binder family. Boltz maintains and may update the
// underlying template lists on behalf of customers.
type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBinder string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBinderBoltzNanobody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBinder = "boltz_nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBinderBoltzAntibody ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationBinder = "boltz_antibody"
)

type ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModality string

const (
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModalityPeptide       ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModality = "peptide"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModalityAntibody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModality = "antibody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModalityNanobody      ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModality = "nanobody"
	ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModalityCustomProtein ShareLinkReadResponsePipelineProteinDesignRunInputBinderSpecificationModality = "custom_protein"
)

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse].
	Structure ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                                           `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	Entities []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	Constraints []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the pipeline run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                                            `json:"type" default:"structure_template"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	NonBindingResidues []int64 `json:"non_binding_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
	Ccd string `json:"ccd"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
	Smiles string `json:"smiles"`
	JSON   struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		Ccd                respjson.Field
		Smiles             respjson.Field
		raw                string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the pipeline run.
	CropResidues ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included. An optional CCD or SMILES override preserves the
// ligand chemistry while retaining its coordinates.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// Original CCD identity for this ligand. Use when the structure stores the
	// coordinates under a generic component ID such as LIG. Mutually exclusive with
	// smiles.
	Ccd string `json:"ccd"`
	// Original SMILES identity for this ligand. Use when the structure stores the
	// coordinates under a generic component ID such as LIG. Mutually exclusive with
	// ccd.
	Smiles string `json:"smiles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		Ccd         respjson.Field
		Smiles      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                                           `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion].
// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                   struct {
		OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                                        string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                                     `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                               `json:"type" default:"contact"`
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinDesignRunInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunProgress struct {
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
func (r ShareLinkReadResponsePipelineProteinDesignRunProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinDesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinDesignRunStatus string

const (
	ShareLinkReadResponsePipelineProteinDesignRunStatusPending   ShareLinkReadResponsePipelineProteinDesignRunStatus = "pending"
	ShareLinkReadResponsePipelineProteinDesignRunStatusRunning   ShareLinkReadResponsePipelineProteinDesignRunStatus = "running"
	ShareLinkReadResponsePipelineProteinDesignRunStatusSucceeded ShareLinkReadResponsePipelineProteinDesignRunStatus = "succeeded"
	ShareLinkReadResponsePipelineProteinDesignRunStatusFailed    ShareLinkReadResponsePipelineProteinDesignRunStatus = "failed"
	ShareLinkReadResponsePipelineProteinDesignRunStatusStopped   ShareLinkReadResponsePipelineProteinDesignRunStatus = "stopped"
)

// A fixed-structure binder sequence redesign run
type ShareLinkReadResponsePipelineProteinRedesignRun struct {
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
	EngineVersion constant.V2026_03_01                                 `json:"engine_version" default:"v2026-03-01"`
	Error         ShareLinkReadResponsePipelineProteinRedesignRunError `json:"error" api:"required"`
	// Input for protein redesign from a complete target/binder CIF complex.
	Input ShareLinkReadResponsePipelineProteinRedesignRunInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein redesign
	Pipeline constant.BoltzProteinRedesign `json:"pipeline" default:"boltz-protein-redesign"`
	// Pipeline version used for protein redesign
	PipelineVersion constant.V2026_03_01                                    `json:"pipeline_version" default:"v2026-03-01"`
	Progress        ShareLinkReadResponsePipelineProteinRedesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                               `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkReadResponsePipelineProteinRedesignRunStatus `json:"status" api:"required"`
	StoppedAt time.Time                                             `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkReadResponsePipelineProteinRedesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinRedesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinRedesignRunError struct {
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
func (r ShareLinkReadResponsePipelineProteinRedesignRunError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinRedesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for protein redesign from a complete target/binder CIF complex.
type ShareLinkReadResponsePipelineProteinRedesignRunInput struct {
	// Complete chain assignment for the input CIF. Every CIF chain must appear exactly
	// once, either under target_chains or binder_chains.
	ChainsConfig ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfig `json:"chains_config" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                         `json:"num_proteins" api:"required"`
	Structure   ShareLinkReadResponsePipelineProteinRedesignRunInputStructure `json:"structure" api:"required"`
	// Optional idempotency key.
	IdempotencyKey string `json:"idempotency_key"`
	// Constraints applied during sequence design
	Rules ShareLinkReadResponsePipelineProteinRedesignRunInputRules `json:"rules"`
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
func (r ShareLinkReadResponsePipelineProteinRedesignRunInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete chain assignment for the input CIF. Every CIF chain must appear exactly
// once, either under target_chains or binder_chains.
type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfig struct {
	// Binder chains keyed by CIF chain ID. Protein binder chains may provide
	// designed_residues; at least 5 total unique designed residues are required across
	// all protein binder chains. Target and binder chains must be disjoint and
	// together cover every chain in the CIF.
	BinderChains map[string]ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion `json:"binder_chains" api:"required"`
	// Target chains keyed by CIF chain ID. Target and binder chains must be disjoint
	// and together cover every chain in the CIF.
	TargetChains map[string]ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChain `json:"target_chains" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinderChains respjson.Field
		TargetChains respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig],
// [ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig].
	DesignedResidues []int64 `json:"designed_residues"`
	JSON             struct {
		ChainType        respjson.Field
		DesignedResidues respjson.Field
		raw              string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) AsShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig() (v ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) AsShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig() (v ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a protein binder chain. Across all protein binder chains, at
// least 5 unique designed residues are required.
type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig struct {
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
func (r ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderProteinChainConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a fixed non-protein binder chain in the input CIF.
type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig struct {
	// Any of "ligand", "rna", "dna".
	ChainType ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType `json:"chain_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType string

const (
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainTypeLigand ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType = "ligand"
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainTypeRna    ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType = "rna"
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainTypeDna    ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigBinderChainProteinRedesignBinderNonProteinChainConfigChainType = "dna"
)

// Configuration for a fixed target chain in the input CIF.
type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChain struct {
	// Molecule type for a chain in the input CIF.
	//
	// Any of "ligand", "protein", "rna", "dna".
	ChainType ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType `json:"chain_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChain) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule type for a chain in the input CIF.
type ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType string

const (
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeLigand  ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "ligand"
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeProtein ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "protein"
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeRna     ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "rna"
	ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainTypeDna     ShareLinkReadResponsePipelineProteinRedesignRunInputChainsConfigTargetChainChainType = "dna"
)

type ShareLinkReadResponsePipelineProteinRedesignRunInputStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinRedesignRunInputStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constraints applied during sequence design
type ShareLinkReadResponsePipelineProteinRedesignRunInputRules struct {
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
func (r ShareLinkReadResponsePipelineProteinRedesignRunInputRules) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinRedesignRunInputRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinRedesignRunProgress struct {
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
func (r ShareLinkReadResponsePipelineProteinRedesignRunProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinRedesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinRedesignRunStatus string

const (
	ShareLinkReadResponsePipelineProteinRedesignRunStatusPending   ShareLinkReadResponsePipelineProteinRedesignRunStatus = "pending"
	ShareLinkReadResponsePipelineProteinRedesignRunStatusRunning   ShareLinkReadResponsePipelineProteinRedesignRunStatus = "running"
	ShareLinkReadResponsePipelineProteinRedesignRunStatusSucceeded ShareLinkReadResponsePipelineProteinRedesignRunStatus = "succeeded"
	ShareLinkReadResponsePipelineProteinRedesignRunStatusFailed    ShareLinkReadResponsePipelineProteinRedesignRunStatus = "failed"
	ShareLinkReadResponsePipelineProteinRedesignRunStatusStopped   ShareLinkReadResponsePipelineProteinRedesignRunStatus = "stopped"
)

// A fixed-structure protein sequence redesign run.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRun struct {
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
	EngineVersion constant.V2026_07_14                                         `json:"engine_version" default:"v2026-07-14"`
	Error         ShareLinkReadResponsePipelineProteinSequenceRedesignRunError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode        bool                                                            `json:"livemode" api:"required"`
	Pipeline        constant.BoltzProteinRedesign                                   `json:"pipeline" default:"boltz-protein-redesign"`
	PipelineVersion constant.V2026_07_14                                            `json:"pipeline_version" default:"v2026-07-14"`
	Progress        ShareLinkReadResponsePipelineProteinSequenceRedesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                                       `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus `json:"status" api:"required"`
	StoppedAt time.Time                                                     `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunError struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunError) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion struct {
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion],
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity]
	Entities    ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionEntities `json:"entities"`
	NumProteins int64                                                                     `json:"num_proteins"`
	// This field is a union of
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseStructure],
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseStructure]
	Structure ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionStructure `json:"structure"`
	Type      string                                                                     `json:"type"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion],
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	GlobalDesignFilters ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionGlobalDesignFilters `json:"global_design_filters"`
	IdempotencyKey      string                                                                               `json:"idempotency_key"`
	WorkspaceID         string                                                                               `json:"workspace_id"`
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

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionEntities is an
// implicit subunion of
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion].
// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionEntities
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntities
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntities]
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionEntities struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntities []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntities []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity `json:",inline"`
	JSON                                                                                                                 struct {
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntities  respjson.Field
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntities respjson.Field
		raw                                                                                                                  string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionEntities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionStructure is an
// implicit subunion of
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion].
// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionStructure
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion].
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionStructure struct {
	URL          string    `json:"url"`
	URLExpiresAt time.Time `json:"url_expires_at"`
	JSON         struct {
		URL          respjson.Field
		URLExpiresAt respjson.Field
		raw          string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionGlobalDesignFilters
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion].
// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionGlobalDesignFilters
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters
// OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters]
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionGlobalDesignFilters struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:",inline"`
	JSON                                                                                                                            struct {
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilters  respjson.Field
		OfShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilters respjson.Field
		raw                                                                                                                             string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputUnionGlobalDesignFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once as target or binder.
	Entities []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                                              `json:"num_proteins" api:"required"`
	Structure   ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Binder                                                                                                    `json:"type" default:"binder"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                                             `json:"idempotency_key"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion struct {
	ChainID string `json:"chain_id"`
	Role    string `json:"role"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse].
	Type constant.FromTemplate `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse].
	DesignMotifs []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
	JSON         struct {
		ChainID      respjson.Field
		Role         respjson.Field
		Type         respjson.Field
		DesignMotifs respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fixed target chain from the input CIF.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignTargetEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse struct {
	ChainID string                `json:"chain_id" api:"required"`
	Role    constant.Binder       `json:"role" default:"binder"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the binder chain fixed.
	DesignMotifs []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif `json:"design_motifs"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion `json:"filters" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseEntityBinderSequenceRedesignBinderEntityResponseDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputBinderProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponse struct {
	// Every chain in the input CIF, assigned exactly once.
	Entities []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity `json:"entities" api:"required"`
	// Number of unique filter-passing redesigned proteins to generate.
	NumProteins int64                                                                                                               `json:"num_proteins" api:"required"`
	Structure   ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseStructure `json:"structure" api:"required"`
	Type        constant.Generic                                                                                                    `json:"type" default:"generic"`
	// Filters applied to every redesigned region. When omitted, cysteine is excluded.
	// Pass [] to disable global filters.
	GlobalDesignFilters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion `json:"global_design_filters"`
	IdempotencyKey      string                                                                                                                              `json:"idempotency_key"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity struct {
	ChainID string                `json:"chain_id" api:"required"`
	Type    constant.FromTemplate `json:"type" default:"from_template"`
	// Residues to redesign. Omit this field to keep the chain fixed.
	DesignMotifs []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif `json:"design_motifs"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif struct {
	// Filters applied to this motif in addition to global_design_filters.
	Filters []ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion `json:"filters" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseEntityDesignMotifFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse],
// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse].
	AminoAcids []string `json:"amino_acids"`
	Type       string   `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse].
	MaxFraction float64 `json:"max_fraction"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse].
	Motifs []string `json:"motifs"`
	JSON   struct {
		AminoAcids  respjson.Field
		Type        respjson.Field
		MaxFraction respjson.Field
		Motifs      respjson.Field
		raw         string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) AsShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse() (v ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedAminoAcidsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterMaxHydrophobicFractionDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunInputGenericProteinSequenceRedesignRunInputResponseGlobalDesignFilterExcludedSequenceMotifsDesignFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunProgress struct {
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
func (r ShareLinkReadResponsePipelineProteinSequenceRedesignRunProgress) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinSequenceRedesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus string

const (
	ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatusPending   ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus = "pending"
	ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatusRunning   ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus = "running"
	ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatusSucceeded ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus = "succeeded"
	ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatusFailed    ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus = "failed"
	ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatusStopped   ShareLinkReadResponsePipelineProteinSequenceRedesignRunStatus = "stopped"
)

// A protein library screening pipeline run
type ShareLinkReadResponsePipelineProteinLibraryScreen struct {
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
	EngineVersion constant.String1_0                                     `json:"engine_version" default:"1.0"`
	Error         ShareLinkReadResponsePipelineProteinLibraryScreenError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkReadResponsePipelineProteinLibraryScreenInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for protein library screen
	Pipeline constant.Boltzprot `json:"pipeline" default:"boltzprot"`
	// Pipeline version used for protein library screen
	PipelineVersion constant.String1_0                                        `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkReadResponsePipelineProteinLibraryScreenProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                                 `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkReadResponsePipelineProteinLibraryScreenStatus `json:"status" api:"required"`
	StoppedAt time.Time                                               `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreen) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinLibraryScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenError struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkReadResponsePipelineProteinLibraryScreenInput struct {
	Proteins ShareLinkReadResponsePipelineProteinLibraryScreenInputProteins `json:"proteins" api:"required"`
	// Target specification (structure template or template-free)
	Target ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Proteins    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputProteins struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputProteins) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputProteins) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse].
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse].
	Structure ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure `json:"structure"`
	Type      string                                                                                               `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	Entities []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	Bonds []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	Constraints []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	EpitopeLigandChains []string `json:"epitope_ligand_chains"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
	EpitopeResidues map[string][]int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse].
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

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by an uploaded 3D structure (CIF or PDB file). Only chains
// included in chain_selection are used.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse struct {
	// Chains selected from the uploaded structure, keyed by chain ID. Only chains
	// listed here are included in the pipeline run — any chains omitted from this
	// mapping are ignored. Each value defines which residues to keep, which are
	// epitope residues, which are non-binding residues, and which are flexible.
	ChainSelection map[string]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion `json:"chain_selection" api:"required"`
	Structure      ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure                      `json:"structure" api:"required"`
	Type           constant.StructureTemplate                                                                                                `json:"type" default:"structure_template"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion struct {
	ChainType string `json:"chain_type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	CropResidues ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	EpitopeResidues []int64 `json:"epitope_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	FlexibleResidues []int64 `json:"flexible_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec].
	NonBindingResidues []int64 `json:"non_binding_residues"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
	Ccd string `json:"ccd"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec].
	Smiles string `json:"smiles"`
	JSON   struct {
		ChainType          respjson.Field
		CropResidues       respjson.Field
		EpitopeResidues    respjson.Field
		FlexibleResidues   respjson.Field
		NonBindingResidues respjson.Field
		Ccd                respjson.Field
		Smiles             respjson.Field
		raw                string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a polymer (protein/RNA/DNA) chain in a structure
// template target.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec struct {
	ChainType constant.Polymer `json:"chain_type" default:"polymer"`
	// 0-indexed residue indices to retain from this chain, or 'all' to keep all
	// residues. Residues not listed are excluded from the pipeline run.
	CropResidues ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion `json:"crop_residues" api:"required"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion
// contains all possible properties and values from [[]int64], [constant.All].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfIntArray OfAll]
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion struct {
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

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsIntArray() (v []int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) AsAll() (v constant.All) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-chain specification for a ligand chain in a structure template target. The
// full ligand is always included. An optional CCD or SMILES override preserves the
// ligand chemistry while retaining its coordinates.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec struct {
	ChainType constant.Ligand `json:"chain_type" default:"ligand"`
	// Original CCD identity for this ligand. Use when the structure stores the
	// coordinates under a generic component ID such as LIG. Mutually exclusive with
	// smiles.
	Ccd string `json:"ccd"`
	// Original SMILES identity for this ligand. Use when the structure stores the
	// coordinates under a generic component ID such as LIG. Mutually exclusive with
	// ccd.
	Smiles string `json:"smiles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainType   respjson.Field
		Ccd         respjson.Field
		Smiles      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseChainSelectionStructureTemplateTargetLigandChainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetStructureTemplateTargetResponseStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target defined by sequences only, without a 3D structure template
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse struct {
	// Entities (proteins, RNA, DNA, ligands) defining the target complex.
	Entities []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion `json:"entities" api:"required"`
	Type     constant.NoTemplate                                                                               `json:"type" default:"no_template"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion `json:"constraints"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification],
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	Modifications ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications
// is an implicit subunion of
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion].
// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications
// OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications]
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                                                       struct {
		OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModifications respjson.Field
		OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModifications     respjson.Field
		OfShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModifications     respjson.Field
		raw                                                                                                                            string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                                                         `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token1 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse].
	Token2 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                                                   `json:"type" default:"contact"`
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenInputTargetNoTemplateTargetResponseConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenProgress struct {
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
func (r ShareLinkReadResponsePipelineProteinLibraryScreenProgress) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineProteinLibraryScreenProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineProteinLibraryScreenStatus string

const (
	ShareLinkReadResponsePipelineProteinLibraryScreenStatusPending   ShareLinkReadResponsePipelineProteinLibraryScreenStatus = "pending"
	ShareLinkReadResponsePipelineProteinLibraryScreenStatusRunning   ShareLinkReadResponsePipelineProteinLibraryScreenStatus = "running"
	ShareLinkReadResponsePipelineProteinLibraryScreenStatusSucceeded ShareLinkReadResponsePipelineProteinLibraryScreenStatus = "succeeded"
	ShareLinkReadResponsePipelineProteinLibraryScreenStatusFailed    ShareLinkReadResponsePipelineProteinLibraryScreenStatus = "failed"
	ShareLinkReadResponsePipelineProteinLibraryScreenStatusStopped   ShareLinkReadResponsePipelineProteinLibraryScreenStatus = "stopped"
)

// A small molecule design pipeline run that generates novel molecules
type ShareLinkReadResponsePipelineSmDesignRun struct {
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
	EngineVersion constant.String1_0                            `json:"engine_version" default:"1.0"`
	Error         ShareLinkReadResponsePipelineSmDesignRunError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkReadResponsePipelineSmDesignRunInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule design
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule design
	PipelineVersion constant.String1_0                               `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkReadResponsePipelineSmDesignRunProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                        `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkReadResponsePipelineSmDesignRunStatus `json:"status" api:"required"`
	StoppedAt time.Time                                      `json:"stopped_at" api:"required" format:"date-time"`
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
func (r ShareLinkReadResponsePipelineSmDesignRun) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmDesignRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunError struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmDesignRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkReadResponsePipelineSmDesignRunInput struct {
	// Number of molecules to generate. Must be between 10 and 1,000,000.
	NumMolecules int64 `json:"num_molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target ShareLinkReadResponsePipelineSmDesignRunInputTarget `json:"target" api:"required"`
	// Chemical space to constrain generated molecules. Use 'enamine_real' for the
	// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
	//
	// Any of "enamine_real", "none".
	ChemicalSpace ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpace `json:"chemical_space"`
	// Client-provided key to prevent duplicate submissions on retries
	IdempotencyKey string `json:"idempotency_key"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFilters `json:"molecule_filters"`
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
func (r ShareLinkReadResponsePipelineSmDesignRunInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmDesignRunInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type ShareLinkReadResponsePipelineSmDesignRunInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ShareLinkReadResponsePipelineSmDesignRunInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion `json:"constraints"`
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTarget) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse].
	Modifications []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion) AsShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion) AsShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                        `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse].
	Token1 ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse].
	Token2 ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion) AsShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion) AsShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chemical space to constrain generated molecules. Use 'enamine_real' for the
// Enamine REAL chemical space or 'none' to disable chemical-space filtering.
type ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpace string

const (
	ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpaceEnamineReal ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpace = "enamine_real"
	ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpaceNone        ShareLinkReadResponsePipelineSmDesignRunInputChemicalSpace = "none"
)

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                                  `json:"patterns"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmDesignRunInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunProgress struct {
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
func (r ShareLinkReadResponsePipelineSmDesignRunProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmDesignRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmDesignRunStatus string

const (
	ShareLinkReadResponsePipelineSmDesignRunStatusPending   ShareLinkReadResponsePipelineSmDesignRunStatus = "pending"
	ShareLinkReadResponsePipelineSmDesignRunStatusRunning   ShareLinkReadResponsePipelineSmDesignRunStatus = "running"
	ShareLinkReadResponsePipelineSmDesignRunStatusSucceeded ShareLinkReadResponsePipelineSmDesignRunStatus = "succeeded"
	ShareLinkReadResponsePipelineSmDesignRunStatusFailed    ShareLinkReadResponsePipelineSmDesignRunStatus = "failed"
	ShareLinkReadResponsePipelineSmDesignRunStatusStopped   ShareLinkReadResponsePipelineSmDesignRunStatus = "stopped"
)

// A small molecule library screening pipeline run
type ShareLinkReadResponsePipelineSmScreen struct {
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
	EngineVersion constant.String1_0                         `json:"engine_version" default:"1.0"`
	Error         ShareLinkReadResponsePipelineSmScreenError `json:"error" api:"required"`
	// Pipeline input (null if data deleted)
	Input ShareLinkReadResponsePipelineSmScreenInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Pipeline used for small molecule library screen
	Pipeline constant.Boltzmol `json:"pipeline" default:"boltzmol"`
	// Pipeline version used for small molecule library screen
	PipelineVersion constant.String1_0                            `json:"pipeline_version" default:"1.0"`
	Progress        ShareLinkReadResponsePipelineSmScreenProgress `json:"progress" api:"required"`
	StartedAt       time.Time                                     `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed", "stopped".
	Status    ShareLinkReadResponsePipelineSmScreenStatus `json:"status" api:"required"`
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
func (r ShareLinkReadResponsePipelineSmScreen) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreen) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenError struct {
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
func (r ShareLinkReadResponsePipelineSmScreenError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreenError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pipeline input (null if data deleted)
type ShareLinkReadResponsePipelineSmScreenInput struct {
	Molecules ShareLinkReadResponsePipelineSmScreenInputMolecules `json:"molecules" api:"required"`
	// Target protein sequences for small molecule design or screening.
	Target ShareLinkReadResponsePipelineSmScreenInputTarget `json:"target" api:"required"`
	// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
	// and custom filters.
	MoleculeFilters ShareLinkReadResponsePipelineSmScreenInputMoleculeFilters `json:"molecule_filters"`
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
func (r ShareLinkReadResponsePipelineSmScreenInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreenInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputMolecules struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMolecules) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreenInputMolecules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target protein sequences for small molecule design or screening.
type ShareLinkReadResponsePipelineSmScreenInputTarget struct {
	// Protein and glycan entities defining the target structure. At least one protein
	// entity is required.
	Entities []ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion `json:"entities" api:"required"`
	// Covalent bond constraints between atoms in the target complex. Ligand atom
	// references support CCD atom names and explicitly atom-mapped SMILES atoms.
	Bonds []ShareLinkReadResponsePipelineSmScreenInputTargetBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints []ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion `json:"constraints"`
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
func (r ShareLinkReadResponsePipelineSmScreenInputTarget) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreenInputTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse].
	Value string `json:"value"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse].
	Cyclic bool `json:"cyclic"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse].
	Modifications []ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponseModification `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseResidue `json:"residues"`
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

func (u ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion) AsShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion) AsShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                     `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePipelineSmScreenInputTargetBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBond) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union) AsShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse].
	Token1 ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse].
	Token2 ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion) AsShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion) AsShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                               `json:"type" default:"contact"`
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputTargetConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Molecule filtering configuration. Controls both Boltz built-in SMARTS filtering
// and custom filters.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFilters struct {
	// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
	// which removes molecules matching known problematic substructures. 'recommended'
	// (default): applies a curated set of alerts balancing safety and hit rate.
	// 'extra': adds additional alerts beyond the recommended set for stricter
	// filtering. 'aggressive': applies the most comprehensive alert set — may reject
	// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
	// custom_filters will be applied.
	//
	// Any of "recommended", "extra", "aggressive", "disabled".
	BoltzSmartsCatalogFilterLevel ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel `json:"boltz_smarts_catalog_filter_level"`
	// Custom filters to apply. Molecules must pass all filters (AND logic).
	CustomFilters []ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion `json:"custom_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoltzSmartsCatalogFilterLevel respjson.Field
		CustomFilters                 respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFilters) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the stringency of Boltz's built-in SMARTS structural alert filtering,
// which removes molecules matching known problematic substructures. 'recommended'
// (default): applies a curated set of alerts balancing safety and hit rate.
// 'extra': adds additional alerts beyond the recommended set for stricter
// filtering. 'aggressive': applies the most comprehensive alert set — may reject
// viable molecules. 'disabled': turns off Boltz SMARTS filtering entirely; only
// custom_filters will be applied.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel string

const (
	ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "recommended"
	ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelExtra       ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "extra"
	ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelAggressive  ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "aggressive"
	ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevelDisabled    ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersBoltzSmartsCatalogFilterLevel = "disabled"
)

// ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse],
// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse],
// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse],
// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse],
// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHba float64 `json:"max_hba"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxHbd float64 `json:"max_hbd"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxLogp float64 `json:"max_logp"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	MaxMw float64 `json:"max_mw"`
	Type  string  `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse].
	AllowSingleViolation bool `json:"allow_single_violation"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	FractionCsp3 ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolLogp ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	MolWt ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumAromaticRings ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHAcceptors ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHDonors ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumHeteroatoms ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRings ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	NumRotatableBonds ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse].
	Tpsa     ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
	Patterns []string                                                                                               `json:"patterns"`
	// This field is from variant
	// [ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse].
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

func (u ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse() (v ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse() (v ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse() (v ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse() (v ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) AsShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse() (v ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lipinski's Rule of Five filter. Rejects molecules that violate drug-likeness
// criteria based on molecular weight, LogP, hydrogen bond donors, and hydrogen
// bond acceptors.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterLipinskiFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by RDKit molecular descriptors. Each descriptor is constrained
// to a min/max range. Only descriptors you provide are checked — omitted
// descriptors are unconstrained.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse struct {
	Type constant.RdkitDescriptorFilter `json:"type" default:"rdkit_descriptor_filter"`
	// Min/max range constraint for an RDKit molecular descriptor
	FractionCsp3 ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 `json:"fraction_csp3"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolLogp ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp `json:"mol_logp"`
	// Min/max range constraint for an RDKit molecular descriptor
	MolWt ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt `json:"mol_wt"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumAromaticRings ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings `json:"num_aromatic_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHAcceptors ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors `json:"num_h_acceptors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHDonors ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors `json:"num_h_donors"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumHeteroatoms ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms `json:"num_heteroatoms"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRings ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings `json:"num_rings"`
	// Min/max range constraint for an RDKit molecular descriptor
	NumRotatableBonds ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds `json:"num_rotatable_bonds"`
	// Min/max range constraint for an RDKit molecular descriptor
	Tpsa ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa `json:"tpsa"`
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3 struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseFractionCsp3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolLogp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseMolWt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumAromaticRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHAcceptors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHDonors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumHeteroatoms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseNumRotatableBonds) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Min/max range constraint for an RDKit molecular descriptor
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterRdkitDescriptorFilterResponseTpsa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by custom SMARTS patterns. Molecules matching any pattern are
// rejected.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCustomFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules using a predefined SMARTS catalog of structural alerts.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmartsCatalogFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter molecules by regex patterns on their SMILES representation.
type ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse struct {
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
func (r ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenInputMoleculeFiltersCustomFilterSmilesRegexFilterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenProgress struct {
	// Number of accepted molecules that reached terminal failure during screening.
	NumMoleculesFailed int64 `json:"num_molecules_failed" api:"required"`
	// Number of accepted molecules that produced usable screening results.
	NumMoleculesScreened int64 `json:"num_molecules_screened" api:"required"`
	// Total number of molecules accepted into screening after server-side validation
	// and filtering.
	TotalMoleculesToScreen int64 `json:"total_molecules_to_screen" api:"required"`
	// ID of the most recently screened result
	LatestResultID   string                                                        `json:"latest_result_id"`
	RejectionSummary ShareLinkReadResponsePipelineSmScreenProgressRejectionSummary `json:"rejection_summary"`
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
func (r ShareLinkReadResponsePipelineSmScreenProgress) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePipelineSmScreenProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenProgressRejectionSummary struct {
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
func (r ShareLinkReadResponsePipelineSmScreenProgressRejectionSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePipelineSmScreenProgressRejectionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePipelineSmScreenStatus string

const (
	ShareLinkReadResponsePipelineSmScreenStatusPending   ShareLinkReadResponsePipelineSmScreenStatus = "pending"
	ShareLinkReadResponsePipelineSmScreenStatusRunning   ShareLinkReadResponsePipelineSmScreenStatus = "running"
	ShareLinkReadResponsePipelineSmScreenStatusSucceeded ShareLinkReadResponsePipelineSmScreenStatus = "succeeded"
	ShareLinkReadResponsePipelineSmScreenStatusFailed    ShareLinkReadResponsePipelineSmScreenStatus = "failed"
	ShareLinkReadResponsePipelineSmScreenStatusStopped   ShareLinkReadResponsePipelineSmScreenStatus = "stopped"
)

type ShareLinkReadResponsePipelineStatus string

const (
	ShareLinkReadResponsePipelineStatusPending   ShareLinkReadResponsePipelineStatus = "pending"
	ShareLinkReadResponsePipelineStatusRunning   ShareLinkReadResponsePipelineStatus = "running"
	ShareLinkReadResponsePipelineStatusSucceeded ShareLinkReadResponsePipelineStatus = "succeeded"
	ShareLinkReadResponsePipelineStatusFailed    ShareLinkReadResponsePipelineStatus = "failed"
	ShareLinkReadResponsePipelineStatusStopped   ShareLinkReadResponsePipelineStatus = "stopped"
)

// ShareLinkReadResponsePredictionUnion contains all possible properties and values
// from [ShareLinkReadResponsePredictionBoltz2Prediction],
// [ShareLinkReadResponsePredictionAdmePrediction].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionUnion struct {
	ID            string    `json:"id"`
	CompletedAt   time.Time `json:"completed_at"`
	CreatedAt     time.Time `json:"created_at"`
	DataDeletedAt time.Time `json:"data_deleted_at"`
	// This field is a union of [ShareLinkReadResponsePredictionBoltz2PredictionError],
	// [ShareLinkReadResponsePredictionAdmePredictionError]
	Error     ShareLinkReadResponsePredictionUnionError `json:"error"`
	ExpiresAt time.Time                                 `json:"expires_at"`
	// This field is a union of [ShareLinkReadResponsePredictionBoltz2PredictionInput],
	// [ShareLinkReadResponsePredictionAdmePredictionInput]
	Input    ShareLinkReadResponsePredictionUnionInput `json:"input"`
	Livemode bool                                      `json:"livemode"`
	Model    string                                    `json:"model"`
	// This field is a union of
	// [ShareLinkReadResponsePredictionBoltz2PredictionOutput],
	// [ShareLinkReadResponsePredictionAdmePredictionOutput]
	Output         ShareLinkReadResponsePredictionUnionOutput `json:"output"`
	StartedAt      time.Time                                  `json:"started_at"`
	Status         string                                     `json:"status"`
	Version        string                                     `json:"version"`
	WorkspaceID    string                                     `json:"workspace_id"`
	IdempotencyKey string                                     `json:"idempotency_key"`
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

func (u ShareLinkReadResponsePredictionUnion) AsShareLinkReadResponsePredictionBoltz2Prediction() (v ShareLinkReadResponsePredictionBoltz2Prediction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionUnion) AsShareLinkReadResponsePredictionAdmePrediction() (v ShareLinkReadResponsePredictionAdmePrediction) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionUnion) RawJSON() string { return u.JSON.raw }

func (r *ShareLinkReadResponsePredictionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionUnionError is an implicit subunion of
// [ShareLinkReadResponsePredictionUnion].
// ShareLinkReadResponsePredictionUnionError provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePredictionUnion].
type ShareLinkReadResponsePredictionUnionError struct {
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

func (r *ShareLinkReadResponsePredictionUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionUnionInput is an implicit subunion of
// [ShareLinkReadResponsePredictionUnion].
// ShareLinkReadResponsePredictionUnionInput provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePredictionUnion].
type ShareLinkReadResponsePredictionUnionInput struct {
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	Entities []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion `json:"entities"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	Binding ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion `json:"binding"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	Bonds []ShareLinkReadResponsePredictionBoltz2PredictionInputBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	Constraints []ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion `json:"constraints"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	ModelOptions ShareLinkReadResponsePredictionBoltz2PredictionInputModelOptions `json:"model_options"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	NumSamples int64 `json:"num_samples"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInput].
	Templates []ShareLinkReadResponsePredictionBoltz2PredictionInputTemplate `json:"templates"`
	// This field is from variant [ShareLinkReadResponsePredictionAdmePredictionInput].
	Molecules []ShareLinkReadResponsePredictionAdmePredictionInputMolecule `json:"molecules"`
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

func (r *ShareLinkReadResponsePredictionUnionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionUnionOutput is an implicit subunion of
// [ShareLinkReadResponsePredictionUnion].
// ShareLinkReadResponsePredictionUnionOutput provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePredictionUnion].
type ShareLinkReadResponsePredictionUnionOutput struct {
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionOutput].
	AllSampleResults []ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResult `json:"all_sample_results"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionOutput].
	BestSample ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSample `json:"best_sample"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionOutput].
	Archive ShareLinkReadResponsePredictionBoltz2PredictionOutputArchive `json:"archive"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionOutput].
	BindingMetrics ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion `json:"binding_metrics"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutput].
	Molecules []ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion `json:"molecules"`
	JSON      struct {
		AllSampleResults respjson.Field
		BestSample       respjson.Field
		Archive          respjson.Field
		BindingMetrics   respjson.Field
		Molecules        respjson.Field
		raw              string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePredictionUnionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2Prediction struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error ShareLinkReadResponsePredictionBoltz2PredictionError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input ShareLinkReadResponsePredictionBoltz2PredictionInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.Boltz2_1 `json:"model" default:"boltz-2.1"`
	// Prediction output when succeeded
	Output    ShareLinkReadResponsePredictionBoltz2PredictionOutput `json:"output" api:"required"`
	StartedAt time.Time                                             `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status ShareLinkReadResponsePredictionBoltz2PredictionStatus `json:"status" api:"required"`
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
func (r ShareLinkReadResponsePredictionBoltz2Prediction) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionBoltz2Prediction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type ShareLinkReadResponsePredictionBoltz2PredictionError struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionBoltz2PredictionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type ShareLinkReadResponsePredictionBoltz2PredictionInput struct {
	// Entities (proteins, RNA, DNA, ligands, and glycans) forming the complex to
	// predict. Order determines chain assignment.
	Entities []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion `json:"entities" api:"required"`
	Binding  ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion  `json:"binding"`
	// Request-level covalent bonds between atoms. Use ccd_atom with a glycan residue
	// ID, smiles_atom with a numeric SMILES atom-map, or ligand_atom for a
	// single-residue ligand. Internal glycan bonds belong in the glycan entity bonds
	// field.
	Bonds []ShareLinkReadResponsePredictionBoltz2PredictionInputBond `json:"bonds"`
	// Structural constraints (pocket and contact). Ligand atom references support CCD
	// atom names and explicitly atom-mapped SMILES atoms.
	Constraints  []ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion `json:"constraints"`
	ModelOptions ShareLinkReadResponsePredictionBoltz2PredictionInputModelOptions      `json:"model_options"`
	// Number of structure samples to generate (1-10)
	NumSamples int64 `json:"num_samples"`
	// Template structure files to guide protein-chain prediction. Supports up to 4 CIF
	// or PDB templates from HTTPS URLs or base64 uploads. Use template_chains to map
	// request chains to template-file chains.
	Templates []ShareLinkReadResponsePredictionBoltz2PredictionInputTemplate `json:"templates"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion struct {
	ChainIDs []string `json:"chain_ids"`
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Cyclic   bool     `json:"cyclic"`
	// This field is a union of
	// [[]ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification],
	// [[]ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification],
	// [[]ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification]
	Modifications ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnionModifications `json:"modifications"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse].
	Msa ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse].
	Bonds []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBond `json:"bonds"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse].
	Residues []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseResidue `json:"residues"`
	JSON     struct {
		ChainIDs      respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		Cyclic        respjson.Field
		Modifications respjson.Field
		Msa           respjson.Field
		Bonds         respjson.Field
		Residues      respjson.Field
		raw           string
	} `json:"-"`
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnionModifications is
// an implicit subunion of
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion].
// ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnionModifications
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModifications
// OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModifications
// OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModifications]
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnionModifications struct {
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModifications []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModifications []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification `json:",inline"`
	// This field will be present if the value is a
	// [[]ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification]
	// instead of an object.
	OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModifications []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification `json:",inline"`
	JSON                                                                                       struct {
		OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModifications respjson.Field
		OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModifications           respjson.Field
		OfShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModifications           respjson.Field
		raw                                                                                                  string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityUnionModifications) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string         `json:"chain_ids" api:"required"`
	Type     constant.Protein `json:"type" default:"protein"`
	// Amino acid sequence (one-letter codes)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD post-translational modifications. Optional; defaults to an empty list when
	// omitted. SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification `json:"modifications"`
	// Optional protein MSA control. Omit msa on all protein entities to use automatic
	// MSA generation. Use custom for user-provided A3M/CSV files, or empty for
	// single-sequence mode. Custom MSA and automatic MSA cannot be mixed in one
	// request.
	Msa ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion `json:"msa"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Format ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse].
	Source ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source"`
	Type   string                                                                                                                `json:"type"`
	JSON   struct {
		Format respjson.Field
		Source respjson.Field
		Type   respjson.Field
		raw    string
	} `json:"-"`
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use a user-provided MSA for this protein entity. If any protein entity uses a
// custom MSA, every other protein entity must use either custom or empty MSA;
// automatic MSA generation cannot be mixed with custom MSAs in the same request.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse struct {
	// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
	// text/csv for CSV.
	//
	// Any of "a3m", "csv".
	Format ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat `json:"format" api:"required"`
	Source ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource `json:"source" api:"required"`
	Type   constant.Custom                                                                                                       `json:"type" default:"custom"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat string

const (
	ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatA3m ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "a3m"
	ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormatCsv ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseFormat = "csv"
)

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2CustomMsaResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run this protein entity in single-sequence mode without an MSA. Use this for
// chains that should not use automatic MSA generation, including non-homologous
// chains in a request that also includes custom MSAs.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse struct {
	Type constant.Empty `json:"type" default:"empty"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaBoltz2EmptyMsaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom MSA file format. Base64 uploads must use media_type text/x-a3m for A3M or
// text/csv for CSV.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormat string

const (
	ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormatA3m ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormat = "a3m"
	ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormatCsv ShareLinkReadResponsePredictionBoltz2PredictionInputEntityBoltz2ProteinEntityResponseMsaFormat = "csv"
)

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Rna `json:"type" default:"rna"`
	// RNA nucleotide sequence (A, C, G, U, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityRnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse struct {
	// Chain IDs for this entity
	ChainIDs []string     `json:"chain_ids" api:"required"`
	Type     constant.Dna `json:"type" default:"dna"`
	// DNA nucleotide sequence (A, C, G, T, N)
	Value string `json:"value" api:"required"`
	// Whether the sequence is cyclic
	Cyclic bool `json:"cyclic"`
	// CCD chemical modifications. Optional; defaults to an empty list when omitted.
	// SMILES modifications are not supported.
	Modifications []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification `json:"modifications"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Polymer residue modification. Only CCD codes are supported; SMILES modifications
// are not accepted.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityDnaEntityResponseModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandCcdEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityLigandSmilesEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Branched glycan represented as an explicit graph of CCD monosaccharide residues.
// Declare internal connectivity in this entity and cross-entity attachments in the
// request-level bonds array.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse struct {
	// Internal covalent bonds connecting the glycan residues. A single-residue glycan
	// uses an empty array.
	Bonds []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBond `json:"bonds" api:"required"`
	// Chain IDs for identical copies of this glycan
	ChainIDs []string `json:"chain_ids" api:"required"`
	// CCD residues in the glycan. Array order is not part of the public residue
	// identity; bonds reference residue IDs.
	Residues []ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseResidue `json:"residues" api:"required"`
	Type     constant.Glycan                                                                         `json:"type" default:"glycan"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Internal covalent bond between atoms in two residues of the glycan graph.
type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBond struct {
	Atom1 ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom1 `json:"atom1" api:"required"`
	Atom2 ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom2 `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBond) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom1 struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom1) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom2 struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom2) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseBondAtom2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseResidue struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseResidue) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputEntityGlycanEntityResponseResidue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse].
	BinderChainID string `json:"binder_chain_id"`
	Type          string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse].
	BinderChainIDs []string `json:"binder_chain_ids"`
	JSON           struct {
		BinderChainID  respjson.Field
		Type           respjson.Field
		BinderChainIDs respjson.Field
		raw            string
	} `json:"-"`
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBindingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse struct {
	// Chain ID of the ligand binder (must have exactly 1 copy, at most 2048 heavy
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBindingLigandProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBindingProteinProteinBindingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-level covalent bond between atoms, including protein-glycan attachments.
// Internal glycan connectivity belongs in the glycan entity bonds field.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBond struct {
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom1 ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union `json:"atom1" api:"required"`
	// Atom reference for a specific CCD residue in a glycan graph.
	Atom2 ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union `json:"atom2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Atom1       respjson.Field
		Atom2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBond) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBond) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union contains all
// possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse].
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

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom1LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union contains all
// possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union struct {
	AtomName string `json:"atom_name"`
	ChainID  string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse].
	AtomID string `json:"atom_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse].
	ResidueID string `json:"residue_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse].
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

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2PolymerAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a specific CCD residue in a glycan graph.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2CcdAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference using an explicit numeric atom-map in the input SMILES.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2SmilesAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Atom reference for a single-residue ligand_ccd or an explicitly atom-mapped
// SMILES ligand. Glycan bonds use ccd_atom; new SMILES bonds should use
// smiles_atom.
type ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputBondAtom2LigandAtomResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion struct {
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse].
	BinderChainID string `json:"binder_chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse].
	ContactResidues     map[string][]int64 `json:"contact_residues"`
	MaxDistanceAngstrom float64            `json:"max_distance_angstrom"`
	Type                string             `json:"type"`
	Force               bool               `json:"force"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse].
	Token1 ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union `json:"token1"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse].
	Token2 ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union `json:"token2"`
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

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion) AsShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constrains the binder to interact with specific pocket residues on the target.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintPocketConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maximum-distance contact constraint between two polymer residues or ligand
// atoms.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse struct {
	// Maximum distance in Angstroms
	MaxDistanceAngstrom float64 `json:"max_distance_angstrom" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token1 ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union `json:"token1" api:"required"`
	// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
	Token2 ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union `json:"token2" api:"required"`
	Type   constant.Contact                                                                                   `json:"type" default:"contact"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union
// contains all possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken1LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union
// contains all possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse],
// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union struct {
	ChainID string `json:"chain_id"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse].
	ResidueIndex int64  `json:"residue_index"`
	Type         string `json:"type"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse].
	AtomName string `json:"atom_name"`
	JSON     struct {
		ChainID      respjson.Field
		ResidueIndex respjson.Field
		Type         respjson.Field
		AtomName     respjson.Field
		raw          string
	} `json:"-"`
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) AsShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse() (v ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2PolymerContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ligand contact token for a CCD atom or an explicitly atom-mapped SMILES atom.
type ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputConstraintContactConstraintResponseToken2LigandContactTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputModelOptions struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputModelOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputModelOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template structure used as an inference-time guide for Boltz-2.1 protein-chain
// geometry. Provide a CIF or PDB file from an HTTPS URL or base64 upload.
type ShareLinkReadResponsePredictionBoltz2PredictionInputTemplate struct {
	// Request-to-template chain mappings. Each input_chain_id and template_chain_id
	// must be unique within this template.
	TemplateChains    []ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateChain   `json:"template_chains" api:"required"`
	TemplateStructure ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateStructure `json:"template_structure" api:"required"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputTemplate) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mapping from one request chain to the corresponding chain in the template
// structure file.
type ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateChain struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateChain) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateChain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateStructure struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionInputTemplateTemplateStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type ShareLinkReadResponsePredictionBoltz2PredictionOutput struct {
	// Per-sample structure results
	AllSampleResults []ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResult   `json:"all_sample_results" api:"required"`
	BestSample       ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSample          `json:"best_sample" api:"required"`
	Archive          ShareLinkReadResponsePredictionBoltz2PredictionOutputArchive             `json:"archive"`
	BindingMetrics   ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion `json:"binding_metrics"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResult struct {
	Metrics         ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics         `json:"metrics" api:"required"`
	Structure       ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultStructure       `json:"structure" api:"required"`
	LigandStructure ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure `json:"ligand_structure"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultStructure struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputAllSampleResultLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSample struct {
	Metrics         ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleMetrics         `json:"metrics" api:"required"`
	Structure       ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleStructure       `json:"structure" api:"required"`
	LigandStructure ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure `json:"ligand_structure"`
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSample) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleMetrics struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleStructure struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBestSampleLigandStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputArchive struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputArchive) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputArchive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion
// contains all possible properties and values from
// [ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics],
// [ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion struct {
	BindingConfidence float64 `json:"binding_confidence"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics].
	OptimizationScore float64 `json:"optimization_score"`
	Type              string  `json:"type"`
	JSON              struct {
		BindingConfidence respjson.Field
		OptimizationScore respjson.Field
		Type              respjson.Field
		raw               string
	} `json:"-"`
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) AsShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics() (v ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) AsShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics() (v ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsLigandProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics struct {
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
func (r ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionBoltz2PredictionOutputBindingMetricsProteinProteinBindingMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionBoltz2PredictionStatus string

const (
	ShareLinkReadResponsePredictionBoltz2PredictionStatusPending   ShareLinkReadResponsePredictionBoltz2PredictionStatus = "pending"
	ShareLinkReadResponsePredictionBoltz2PredictionStatusRunning   ShareLinkReadResponsePredictionBoltz2PredictionStatus = "running"
	ShareLinkReadResponsePredictionBoltz2PredictionStatusSucceeded ShareLinkReadResponsePredictionBoltz2PredictionStatus = "succeeded"
	ShareLinkReadResponsePredictionBoltz2PredictionStatusFailed    ShareLinkReadResponsePredictionBoltz2PredictionStatus = "failed"
)

type ShareLinkReadResponsePredictionAdmePrediction struct {
	// Unique prediction identifier
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the input/output data was deleted, or null if still available
	DataDeletedAt time.Time `json:"data_deleted_at" api:"required" format:"date-time"`
	// Error details when failed
	Error ShareLinkReadResponsePredictionAdmePredictionError `json:"error" api:"required"`
	// When this resource and its associated data will be permanently deleted. Null
	// while still in progress.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Prediction input (null if data deleted)
	Input ShareLinkReadResponsePredictionAdmePredictionInput `json:"input" api:"required"`
	// Whether this resource was created with a live API key.
	Livemode bool `json:"livemode" api:"required"`
	// Model used for prediction
	Model constant.AdmeV1 `json:"model" default:"adme-v1"`
	// Prediction output when succeeded
	Output    ShareLinkReadResponsePredictionAdmePredictionOutput `json:"output" api:"required"`
	StartedAt time.Time                                           `json:"started_at" api:"required" format:"date-time"`
	// Any of "pending", "running", "succeeded", "failed".
	Status ShareLinkReadResponsePredictionAdmePredictionStatus `json:"status" api:"required"`
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
func (r ShareLinkReadResponsePredictionAdmePrediction) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionAdmePrediction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when failed
type ShareLinkReadResponsePredictionAdmePredictionError struct {
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
func (r ShareLinkReadResponsePredictionAdmePredictionError) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionAdmePredictionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction input (null if data deleted)
type ShareLinkReadResponsePredictionAdmePredictionInput struct {
	// Molecules to score (1-128 per request). Results are returned in the same order
	// as this list.
	Molecules []ShareLinkReadResponsePredictionAdmePredictionInputMolecule `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePredictionAdmePredictionInput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionAdmePredictionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionAdmePredictionInputMolecule struct {
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
func (r ShareLinkReadResponsePredictionAdmePredictionInputMolecule) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionAdmePredictionInputMolecule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prediction output when succeeded
type ShareLinkReadResponsePredictionAdmePredictionOutput struct {
	// Per-molecule results in the same order as the request. Successful molecules
	// carry an `adme` summary. Failed molecules carry `status: "failed"` and a
	// non-null `error`.
	Molecules []ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion `json:"molecules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Molecules   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLinkReadResponsePredictionAdmePredictionOutput) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkReadResponsePredictionAdmePredictionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion contains all
// possible properties and values from
// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded],
// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion struct {
	ID string `json:"id"`
	// This field is a union of
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme],
	// [any]
	Adme ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionAdme `json:"adme"`
	// This field is a union of [any],
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError]
	Error      ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionError `json:"error"`
	Smiles     string                                                                `json:"smiles"`
	Status     string                                                                `json:"status"`
	ExternalID string                                                                `json:"external_id"`
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

func (u ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion) AsShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded() (v ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion) AsShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed() (v ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionAdme is an
// implicit subunion of
// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion].
// ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionAdme provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedAdme]
type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionAdme struct {
	// This field will be present if the value is a [any] instead of an object.
	OfShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedAdme any `json:",inline"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme].
	Lipophilicity float64 `json:"lipophilicity"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme].
	Permeability float64 `json:"permeability"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme].
	Solubility ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility"`
	JSON       struct {
		OfShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedAdme respjson.Field
		Lipophilicity                                                                       respjson.Field
		Permeability                                                                        respjson.Field
		Solubility                                                                          respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionError is an
// implicit subunion of
// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion].
// ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionError provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededError]
type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionError struct {
	// This field will be present if the value is a [any] instead of an object.
	OfShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededError any `json:",inline"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError].
	Code string `json:"code"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError].
	Message string `json:"message"`
	// This field is from variant
	// [ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError].
	Details any `json:"details"`
	JSON    struct {
		OfShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededError respjson.Field
		Code                                                                                    respjson.Field
		Message                                                                                 respjson.Field
		Details                                                                                 respjson.Field
		raw                                                                                     string
	} `json:"-"`
}

func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded struct {
	// Internally generated molecule identifier.
	ID string `json:"id" api:"required"`
	// Tier 1 ADME summary values for this molecule.
	Adme  ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme `json:"adme" api:"required"`
	Error any                                                                                  `json:"error" api:"required"`
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
func (r ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceeded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tier 1 ADME summary values for this molecule.
type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme struct {
	// Lipophilicity score from the internal LogD prediction.
	Lipophilicity float64 `json:"lipophilicity" api:"required"`
	// Permeability score for this molecule.
	Permeability float64 `json:"permeability" api:"required"`
	// Solubility judgement for this molecule.
	//
	// Any of "high-confidence", "medium-confidence", "high-risk".
	Solubility ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility `json:"solubility" api:"required"`
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
func (r ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Solubility judgement for this molecule.
type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility string

const (
	ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighConfidence   ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-confidence"
	ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityMediumConfidence ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "medium-confidence"
	ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubilityHighRisk         ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeSucceededAdmeSolubility = "high-risk"
)

type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed struct {
	// Internally generated molecule identifier.
	ID    string                                                                             `json:"id" api:"required"`
	Adme  any                                                                                `json:"adme" api:"required"`
	Error ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError `json:"error" api:"required"`
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
func (r ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError struct {
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
func (r ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError) RawJSON() string {
	return r.JSON.raw
}
func (r *ShareLinkReadResponsePredictionAdmePredictionOutputMoleculeAdmeMoleculeFailedError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkReadResponsePredictionAdmePredictionStatus string

const (
	ShareLinkReadResponsePredictionAdmePredictionStatusPending   ShareLinkReadResponsePredictionAdmePredictionStatus = "pending"
	ShareLinkReadResponsePredictionAdmePredictionStatusRunning   ShareLinkReadResponsePredictionAdmePredictionStatus = "running"
	ShareLinkReadResponsePredictionAdmePredictionStatusSucceeded ShareLinkReadResponsePredictionAdmePredictionStatus = "succeeded"
	ShareLinkReadResponsePredictionAdmePredictionStatusFailed    ShareLinkReadResponsePredictionAdmePredictionStatus = "failed"
)

type ShareLinkReadResponsePredictionStatus string

const (
	ShareLinkReadResponsePredictionStatusPending   ShareLinkReadResponsePredictionStatus = "pending"
	ShareLinkReadResponsePredictionStatusRunning   ShareLinkReadResponsePredictionStatus = "running"
	ShareLinkReadResponsePredictionStatusSucceeded ShareLinkReadResponsePredictionStatus = "succeeded"
	ShareLinkReadResponsePredictionStatusFailed    ShareLinkReadResponsePredictionStatus = "failed"
)

type ShareLinkNewParams struct {
	ExpiresAt string `json:"expires_at" api:"required" format:"date-time-string"`
	// Workspace to target. Admin API keys and OAuth callers may select an authorized
	// workspace; for workspace-scoped keys the value must match the key assignment.
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
