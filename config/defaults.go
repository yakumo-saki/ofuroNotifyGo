package config

// 未設定ならデフォルト値をセットするもの
func SetDefaultConfig(c *ConfigStruct) {

	if c.LogType == "" {
		c.LogType = "PLAIN"
	}
	if c.LogLevel == "" {
		c.LogLevel = "WARN"
	}

	// slack
	if c.SlackChannel == "" {
		c.SlackChannel = "#general"
	}
	if c.SlackDisplayName == "" {
		c.SlackDisplayName = "bot"
	}
	if c.SlackIconEmoji == "" {
		c.SlackIconEmoji = "ghost"
	}

}
