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

	conf.SlackApiKey = os.Getenv("SLACK_API_KEY")

	conf.MastodonKey = os.Getenv("MASTODON_KEY")
	conf.MastodonUrl = os.Getenv("MASTODON_URL")

	conf.WebhookUrl = os.Getenv("WEBHOOK_URL")

	return &conf
}
