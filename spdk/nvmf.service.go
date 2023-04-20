// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

type NvmfService interface {
	CreateSubsystem(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	DeleteSubsystem(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	GetSubsystems(page int, limit int) (*NvmfGetSubsystemsResult, error)
	GetStats(page int, limit int) (*NvmfGetSubsystemStatsResult, error)
 	AddListener(*NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error)
 	RemoveListener(*NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error)
 	AddNamespace(*NvmfSubsystemAddNsParams) (*NvmfSubsystemAddNsResult, error)
 	RemoveNamespace(*NvmfSubsystemRemoveNsParams) (*NvmfSubsystemRemoveNsResult, error)
}
