#!/bin/sh
# Run migration
go run internal/app/migrations/migrations.go
if [ $? -ne 0 ]; then
    echo "Migration failed"
    exit 1
fi
# Start consumer
/app/quixzap-consumer