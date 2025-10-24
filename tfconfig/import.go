// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfconfig

// Import represents a single import block from a Terraform module.
type Import struct {
	To  string    `json:"to"`
	Id  string    `json:"id"`
	Pos SourcePos `json:"pos"`
}
