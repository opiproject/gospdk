// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
)

type AccelServiceImpl struct {
	ctx    context.Context
	client JSONRPC
}

func NewAccelService(ctx context.Context) AccelService {
	// client := spdk.NewSpdkJSONRPC(spdkAddress)
	return &AccelServiceImpl{ctx, nil}
}

func (p *AccelServiceImpl) CryptoKeyCreate(name string, cipher string, key []byte) (*AccelCryptoKeyCreateResult, error) {
	// TBD
	keyHalf := len(key) / 2
	params := AccelCryptoKeyCreateParams{
		Cipher: cipher,
		Key:    hex.EncodeToString(key[:keyHalf]),
		Key2:   hex.EncodeToString(key[keyHalf:]),
		Name:   name,
	}
	var result AccelCryptoKeyCreateResult
	err := p.client.Call("accel_crypto_key_create", &params, &result)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	log.Printf("Received from SPDK: %v", result)
	if !result {
		msg := fmt.Sprintf("Could not create Crypto Key with name: %s", name)
		log.Print(msg)
		return nil, ErrUnexpectedSpdkCallResult
	}
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
