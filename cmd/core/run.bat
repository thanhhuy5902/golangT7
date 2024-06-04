@echo off
SETLOCAL
SET SERVICEWEAVER_CONFIG=config.toml
weaver generate
go mod tidy
go run .