// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// Configuration is the top-level type representing a parsed Terraform
// configuration after terraform init has been run.
type Configuration struct {
	// Path is the local filesystem directory where the configuration was loaded from.
	Path string `json:"path"`
	// Providers contains the resolved provider versions from .terraform.lock.hcl
	Providers map[string]*LiveProviderInstance `json:"providers"`
	// Modules contains the resolved module versions from .terraform/modules/modules.json
	Modules map[string]*LiveModuleInstance `json:"modules"`

	// Diagnostics records any errors and warnings that were detected during
	// loading, primarily for inclusion in serialized forms of the configuration.
	Diagnostics Diagnostics `json:"diagnostics,omitempty"`
}

// LiveProviderInstance represents a provider with its resolved version
// as recorded in the terraform lock file.
type LiveProviderInstance struct {
	Source  string `json:"source"`
	Version string `json:"version"`
}

// LiveModuleInstance represents a module with its resolved version
// as recorded in the terraform modules manifest.
type LiveModuleInstance struct {
	Source  string `json:"source"`
	Version string `json:"version"`
}

// LoadPostInit reads the terraform lock file and modules manifest from a
// directory where terraform init has already been run.
// workingDir is the directory containing .terraform.lock.hcl
// dataDir is the .terraform directory (or TF_DATA_DIR) containing the modules manifest
// See: https://developer.hashicorp.com/terraform/language/files/dependency-lock#lock-file-location
// See: https://developer.hashicorp.com/terraform/cli/config/environment-variables#tf_data_dir
func LoadPostInit(workingDir string, dataDir string) *Configuration {
	osfs := NewOsFs()
	return LoadPostInitFromFilesystem(osfs, workingDir, osfs, dataDir)
}

// LoadPostInitFromFilesystem reads the terraform lock file and modules manifest
// from the given filesystems and directories.
func LoadPostInitFromFilesystem(workingFs FS, workingDir string, dataFs FS, dataDir string) *Configuration {
	cfg := &Configuration{
		Path:        workingDir,
		Providers:   make(map[string]*LiveProviderInstance),
		Modules:     make(map[string]*LiveModuleInstance),
		Diagnostics: make(Diagnostics, 0),
	}

	loadTFLockFile(workingFs, workingDir, cfg)
	loadTFDataDir(dataFs, dataDir, cfg)

	return cfg
}

type modulesRootSchema struct {
	Modules *[]moduleRootSchema `json:"Modules"`
}

type moduleRootSchema struct {
	Key     string `json:"Key"`
	Source  string `json:"Source"`
	Version string `json:"Version"`
}

func loadTFDataDir(fs FS, dataDir string, cfg *Configuration) {
	filename := filepath.Join(dataDir, "modules", "modules.json")

	content, err := fs.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagWarning,
			Summary:  "Module manifest file does not exist",
			Detail:   fmt.Sprintf("Module manifest file %s does not exist", filename),
		})
		return
	}
	if err != nil {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagError,
			Summary:  "Failed to read module manifest",
			Detail:   fmt.Sprintf("Failed to read module manifest %s: %s", filename, err),
		})
		return
	}

	var cachedModules = modulesRootSchema{}
	if err = json.Unmarshal(content, &cachedModules); err != nil {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagError,
			Summary:  "Failed to parse module manifest",
			Detail:   fmt.Sprintf("Failed to parse module manifest %s: %s", filename, err),
		})
		return
	}

	if cachedModules.Modules == nil {
		return
	}

	for _, mod := range *cachedModules.Modules {
		if mod.Key != "" {
			cfg.Modules[mod.Key] = &LiveModuleInstance{
				Source:  mod.Source,
				Version: mod.Version,
			}
		}
	}
}

func loadTFLockFile(fs FS, workingDir string, cfg *Configuration) {
	filename := filepath.Join(workingDir, ".terraform.lock.hcl")

	content, err := fs.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagWarning,
			Summary:  "Lock file does not exist",
			Detail:   fmt.Sprintf("Lock file %s does not exist", filename),
		})
		return
	}
	if err != nil {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagError,
			Summary:  "Failed to read lock file",
			Detail:   fmt.Sprintf("Failed to read lock file %s: %s", filename, err),
		})
		return
	}

	inspectTFLockFile(content, filename, cfg)
}

var lockFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
	},
}

var lockFileProviderBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "version",
		},
	},
}

func inspectTFLockFile(b []byte, filename string, cfg *Configuration) {
	parser := hclparse.NewParser()

	file, fileDiags := parser.ParseHCL(b, filename)
	if fileDiags.HasErrors() {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagError,
			Summary:  "Failed to parse lock file",
			Detail:   fmt.Sprintf("Failed to parse lock file %s: %s", filename, fileDiags.Error()),
		})
		return
	}

	content, _, contentDiags := file.Body.PartialContent(lockFileSchema)
	if contentDiags.HasErrors() {
		cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
			Severity: DiagError,
			Summary:  "Failed to parse lock file",
			Detail:   fmt.Sprintf("Failed to parse lock file %s: %s", filename, contentDiags.Error()),
		})
		return
	}

	for _, block := range content.Blocks {
		if block.Type == "provider" && len(block.Labels) > 0 {
			blockContent, _, contentDiags := block.Body.PartialContent(lockFileProviderBlockSchema)
			if contentDiags.HasErrors() {
				continue
			}

			if attr, defined := blockContent.Attributes["version"]; defined {
				var version string
				valDiags := gohcl.DecodeExpression(attr.Expr, nil, &version)
				if valDiags.HasErrors() {
					continue
				}

				cfg.Providers[block.Labels[0]] = &LiveProviderInstance{
					Source:  block.Labels[0],
					Version: version,
				}
			} else {
				cfg.Diagnostics = append(cfg.Diagnostics, Diagnostic{
					Severity: DiagWarning,
					Summary:  "Provider missing version",
					Detail:   fmt.Sprintf("Provider %s has no version in lock file", block.Labels[0]),
				})
			}
		}
	}
}
