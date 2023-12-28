[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rquadling/terraform-config-inspect/push.yml?style=for-the-badge&logo=github)](https://github.com/rquadling/terraform-config-inspect/actions/workflows/push.yml)
[![GitHub issues](https://img.shields.io/github/issues/rquadling/terraform-config-inspect.svg?style=for-the-badge&logo=github)](https://github.com/rquadling/terraform-config-inspect/issues)

# terraform-config-inspect

From the Hashicorp Terraform Config [README.md](https://github.com/hashicorp/terraform-config-inspect/blob/master/README.md).
>This repository contains a helper library for extracting high-level metadata
about Terraform modules from their source code. It processes only a subset
of the information Terraform itself would process, and in return it's able
to be broadly compatible with modules written for many different versions of
Terraform.
> 
> This library can also interpret valid Terraform configurations targeting
Terraform v0.10 through v0.15, although the level of detail returned may
be lower in older language versions.

Seems good enough, yeah? So why does this fork exist?

Primarily the Terraform language moves forward. For example, we now have input validation (yes that was NOT there in
V1.0 of Terraform), and, more recently, checks.

But, due to the limitations Hashicorp have on their helper library, there are a lot of things that able to be accessed.

And so this fork was created for that reason. By having a more uptodate helper library, tools that are currently unable
to access things like input validation or checks, can now do so.

Hopefully.

As with many open source repositories, the code presented here is as is. I'm not a lawyer. I'm also not a business or a
hacker. The primary purpose is to add support for features that are native to Terraform but not accessible in
Hashicorp's helper library. That's it really.

I hope you find it useful.

If anyone wants to make subsequent pull requests, or

```
$ go install github.com/rquadling/terraform-config-inspect@latest
```

```go
import "github.com/rquadling/terraform-config-inspect/tfconfig"

// ...

module, diags := tfconfig.LoadModule(dir)

// ...
```

## Command Line Tool

The primary way to use this repository is as a Go library, but as a convenience
it also contains a CLI tool called `terraform-config-inspect`, installed
automatically by the `go get` command above, that allows viewing module
information in either a Markdown-like format or in JSON format.

```sh
$ terraform-config-inspect path/to/module
```

```markdown
# Module `path/to/module`

Provider Requirements:

- **null:** (any version)

## Input Variables

- `a` (default `"a default"`)
- `b` (required): The b variable

## Output Values

- `a`
- `b`: I am B

## Managed Resources

- `null_resource.a` from `null`
- `null_resource.b` from `null`
```

```sh
$ terraform-config-inspect --json path/to/module
```

```json
{
  "path": "path/to/module",
  "variables": {
    "A": {
      "name": "A",
      "default": "A default",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 1
      }
    },
    "B": {
      "name": "B",
      "description": "The B variable",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 5
      }
    }
  },
  "outputs": {
    "A": {
      "name": "A",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 9
      }
    },
    "B": {
      "name": "B",
      "description": "I am B",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 13
      }
    }
  },
  "required_providers": {
    "null": []
  },
  "managed_resources": {
    "null_resource.A": {
      "mode": "managed",
      "type": "null_resource",
      "name": "A",
      "provider": {
        "name": "null"
      },
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 18
      }
    },
    "null_resource.B": {
      "mode": "managed",
      "type": "null_resource",
      "name": "B",
      "provider": {
        "name": "null"
      },
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 19
      }
    }
  },
  "data_resources": {},
  "module_calls": {}
}
```

## Contributing

From the Hashicorp Terraform Config [README.md](https://github.com/hashicorp/terraform-config-inspect/blob/master/README.md).
> This library and tool are intentionally focused on only extracting simple
top-level metadata about a single Terraform module. This is to reduce the
maintenance burden of keeping this codebase synchronized with changes to
Terraform itself: the features extracted by this package are unlikely to change
significantly in future versions.
> 
> For that reason, **we cannot accept external PRs for this codebase that add support for additional Terraform language features**.
>
> Furthermore, we consider this package feature-complete; if there is a feature
you wish to see added, please open a GitHub issue first so we can discuss the
feasibility and design before submitting a pull request. We are unlikely to
accept PRs that add features without discussion first.
>
> We would be happy to review PRs to fix bugs in existing functionality or to
improve the usability of the Go package API, however. We will be hesitant about
any breaking changes to the API, since this library is used by a number of
existing tools and systems.
>
> To work on this codebase you will need a recent version of Go installed. Please
ensure all files match the formatting rules applied by `go fmt` and that all
unit tests are passing.

So, with all of that said from Hashicorp, this repository will accept pull requests, but they really must be related to
the functionality, and output, of the helper library, and as best as possible, maintain compatibility so that any
upstream changes from Hashicorp are easy to merge.
