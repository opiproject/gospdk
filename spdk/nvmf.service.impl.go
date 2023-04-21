// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
	"fmt"
	"log"
)

type NvmfServiceImpl struct {
	ctx    context.Context
	client JSONRPC
}

func NewNvmfService(ctx context.Context) NvmfService {
	// client := spdk.NewSpdkJSONRPC(spdkAddress)
	return &NvmfServiceImpl{ctx, nil}
}

func (p *NvmfServiceImpl) CreateSubsystem(nqn string, serial string, model string, ns int) (*NvmfCreateSubsystemResult, error) {
	// TBD
	params := NvmfCreateSubsystemParams{
		Nqn:           nqn,
		SerialNumber:  serial,
		ModelNumber:   model,
		AllowAnyHost:  true,
		MaxNamespaces: ns,
	}
	var result NvmfCreateSubsystemResult
	err := p.client.Call("nvmf_create_subsystem", &params, &result)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	log.Printf("Received from SPDK: %v", result)
	if !result {
		msg := fmt.Sprintf("Could not create NQN: %s", nqn)
		log.Print(msg)
		return nil, ErrUnexpectedSpdkCallResult
	}
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
