#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$DB_USER"  <<-EOSQL
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE USER postgres ;
    CREATE DATABASE golang-gorm;
    GRANT ALL PRIVILEGES ON DATABASE golang-gorm TO postgres;
EOSQL