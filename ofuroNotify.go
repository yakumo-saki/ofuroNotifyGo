package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yakumo-saki/ofuroNotifyGo/CS"
	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
	"github.com/yakumo-saki/ofuroNotifyGo/hook"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

var Config config.ConfigStruct // dotenv + flags

const IN = "In"
const OUT = "Out"

type DeviceEvent struct {
	buttonClicked IoTButtonEvent
}
type IoTButtonEvent struct {
	ClickType    string `json:"clickType"`
	ReportedTime string `json:"ReportedTime"`
}

type ConfigError struct {
	Error string
}

// ctx context.Context
func HandleRequest(ctx context.Context, event DeviceEvent) (string, error) {
	logger := ylog.GetLogger()
	ylog.SetLogLevel("DEBUG")
	ylog.SetLogOutput("STDERR")

	cfg := config.LoadConfig()
	cfgerr := config.CheckConfig(cfg)
	if cfgerr != nil {
		logger.F(cfg)
		logger.F(cfgerr)
		os.Exit(10)
		return "Config Error", errors.New("Config Error")
	}

	ylog.SetLogLevel(cfg.LogLevel)
	ylog.SetLogType(cfg.LogType)

	// logger.I(event)
	// bytes, _ := json.MarshalIndent(event, "", "\t")
	// logger.I(string(bytes))

	db.Init(cfg)
	db.MakeSureTableExist()

	// お風呂 In Out 判定
	inOut := CS.OFURO_IN
	last := db.GetLastOfuro()
	if last != nil && last.InOut == CS.OFURO_IN {
		inOut = CS.OFURO_OUT
	}

	//
	clickType := event.buttonClicked.ClickType

	//
	var newOfuro *db.LastOfuro
	switch inOut {
	case IN:
		newOfuro = db.CreateLastOfuro(inOut, clickType, "")
	case OUT:
		newOfuro = db.CreateLastOfuro(inOut, clickType, last.DateTime)
	}

	newHistory := db.LastOfuroToHistory(*newOfuro)
	err := db.PutLastOfuro(newOfuro)
	if err != nil {
		return "PutLastOfuro fail", err
	}
	db.PutHistory(&newHistory)

	// Do hooks
	hook.Init(cfg)
	hook.Exec(*newOfuro)

	return clickType, nil
}

func main() {
	ylog.Init()
	cfg := config.LoadConfig() // DEBUGモード判定のために一度読んでしまう

	if cfg.DebugNoLambda {
		logger := ylog.GetLogger()
		ylog.SetLogLevel("DEBUG")
		ylog.SetLogOutput("STDERR")

		now := time.Now()
		ev := DeviceEvent{IoTButtonEvent{ClickType: "SINGLE"}}
		HandleRequest(context.TODO(), ev)

		logger.I(fmt.Sprintf("Overall took %v ms", time.Since(now).Milliseconds()))
	} else {
		lambda.Start(HandleRequest)
	}
}
