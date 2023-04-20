// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

type AccelService interface {
	CryptoKeyCreate(*AccelCryptoKeyCreateParams) (*AccelCryptoKeyCreateResult, error)
	CryptoKeyDestroy(*AccelCryptoKeyDestroyParams) (*AccelCryptoKeyDestroyResult, error)
	CryptoKeyList(*AccelCryptoKeyGetParams) (*AccelCryptoKeyGetResult, error)
	GetStats(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
}
