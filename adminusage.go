// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"context"
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
)

// Retrieve aggregated usage data for the organization. Usage can be grouped by
// workspace and/or application, and filtered by time range, workspace, and
// application.
//
// AdminUsageService contains methods and other services that help with interacting
// with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminUsageService] method instead.
type AdminUsageService struct {
	Options []option.RequestOption
}

// NewAdminUsageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminUsageService(opts ...option.RequestOption) (r AdminUsageService) {
	r = AdminUsageService{}
	r.Options = opts
	return
}

// Retrieve aggregated usage data across the organization, optionally grouped by
// workspace and/or application.
func (r *AdminUsageService) List(ctx context.Context, query AdminUsageListParams, opts ...option.RequestOption) (res *pagination.OpaqueCursorPage[AdminUsageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/admin/usage"
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

// Retrieve aggregated usage data across the organization, optionally grouped by
// workspace and/or application.
func (r *AdminUsageService) ListAutoPaging(ctx context.Context, query AdminUsageListParams, opts ...option.RequestOption) *pagination.OpaqueCursorPageAutoPager[AdminUsageListResponse] {
	return pagination.NewOpaqueCursorPageAutoPager(r.List(ctx, query, opts...))
}

type AdminUsageListResponse struct {
	EndTime time.Time `json:"end_time" api:"required" format:"date-time"`
	// Aggregated billed quantity for this bucket
	Quantity  int64     `json:"quantity" api:"required"`
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// Any of "structure_and_binding", "small_molecule_design",
	// "small_molecule_library_screen", "protein_design", "protein_redesign",
	// "protein_library_screen", "adme".
	Application AdminUsageListResponseApplication `json:"application"`
	// Present when grouped by workspace_id
	WorkspaceID string `json:"workspace_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		Quantity    respjson.Field
		StartTime   respjson.Field
		Application respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminUsageListResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminUsageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminUsageListResponseApplication string

const (
	AdminUsageListResponseApplicationStructureAndBinding        AdminUsageListResponseApplication = "structure_and_binding"
	AdminUsageListResponseApplicationSmallMoleculeDesign        AdminUsageListResponseApplication = "small_molecule_design"
	AdminUsageListResponseApplicationSmallMoleculeLibraryScreen AdminUsageListResponseApplication = "small_molecule_library_screen"
	AdminUsageListResponseApplicationProteinDesign              AdminUsageListResponseApplication = "protein_design"
	AdminUsageListResponseApplicationProteinRedesign            AdminUsageListResponseApplication = "protein_redesign"
	AdminUsageListResponseApplicationProteinLibraryScreen       AdminUsageListResponseApplication = "protein_library_screen"
	AdminUsageListResponseApplicationAdme                       AdminUsageListResponseApplication = "adme"
)

type AdminUsageListParams struct {
	// End of the time range as an ISO 8601 date-time with timezone, for example
	// 2026-04-08T18:56:46Z
	EndingAt time.Time `query:"ending_at" api:"required" format:"date-time" json:"-"`
	// Start of the time range as an ISO 8601 date-time with timezone, for example
	// 2026-04-08T18:56:46Z
	StartingAt time.Time `query:"starting_at" api:"required" format:"date-time" json:"-"`
	// Time window size. HOUR supports up to 31 days per query; DAY supports up to 365
	// days per query.
	//
	// Any of "HOUR", "DAY".
	WindowSize AdminUsageListParamsWindowSize `query:"window_size,omitzero" api:"required" json:"-"`
	// Maximum number of buckets to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Cursor for pagination
	Page param.Opt[string] `query:"page,omitzero" json:"-"`
	// Filter to specific applications
	Applications AdminUsageListParamsApplicationsUnion `query:"applications[],omitzero" json:"-"`
	// Group results by workspace_id and/or application
	GroupBy AdminUsageListParamsGroupByUnion `query:"group_by[],omitzero" json:"-"`
	// Filter to specific workspace IDs
	WorkspaceIDs AdminUsageListParamsWorkspaceIDsUnion `query:"workspace_ids[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminUsageListParams]'s query parameters as `url.Values`.
func (r AdminUsageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time window size. HOUR supports up to 31 days per query; DAY supports up to 365
// days per query.
type AdminUsageListParamsWindowSize string

const (
	AdminUsageListParamsWindowSizeHour AdminUsageListParamsWindowSize = "HOUR"
	AdminUsageListParamsWindowSizeDay  AdminUsageListParamsWindowSize = "DAY"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdminUsageListParamsApplicationsUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAdminUsageListsApplicationsString)
	OfAdminUsageListsApplicationsString         param.Opt[AdminUsageListParamsApplicationsString] `query:",omitzero,inline"`
	OfAdminUsageListsApplicationsArrayItemArray []AdminUsageListParamsApplicationsArrayItem       `query:",omitzero,inline"`
	paramUnion
}

type AdminUsageListParamsApplicationsString string

const (
	AdminUsageListParamsApplicationsStringStructureAndBinding        AdminUsageListParamsApplicationsString = "structure_and_binding"
	AdminUsageListParamsApplicationsStringSmallMoleculeDesign        AdminUsageListParamsApplicationsString = "small_molecule_design"
	AdminUsageListParamsApplicationsStringSmallMoleculeLibraryScreen AdminUsageListParamsApplicationsString = "small_molecule_library_screen"
	AdminUsageListParamsApplicationsStringProteinDesign              AdminUsageListParamsApplicationsString = "protein_design"
	AdminUsageListParamsApplicationsStringProteinRedesign            AdminUsageListParamsApplicationsString = "protein_redesign"
	AdminUsageListParamsApplicationsStringProteinLibraryScreen       AdminUsageListParamsApplicationsString = "protein_library_screen"
	AdminUsageListParamsApplicationsStringAdme                       AdminUsageListParamsApplicationsString = "adme"
)

type AdminUsageListParamsApplicationsArrayItem string

const (
	AdminUsageListParamsApplicationsArrayItemStructureAndBinding        AdminUsageListParamsApplicationsArrayItem = "structure_and_binding"
	AdminUsageListParamsApplicationsArrayItemSmallMoleculeDesign        AdminUsageListParamsApplicationsArrayItem = "small_molecule_design"
	AdminUsageListParamsApplicationsArrayItemSmallMoleculeLibraryScreen AdminUsageListParamsApplicationsArrayItem = "small_molecule_library_screen"
	AdminUsageListParamsApplicationsArrayItemProteinDesign              AdminUsageListParamsApplicationsArrayItem = "protein_design"
	AdminUsageListParamsApplicationsArrayItemProteinRedesign            AdminUsageListParamsApplicationsArrayItem = "protein_redesign"
	AdminUsageListParamsApplicationsArrayItemProteinLibraryScreen       AdminUsageListParamsApplicationsArrayItem = "protein_library_screen"
	AdminUsageListParamsApplicationsArrayItemAdme                       AdminUsageListParamsApplicationsArrayItem = "adme"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdminUsageListParamsGroupByUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAdminUsageListsGroupByString)
	OfAdminUsageListsGroupByString         param.Opt[AdminUsageListParamsGroupByString] `query:",omitzero,inline"`
	OfAdminUsageListsGroupByArrayItemArray []AdminUsageListParamsGroupByArrayItem       `query:",omitzero,inline"`
	paramUnion
}

type AdminUsageListParamsGroupByString string

const (
	AdminUsageListParamsGroupByStringWorkspaceID AdminUsageListParamsGroupByString = "workspace_id"
	AdminUsageListParamsGroupByStringApplication AdminUsageListParamsGroupByString = "application"
)

type AdminUsageListParamsGroupByArrayItem string

const (
	AdminUsageListParamsGroupByArrayItemWorkspaceID AdminUsageListParamsGroupByArrayItem = "workspace_id"
	AdminUsageListParamsGroupByArrayItemApplication AdminUsageListParamsGroupByArrayItem = "application"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AdminUsageListParamsWorkspaceIDsUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}
