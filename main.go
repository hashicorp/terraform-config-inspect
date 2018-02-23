package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	flag "github.com/spf13/pflag"
)

var showJSON = flag.Bool("json", false, "produce JSON-formatted output")

func main() {
	flag.Parse()

	var dir string
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	} else {
		dir = "."
	}

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
	tmpl := template.New("md")
	tmpl.Funcs(template.FuncMap{
		"tt": func(s string) string {
			return "`" + s + "`"
		},
		"commas": func(s []string) string {
			return strings.Join(s, ", ")
		},
		"json": func(v interface{}) (string, error) {
			j, err := json.Marshal(v)
			return string(j), err
		},
		"severity": func(s tfconfig.DiagSeverity) string {
			switch s {
			case tfconfig.DiagError:
				return "Error: "
			case tfconfig.DiagWarning:
				return "Warning: "
			default:
				return ""
			}
		},
	})
	template.Must(tmpl.Parse(markdownTemplate))
	err := tmpl.Execute(os.Stdout, module)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template: %s\n", err)
		os.Exit(2)
	}
}

const markdownTemplate = `
# Module {{ tt .Path }}

{{- if .RequiredCore}}

Core Version Constraints:
{{- range .RequiredCore }}
* {{ tt . }}
{{- end}}{{end}}

{{- if .RequiredProviders}}

Provider Requirements:
{{- range $name, $versions := .RequiredProviders }}
* **{{ $name }}:** {{ if $versions }}{{ commas $versions | tt }}{{ else }}(any version){{ end }}
{{- end}}{{end}}

{{- if .Variables}}

## Input Variables
{{- range .Variables }}
* {{ tt .Name }}{{ if .Default }} (default {{ json .Default | tt }}){{else}} (required){{end}}
{{- if .Description}}: {{ .Description }}{{ end }}
{{- end}}{{end}}

{{- if .Outputs}}

## Output Values
{{- range .Outputs }}
* {{ tt .Name }}{{ if .Description}}: {{ .Description }}{{ end }}
{{- end}}{{end}}

{{- if .ManagedResources}}

## Managed Resources
{{- range .ManagedResources }}
* {{ printf "%s.%s" .Type .Name | tt }} from {{ tt .Provider.Name }}
{{- end}}{{end}}

{{- if .DataResources}}

## Data Resources
{{- range .DataResources }}
* {{ printf "data.%s.%s" .Type .Name | tt }} from {{ tt .Provider.Name }}
{{- end}}{{end}}

{{- if .ModuleCalls}}

## Child Modules
{{- range .ModuleCalls }}
* {{ tt .Name }} from {{ tt .Source }}{{ if .Version }} ({{ tt .Version }}){{ end }}
{{- end}}{{end}}

{{- if .Diagnostics}}

## Problems
{{- range .Diagnostics }}

## {{ severity .Severity }}{{ .Summary }}{{ if .Pos }}

(at {{ tt .Pos.Filename }} line {{ .Pos.Line }}{{ end }})
{{ if .Detail }}
{{ .Detail }}
{{- end }}

{{- end}}{{end}}

`
