// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Output represents a single output from a Terraform module.
type Output struct {
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Sensitive   bool      `json:"sensitive,omitempty"`
	Pos         SourcePos `json:"pos"`
	Type        string    `json:"type,omitempty"`
}
