#!/bin/sh

# Run staticcheck
echo "Running staticcheck..."
staticcheck ./...

# Run gosec
echo "Running gosec..."
gosec ./...