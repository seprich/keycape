#!/usr/bin/env sh
set -e
cd "$(dirname "$0")"

docker compose --project-name be_devcontainer -f .devcontainer/compose.yaml down
