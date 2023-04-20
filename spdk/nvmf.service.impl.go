// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
)

type NvmfServiceImpl struct {
	ctx context.Context
}

func NewNvmfService(ctx context.Context) NvmfService {
	return &NvmfServiceImpl{ctx}
}

func (p *NvmfServiceImpl) CreateSubsystem(params *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) DeleteSubsystem(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) GetSubsystems(page int, limit int) (*NvmfGetSubsystemsResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) GetStats(page int, limit int) (*NvmfGetSubsystemStatsResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) AddListener(*NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) RemoveListener(*NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) AddNamespace(*NvmfSubsystemAddNsParams) (*NvmfSubsystemAddNsResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) RemoveNamespace(*NvmfSubsystemRemoveNsParams) (*NvmfSubsystemRemoveNsResult, error) {
	// TBD
	return nil, nil
}
