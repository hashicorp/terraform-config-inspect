package tfconfig

import (
	"fmt"
)

func loadModuleLegacyHCL(dir string) (*Module, Diagnostics) {
	mod := newModule(dir)
	diags := diagnosticsError(fmt.Errorf("legacy HCL parser not yet implemented"))

	return mod, diags
}
