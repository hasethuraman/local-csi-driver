// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build manageddisk

package probe

// EphemeralDiskFilter is a filter for Azure managed disks used in scale testing.
// Build with: go build -tags manageddisk ./cmd/driver
var EphemeralDiskFilter = &Filter{
	Filters: []FilterPredicate{
		&PathFilter{Path: "/dev/sd"},
		NewModelFilter("Msft Virtual Disk", "Virtual Disk"),
		&TypeFilter{Type: "disk"},
	},
}
