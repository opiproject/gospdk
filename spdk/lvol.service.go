// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

type LvolService interface {
	CreateLvstore(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	DeleteLvstore(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	GetLvstores(page int, limit int) (*NvmfGetSubsystemsResult, error)
	RenameLvstore(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	GrowLvstore(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)

 	CreateLvol(*NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error)
	SnapshotLvol(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	CloneLvol(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	RenameLvol(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	ResizeLvol(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
	DeleteLvol(*NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error)
}
