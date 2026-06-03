// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build !manageddisk

package probe

// EphemeralDiskFilter is a filter for Azure local NVMe ephemeral disks.
var EphemeralDiskFilter = &Filter{
	Filters: []FilterPredicate{
		&PathFilter{Path: "/dev/nvme"},
		NewModelFilter("Microsoft NVMe Direct Disk", "Microsoft NVMe Direct Disk v2"),
		&TypeFilter{Type: "disk"},
	},
}
