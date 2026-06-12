#!/usr/bin/env bash

OUTPUT_BIN="/home/primevoid/.local/bin/spymux"
# OUTPUT_BIN="./build/spymux"

# Remove existing build file if it exists
if [ -f "$OUTPUT_BIN" ]; then
  rm "$OUTPUT_BIN"
fi

# Build the new binary
go build -o "$OUTPUT_BIN" main.go
