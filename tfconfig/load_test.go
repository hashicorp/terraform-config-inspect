package tfconfig

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
)

func TestLoadModule(t *testing.T) {
	fixturesDir := "test-fixtures"
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

			wantDiags := map[string]struct{}{}
			wantDiagsFile, err := os.Open(filepath.Join(path, name+".diags.txt"))
			if err == nil {
				sc := bufio.NewScanner(wantDiagsFile)
				for sc.Scan() {
					wantDiags[sc.Text()] = struct{}{}
				}
				wantDiagsFile.Close()
			}

			gotObj, gotDiagsObjs := LoadModule(path)
			if gotObj == nil {
				t.Fatalf("result object is nil; want a real object")
			}

			gotDiags := map[string]struct{}{}
			for _, diag := range gotDiagsObjs {
				gotDiags[diag.Summary] = struct{}{}
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

			if diff := deep.Equal(gotDiags, wantDiags); diff != nil {
				for _, problem := range diff {
					t.Errorf("mismatching diagnostic: %s", problem)
				}
			}

			if diff := deep.Equal(got, want); diff != nil {
				for _, problem := range diff {
					t.Errorf("%s", problem)
				}
			}
		})
	}
}
