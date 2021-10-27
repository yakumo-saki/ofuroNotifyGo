package config

import (
	"encoding/json"
)

type ConfigStruct struct {
	Region     string
	Endpoint   string
	DisableSSL bool

	LogLevel string
	LogType  string

	SlackHookUrl     string
	SlackChannel     string
	SlackDisplayName string
	SlackIconEmoji   string

	MastodonUrl string
	MastodonKey string
	WebhookUrl  string

	DebugNoHooks  bool
	DebugNoLambda bool
}

func (c ConfigStruct) String() string {
	bytes, _ := json.MarshalIndent(c, "", "\t")
	return string(bytes)
}
