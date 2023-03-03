#!/bin/zsh
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

cd ../../build
docker-compose up -d
