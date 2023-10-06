// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (c) 2022 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// Copyright (C) 2023 Intel Corporation

// main is the main package of the application
package main

import (
	"context"
	"flag"
	"log"

	"github.com/opiproject/gospdk/spdk"
)

func main() {
	ctx := context.Background()

	var spdkAddress string
	flag.StringVar(&spdkAddress, "spdk_addr", "/var/tmp/spdk.sock", "Points to SPDK unix socket/tcp socket to interact with")
	flag.Parse()

	// use like this:
	jsonRPC := spdk.NewClient(spdkAddress)
	version := jsonRPC.GetVersion(ctx)
	log.Printf("Received from SPDK: %v", version)

	// or like this:
	var ver spdk.GetVersionResult
	err := jsonRPC.Call(ctx, "spdk_get_version", nil, &ver)
	if err != nil {
		log.Fatalf("failed to get SPDK version: %v", err)
	}

	log.Printf("Received from SPDK: %v", ver)
}
