// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	flag "github.com/spf13/pflag"
)

var showJSON = flag.Bool("json", false, "produce JSON-formatted output")
var parseStack = flag.Bool("stack", false, "parse a Terraform stack")

func main() {
	flag.Parse()

	var dir string
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	} else {
		dir = "."
	}

	if *parseStack {
		stack, diags := tfconfig.LoadStack(dir)
		stack.Diagnostics = diags

		if *showJSON {
			showStackJSON(stack)
		} else {
			showStackMarkdown(stack)
		}

		if stack.Diagnostics.HasErrors() {
			os.Exit(1)
		}
	} else {
		module, _ := tfconfig.LoadModule(dir)

		if *showJSON {
			showModuleJSON(module)
		} else {
			showModuleMarkdown(module)
		}

		if module.Diagnostics.HasErrors() {
			os.Exit(1)
		}
	}
}

func showModuleJSON(module *tfconfig.Module) {
	j, err := json.MarshalIndent(module, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error producing JSON: %s\n", err)
		os.Exit(2)
	}
	os.Stdout.Write(j)
	os.Stdout.Write([]byte{'\n'})
}

func showModuleMarkdown(module *tfconfig.Module) {
	err := tfconfig.RenderMarkdown(os.Stdout, module)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template: %s\n", err)
		os.Exit(2)
	}
}

func showStackJSON(stack *tfconfig.Stack) {
	j, err := json.MarshalIndent(stack, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error producing JSON: %s\n", err)
		os.Exit(2)
	}
	os.Stdout.Write(j)
	os.Stdout.Write([]byte{'\n'})
}

func showStackMarkdown(stack *tfconfig.Stack) {
	fmt.Printf("# Terraform Stack: %s\n\n", stack.Path)

	if len(stack.Variables) > 0 {
		fmt.Printf("## Variables\n\n")
		for name, variable := range stack.Variables {
			fmt.Printf("- **%s** (%s)", name, variable.Type)
			if variable.Description != "" {
				fmt.Printf(": %s", variable.Description)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	if len(stack.Outputs) > 0 {
		fmt.Printf("## Outputs\n\n")
		for name, output := range stack.Outputs {
			fmt.Printf("- **%s**", name)
			if output.Description != "" {
				fmt.Printf(": %s", output.Description)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	if len(stack.Components) > 0 {
		fmt.Printf("## Components\n\n")
		for name, component := range stack.Components {
			fmt.Printf("- **%s** (source: `%s`)\n", name, component.Source)
		}
		fmt.Printf("\n")
	}

	if len(stack.RequiredProviders) > 0 {
		fmt.Printf("## Required Providers\n\n")
		for name, provider := range stack.RequiredProviders {
			fmt.Printf("- **%s**", name)
			if provider.Source != "" {
				fmt.Printf(" (source: `%s`)", provider.Source)
			}
			if len(provider.VersionConstraints) > 0 {
				fmt.Printf(" version: %s", provider.VersionConstraints[0])
			}
			fmt.Printf("\n")
		}
	}

	if len(stack.Diagnostics) > 0 {
		fmt.Printf("## Problems\n\n")
		for _, diag := range stack.Diagnostics {
			severity := ""
			switch diag.Severity {
			case tfconfig.DiagError:
				severity = "Error: "
			case tfconfig.DiagWarning:
				severity = "Warning: "
			}
			fmt.Printf("## %s%s", severity, diag.Summary)
			if diag.Pos != nil {
				fmt.Printf("\n\n(at `%s` line %d)", diag.Pos.Filename, diag.Pos.Line)
			}
			if diag.Detail != "" {
				fmt.Printf("\n\n%s", diag.Detail)
			}
			fmt.Printf("\n\n")
		}
	}
}
