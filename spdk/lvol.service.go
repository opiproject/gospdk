// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
)

// LvolService is interface to all logical volumes functions in spdk
type LvolService interface {
	CreateLvstore(context.Context, *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	DeleteLvstore(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	GetLvstores(ctx context.Context, page int, limit int) (*NvmfGetSubsystemsResult, error)
	RenameLvstore(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	GrowLvstore(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)

	CreateLvol(context.Context, *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	SnapshotLvol(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	CloneLvol(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	RenameLvol(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	ResizeLvol(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	DeleteLvol(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
}
