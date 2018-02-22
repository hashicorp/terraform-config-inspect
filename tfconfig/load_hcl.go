package tfconfig

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"
)

func loadModule(dir string) (*Module, Diagnostics) {
	mod := newModule(dir)
	primaryPaths, overridePaths, diags := dirFiles(dir)

	parser := hclparse.NewParser()

	if len(primaryPaths) == 0 && len(overridePaths) == 0 {
		diags = append(diags, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "No Terraform configuration files",
			Detail:   fmt.Sprintf("Module directory %s does not contain any .tf or .tf.json files.", dir),
		})
	}

	for _, filename := range primaryPaths {
		var file *hcl.File
		var fileDiags hcl.Diagnostics
		if strings.HasSuffix(filename, ".json") {
			file, fileDiags = parser.ParseJSONFile(filename)
		} else {
			file, fileDiags = parser.ParseHCLFile(filename)
		}
		diags = append(diags, fileDiags...)
		if file == nil {
			continue
		}

		content, _, contentDiags := file.Body.PartialContent(rootSchema)
		diags = append(diags, contentDiags...)

		for _, block := range content.Blocks {
			switch block.Type {

			case "terraform":
				content, _, contentDiags := block.Body.PartialContent(terraformBlockSchema)
				diags = append(diags, contentDiags...)

				if attr, defined := content.Attributes["required_version"]; defined {
					var version string
					valDiags := gohcl.DecodeExpression(attr.Expr, nil, &version)
					diags = append(diags, valDiags...)
					if !valDiags.HasErrors() {
						mod.RequiredCore = append(mod.RequiredCore, version)
					}
				}

				for _, block := range content.Blocks {
					// Our schema only allows required_providers here, so we
					// assume that we'll only get that block type.
					attrs, attrDiags := block.Body.JustAttributes()
					diags = append(diags, attrDiags...)

					for name, attr := range attrs {
						var version string
						valDiags := gohcl.DecodeExpression(attr.Expr, nil, &version)
						diags = append(diags, valDiags...)
						if !valDiags.HasErrors() {
							mod.RequiredProviders[name] = append(mod.RequiredProviders[name], version)
						}
					}
				}

			default:
				// Should never happen because our cases above should be
				// exhaustive for our schema.
				panic(fmt.Errorf("unhandled block type %q", block.Type))
			}
		}
	}

	return mod, diagnosticsHCL(diags)
}
