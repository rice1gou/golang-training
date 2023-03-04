#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

cd ../../../build
docker-compose down --rmi all --volumes --remove-orphans
