package tfconfig

import "github.com/hashicorp/hcl/v2"

// Variable represents a single variable from a Terraform module.
type Variable struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`

	// Default is an approximate representation of the default value in
	// the native Go type system. The conversion from the value given in
	// configuration may be slightly lossy. Only values that can be
	// serialized by json.Marshal will be included here.
	Default   interface{} `json:"default"`
	Required  bool        `json:"required"`
	Sensitive bool        `json:"sensitive,omitempty"`

	Validations []VariableValidation `json:"validations,omitempty"`

	Pos SourcePos `json:"pos"`
}

// VariableValidation represents a configuration-defined validation rule
// for a particular input variable, given as a "validation" block inside
// a "variable" block.
type VariableValidation struct {
	// Condition is an expression that refers to the variable being tested
	// and contains no other references. The expression must return true
	// to indicate that the value is valid or false to indicate that it is
	// invalid. If the expression produces an error, that's considered a bug
	// in the module defining the validation rule, not an error in the caller.
	Condition hcl.Expression `json:"-"`

	// ErrorMessage is one or more full sentences, which would need to be in
	// English for consistency with the rest of the error message output but
	// can in practice be in any language as long as it ends with a period.
	// The message should describe what is required for the condition to return
	// true in a way that would make sense to a caller of the module.
	ErrorMessage string `json:"error_message"`
}
