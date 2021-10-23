#!/bin/bash -eu

export AWS_REGION=ap-northeast-3

# WebHooks
export SLACK_API_KEY=abc123

export MASTODON_KEY=bcd123
export MASTODON_URL=http://localhost

export WEBHOOK_URL=http://localhost

# for test environment
# comment out on production
export ENDPOINT=http://localhost:8000
export DISABLE_SSL=true

# LOGLEVEL (default=WARN)
export LOG_LEVEL=DEBUG

# DEBUG DB R/W only. no webhooks
export DEBUG_NO_HOOKS=0

go run ofuroNotify.go