// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
	"errors"
)

type NvmfServiceImpl struct {
	ctx            context.Context
}

func NewNvmfService(ctx context.Context) NvmfService {
	return &NvmfServiceImpl{ctx}
}

func (p *NvmfServiceImpl) CreateSubsystem(params *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error) {
    // TBD
}
