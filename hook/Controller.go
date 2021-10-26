package hook

import (
	"github.com/yakumo-saki/ofuroNotifyGo/config"
)

var hooks []externalHooks

func Init(cfg *config.ConfigStruct) {
	var slack slackHook
	if slack.init(cfg) {
		hooks = append(hooks, &slack)
	}

	var mast mastodonHook
	if mast.init(cfg) {
		hooks = append(hooks, &mast)
	}

	var wh webHook
	if wh.init(cfg) {
		hooks = append(hooks, &wh)
	}

}

func Exec(inOut, clickType, message string) {
	for _, v := range hooks {
		v.exec(inOut, clickType, message)
	}
}
