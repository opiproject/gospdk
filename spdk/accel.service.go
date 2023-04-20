// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

type AccelService interface {
	CryptoKeyCreate(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	CryptoKeyDestroy(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	CryptoKeyList(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	GetStats(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
}
