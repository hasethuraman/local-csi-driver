// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build manageddisk

package probe

import (
	"testing"

	"local-csi-driver/internal/pkg/block"
)

func TestEphemeralDiskFilter(t *testing.T) {
	tests := []struct {
		name     string
		device   block.Device
		expected bool
	}{
		{
			name: "match managed disk",
			device: block.Device{
				Path:  "/dev/sdc",
				Type:  "disk",
				Model: "Msft Virtual Disk",
			},
			expected: true,
		},
		{
			name: "match managed disk with Virtual Disk model",
			device: block.Device{
				Path:  "/dev/sdd",
				Type:  "disk",
				Model: "Virtual Disk",
			},
			expected: true,
		},
		{
			name: "reject NVMe disk under managed disk build",
			device: block.Device{
				Path:  "/dev/nvme0n1",
				Type:  "disk",
				Model: "Microsoft NVMe Direct Disk",
			},
			expected: false,
		},
		{
			name: "reject OS disk",
			device: block.Device{
				Path:  "/dev/sda",
				Type:  "disk",
				Model: "Msft Virtual Disk",
			},
			expected: true, // path matches /dev/sd prefix; OS disk exclusion is handled elsewhere
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EphemeralDiskFilter.Match(tt.device)
			if result != tt.expected {
				t.Errorf("Match(%v) = %v, want %v", tt.device, result, tt.expected)
			}
		})
	}
}
