// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
)

type NvmfService interface {
	CreateSubsystem(ctx context.Context, nqn string, serial string, model string, ns int) (*NvmfCreateSubsystemResult, error)
	DeleteSubsystem(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	GetSubsystems(ctx context.Context, page int, limit int) (*NvmfGetSubsystemsResult, error)
	GetStats(ctx context.Context, page int, limit int) (*NvmfGetSubsystemStatsResult, error)
	AddListener(context.Context, *NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error)
	RemoveListener(context.Context, *NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error)
	AddNamespace(context.Context, *NvmfSubsystemAddNsParams) (*NvmfSubsystemAddNsResult, error)
	RemoveNamespace(context.Context, *NvmfSubsystemRemoveNsParams) (*NvmfSubsystemRemoveNsResult, error)
}
