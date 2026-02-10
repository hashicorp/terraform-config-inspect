// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"github.com/hashicorp/hcl/v2"
)

var rootSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "terraform",
			LabelNames: nil,
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
		{
			Type:       "component",
			LabelNames: []string{"name"},
		},
		{
			Type:       "required_providers",
			LabelNames: nil,
		},
	},
}

var terraformBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "required_version",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "required_providers",
		},
	},
}

var providerConfigSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "version",
		},
		{
			Name: "alias",
		},
	},
}

var variableSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "type",
		},
		{
			Name: "description",
		},
		{
			Name: "default",
		},
		{
			Name: "sensitive",
		},
		{
			Name: "deprecated",
		},
	},
}

var outputSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "sensitive",
		},
		{
			Name: "type",
		},
		{
			Name: "deprecated",
		},
	},
}

var moduleCallSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "source",
		},
		{
			Name: "version",
		},
		{
			Name: "providers",
		},
	},
}

var resourceSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "provider",
		},
	},
}

var componentSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "source",
		},
	},
}

// stackSchema defines the schema for Terraform Stacks files based on Terraform core's rootConfigSchema
// https://github.com/hashicorp/terraform/blob/8b65426ecfac58a6937c1c26297c8e6a0db57a35/internal/stacks/stackconfig/file.go
var stackSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "language",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "stack",
			LabelNames: []string{"name"},
		},
		{
			Type:       "component",
			LabelNames: []string{"name"},
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type: "locals",
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "provider",
			LabelNames: []string{"type", "name"},
		},
		{
			Type: "required_providers",
		},
		{
			Type: "removed",
		},
	},
}
