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
		"tt": func(i interface{}) string {
			s := fmt.Sprintf("%v", i)
			return "`" + s + "`"
		},
		"commas": func(s []string) string {
			return strings.Join(s, ", ")
		},
		"json": func(v interface{}) (string, error) {
			j, err := json.Marshal(v)
			return string(j), err
		},
		"skip": func(p tfconfig.SourcePos) bool {
			blacklist := []string{"environment.tf.json", "global-variables.tf.json", "account-variables.tf.json", "variables.tf"}

			for _, b := range blacklist {
				if p.Filename == b {
					return false
				}
			}
			return true
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
## Inputs
| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
{{- range .Variables }}{{if skip .Pos }}
| {{ tt .Name }} | {{- if .Description}}{{ .Description }}{{ end }} | | {{ tt .Default }} | {{if tt .Default}}no{{else}}yes{{end}} |{{end}}{{end}}

{{- if .Outputs}}

## Outputs
| Name | Description |
|------|-------------|
{{- range .Outputs }}
| {{ tt .Name }} | {{ if .Description}}{{ .Description }}{{ end }} |
{{- end}}{{end}}

{{- if .ManagedResources}}

Managed Resources
-----------------
{{- range .ManagedResources }}
* {{ printf "%s.%s" .Type .Name | tt }}
{{- end}}{{end}}

{{- if .DataResources}}

Data Resources
--------------
{{- range .DataResources }}
* {{ printf "data.%s.%s" .Type .Name | tt }}
{{- end}}{{end}}

{{- if .ModuleCalls}}

Child Modules
-------------
{{- range .ModuleCalls }}
* {{ tt .Name }} from {{ tt .Source }}{{ if .Version }} ({{ tt .Version }}){{ end }}
{{- end}}{{end}}

{{- if .Diagnostics}}

Problems
-------------
{{- range .Diagnostics }}

{{ severity .Severity }}{{ .Summary }}{{ if .Pos }}
-------------

(at {{ tt .Pos.Filename }} line {{ .Pos.Line }}{{ end }})
{{ if .Detail }}
{{ .Detail }}
{{- end }}

{{- end}}{{end}} 
`
