#!/bin/bash -eu

export AWS_REGION=ap-northeast-3

# WebHooks - Slack
export SLACK_CHANNEL="#general"
export SLACK_URL=https://hooks.slack.com/services/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
export SLACK_DISP_NAME=ofuroNotify
export SlackIconEmoji=ghost

# WebHooks - Mastodon
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