package tfconfig

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
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

			if diff := deep.Equal(buf.String(), string(expected)); diff != nil {
				for _, problem := range diff {
					t.Errorf("%s", problem)
				}
			}
		})
	}
}
