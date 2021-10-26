package config

import (
	"os"
	"strconv"
)

// 実行時の環境変数からconfigを生成
//
func LoadFromEnvValue() *ConfigStruct {

	var conf ConfigStruct

	conf.Region = os.Getenv("AWS_REGION")
	conf.Endpoint = os.Getenv("ENDPOINT")
	conf.DisableSSL, _ = strconv.ParseBool(os.Getenv("DISABLE_SSL"))
	conf.Output = os.Getenv("OUTPUT")
	conf.Loglevel = os.Getenv("LOGLEVEL")

	conf.SlackHookUrl = os.Getenv("SLACK_URL")
	conf.SlackChannel = os.Getenv("SLACK_CHANNEL")
	conf.SlackDisplayName = os.Getenv("SLACK_DISP_NAME")
	conf.SlackIconEmoji = os.Getenv("SLACK_ICON_EMOJI")

	conf.MastodonKey = os.Getenv("MASTODON_KEY")
	conf.MastodonUrl = os.Getenv("MASTODON_URL")

	conf.WebhookUrl = os.Getenv("WEBHOOK_URL")

	conf.DebugNoHooks, _ = strconv.ParseBool(os.Getenv("DEBUG_NO_HOOKS"))
	conf.DebugNoLambda, _ = strconv.ParseBool(os.Getenv("DEBUG_NO_LAMBDA"))

	return &conf
}
