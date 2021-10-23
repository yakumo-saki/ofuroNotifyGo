package hook

import (
	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type mastodonHook struct {
	apiKey string
	url    string
}

func (sh *mastodonHook) init(cfg *config.ConfigStruct) bool {

	if cfg.MastodonKey == "" || cfg.MastodonUrl == "" {
		return false
	}

	sh.apiKey = cfg.MastodonKey
	sh.url = cfg.MastodonUrl

	return true
}

func (sh mastodonHook) exec() {
	logger := ylog.GetLogger("mastodonHook")
	logger.D("exec " + sh.url + " " + sh.apiKey)
}
