package hook

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/yakumo-saki/ofuroNotifyGo/config"
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

type mastodonHook struct {
	apiKey string
	url    string
}

func (sh *mastodonHook) init(cfg *config.ConfigStruct) bool {

	if cfg.MastodonKey == "" || cfg.MastodonUrl == "" {
		return false
	}

	sh.apiKey = cfg.MastodonKey
	sh.url = cfg.MastodonUrl

	return true
}

func (sh *mastodonHook) exec(_, _, message string) {
	logger := ylog.GetLogger("mastodonHook")
	logger.D("exec " + sh.url + " " + sh.apiKey)
	err := sh.post(message)
	if err != nil {
		logger.E("Mastodon POST failed: ", err)
		return
	}

	logger.D("Mastodon POST OK")

}

// curl -X POST -d "status=test message" --header "Authorization: Bearer $ACCESS_TOKEN" -sS http://localhost:3000/api/v1/statuses; echo $?
func (sh *mastodonHook) post(message string) error {
	logger := ylog.GetLogger("mastodonHook")

	params := url.Values{}
	params.Set("status", message+" 👁")

	req, err := http.NewRequest(
		"POST",
		sh.url,
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		logger.E("NewReader failed")
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+sh.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.TeeReader(resp.Body, os.Stderr)

	return err
}
