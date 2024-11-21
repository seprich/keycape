#!/usr/bin/env sh
set -e
cd "$(dirname "$0")"

# go install github.com/cespare/reflex@latest
# (and make sure ~/go/bin is on your PATH)

reflex -c .reflex.conf
