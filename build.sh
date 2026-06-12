#!/usr/bin/env bash

OUTPUT_BIN="./build/spymux"

if [ -f "$OUTPUT_BIN" ]; then
  rm "$OUTPUT_BIN"
fi

go build -o "$OUTPUT_BIN" main.go
