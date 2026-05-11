// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi

import (
	"github.com/boltz-bio/boltz-api-go/option"
)

// Run prediction models on molecular inputs. Each application is available as its
// own endpoint with application-specific inputs and outputs.
//
// PredictionService contains methods and other services that help with interacting
// with the boltz API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPredictionService] method instead.
type PredictionService struct {
	Options []option.RequestOption
	// Predict 3D structure coordinates, per-residue confidence scores, and binding
	// metrics for a molecular complex.
	StructureAndBinding PredictionStructureAndBindingService
	// Predict Tier 1 ADME summary values for a batch of small molecules specified by
	// SMILES.
	Adme PredictionAdmeService
}

// NewPredictionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPredictionService(opts ...option.RequestOption) (r PredictionService) {
	r = PredictionService{}
	r.Options = opts
	r.StructureAndBinding = NewPredictionStructureAndBindingService(opts...)
	r.Adme = NewPredictionAdmeService(opts...)
	return
}
