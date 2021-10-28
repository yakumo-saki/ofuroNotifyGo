package hook

import (
	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
	"github.com/yakumo-saki/ofuroNotifyGo/util"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

var hooks []externalHooks
var dontExecHooks bool

func Init(cfg *config.ConfigStruct) {
	dontExecHooks = cfg.DebugNoHooks

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

func Exec(last db.LastOfuro) {
	logger := ylog.GetLogger()

	if dontExecHooks {
		logger.I("DontExecHooks enabled. dump parameters.")
		logger.Add("LastOfuro", last).I()

		logger.Add("Message", util.CreateMessage(last)).I()
		return
	}
	for _, v := range hooks {
		v.exec(last)
	}
}
