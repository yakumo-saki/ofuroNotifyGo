#!/bin/bash -eu

export AWS_REGION=ap-northeast-3    # change to your dynamodb region

# WebHooks - Slack
export SLACK_CHANNEL="#general"
export SLACK_URL="Slack incoming webhook url like https://hooks.slack.com/services/~~~~"
export SLACK_DISP_NAME=ofuroNotify
export SLACK_ICON_EMOJI=bath        # Slack Emoji Icon

# WebHooks - Mastodon
export MASTODON_KEY=dd964bf1a315eaab3d3d86c98f7186c1fff365894624830e20ad2e5b2530c928
export MASTODON_URL=https://example.com/api/v1/statuses

# WebHooks - Closed API
export WEBHOOK_URL=http://example.com/api/beep

# LOGLEVEL (default=WARN, good for production)
export LOG_LEVEL=DEBUG

# DEBUG for dynamodb local environment.
# comment out on production
export ENDPOINT=http://localhost:8000  # DynamoDB local endpoint.
export DISABLE_SSL=true    # DynamoDB local needs this
export DEBUG_NO_LAMBDA=1   # DEBUG execute in your pc

export DEBUG_NO_HOOKS=0 # DEBUG DB R/W only. dont exec webhooks
export LOG_TYPE=JSON  # JSON or PLAIN or LTSV, recommend JSON for production

go run ofuroNotify.go
