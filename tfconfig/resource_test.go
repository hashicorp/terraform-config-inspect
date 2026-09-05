// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package tfconfig

import (
	"encoding/json"
	"testing"
)

func TestResourceModeJSONRoundTrip(t *testing.T) {
	cases := []struct {
		name string
		mode ResourceMode
		want string
	}{
		{"managed", ManagedResourceMode, `"managed"`},
		{"data", DataResourceMode, `"data"`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.mode)
			if err != nil {
				t.Fatalf("marshal: %s", err)
			}
			if string(got) != tc.want {
				t.Fatalf("marshal = %s, want %s", got, tc.want)
			}

			var mode ResourceMode
			if err := json.Unmarshal(got, &mode); err != nil {
				t.Fatalf("unmarshal: %s", err)
			}
			if mode != tc.mode {
				t.Errorf("unmarshal = %v (%q), want %v (%q)", mode, mode, tc.mode, tc.mode)
			}
		})
	}
}

func TestModuleJSONUnmarshalResourceMode(t *testing.T) {
	// MarshalJSON encodes ResourceMode as "managed"/"data" strings. Unmarshal
	// must accept those strings so a serialized Module can be loaded back.
	src := []byte(`{
		"path": ".",
		"data_resources": {
			"data.null_data_source.x": {
				"mode": "data",
				"type": "null_data_source",
				"name": "x",
				"provider": {"name": "null"},
				"pos": {"filename": "main.tf", "line": 1}
			}
		},
		"managed_resources": {
			"null_resource.y": {
				"mode": "managed",
				"type": "null_resource",
				"name": "y",
				"provider": {"name": "null"},
				"pos": {"filename": "main.tf", "line": 5}
			}
		}
	}`)

	var mod Module
	if err := json.Unmarshal(src, &mod); err != nil {
		t.Fatalf("unmarshal Module: %s", err)
	}

	data := mod.DataResources["data.null_data_source.x"]
	if data == nil {
		t.Fatal("missing data resource")
	}
	if data.Mode != DataResourceMode {
		t.Errorf("data resource mode = %v (%q), want data", data.Mode, data.Mode)
	}

	managed := mod.ManagedResources["null_resource.y"]
	if managed == nil {
		t.Fatal("missing managed resource")
	}
	if managed.Mode != ManagedResourceMode {
		t.Errorf("managed resource mode = %v (%q), want managed", managed.Mode, managed.Mode)
	}
}
