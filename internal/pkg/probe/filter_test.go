// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package probe

import (
	"testing"

	"local-csi-driver/internal/pkg/block"
)

func TestPathFilter(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		device   block.Device
		expected bool
	}{
		{
			name:     "match path prefix",
			path:     "/dev/sd",
			device:   block.Device{Path: "/dev/sda"},
			expected: true,
		},
		{
			name:     "no match path prefix",
			path:     "/dev/sd",
			device:   block.Device{Path: "/dev/nvme0n1"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &PathFilter{Path: tt.path}
			result := filter.Match(tt.device)
			if result != tt.expected {
				t.Errorf("Match(%v) = %v, want %v", tt.device, result, tt.expected)
			}
		})
	}
}

func TestTypeFilter(t *testing.T) {
	tests := []struct {
		name     string
		typeStr  string
		device   block.Device
		expected bool
	}{
		{
			name:     "match type",
			typeStr:  "SSD",
			device:   block.Device{Type: "ssd"},
			expected: true,
		},
		{
			name:     "no match type",
			typeStr:  "HDD",
			device:   block.Device{Type: "ssd"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &TypeFilter{Type: tt.typeStr}
			result := filter.Match(tt.device)
			if result != tt.expected {
				t.Errorf("Match(%v) = %v, want %v", tt.device, result, tt.expected)
			}
		})
	}
}

func TestModelFilter(t *testing.T) {
	tests := []struct {
		name     string
		models   []string
		device   block.Device
		expected bool
	}{
		{
			name:     "match model",
			models:   []string{"Samsung v2", "Samsung"},
			device:   block.Device{Model: " Samsung "},
			expected: true,
		},
		{
			name:     "match model",
			models:   []string{"Samsung v2", "Samsung"},
			device:   block.Device{Model: " Samsung v2 "},
			expected: true,
		},
		{
			name:     "no match model",
			models:   []string{"Intel"},
			device:   block.Device{Model: "Samsung"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := NewModelFilter(tt.models...)
			result := filter.Match(tt.device)
			if result != tt.expected {
				t.Errorf("Match(%v) = %v, want %v", tt.device, result, tt.expected)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		filters  []FilterPredicate
		device   block.Device
		expected bool
	}{
		{
			name: "all filters match",
			filters: []FilterPredicate{
				&PathFilter{Path: "/dev/sd"},
				&TypeFilter{Type: "SSD"},
				NewModelFilter("Samsung", "Samsung v2"),
			},
			device:   block.Device{Path: "/dev/sda", Type: "ssd", Model: "Samsung"},
			expected: true,
		},
		{
			name: "one filter does not match",
			filters: []FilterPredicate{
				&PathFilter{Path: "/dev/sd"},
				&TypeFilter{Type: "SSD"},
				NewModelFilter("Intel"),
			},
			device:   block.Device{Path: "/dev/sda", Type: "ssd", Model: "Samsung"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &Filter{Filters: tt.filters}
			result := filter.Match(tt.device)
			if result != tt.expected {
				t.Errorf("Match(%v) = %v, want %v", tt.device, result, tt.expected)
			}
		})
	}
}

