package tfconfig

// Module is the top-level type representing a parsed and processed Terraform
// module.
type Module struct {
	// Path is the local filesystem directory where the module was loaded from.
	Path string `json:"path"`

	Variables map[string]*Variable `json:"variables"`
	Outputs   map[string]*Output   `json:"outputs"`

	RequiredCore      []string            `json:"required_core,omitempty"`
	RequiredProviders map[string][]string `json:"required_providers"`

	ManagedResources map[string]*Resource   `json:"managed_resources"`
	DataResources    map[string]*Resource   `json:"data_resources"`
	ModuleCalls      map[string]*ModuleCall `json:"module_calls"`
}

func newModule(path string) *Module {
	return &Module{
		Path:              path,
		Variables:         make(map[string]*Variable),
		Outputs:           make(map[string]*Output),
		RequiredProviders: make(map[string][]string),
		ManagedResources:  make(map[string]*Resource),
		DataResources:     make(map[string]*Resource),
		ModuleCalls:       make(map[string]*ModuleCall),
	}
}
