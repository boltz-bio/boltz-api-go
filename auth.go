// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/boltz-bio/boltz-api-go/internal/apijson"
	"github.com/boltz-bio/boltz-api-go/internal/requestconfig"
	"github.com/boltz-bio/boltz-api-go/option"
	"github.com/boltz-bio/boltz-api-go/packages/respjson"
	"github.com/boltz-bio/boltz-api-go/shared/constant"
)

// Inspect the authentication context for the current credential, including the
// organization or workspace scope for API keys and the available organization
// memberships for OAuth bearer tokens. OAuth callers can use this information to
// choose which organization to send with future requests.
//
// AuthService contains methods and other services that help with interacting with
// the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	return
}

// Returns the organization context available to the current API key or OAuth
// bearer token. OAuth callers can use X-Boltz-Organization-Id to select one of
// their organizations.
func (r *AuthService) Me(ctx context.Context, opts ...option.RequestOption) (res *AuthMeResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/auth/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// AuthMeResponseUnion contains all possible properties and values from
// [AuthMeResponseAuthMeAPIKeyResponse], [AuthMeResponseAuthMeUserResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AuthMeResponseUnion struct {
	// This field is from variant [AuthMeResponseAuthMeAPIKeyResponse].
	APIKeyID string `json:"api_key_id"`
	// This field is from variant [AuthMeResponseAuthMeAPIKeyResponse].
	KeyType AuthMeResponseAuthMeAPIKeyResponseKeyType `json:"key_type"`
	// This field is from variant [AuthMeResponseAuthMeAPIKeyResponse].
	Mode AuthMeResponseAuthMeAPIKeyResponseMode `json:"mode"`
	// This field is from variant [AuthMeResponseAuthMeAPIKeyResponse].
	OrganizationID         string `json:"organization_id"`
	PrincipalType          string `json:"principal_type"`
	SelectedOrganizationID string `json:"selected_organization_id"`
	// This field is from variant [AuthMeResponseAuthMeAPIKeyResponse].
	WorkspaceID string `json:"workspace_id"`
	// This field is from variant [AuthMeResponseAuthMeAPIKeyResponse].
	OrganizationName string `json:"organization_name"`
	// This field is from variant [AuthMeResponseAuthMeUserResponse].
	ActiveOrganizationID string `json:"active_organization_id"`
	// This field is from variant [AuthMeResponseAuthMeUserResponse].
	OrganizationMemberships []AuthMeResponseAuthMeUserResponseOrganizationMembership `json:"organization_memberships"`
	// This field is from variant [AuthMeResponseAuthMeUserResponse].
	UserID string `json:"user_id"`
	JSON   struct {
		APIKeyID                respjson.Field
		KeyType                 respjson.Field
		Mode                    respjson.Field
		OrganizationID          respjson.Field
		PrincipalType           respjson.Field
		SelectedOrganizationID  respjson.Field
		WorkspaceID             respjson.Field
		OrganizationName        respjson.Field
		ActiveOrganizationID    respjson.Field
		OrganizationMemberships respjson.Field
		UserID                  respjson.Field
		raw                     string
	} `json:"-"`
}

func (u AuthMeResponseUnion) AsAuthMeResponseAuthMeAPIKeyResponse() (v AuthMeResponseAuthMeAPIKeyResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuthMeResponseUnion) AsAuthMeResponseAuthMeUserResponse() (v AuthMeResponseAuthMeUserResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuthMeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AuthMeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthMeResponseAuthMeAPIKeyResponse struct {
	APIKeyID string `json:"api_key_id" api:"required"`
	// Any of "admin", "workspace".
	KeyType AuthMeResponseAuthMeAPIKeyResponseKeyType `json:"key_type" api:"required"`
	// Any of "live", "test".
	Mode                   AuthMeResponseAuthMeAPIKeyResponseMode `json:"mode" api:"required"`
	OrganizationID         string                                 `json:"organization_id" api:"required"`
	PrincipalType          constant.APIKey                        `json:"principal_type" default:"api_key"`
	SelectedOrganizationID string                                 `json:"selected_organization_id" api:"required"`
	WorkspaceID            string                                 `json:"workspace_id" api:"required"`
	OrganizationName       string                                 `json:"organization_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyID               respjson.Field
		KeyType                respjson.Field
		Mode                   respjson.Field
		OrganizationID         respjson.Field
		PrincipalType          respjson.Field
		SelectedOrganizationID respjson.Field
		WorkspaceID            respjson.Field
		OrganizationName       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthMeResponseAuthMeAPIKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthMeResponseAuthMeAPIKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthMeResponseAuthMeAPIKeyResponseKeyType string

const (
	AuthMeResponseAuthMeAPIKeyResponseKeyTypeAdmin     AuthMeResponseAuthMeAPIKeyResponseKeyType = "admin"
	AuthMeResponseAuthMeAPIKeyResponseKeyTypeWorkspace AuthMeResponseAuthMeAPIKeyResponseKeyType = "workspace"
)

type AuthMeResponseAuthMeAPIKeyResponseMode string

const (
	AuthMeResponseAuthMeAPIKeyResponseModeLive AuthMeResponseAuthMeAPIKeyResponseMode = "live"
	AuthMeResponseAuthMeAPIKeyResponseModeTest AuthMeResponseAuthMeAPIKeyResponseMode = "test"
)

type AuthMeResponseAuthMeUserResponse struct {
	ActiveOrganizationID    string                                                   `json:"active_organization_id" api:"required"`
	OrganizationMemberships []AuthMeResponseAuthMeUserResponseOrganizationMembership `json:"organization_memberships" api:"required"`
	PrincipalType           constant.User                                            `json:"principal_type" default:"user"`
	SelectedOrganizationID  string                                                   `json:"selected_organization_id" api:"required"`
	UserID                  string                                                   `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveOrganizationID    respjson.Field
		OrganizationMemberships respjson.Field
		PrincipalType           respjson.Field
		SelectedOrganizationID  respjson.Field
		UserID                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthMeResponseAuthMeUserResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthMeResponseAuthMeUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthMeResponseAuthMeUserResponseOrganizationMembership struct {
	// Any of "admin", "member".
	ComputeRole      AuthMeResponseAuthMeUserResponseOrganizationMembershipComputeRole `json:"compute_role" api:"required"`
	OrganizationID   string                                                            `json:"organization_id" api:"required"`
	OrganizationName string                                                            `json:"organization_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComputeRole      respjson.Field
		OrganizationID   respjson.Field
		OrganizationName respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthMeResponseAuthMeUserResponseOrganizationMembership) RawJSON() string { return r.JSON.raw }
func (r *AuthMeResponseAuthMeUserResponseOrganizationMembership) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthMeResponseAuthMeUserResponseOrganizationMembershipComputeRole string

const (
	AuthMeResponseAuthMeUserResponseOrganizationMembershipComputeRoleAdmin  AuthMeResponseAuthMeUserResponseOrganizationMembershipComputeRole = "admin"
	AuthMeResponseAuthMeUserResponseOrganizationMembershipComputeRoleMember AuthMeResponseAuthMeUserResponseOrganizationMembershipComputeRole = "member"
)

type AuthMeResponseKeyType string

const (
	AuthMeResponseKeyTypeAdmin     AuthMeResponseKeyType = "admin"
	AuthMeResponseKeyTypeWorkspace AuthMeResponseKeyType = "workspace"
)

type AuthMeResponseMode string

const (
	AuthMeResponseModeLive AuthMeResponseMode = "live"
	AuthMeResponseModeTest AuthMeResponseMode = "test"
)
