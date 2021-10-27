#!/bin/bash -eu

export AWS_REGION=ap-northeast-3

# WebHooks - Slack
export SLACK_CHANNEL="#general"
export SLACK_URL="Slack incoming webhook url like https://hooks.slack.com/services/~~~~"
export SLACK_DISP_NAME=ofuroNotify
export SLACK_ICON_EMOJI=ghost

# WebHooks - Mastodon
export MASTODON_KEY=dd964bf1a315eaab3d3d86c98f7186c1fff365894624830e20ad2e5b2530c928
export MASTODON_URL=https://example.com/api/v1/statuses

# WebHooks - Closed API
export WEBHOOK_URL=http://example.com/api/beep

# LOGLEVEL (default=WARN, good for production)
export LOG_LEVEL=DEBUG

# DEBUG for dynamodb local environment.
# comment out on production
export ENDPOINT=http://localhost:8000
export DISABLE_SSL=true
export DEBUG_NO_LAMBDA=1 # DEBUG no lambda function

# DEBUG DB R/W only. dont exec webhooks
export DEBUG_NO_HOOKS=0


go run ofuroNotify.go