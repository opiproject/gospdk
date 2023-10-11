// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package spdk implements the spdk json-rpc protocol
package spdk

import (
	"reflect"
	"testing"

	"go.opentelemetry.io/otel"
)

func TestSpdk_NewClient(t *testing.T) {
	tests := map[string]struct {
		address   string
		transport string
		wantPanic bool
	}{
		"testing unix": {
			"/var/tmp/spdk.sock",
			"unix",
			false,
		},
		"testing tcp": {
			"10.10.10.1:1234",
			"tcp",
			false,
		},
		"testing empty": {
			"",
			"",
			true,
		},
		"testing nonsense assuming unix": {
			"nonsense",
			"unix",
			false,
		},
	}

	// run tests
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("NewClient() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			before := NewClient(tt.address)
			after := &Client{
				transport: tt.transport,
				socket:    tt.address,
				id:        0,
				tracer:    otel.Tracer(""),
			}
			if !reflect.DeepEqual(before, after) {
				t.Error("response: expected", after, "received", before)
			}
		})
	}
}

func TestSpdk_Call(_ *testing.T) {

}
