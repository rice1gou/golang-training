#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

cd ../../

golangci-lint run -c ./golangci.yaml
