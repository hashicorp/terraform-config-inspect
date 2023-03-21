// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// ModuleCall represents a "module" block within a module. That is, a
// declaration of a child module from inside its parent.
type ModuleCall struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Version string `json:"version,omitempty"`

	Pos SourcePos `json:"pos"`

	Providers []PassedProviderConfig `json:"providers,omitempty"`
}

// PassedProviderConfig represents a provider config explicitly passed down to
// a child module, possibly giving it a new local address in the process.
type PassedProviderConfig struct {
	InChild  *ProviderRef `json:"in_child"`
	InParent *ProviderRef `json:"in_parent"`
}
