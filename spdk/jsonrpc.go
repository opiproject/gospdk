// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023 Intel Corporation

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	// ErrFailedSpdkCall indicates that the bridge failed to execute SPDK call
	ErrFailedSpdkCall = status.Error(codes.Unknown, "Failed to execute SPDK call")
	// ErrUnexpectedSpdkCallResult indicates that the bridge got an error from SPDK
	ErrUnexpectedSpdkCallResult = status.Error(codes.FailedPrecondition, "Unexpected SPDK call result.")
)

// JSONRPC represents an interface to execute JSON RPC to SPDK
type JSONRPC interface {
	GetID() uint64
	GetVersion(context.Context) string
	StartUnixListener() net.Listener
	Call(ctx context.Context, method string, args, result interface{}) error
}

// Client implements JSONRPC interface
type Client struct {
	transport string
	socket    string
	id        uint64
	tracer    trace.Tracer
}

// build time check that struct implements interface
var _ JSONRPC = (*Client)(nil)

// NewClient creates a new instance of JSONRPC which is capable to
// interact with either unix domain socket, e.g.: /var/tmp/spdk.sock
// or with tcp connection ip and port tuple, e.g.: 10.1.1.2:1234
func NewClient(socketPath string) *Client {
	if socketPath == "" {
		log.Panic("empty socketPath is not allowed")
	}
	protocol := "tcp"
	if _, _, err := net.SplitHostPort(socketPath); err != nil {
		protocol = "unix"
	}
	log.Printf("Connection to SPDK will be via: %s detected from %s", protocol, socketPath)
	return &Client{
		transport: protocol,
		socket:    socketPath,
		id:        0,
		tracer:    otel.Tracer(""),
	}
}

// GetID implements low level rpc request/response handling
func (r *Client) GetID() uint64 {
	return r.id
}

// GetVersion implements low level rpc request/response handling
func (r *Client) GetVersion(ctx context.Context) string {
	var ver GetVersionResult
	err := r.Call(ctx, "spdk_get_version", nil, &ver)
	if err != nil {
		msg := fmt.Sprintf("Could not get spdk version: %v", err)
		log.Print(msg)
		return ""
	}
	log.Printf("Received from SPDK: %v", ver)
	return ver.Version
}

// StartUnixListener is utility function used to create new listener in tests
func (r *Client) StartUnixListener() net.Listener {
	if err := os.RemoveAll(r.socket); err != nil {
		log.Fatal(err)
	}
	ln, err := net.Listen("unix", r.socket)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	return ln
}

// Call implements low level rpc request/response handling
func (r *Client) Call(ctx context.Context, method string, args, result interface{}) error {
	id := atomic.AddUint64(&r.id, 1)

	_, childSpan := r.tracer.Start(ctx, "spdk."+method)
	childSpan.SetAttributes(attribute.Int64("request.id", int64(id)))
	defer childSpan.End()

	request := RPCRequest{
		RPCVersion: JSONRPCVersion,
		ID:         id,
		Method:     method,
		Params:     args,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("%s: %s", method, err)
	}

	log.Printf("Sending to SPDK: %s", data)

	resp, _ := r.communicate(data)

	var response RPCResponse
	err = json.NewDecoder(resp).Decode(&response)
	jsonresponse, _ := json.Marshal(response)
	log.Printf("Received from SPDK: %s", jsonresponse)
	if err != nil {
		return fmt.Errorf("%s: %s", method, err)
	}
	if response.ID != id {
		return fmt.Errorf("%s: json response ID mismatch", method)
	}
	if response.Error.Code != 0 {
		return fmt.Errorf("%s: json response error: %s", method, response.Error.Message)
	}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return fmt.Errorf("%s: %s", method, err)
	}
	return nil
}

func (r *Client) communicate(buf []byte) (io.Reader, error) {
	// connect
	conn, err := net.Dial(r.transport, r.socket)
	if err != nil {
		log.Fatal(err)
	}
	// write
	_, err = conn.Write(buf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// close
	switch conn := conn.(type) {
	case *net.TCPConn:
		err = conn.CloseWrite()
	case *net.UnixConn:
		err = conn.CloseWrite()
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// read
	return bufio.NewReader(conn), nil
}
