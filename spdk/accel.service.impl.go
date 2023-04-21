// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
)

type AccelServiceImpl struct {
	ctx context.Context
}

func NewAccelService(ctx context.Context) AccelService {
	return &AccelServiceImpl{ctx}
}

func (p *AccelServiceImpl) CryptoKeyCreate(params *AccelCryptoKeyCreateParams) (*AccelCryptoKeyCreateResult, error) {
	// TBD
	return nil, nil
}

func (p *AccelServiceImpl) CryptoKeyDestroy(params *AccelCryptoKeyDestroyParams) (*AccelCryptoKeyDestroyResult, error) {
	// TBD
	return nil, nil
}

func (p *AccelServiceImpl) CryptoKeyList(params *AccelCryptoKeyGetParams) (*AccelCryptoKeyGetResult, error) {
	// TBD
	return nil, nil
}

func (p *AccelServiceImpl) GetStats(params *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error) {
	// TBD
	return nil, nil
}
