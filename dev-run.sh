#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

command -v reflex >/dev/null 2>&1 || { echo >&2 "Required tool 'reflex' is missing"; exit 1; }
command -v awslocal --version >/dev/null 2>&1 || { echo >&2 "Required tool 'awslocal' is missing"; exit 1; }

initialize_localstack()
{
  awslocal cognito-idp create-user-pool --pool-name test
}

# go install github.com/cespare/reflex@latest
# (and make sure ~/go/bin is on your PATH)

container_name="keycape_postgres"

exit_cleanup()
{
  docker compose down
}
trap exit_cleanup SIGINT
docker compose up -d

$(cd be/; go mod tidy; go generate ./...)
reflex -c .reflex.conf
