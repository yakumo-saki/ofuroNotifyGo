package config

import "github.com/yakumo-saki/ofuroNotifyGo/ylog"

// var logger = YLogger.GetLogger("zabbix")

// 設定をロードします
func LoadConfig() *ConfigStruct {
	logger := ylog.GetLogger()

	config := LoadFromEnvValue()

	SetDefaultConfig(config)

	logger.T("Config = ", config)

	return config
}
