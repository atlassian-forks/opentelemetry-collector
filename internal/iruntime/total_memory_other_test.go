// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package iruntime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalMemory(t *testing.T) {
	totalMemory, err := TotalMemory()
	assert.NoError(t, err)
	assert.Positive(t, totalMemory)
}
