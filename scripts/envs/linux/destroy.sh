#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

cd ../../../build
docker-compose down --volumes --remove-orphans
