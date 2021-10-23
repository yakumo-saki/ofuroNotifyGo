package hook

import (
	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type slackHook struct {
	apiKey string
}

func (sh *slackHook) init(cfg *config.ConfigStruct) bool {

	// logger = ylog.GetLogger("slackHook")

	if cfg.SlackApiKey == "" {
		return false
	}
	sh.apiKey = cfg.SlackApiKey

	return true
}

func (sh slackHook) exec() {
	logger := ylog.GetLogger("slackHook")
	logger.D("exec " + sh.apiKey)
}
