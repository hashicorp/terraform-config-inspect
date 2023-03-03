variable "variable_with_no_validation" {
    description = "This variable has no validation"
    type        = string
    default     = ""
}

variable "variable_with_one_validation" {
    description = "This variable has one validation"
    type        = string
    default     = ""

    validation {
        condition = (length(var.variable_with_one_validation) == 0 || length(var.variable_with_one_validation) == 10)
        error_message = "var.variable_with_one_validation must be empty or 10 characters long."
    }
}

variable "variable_with_two_validations" {
    description = "This variable has two validations"
    type        = string

    validation {
        condition = (length(var.variable_with_two_validations) == 10)
        error_message = "var.variable_with_two_validations must be 10 characters long."
    }

    validation {
        condition = (startswith(var.variable_with_two_validations, "magic"))
        error_message = "var.variable_with_two_validations must start with 'magic'."
    }
}

