#!/usr/bin/env sh
set -eu
image="${1:-goletalab-toml}"
docker build --platform linux/amd64 -f goletalab.Dockerfile -t "$image" .
docker run --rm "$image" go test ./...
