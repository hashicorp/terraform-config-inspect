// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	fixturesDir := "testdata"
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

			fullPath := filepath.Join(path, name+".out.md")
			expected, err := ioutil.ReadFile(fullPath)
			if err != nil {
				t.Skipf("%q not found, skipping test", fullPath)
			}

			module, _ := LoadModule(path)
			if module == nil {
				t.Fatalf("result object is nil; want a real object")
			}

			var b bytes.Buffer
			buf := &b
			err = RenderMarkdown(buf, module)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(buf.String(), string(expected)); len(diff) > 0 {
				t.Errorf("got:\n%s\nwant:\n%s\ndiff:\n%s", buf.String(), expected, diff)
			}
		})
	}
}
