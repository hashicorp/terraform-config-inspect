
# Module `testdata/type-errors`

Provider Requirements:
* **:** (any version)
* **foo:** (any version)

## Input Variables
* `foo` (required)

## Output Values
* `foo`

## Managed Resources
* `foo.foo` from ``

## Child Modules
* `foo` from ``

## Problems

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 3)

Unsuitable value: string required

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 7)

Unsuitable value: string required

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 8)

Unsuitable value: a bool is required

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 12)

Unsuitable value: string required

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 13)

Unsuitable value: string required

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 17)

Unsuitable value: string required

## Error: Invalid provider reference

(at `testdata/type-errors/type-errors.tf` line 21)

Provider argument requires a provider name followed by an optional alias, like "aws.foo".

## Error: Unsuitable value type

(at `testdata/type-errors/type-errors.tf` line 25)

Unsuitable value: string required

## Error: Invalid required_providers object

(at `testdata/type-errors/type-errors.tf` line 27)

Required providers entries must be strings or objects.

