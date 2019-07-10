package main

import (
	"bytes"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestDefaultTemplate(t *testing.T) {
	module, _ := tfconfig.LoadModule("test-fixtures")
	buf := new(bytes.Buffer)
	err := showModuleMarkdown(module, defaultMarkdownTemplate, buf)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	expected, err := ioutil.ReadFile(filepath.Join("test-fixtures", "default.md"))
	actual := buf.String()
	if actual != string(expected) {
		t.Errorf("default template produced unexpected output:\n%s", actual)
	}
}

func TestValidCustomTemplate(t *testing.T) {
	module, _ := tfconfig.LoadModule("test-fixtures")
	buf := new(bytes.Buffer)
	customTemplate, err := ioutil.ReadFile(filepath.Join("test-fixtures", "valid-custom.tmpl"))
	if err != nil {
		t.Fatalf("failed to read template file: %s", err)
	}

	err = showModuleMarkdown(module, string(customTemplate), buf)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	expected, err := ioutil.ReadFile(filepath.Join("test-fixtures", "valid-custom.md"))
	actual := buf.String()
	if actual != string(expected) {
		t.Errorf("custom template produced unexpected output:\n%s", actual)
	}
}

func TestInvalidTemplate(t *testing.T) {
	module, _ := tfconfig.LoadModule("test-fixtures")
	buf := new(bytes.Buffer)
	customTemplate, err := ioutil.ReadFile(filepath.Join("test-fixtures", "invalid-custom.tmpl"))
	if err != nil {
		t.Fatalf("failed to read template file: %s", err)
	}

	err = showModuleMarkdown(module, string(customTemplate), buf)
	if err == nil {
		t.Fatalf("err was nil; expected an error for an invalid template")
	}
	if err.Error() != "template: md:9: unexpected EOF" {
		t.Errorf("invalid template produced an unexpected error: %s", err)
	}
}
