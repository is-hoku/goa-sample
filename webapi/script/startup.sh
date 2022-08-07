#!/bin/bash
chmod 777 ./script/migrate-up
chmod 777 ./script/migrate-down
chmod 777 ./script/migrate-create
chmod 777 ./script/sql-generate
air -c .air.toml
