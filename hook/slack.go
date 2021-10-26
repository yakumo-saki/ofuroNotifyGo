package hook

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/external"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type slackHook struct {
	url         string
	channel     string
	displayName string
	iconEmoji   string
}

func (sh *slackHook) init(cfg *config.ConfigStruct) bool {

	// logger = ylog.GetLogger("slackHook")

	if cfg.SlackHookUrl == "" {
		return false
	}
	sh.url = cfg.SlackHookUrl
	sh.channel = cfg.SlackChannel
	sh.displayName = cfg.SlackDisplayName
	sh.iconEmoji = cfg.SlackIconEmoji

	return true
}

func (sh *slackHook) exec(_, _, message string) {
	logger := ylog.GetLogger("slackHook")
	logger.D("exec " + sh.url)

	var body external.SlackMessage
	body.Channel = sh.channel
	body.IconEmoji = sh.iconEmoji

	body.Text = message

	json, _ := json.Marshal(body)
	err := sh.post(sh.url, json)
	if err != nil {
		logger.E("Slack POST Failed: ", err)
		return
	}
	logger.D("Slack POST OK")
}

func (sh *slackHook) post(url string, body []byte) error {
	logger := ylog.GetLogger("slackHook")
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		logger.E("NewReader failed")
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
