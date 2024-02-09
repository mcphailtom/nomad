// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build linux

package proclib

import (
	"github.com/hashicorp/nomad/client/lib/cgroupslib"
)

// New creates a Wranglers factory for creating ProcessWrangler's appropriate
// for the given system (i.e. cgroups v1 or cgroups v2).
func New(configs *Configs) (*Wranglers, error) {
	w := &Wranglers{
		configs: configs,
		m:       make(map[Task]ProcessWrangler),
	}

	var err error
	switch cgroupslib.GetMode() {
	case cgroupslib.CG1:
		w.create, err = newCG1(configs)
	default:
		w.create, err = newCG2(configs)
	}

	return w, err
}
