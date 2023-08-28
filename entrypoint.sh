#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${CULTURA_STORE:-}" ]; then
  export CULTURA_CRUD_PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/cultura_crud/db/username --output text --query Parameter.Value)"
  export CULTURA_CRUD_PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/cultura_crud/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
