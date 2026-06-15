// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"github.com/boltz-bio/boltz-api-go/option"
)

// Manage workspaces and API keys. Requires an admin API key. Admin keys have full
// access to all management and compute operations across all workspaces in the
// organization.
//
// AdminService contains methods and other services that help with interacting with
// the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	Options []option.RequestOption
	// Workspaces provide isolated environments for organizing predictions and pipeline
	// runs across teams, projects, or customers. Each workspace has independent data
	// retention settings and can be associated with workspace API keys.
	Workspaces AdminWorkspaceService
	// API keys authenticate requests to the Boltz API. There are two key types: admin
	// keys have full access to all management and compute operations across the
	// organization, while workspace keys are scoped to a single workspace and can only
	// perform compute operations (predictions, protein design, small molecule design)
	// within that workspace. Keys can be created in live or test mode. Test keys
	// (prefixed `sk_bc_*_test_`) create test-mode resources with synthetic data and no
	// GPU cost. Every resource includes a `livemode` field indicating its mode.
	APIKeys AdminAPIKeyService
	// Retrieve aggregated usage data for the organization. Usage can be grouped by
	// workspace and/or application, and filtered by time range, workspace, and
	// application.
	Usage AdminUsageService
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r AdminService) {
	r = AdminService{}
	r.Options = opts
	r.Workspaces = NewAdminWorkspaceService(opts...)
	r.APIKeys = NewAdminAPIKeyService(opts...)
	r.Usage = NewAdminUsageService(opts...)
	return
}
