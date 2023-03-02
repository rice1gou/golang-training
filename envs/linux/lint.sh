#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

cd ../../

GOOS="linux"
GOARCH="amd64"

golangci-lint run -c ./envs/lint/.golangci.yaml
