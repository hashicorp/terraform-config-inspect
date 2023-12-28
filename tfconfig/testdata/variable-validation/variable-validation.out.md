
# Module `testdata/variable-validation`

## Input Variables
* `variable_with_no_validation` (default `""`): This variable has no validation
* `variable_with_one_validation` (default `""`): This variable has one validation  
  Validation error messages
  * `var.variable_with_one_validation must be empty or 10 characters long.`
* `variable_with_two_validations` (required): This variable has two validations  
  Validation error messages
  * `var.variable_with_two_validations must be 10 characters long.`
  * `var.variable_with_two_validations must start with 'magic'.`

