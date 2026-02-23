// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"os"
	"path/filepath"
	"testing"
)

func assertEqual[T comparable](t *testing.T, got, want T, msg string) {
	t.Helper()
	if got != want {
		t.Errorf("%s: got %v, want %v", msg, got, want)
	}
}

func TestLoadPostInit(t *testing.T) {
	path := filepath.Join("testdata-post-init", "basic")
	dataDir := filepath.Join(path, ".terraform")

	cfg := LoadPostInit(path, dataDir)

	if cfg == nil {
		t.Fatal("result is nil")
	}

	assertEqual(t, cfg.Path, path, "path")
	assertEqual(t, len(cfg.Providers), 2, "providers count")
	assertEqual(t, len(cfg.Modules), 4, "modules count")

	// Verify providers
	expectedProviders := map[string]LiveProviderInstance{
		"registry.terraform.io/hashicorp/aws":    {Source: "registry.terraform.io/hashicorp/aws", Version: "5.31.0"},
		"registry.terraform.io/hashicorp/random": {Source: "registry.terraform.io/hashicorp/random", Version: "3.6.0"},
	}

	for name, want := range expectedProviders {
		got, ok := cfg.Providers[name]
		if !ok {
			t.Errorf("provider %q: not found", name)
			continue
		}
		assertEqual(t, got.Source, want.Source, "provider "+name+" source")
		assertEqual(t, got.Version, want.Version, "provider "+name+" version")
	}

	// Verify modules
	expectedModules := map[string]LiveModuleInstance{
		"vpc":          {Source: "registry.terraform.io/terraform-aws-modules/vpc/aws", Version: "5.1.0"},
		"ec2":          {Source: "registry.terraform.io/terraform-aws-modules/ec2-instance/aws", Version: "5.5.0"},
		"local_module": {Source: "./modules/local", Version: ""},
		"git_module":   {Source: "git::https://example.com/module.git", Version: ""},
	}

	for name, want := range expectedModules {
		got, ok := cfg.Modules[name]
		if !ok {
			t.Errorf("module %q: not found", name)
			continue
		}
		assertEqual(t, got.Source, want.Source, "module "+name+" source")
		assertEqual(t, got.Version, want.Version, "module "+name+" version")
	}

	// Verify root module (empty key) is skipped
	if _, ok := cfg.Modules[""]; ok {
		t.Error("root module (empty key) should not be included")
	}
}

func TestLoadPostInit_Empty(t *testing.T) {
	path := filepath.Join("testdata-post-init", "empty")
	dataDir := filepath.Join(path, ".terraform")

	cfg := LoadPostInit(path, dataDir)

	if cfg == nil {
		t.Fatal("result is nil")
	}

	assertEqual(t, cfg.Path, path, "path")
	assertEqual(t, len(cfg.Providers), 0, "providers count")
	assertEqual(t, len(cfg.Modules), 0, "modules count")
	assertEqual(t, cfg.Diagnostics.HasErrors(), false, "diagnostic errors")
	assertEqual(t, len(cfg.Diagnostics), 2, "diagnostics count")
}

func TestLoadPostInitFromFilesystem(t *testing.T) {
	path := filepath.Join("testdata-post-init", "basic")
	dataDir := filepath.Join(path, ".terraform")

	fs := os.DirFS(".")
	cfg := LoadPostInitFromFilesystem(WrapFS(fs), path, WrapFS(fs), dataDir)

	if cfg == nil {
		t.Fatal("result is nil")
	}

	assertEqual(t, len(cfg.Providers), 2, "providers count")
	assertEqual(t, len(cfg.Modules), 4, "modules count")
}
