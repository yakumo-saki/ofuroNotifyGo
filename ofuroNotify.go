package main

import (
	"os"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
	"github.com/yakumo-saki/ofuroNotifyGo/hook"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

var logger = ylog.GetLogger("main")
var Config config.ConfigStruct // dotenv + flags

const IN = "In"
const OUT = "Out"

func main() {

	logger = ylog.GetLogger("main")
	ylog.SetLogLevel("DEBUG")
	ylog.SetLogOutput("STDERR")

	cfg := config.LoadConfig()
	cfgerr := config.CheckConfig(cfg)
	logger.D(cfg)
	if cfgerr != nil {
		logger.F(cfgerr)
		os.Exit(10)
		return
	}

	db.Init(cfg)
	db.MakeSureTableExist()

	// お風呂 In Out 判定
	inOut := "In"
	last := db.GetLastOfuro()
	if last != nil && last.InOut == "In" {
		inOut = "Out"
	}

	//
	var newOfuro *db.LastOfuro
	switch inOut {
	case IN:
		newOfuro = db.CreateLastOfuro(inOut, "")
	case OUT:
		newOfuro = db.CreateLastOfuro(inOut, last.DateTime)
	}

	newHistory := db.LastOfuroToHistory(*newOfuro)
	err := db.PutLastOfuro(newOfuro)
	if err != nil {

	}
	db.PutHistory(&newHistory)

	// Do hooks
	hook.Init(cfg)
	hook.Exec()
}
