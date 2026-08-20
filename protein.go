// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"github.com/boltz-bio/boltz-api-go/option"
)

// Design novel protein binders, redesign selected residues in fixed structures,
// and screen protein libraries against targets.
//
// ProteinService contains methods and other services that help with interacting
// with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProteinService] method instead.
type ProteinService struct {
	Options []option.RequestOption
	// Generate binder or generic protein designs. New requests use the top-level type
	// discriminator (`binder` or `generic`), while the legacy target plus
	// binder_specification body remains accepted for migration. Binder requests can
	// share one CIF across target and binder, sample uniformly across multiple
	// specifications, or use Boltz-managed curated antibody and nanobody defaults.
	// Results are discriminated by type: binder runs include binding metrics, while
	// generic runs return structure and secondary-structure metrics only. A generic
	// request can use a `fusion_protein` entity to concatenate two or more ordered
	// fixed, designed, or template-backed protein segments into one output chain.
	Design ProteinDesignService
	// Redesign selected protein residues in one fixed CIF structure. Use the top-level
	// type discriminator to choose binder redesign, with target and binder chain
	// roles, or generic redesign. Every chain in the input structure must be assigned
	// exactly once. Binder results include binding and structure metrics; generic
	// results include structure and secondary-structure metrics.
	SequenceRedesign ProteinSequenceRedesignService
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
	r.SequenceRedesign = NewProteinSequenceRedesignService(opts...)
	r.LibraryScreen = NewProteinLibraryScreenService(opts...)
	return
}
