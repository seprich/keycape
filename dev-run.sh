#!/usr/bin/env sh
set -e
cd "$(dirname "$0")"

command -v reflex >/dev/null 2>&1 || { echo >&2 "Required tool 'reflex' is missing"; exit 1; }

# go install github.com/cespare/reflex@latest
# (and make sure ~/go/bin is on your PATH)

$(cd be/; go mod tidy; go generate ./...)

reflex -c .reflex.conf
