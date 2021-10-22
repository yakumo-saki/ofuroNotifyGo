package config

// var logger = YLogger.GetLogger("zabbix")

// 設定をロードします
func LoadConfig() *ConfigStruct {
	// var conf ConfigStruct
	config := LoadFromEnvValue()

	SetDefaultConfig(config)

	logger.T("Config = ", config)

	return config
}
