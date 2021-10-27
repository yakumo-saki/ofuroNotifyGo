package hook

import (
	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
)

type externalHooks interface {
	init(cfg *config.ConfigStruct) bool
	exec(last db.LastOfuro)
}
