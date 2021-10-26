package hook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/external"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type webHook struct {
	url string
}

func (wh *webHook) init(cfg *config.ConfigStruct) bool {

	if cfg.WebhookUrl == "" {
		return false
	}

	wh.url = cfg.WebhookUrl

	return true
}

func (wh *webHook) exec(inOut, clickType, message string) {
	logger := ylog.GetLogger("webHook")
	logger.D("exec " + wh.url)

	wh.post(inOut, clickType, message)
}

func (wh *webHook) post(inOut, clickType, message string) error {
	logger := ylog.GetLogger("webHook")

	var payload external.WebhookApi
	payload.InOut = inOut
	payload.ClickType = clickType
	payload.Message = message

	jsonBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		"POST",
		wh.url,
		bytes.NewReader(jsonBytes),
	)
	if err != nil {
		logger.E("NewReader failed")
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.TeeReader(resp.Body, os.Stderr)

	return nil
}
