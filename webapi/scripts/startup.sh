#!/bin/bash
chmod 777 ./scripts/migrate-up
chmod 777 ./scripts/migrate-down
chmod 777 ./scripts/migrate-create
air -c .air.toml
