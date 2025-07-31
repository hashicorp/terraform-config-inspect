// Copyright (c) HashiCorp, Inc.
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
