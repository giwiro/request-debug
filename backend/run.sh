#!/usr/bin/env bash

export REQUEST_DEBUG_CONFIG_PATH="$(pwd)/config.yaml"

go run ./cmd/api/main.go
