#!/usr/bin/env sh
find . -maxdepth 1 -type d \( ! -name . \) -exec sh -c "cd '{}' &&   go build -buildmode=plugin" \;
