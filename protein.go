// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"github.com/boltz-bio/boltz-api-go/option"
)

// Protein Pipeline: design novel protein binders and screen protein libraries
// against targets. Includes de novo protein design and library screening.
//
// ProteinService contains methods and other services that help with interacting
// with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProteinService] method instead.
type ProteinService struct {
	Options []option.RequestOption
	// Generate novel protein binders optimized for binding to a target structure.
	// Binder specifications can be provided directly, uploaded as structure templates,
	// or selected from Boltz-managed curated nanobody and antibody defaults. Results
	// are scored by binding confidence (likelihood of protein-protein interaction) and
	// structure confidence.
	Design ProteinDesignService
	// Screen an existing library of proteins against a target structure. Results are
	// scored by binding confidence (likelihood of protein-protein interaction) and
	// structure confidence.
	LibraryScreen ProteinLibraryScreenService
}

// NewProteinService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProteinService(opts ...option.RequestOption) (r ProteinService) {
	r = ProteinService{}
	r.Options = opts
	r.Design = NewProteinDesignService(opts...)
	r.LibraryScreen = NewProteinLibraryScreenService(opts...)
	return
}
