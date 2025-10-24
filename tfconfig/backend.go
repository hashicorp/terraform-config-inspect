package tfconfig

import "github.com/hashicorp/hcl/v2"

// Backend represents a backend definition from a Terraform module.
type Backend struct {
	Type string    `json:"type"`
	Pos  SourcePos `json:"pos"`
}

func decodeBackendBlock(block *hcl.Block) (*Backend, hcl.Diagnostics) {
	backend := new(Backend)
	backend.Type = block.Labels[0]
	return backend, nil
}
