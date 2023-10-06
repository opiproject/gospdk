// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
)

type AccelService interface {
	CryptoKeyCreate(ctx context.Context, name string, cipher string, key []byte) (*AccelCryptoKeyCreateResult, error)
	CryptoKeyDestroy(context.Context, *AccelCryptoKeyDestroyParams) (*AccelCryptoKeyDestroyResult, error)
	CryptoKeyList(context.Context, *AccelCryptoKeyGetParams) (*AccelCryptoKeyGetResult, error)
	GetStats(context.Context, *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
}
