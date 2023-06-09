// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"encoding/json"
	"fmt"
)

// JSONRPCVersion holds the current version of json RPC protocol
const JSONRPCVersion = "2.0"

// RPCRequest holds the parameters required to request struct
type RPCRequest struct {
	RPCVersion string      `json:"jsonrpc"`
	Method     string      `json:"method"`
	ID         uint64      `json:"id"`
	Params     interface{} `json:"params,omitempty"`
}

// RPCResponse holds the parameters of the response struct
type RPCResponse struct {
	JSONRPCVersion string          `json:"jsonrpc"`
	ID             uint64          `json:"id"`
	Result         json.RawMessage `json:"result"`
	Error          RPCError        `json:"error"`
}

// RPCError holds the parameters of the error structs
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns formatted string of RPC error
func (e RPCError) Error() string {
	return fmt.Sprintf("Code=%d Msg=%s", e.Code, e.Message)
}
