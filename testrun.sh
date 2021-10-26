#!/bin/bash -eu

export AWS_REGION=ap-northeast-3

# WebHooks - Slack
export SLACK_CHANNEL="#general"
export SLACK_URL=https://hooks.slack.com/services/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
export SLACK_DISP_NAME=ofuroNotify
export SlackIconEmoji=ghost

# WebHooks - Mastodon
export MASTODON_KEY=fe964bf1a315eaab3d3d86c98f7186c1fff365894624830e20ad2e5b2530c928
export MASTODON_URL=https://example.com/api/v1/statuses

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