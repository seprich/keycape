#!/usr/bin/env sh
set -e
cd "$(dirname "$0")"

devcontainer up --workspace-folder .

docker exec --user "$(id -u):$(id -g)" -it be_devcontainer-devcontainer-1 /bin/bash
