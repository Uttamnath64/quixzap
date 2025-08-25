#!/bin/sh

# Run migration
go run internal/migrations/migrations.go

# Start API
go run cmd/server/main.go