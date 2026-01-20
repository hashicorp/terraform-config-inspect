// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Stack is the top-level type representing a parsed and processed Terraform
// stack.
type Stack struct {
	// Path is the local filesystem directory where the stack was loaded from.
	Path string `json:"path"`

	Variables map[string]*Variable `json:"variables"`
	Outputs   map[string]*Output   `json:"outputs"`

	RequiredProviders map[string]*ProviderRequirement `json:"required_providers"`
	Components        map[string]*Component           `json:"components"`

	// Diagnostics records any errors and warnings that were detected during
	// loading, primarily for inclusion in serialized forms of the stack
	// since this slice is also returned as a second argument from LoadStack.
	Diagnostics Diagnostics `json:"diagnostics,omitempty"`
}

// NewStack creates new Stack representing Terraform stack at the given path
func NewStack(path string) *Stack {
	return &Stack{
		Path:              path,
		Variables:         make(map[string]*Variable),
		Outputs:           make(map[string]*Output),
		RequiredProviders: make(map[string]*ProviderRequirement),
		Components:        make(map[string]*Component),
	}
}
