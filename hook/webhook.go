package hook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
	"github.com/yakumo-saki/ofuroNotifyGo/external"
	"github.com/yakumo-saki/ofuroNotifyGo/util"
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

func (wh *webHook) exec(last db.LastOfuro) {
	logger := ylog.GetLogger()

	now := time.Now()
	logger.D("Webhook POST start")

	wh.post(last.InOut, last.ClickType, wh.createMessage(last))

	logger.I(fmt.Sprintf("Webhook POST took %v ms", time.Since(now).Milliseconds()))
}

func (wh *webHook) createMessage(last db.LastOfuro) string {
	msg := util.CreateMessage(last)
	return msg
}

func (wh *webHook) post(inOut, clickType, message string) error {
	logger := ylog.GetLogger()

	var payload external.WebhookApi
	payload.InOut = "bath" + inOut
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
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.TeeReader(resp.Body, os.Stderr)
	logger.D(resp.StatusCode, resp.Body)

	return nil
}
