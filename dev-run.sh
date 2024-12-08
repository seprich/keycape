#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

command -v reflex >/dev/null 2>&1 || { echo >&2 "Required tool 'reflex' is missing"; exit 1; }

# go install github.com/cespare/reflex@latest
# (and make sure ~/go/bin is on your PATH)

container_name="keycape_postgres"

exit_cleanup()
{
  docker stop "$container_name" 2>/dev/null || true
  docker rm "$container_name" 2>/dev/null || true
}
trap exit_cleanup SIGINT
docker run --name "$container_name" -p "5432:5432" \
       -e "POSTGRES_USER=test_user" -e "POSTGRES_PASSWORD=test_pw" -e "POSTGRES_DB=keycape" \
       -d postgres:17

$(cd be/; go mod tidy; go generate ./...)
reflex -c .reflex.conf
