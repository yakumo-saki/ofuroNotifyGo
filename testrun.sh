#!/bin/bash -eu

export AWS_REGION=ap-northeast-3

# for test environment
# comment out on production
export ENDPOINT=http://localhost:8000
export DISABLE_SSL=true

# LOGLEVEL (default=WARN)
export LOG_LEVEL=DEBUG

go run ofuroNotify.go