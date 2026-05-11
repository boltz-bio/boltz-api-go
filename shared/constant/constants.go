// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/boltz-bio/boltz-api-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type AdmeV1 string                       // Always "adme-v1"
type All string                          // Always "all"
type APIKey string                       // Always "api_key"
type Base64 string                       // Always "base64"
type BoltzCurated string                 // Always "boltz_curated"
type Boltz2_1 string                     // Always "boltz-2.1"
type BoltzProteinDesign string           // Always "boltz-protein-design"
type BoltzProteinScreen string           // Always "boltz-protein-screen"
type BoltzSmDesign string                // Always "boltz-sm-design"
type BoltzSmScreen string                // Always "boltz-sm-screen"
type Ccd string                          // Always "ccd"
type ChemicalXCif string                 // Always "chemical/x-cif"
type Contact string                      // Always "contact"
type DesignedProtein string              // Always "designed_protein"
type Dna string                          // Always "dna"
type Failed string                       // Always "failed"
type Insertion string                    // Always "insertion"
type Ligand string                       // Always "ligand"
type LigandAtom string                   // Always "ligand_atom"
type LigandCcd string                    // Always "ligand_ccd"
type LigandContact string                // Always "ligand_contact"
type LigandProteinBinding string         // Always "ligand_protein_binding"
type LigandProteinBindingMetrics string  // Always "ligand_protein_binding_metrics"
type LigandSmiles string                 // Always "ligand_smiles"
type LipinskiFilter string               // Always "lipinski_filter"
type NoTemplate string                   // Always "no_template"
type Pocket string                       // Always "pocket"
type Polymer string                      // Always "polymer"
type PolymerAtom string                  // Always "polymer_atom"
type PolymerContact string               // Always "polymer_contact"
type Protein string                      // Always "protein"
type ProteinProteinBinding string        // Always "protein_protein_binding"
type ProteinProteinBindingMetrics string // Always "protein_protein_binding_metrics"
type RdkitDescriptorFilter string        // Always "rdkit_descriptor_filter"
type Replacement string                  // Always "replacement"
type Rna string                          // Always "rna"
type SmartsCatalogFilter string          // Always "smarts_catalog_filter"
type SmartsCustomFilter string           // Always "smarts_custom_filter"
type Smiles string                       // Always "smiles"
type SmilesRegexFilter string            // Always "smiles_regex_filter"
type StructureTemplate string            // Always "structure_template"
type Succeeded string                    // Always "succeeded"
type URL string                          // Always "url"
type User string                         // Always "user"
type Workspace string                    // Always "workspace"

func (c AdmeV1) Default() AdmeV1                             { return "adme-v1" }
func (c All) Default() All                                   { return "all" }
func (c APIKey) Default() APIKey                             { return "api_key" }
func (c Base64) Default() Base64                             { return "base64" }
func (c BoltzCurated) Default() BoltzCurated                 { return "boltz_curated" }
func (c Boltz2_1) Default() Boltz2_1                         { return "boltz-2.1" }
func (c BoltzProteinDesign) Default() BoltzProteinDesign     { return "boltz-protein-design" }
func (c BoltzProteinScreen) Default() BoltzProteinScreen     { return "boltz-protein-screen" }
func (c BoltzSmDesign) Default() BoltzSmDesign               { return "boltz-sm-design" }
func (c BoltzSmScreen) Default() BoltzSmScreen               { return "boltz-sm-screen" }
func (c Ccd) Default() Ccd                                   { return "ccd" }
func (c ChemicalXCif) Default() ChemicalXCif                 { return "chemical/x-cif" }
func (c Contact) Default() Contact                           { return "contact" }
func (c DesignedProtein) Default() DesignedProtein           { return "designed_protein" }
func (c Dna) Default() Dna                                   { return "dna" }
func (c Failed) Default() Failed                             { return "failed" }
func (c Insertion) Default() Insertion                       { return "insertion" }
func (c Ligand) Default() Ligand                             { return "ligand" }
func (c LigandAtom) Default() LigandAtom                     { return "ligand_atom" }
func (c LigandCcd) Default() LigandCcd                       { return "ligand_ccd" }
func (c LigandContact) Default() LigandContact               { return "ligand_contact" }
func (c LigandProteinBinding) Default() LigandProteinBinding { return "ligand_protein_binding" }
func (c LigandProteinBindingMetrics) Default() LigandProteinBindingMetrics {
	return "ligand_protein_binding_metrics"
}
func (c LigandSmiles) Default() LigandSmiles                   { return "ligand_smiles" }
func (c LipinskiFilter) Default() LipinskiFilter               { return "lipinski_filter" }
func (c NoTemplate) Default() NoTemplate                       { return "no_template" }
func (c Pocket) Default() Pocket                               { return "pocket" }
func (c Polymer) Default() Polymer                             { return "polymer" }
func (c PolymerAtom) Default() PolymerAtom                     { return "polymer_atom" }
func (c PolymerContact) Default() PolymerContact               { return "polymer_contact" }
func (c Protein) Default() Protein                             { return "protein" }
func (c ProteinProteinBinding) Default() ProteinProteinBinding { return "protein_protein_binding" }
func (c ProteinProteinBindingMetrics) Default() ProteinProteinBindingMetrics {
	return "protein_protein_binding_metrics"
}
func (c RdkitDescriptorFilter) Default() RdkitDescriptorFilter { return "rdkit_descriptor_filter" }
func (c Replacement) Default() Replacement                     { return "replacement" }
func (c Rna) Default() Rna                                     { return "rna" }
func (c SmartsCatalogFilter) Default() SmartsCatalogFilter     { return "smarts_catalog_filter" }
func (c SmartsCustomFilter) Default() SmartsCustomFilter       { return "smarts_custom_filter" }
func (c Smiles) Default() Smiles                               { return "smiles" }
func (c SmilesRegexFilter) Default() SmilesRegexFilter         { return "smiles_regex_filter" }
func (c StructureTemplate) Default() StructureTemplate         { return "structure_template" }
func (c Succeeded) Default() Succeeded                         { return "succeeded" }
func (c URL) Default() URL                                     { return "url" }
func (c User) Default() User                                   { return "user" }
func (c Workspace) Default() Workspace                         { return "workspace" }

func (c AdmeV1) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c All) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c APIKey) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Base64) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c BoltzCurated) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Boltz2_1) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c BoltzProteinDesign) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c BoltzProteinScreen) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c BoltzSmDesign) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c BoltzSmScreen) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Ccd) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c ChemicalXCif) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Contact) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c DesignedProtein) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Dna) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Failed) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Insertion) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Ligand) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c LigandAtom) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c LigandCcd) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c LigandContact) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c LigandProteinBinding) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c LigandProteinBindingMetrics) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c LigandSmiles) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c LipinskiFilter) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c NoTemplate) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Pocket) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Polymer) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c PolymerAtom) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c PolymerContact) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Protein) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ProteinProteinBinding) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ProteinProteinBindingMetrics) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c RdkitDescriptorFilter) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Replacement) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Rna) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c SmartsCatalogFilter) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c SmartsCustomFilter) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Smiles) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c SmilesRegexFilter) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c StructureTemplate) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Succeeded) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c URL) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Workspace) MarshalJSON() ([]byte, error)                    { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
