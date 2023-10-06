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

// AccelServiceImpl implements AccelService interface
type AccelServiceImpl struct {
	client JSONRPC
}

// build time check that struct implements interface
var _ AccelService = (*AccelServiceImpl)(nil)

// NewAccelService is a constructor for AccelServiceImpl
func NewAccelService() *AccelServiceImpl {
	// client := spdk.NewClient(spdkAddress)
	return &AccelServiceImpl{nil}
}

// CryptoKeyCreate creates crypto key
func (p *AccelServiceImpl) CryptoKeyCreate(ctx context.Context, name string, cipher string, key []byte) (*AccelCryptoKeyCreateResult, error) {
	// TBD
	keyHalf := len(key) / 2
	params := AccelCryptoKeyCreateParams{
		Cipher: cipher,
		Key:    hex.EncodeToString(key[:keyHalf]),
		Key2:   hex.EncodeToString(key[keyHalf:]),
		Name:   name,
	}
	var result AccelCryptoKeyCreateResult
	err := p.client.Call(ctx, "accel_crypto_key_create", &params, &result)
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

// CryptoKeyDestroy destroys crypto key
func (p *AccelServiceImpl) CryptoKeyDestroy(_ context.Context, _ *AccelCryptoKeyDestroyParams) (*AccelCryptoKeyDestroyResult, error) {
	// TBD
	return nil, nil
}

// CryptoKeyList lists crypto keys
func (p *AccelServiceImpl) CryptoKeyList(_ context.Context, _ *AccelCryptoKeyGetParams) (*AccelCryptoKeyGetResult, error) {
	// TBD
	return nil, nil
}

// GetStats gets crypto key stats
func (p *AccelServiceImpl) GetStats(_ context.Context, _ *NvmfCreateSubsystemParams) (*NvmfCreateSubsystemResult, error) {
	// TBD
	return nil, nil
}
