package tfconfig

// CheckRule details a single assert block within a Check block.
type CheckRule struct {
	Condition    string `json:"condition"`
	ErrorMessage string `json:"error_message"`

	Pos SourcePos `json:"pos"`
}

// Check is the representation of a check block within a module.
type Check struct {
	Name string `json:"name"`

	DataResource *Resource `json:"data_resource,omitempty"`

	Pos SourcePos `json:"pos"`
}
