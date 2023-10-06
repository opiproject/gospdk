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
	client JSONRPC
}

func NewNvmfService() NvmfService {
	// client := spdk.NewSpdkJSONRPC(spdkAddress)
	return &NvmfServiceImpl{nil}
}

func (p *NvmfServiceImpl) CreateSubsystem(ctx context.Context, nqn string, serial string, model string, ns int) (*NvmfCreateSubsystemResult, error) {
	// TBD
	params := NvmfCreateSubsystemParams{
		Nqn:           nqn,
		SerialNumber:  serial,
		ModelNumber:   model,
		AllowAnyHost:  true,
		MaxNamespaces: ns,
	}
	var result NvmfCreateSubsystemResult
	err := p.client.Call(ctx, "nvmf_create_subsystem", &params, &result)
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

func (p *NvmfServiceImpl) DeleteSubsystem(context.Context, *NvmfDeleteSubsystemParams) (*NvmfDeleteSubsystemResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) GetSubsystems(_ context.Context, page int, limit int) (*NvmfGetSubsystemsResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) GetStats(_ context.Context, page int, limit int) (*NvmfGetSubsystemStatsResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) AddListener(context.Context, *NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) RemoveListener(context.Context, *NvmfSubsystemAddListenerParams) (*NvmfSubsystemAddListenerResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) AddNamespace(context.Context, *NvmfSubsystemAddNsParams) (*NvmfSubsystemAddNsResult, error) {
	// TBD
	return nil, nil
}

func (p *NvmfServiceImpl) RemoveNamespace(context.Context, *NvmfSubsystemRemoveNsParams) (*NvmfSubsystemRemoveNsResult, error) {
	// TBD
	return nil, nil
}
