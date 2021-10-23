package hook

import (
	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type webHook struct {
	url string
}

func (sh *webHook) init(cfg *config.ConfigStruct) bool {

	if cfg.WebhookUrl == "" {
		return false
	}

	sh.url = cfg.WebhookUrl

	return true
}

func (sh webHook) exec() {
	logger := ylog.GetLogger("webHook")
	logger.D("exec " + sh.url)
}
