#!/usr/bin/env bash

OUTPUT_BIN="./bin/spymux"

if [ -f "$OUTPUT_BIN" ]; then
  rm "$OUTPUT_BIN"
fi

go bin -o "$OUTPUT_BIN" main.go
