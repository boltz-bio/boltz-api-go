// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"github.com/boltz-bio/boltz-api-go/option"
)

// Small Molecule Pipeline — design novel small molecules and screen compound
// libraries against protein targets. Includes de novo generation and virtual
// screening.
//
// SmallMoleculeService contains methods and other services that help with
// interacting with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmallMoleculeService] method instead.
type SmallMoleculeService struct {
	Options []option.RequestOption
	// Generate novel small molecules optimized for binding to a protein target.
	// Results are scored by binding confidence (likelihood of binding, for hit
	// discovery), optimization score (binding strength ranking, for lead
	// optimization), and structure confidence.
	Design SmallMoleculeDesignService
	// Screen an existing library of small molecules against a protein target. Results
	// are scored by binding confidence (likelihood of binding, for hit discovery),
	// optimization score (binding strength ranking, for lead optimization), and
	// structure confidence.
	LibraryScreen SmallMoleculeLibraryScreenService
}

// NewSmallMoleculeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSmallMoleculeService(opts ...option.RequestOption) (r SmallMoleculeService) {
	r = SmallMoleculeService{}
	r.Options = opts
	r.Design = NewSmallMoleculeDesignService(opts...)
	r.LibraryScreen = NewSmallMoleculeLibraryScreenService(opts...)
	return
}
