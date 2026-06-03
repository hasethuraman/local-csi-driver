// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build !manageddisk

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
			// Standard_NC40ads_H100_v5 nodes have a different model name
			// for the ephemeral disk. This test case is to ensure that
			// the filter matches the model name for these nodes.
			name: "match disk for v2 direct disk nodes",
			device: block.Device{
				Path:  "/dev/nvme0n1",
				Type:  "disk",
				Model: "Microsoft NVMe Direct Disk v2           ",
			},
			expected: true,
		},
		{
			name: "match disk for direct disk nodes",
			device: block.Device{
				Path:  "/dev/nvme0n1",
				Type:  "disk",
				Model: "Microsoft NVMe Direct Disk           ",
			},
			expected: true,
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
