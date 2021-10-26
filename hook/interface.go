package hook

import "github.com/yakumo-saki/ofuroNotifyGo/config"

type externalHooks interface {
	init(cfg *config.ConfigStruct) bool
	exec(inOut, clickType, message string)
}
