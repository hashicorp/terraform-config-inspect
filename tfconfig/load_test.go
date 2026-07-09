// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// testLoadHelper is the common testing logic for loading functions
func testLoadHelper(t *testing.T, fixturesDir string, loadFunc func(string) interface{}) {
	testDirs, err := ioutil.ReadDir(fixturesDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, info := range testDirs {
		if !info.IsDir() {
			continue
		}
		t.Run(info.Name(), func(t *testing.T) {
			name := info.Name()
			path := filepath.Join(fixturesDir, name)

			wantSrc, err := ioutil.ReadFile(filepath.Join(path, name+".out.json"))
			if err != nil {
				t.Fatalf("failed to read result file: %s", err)
			}
			var want map[string]interface{}
			err = json.Unmarshal(wantSrc, &want)
			if err != nil {
				t.Fatalf("failed to parse result file: %s", err)
			}

			gotObj := loadFunc(path)
			if gotObj == nil {
				t.Fatalf("result object is nil; want a real object")
			}

			gotSrc, err := json.Marshal(gotObj)
			if err != nil {
				t.Fatalf("result is not JSON-able: %s", err)
			}
			var got map[string]interface{}
			err = json.Unmarshal(gotSrc, &got)
			if err != nil {
				t.Fatalf("failed to parse the actual result (!?): %s", err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("wrong result\n%s", diff)
			}
		})
	}
}

func TestLoadModule(t *testing.T) {
	testLoadHelper(t, "testdata", func(path string) interface{} {
		module, _ := LoadModule(path)
		return module
	})
}

func TestLoadModuleFromFilesystem(t *testing.T) {
	testLoadHelper(t, "testdata", func(path string) interface{} {
		fs := os.DirFS(".")
		module, _ := LoadModuleFromFilesystem(WrapFS(fs), path)
		return module
	})
}

func TestLoadStack(t *testing.T) {
	testLoadHelper(t, "testdata-stack", func(path string) interface{} {
		stack, _ := LoadStack(path)
		return stack
	})
}

func TestLoadModuleDynamicSource(t *testing.T) {
	// Terraform 1.15 allows a module's source and version to reference const
	// input variables and local values. These can't be evaluated with an empty
	// context, so the loader must fall back to the raw expression rather than
	// reporting an error.
	module, diags := LoadModule("testdata/module-dynamic-source")
	if diags.HasErrors() {
		t.Fatalf("unexpected errors: %v", diags)
	}

	cases := map[string]struct {
		source  string
		version string
	}{
		"child":   {source: "var.child_source", version: "1.0.0"},
		"sibling": {source: "local.sibling_source", version: "local.sibling_version"},
		"static":  {source: "app.terraform.io/example-org/static/aws", version: "3.0.0"},
	}

	for name, want := range cases {
		mc, exists := module.ModuleCalls[name]
		if !exists {
			t.Errorf("module call %q not found", name)
			continue
		}
		if mc.Source != want.source {
			t.Errorf("module %q source = %q, want %q", name, mc.Source, want.source)
		}
		if mc.Version != want.version {
			t.Errorf("module %q version = %q, want %q", name, mc.Version, want.version)
		}
	}
}

func TestProviderLabels(t *testing.T) {
	// Test that provider blocks with two labels are correctly parsed
	stack, diags := LoadStack("testdata-stack/provider-labels")

	if diags.HasErrors() {
		t.Fatalf("unexpected errors: %v", diags)
	}

	expectedProviders := map[string]bool{
		"aws":    true,
		"random": true,
		"null":   true,
	}

	if len(stack.RequiredProviders) != len(expectedProviders) {
		t.Errorf("expected %d required providers, got %d", len(expectedProviders), len(stack.RequiredProviders))
	}

	for providerName := range expectedProviders {
		if _, exists := stack.RequiredProviders[providerName]; !exists {
			t.Errorf("expected provider %q to be in required_providers", providerName)
		}
	}

	if awsProvider, exists := stack.RequiredProviders["aws"]; exists {
		if awsProvider.Source != "hashicorp/aws" {
			t.Errorf("expected aws provider source to be 'hashicorp/aws', got %q", awsProvider.Source)
		}
	}
}
