#!/usr/bin/env bash
find . -maxdepth 1 -type d \( ! -name . \) -exec bash -c "cd '{}' &&   go build -buildmode=plugin" \;
