#!/bin/sh

# Run migration
go run app/migrations/migrations.go

# Start API
go run cmd/fin-api/main.go