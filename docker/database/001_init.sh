#!/bin/bash

set -o errexit

readonly REQUIRED_ENV_VARS=(
  "APP_DB_USER"
  "APP_DB_PASSWORD"
  "APP_DB_NAME"
  "POSTGRES_USER")

main() {
  check_env_vars_set
  init_user_and_db
}

check_env_vars_set() {
  for required_env_var in ${REQUIRED_ENV_VARS[@]}; do
    if [[ -z "${!required_env_var}" ]]; then
      echo "Error:
    Environment variable '$required_env_var' not set.
    Make sure you have the following environment variables set:
      ${REQUIRED_ENV_VARS[@]}
Aborting."
      exit 1
    fi
  done
}

init_user_and_db() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
     CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASSWORD';
     CREATE DATABASE $APP_DB_NAME;
     GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
     \c $APP_DB_NAME "$POSTGRES_USER"
     GRANT ALL ON SCHEMA public TO $APP_DB_USER
EOSQL
}

main "$@"