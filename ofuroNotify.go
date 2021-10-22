package main

import (
	"os"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

var logger = ylog.GetLogger("main")
var Config config.ConfigStruct // dotenv + flags

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
	db.Test()

}
