package hook

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/db"
	"github.com/yakumo-saki/ofuroNotifyGo/util"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type mastodonHook struct {
	apiKey    string
	url       string
	msgSuffix string
}

func (mh *mastodonHook) init(cfg *config.ConfigStruct) bool {

	if cfg.MastodonKey == "" || cfg.MastodonUrl == "" {
		return false
	}

	mh.apiKey = cfg.MastodonKey
	mh.url = cfg.MastodonUrl
	mh.msgSuffix = cfg.MastodonMsgSuffix

	return true
}

func (mh *mastodonHook) exec(last db.LastOfuro) {
	logger := ylog.GetLogger()

	now := time.Now()
	logger.D("Mastodon POST start")

	err := mh.post(mh.createMessage(last))
	if err != nil {
		logger.E("Mastodon POST failed: ", err)
		return
	}

	logger.I(fmt.Sprintf("Mastodon POST took %v ms", time.Since(now).Milliseconds()))
}

func (mh *mastodonHook) createMessage(last db.LastOfuro) string {
	msg := util.CreateMessage(last)
	return msg + mh.msgSuffix
}

// curl -X POST -d "status=test message" --header "Authorization: Bearer $ACCESS_TOKEN" -sS http://localhost:3000/api/v1/statuses; echo $?
func (mh *mastodonHook) post(message string) error {
	logger := ylog.GetLogger()

	params := url.Values{}
	params.Set("status", message)

	req, err := http.NewRequest(
		"POST",
		mh.url,
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		logger.Add("err", err).E("NewReader failed")
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+mh.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.TeeReader(resp.Body, os.Stderr)

	return err
}
