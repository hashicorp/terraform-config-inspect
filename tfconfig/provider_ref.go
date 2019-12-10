package tfconfig

import (
	"github.com/hashicorp/hcl/v2"
)

// ProviderRef is a reference to a provider configuration within a module.
// It represents the contents of a "provider" argument in a resource, or
// a value in the "providers" map for a module call.
type ProviderRef struct {
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"` // Empty if the default provider configuration is referenced
}

type ProviderRequirement struct {
	Source             string   `json:"source"`
	VersionConstraints []string `json:"version_constraints"`
}

func decodeRequiredProvidersBlock(block *hcl.Block) (map[string]*ProviderRequirement, hcl.Diagnostics) {
	attrs, diags := block.Body.JustAttributes()
	reqs := make(map[string]*ProviderRequirement)
	for name, attr := range attrs {
		expr, err := attr.Expr.Value(nil)
		if err != nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Invalid required_providers reference",
				Detail:   err.Error(),
				Subject:  attr.Expr.Range().Ptr(),
			})
			continue
		}
		if expr.Type().IsPrimitiveType() {
			reqs[name] = &ProviderRequirement{
				VersionConstraints: []string{expr.AsString()},
			}
		} else if expr.Type().IsObjectType() {
			pr := &ProviderRequirement{}
			if expr.Type().HasAttribute("version") {
				version := expr.GetAttr("version").AsString()
				pr.VersionConstraints = append(pr.VersionConstraints, version)
			}
			if expr.Type().HasAttribute("source") {
				pr.Source = expr.GetAttr("source").AsString()
			}
			reqs[name] = pr
		}
	}

	return reqs, diags
}
