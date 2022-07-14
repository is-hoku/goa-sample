#!/bin/bash
chmod 777 ./script/migrate-up
chmod 777 ./script/migrate-down
chmod 777 ./script/migrate-create
air -c .air.toml
