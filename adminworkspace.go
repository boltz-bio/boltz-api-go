// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"context"
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

// Workspaces provide isolated environments for organizing predictions and engine
// runs across teams, projects, or customers. Each workspace has independent data
// retention settings and can be associated with workspace API keys.
//
// AdminWorkspaceService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminWorkspaceService] method instead.
type AdminWorkspaceService struct {
	Options []option.RequestOption
}

// NewAdminWorkspaceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminWorkspaceService(opts ...option.RequestOption) (r AdminWorkspaceService) {
	r = AdminWorkspaceService{}
	r.Options = opts
	return
}

// Create a workspace
func (r *AdminWorkspaceService) New(ctx context.Context, body AdminWorkspaceNewParams, opts ...option.RequestOption) (res *AdminWorkspaceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/admin/workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a workspace
func (r *AdminWorkspaceService) Get(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *AdminWorkspaceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/admin/workspaces/%s", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a workspace
func (r *AdminWorkspaceService) Update(ctx context.Context, workspaceID string, body AdminWorkspaceUpdateParams, opts ...option.RequestOption) (res *AdminWorkspaceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/admin/workspaces/%s", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List workspaces
func (r *AdminWorkspaceService) List(ctx context.Context, query AdminWorkspaceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AdminWorkspaceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/admin/workspaces"
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

// List workspaces
func (r *AdminWorkspaceService) ListAutoPaging(ctx context.Context, query AdminWorkspaceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AdminWorkspaceListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Archives a workspace and deactivates all its API keys. This action is
// irreversible.
func (r *AdminWorkspaceService) Archive(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *AdminWorkspaceArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/admin/workspaces/%s/archive", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get a workspace spending limit
func (r *AdminWorkspaceService) GetSpendingLimit(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *AdminWorkspaceGetSpendingLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/admin/workspaces/%s/spending-limit", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set a workspace spending limit
func (r *AdminWorkspaceService) SetSpendingLimit(ctx context.Context, workspaceID string, body AdminWorkspaceSetSpendingLimitParams, opts ...option.RequestOption) (res *AdminWorkspaceSetSpendingLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/admin/workspaces/%s/spending-limit", url.PathEscape(workspaceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AdminWorkspaceNewResponse struct {
	// Workspace ID
	ID         string    `json:"id" api:"required"`
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceNewResponseDataRetention `json:"data_retention" api:"required"`
	// Whether this is the default workspace
	IsDefault bool `json:"is_default" api:"required"`
	// Workspace name
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		DataRetention respjson.Field
		IsDefault     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
type AdminWorkspaceNewResponseDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceNewResponseDataRetentionUnit `json:"unit" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceNewResponseDataRetention) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceNewResponseDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceNewResponseDataRetentionUnit string

const (
	AdminWorkspaceNewResponseDataRetentionUnitHours AdminWorkspaceNewResponseDataRetentionUnit = "hours"
	AdminWorkspaceNewResponseDataRetentionUnitDays  AdminWorkspaceNewResponseDataRetentionUnit = "days"
)

type AdminWorkspaceGetResponse struct {
	// Workspace ID
	ID         string    `json:"id" api:"required"`
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceGetResponseDataRetention `json:"data_retention" api:"required"`
	// Whether this is the default workspace
	IsDefault bool `json:"is_default" api:"required"`
	// Workspace name
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		DataRetention respjson.Field
		IsDefault     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
type AdminWorkspaceGetResponseDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceGetResponseDataRetentionUnit `json:"unit" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceGetResponseDataRetention) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceGetResponseDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceGetResponseDataRetentionUnit string

const (
	AdminWorkspaceGetResponseDataRetentionUnitHours AdminWorkspaceGetResponseDataRetentionUnit = "hours"
	AdminWorkspaceGetResponseDataRetentionUnitDays  AdminWorkspaceGetResponseDataRetentionUnit = "days"
)

type AdminWorkspaceUpdateResponse struct {
	// Workspace ID
	ID         string    `json:"id" api:"required"`
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceUpdateResponseDataRetention `json:"data_retention" api:"required"`
	// Whether this is the default workspace
	IsDefault bool `json:"is_default" api:"required"`
	// Workspace name
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		DataRetention respjson.Field
		IsDefault     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
type AdminWorkspaceUpdateResponseDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceUpdateResponseDataRetentionUnit `json:"unit" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceUpdateResponseDataRetention) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceUpdateResponseDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceUpdateResponseDataRetentionUnit string

const (
	AdminWorkspaceUpdateResponseDataRetentionUnitHours AdminWorkspaceUpdateResponseDataRetentionUnit = "hours"
	AdminWorkspaceUpdateResponseDataRetentionUnitDays  AdminWorkspaceUpdateResponseDataRetentionUnit = "days"
)

type AdminWorkspaceListResponse struct {
	// Workspace ID
	ID         string    `json:"id" api:"required"`
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceListResponseDataRetention `json:"data_retention" api:"required"`
	// Whether this is the default workspace
	IsDefault bool `json:"is_default" api:"required"`
	// Workspace name
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		DataRetention respjson.Field
		IsDefault     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
type AdminWorkspaceListResponseDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceListResponseDataRetentionUnit `json:"unit" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceListResponseDataRetention) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceListResponseDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceListResponseDataRetentionUnit string

const (
	AdminWorkspaceListResponseDataRetentionUnitHours AdminWorkspaceListResponseDataRetentionUnit = "hours"
	AdminWorkspaceListResponseDataRetentionUnitDays  AdminWorkspaceListResponseDataRetentionUnit = "days"
)

type AdminWorkspaceArchiveResponse struct {
	// Workspace ID
	ID         string    `json:"id" api:"required"`
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceArchiveResponseDataRetention `json:"data_retention" api:"required"`
	// Whether this is the default workspace
	IsDefault bool `json:"is_default" api:"required"`
	// Workspace name
	Name      string    `json:"name" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ArchivedAt    respjson.Field
		CreatedAt     respjson.Field
		DataRetention respjson.Field
		IsDefault     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
type AdminWorkspaceArchiveResponseDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceArchiveResponseDataRetentionUnit `json:"unit" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceArchiveResponseDataRetention) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceArchiveResponseDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceArchiveResponseDataRetentionUnit string

const (
	AdminWorkspaceArchiveResponseDataRetentionUnitHours AdminWorkspaceArchiveResponseDataRetentionUnit = "hours"
	AdminWorkspaceArchiveResponseDataRetentionUnitDays  AdminWorkspaceArchiveResponseDataRetentionUnit = "days"
)

// Configured lifetime workspace spending limit, or null if unset. Unset workspaces
// have no workspace-level cap and continue to use organization-level billing.
type AdminWorkspaceGetSpendingLimitResponse struct {
	Limit AdminWorkspaceGetSpendingLimitResponseLimit `json:"limit" api:"required"`
	Type  constant.Lifetime                           `json:"type" default:"lifetime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceGetSpendingLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceGetSpendingLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminWorkspaceGetSpendingLimitResponseLimit struct {
	// Workspace spending limit amount in milli-USD. Tracking starts when the limit is
	// configured; prior or already-committed unreserved work is not counted in this
	// workspace cap ledger.
	Amount int64 `json:"amount" api:"required"`
	// Workspace spending limits currently support milli-USD only.
	Currency constant.MilliUsd `json:"currency" default:"MILLI_USD"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceGetSpendingLimitResponseLimit) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceGetSpendingLimitResponseLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured lifetime workspace spending limit, or null if unset. Unset workspaces
// have no workspace-level cap and continue to use organization-level billing.
type AdminWorkspaceSetSpendingLimitResponse struct {
	Limit AdminWorkspaceSetSpendingLimitResponseLimit `json:"limit" api:"required"`
	Type  constant.Lifetime                           `json:"type" default:"lifetime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceSetSpendingLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceSetSpendingLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminWorkspaceSetSpendingLimitResponseLimit struct {
	// Workspace spending limit amount in milli-USD. Tracking starts when the limit is
	// configured; prior or already-committed unreserved work is not counted in this
	// workspace cap ledger.
	Amount int64 `json:"amount" api:"required"`
	// Workspace spending limits currently support milli-USD only.
	Currency constant.MilliUsd `json:"currency" default:"MILLI_USD"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminWorkspaceSetSpendingLimitResponseLimit) RawJSON() string { return r.JSON.raw }
func (r *AdminWorkspaceSetSpendingLimitResponseLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminWorkspaceNewParams struct {
	// Workspace name
	Name param.Opt[string] `json:"name,omitzero"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceNewParamsDataRetention `json:"data_retention,omitzero"`
	SpendingLimit AdminWorkspaceNewParamsSpendingLimit `json:"spending_limit,omitzero"`
	paramObj
}

func (r AdminWorkspaceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
//
// The properties Unit, Value are required.
type AdminWorkspaceNewParamsDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceNewParamsDataRetentionUnit `json:"unit,omitzero" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	paramObj
}

func (r AdminWorkspaceNewParamsDataRetention) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceNewParamsDataRetention
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceNewParamsDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceNewParamsDataRetentionUnit string

const (
	AdminWorkspaceNewParamsDataRetentionUnitHours AdminWorkspaceNewParamsDataRetentionUnit = "hours"
	AdminWorkspaceNewParamsDataRetentionUnitDays  AdminWorkspaceNewParamsDataRetentionUnit = "days"
)

// The properties Limit, Type are required.
type AdminWorkspaceNewParamsSpendingLimit struct {
	Limit AdminWorkspaceNewParamsSpendingLimitLimit `json:"limit,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "lifetime".
	Type constant.Lifetime `json:"type" default:"lifetime"`
	paramObj
}

func (r AdminWorkspaceNewParamsSpendingLimit) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceNewParamsSpendingLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceNewParamsSpendingLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Currency are required.
type AdminWorkspaceNewParamsSpendingLimitLimit struct {
	// Workspace spending limit amount in milli-USD. Tracking starts when the limit is
	// configured; prior or already-committed unreserved work is not counted in this
	// workspace cap ledger.
	Amount int64 `json:"amount" api:"required"`
	// Workspace spending limits currently support milli-USD only.
	//
	// This field can be elided, and will marshal its zero value as "MILLI_USD".
	Currency constant.MilliUsd `json:"currency" default:"MILLI_USD"`
	paramObj
}

func (r AdminWorkspaceNewParamsSpendingLimitLimit) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceNewParamsSpendingLimitLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceNewParamsSpendingLimitLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminWorkspaceUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// How long result data is retained before automatic deletion. Defaults to 7 days
	// if not specified. Maximum retention is 14 days (336 hours).
	DataRetention AdminWorkspaceUpdateParamsDataRetention `json:"data_retention,omitzero"`
	paramObj
}

func (r AdminWorkspaceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How long result data is retained before automatic deletion. Defaults to 7 days
// if not specified. Maximum retention is 14 days (336 hours).
//
// The properties Unit, Value are required.
type AdminWorkspaceUpdateParamsDataRetention struct {
	// Time unit for retention duration
	//
	// Any of "hours", "days".
	Unit AdminWorkspaceUpdateParamsDataRetentionUnit `json:"unit,omitzero" api:"required"`
	// Duration value. Maximum retention is 14 days (or 336 hours).
	Value int64 `json:"value" api:"required"`
	paramObj
}

func (r AdminWorkspaceUpdateParamsDataRetention) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceUpdateParamsDataRetention
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceUpdateParamsDataRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time unit for retention duration
type AdminWorkspaceUpdateParamsDataRetentionUnit string

const (
	AdminWorkspaceUpdateParamsDataRetentionUnitHours AdminWorkspaceUpdateParamsDataRetentionUnit = "hours"
	AdminWorkspaceUpdateParamsDataRetentionUnitDays  AdminWorkspaceUpdateParamsDataRetentionUnit = "days"
)

type AdminWorkspaceListParams struct {
	// Return results after this ID
	AfterID param.Opt[string] `query:"after_id,omitzero" json:"-"`
	// Return results before this ID
	BeforeID param.Opt[string] `query:"before_id,omitzero" json:"-"`
	// Max items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Case-insensitive workspace name prefix to filter by
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminWorkspaceListParams]'s query parameters as
// `url.Values`.
func (r AdminWorkspaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminWorkspaceSetSpendingLimitParams struct {
	Limit AdminWorkspaceSetSpendingLimitParamsLimit `json:"limit,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "lifetime".
	Type constant.Lifetime `json:"type" default:"lifetime"`
	paramObj
}

func (r AdminWorkspaceSetSpendingLimitParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceSetSpendingLimitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceSetSpendingLimitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Currency are required.
type AdminWorkspaceSetSpendingLimitParamsLimit struct {
	// Workspace spending limit amount in milli-USD. Tracking starts when the limit is
	// configured; prior or already-committed unreserved work is not counted in this
	// workspace cap ledger.
	Amount int64 `json:"amount" api:"required"`
	// Workspace spending limits currently support milli-USD only.
	//
	// This field can be elided, and will marshal its zero value as "MILLI_USD".
	Currency constant.MilliUsd `json:"currency" default:"MILLI_USD"`
	paramObj
}

func (r AdminWorkspaceSetSpendingLimitParamsLimit) MarshalJSON() (data []byte, err error) {
	type shadow AdminWorkspaceSetSpendingLimitParamsLimit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminWorkspaceSetSpendingLimitParamsLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
