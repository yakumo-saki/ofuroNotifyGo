package config

import "errors"

// 設定のチェック、だめならerror。OKならnil
func CheckConfig(c *ConfigStruct) error {

	if c.Region == "" {
		return errors.New("please specify AWS_REGION env value")
	}

	return nil
}
