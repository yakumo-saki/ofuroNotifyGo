package util

import (
	"fmt"

	"github.com/yakumo-saki/ofuroNotifyGo/db"
)

const CLICK_SINGLE = "SINGLE"
const CLICK_LONG = "LONG"
const CLICK_DOUBLE = "DOUBLE"

const MIN_SEC = 60
const HOUR_SEC = MIN_SEC * 60
const DAY_SEC = HOUR_SEC * 24

func divmod(target, divider int64) (ans, amari int64) {
	ans = target / divider
	amari = target % divider
	return ans, amari
}

func secondToText(second int64) string {

	ret := ""

	days, daysmod := divmod(second, DAY_SEC)

	if days > 0 {
		ret = fmt.Sprintf("%d日", days)
	}

	hour, hourmod := divmod(daysmod, HOUR_SEC)
	minute, second := divmod(hourmod, MIN_SEC)

	if hour > 0 {
		ret = fmt.Sprintf("%s %d時間", ret, hour)
	}

	if minute > 0 {
		ret = fmt.Sprintf("%s %d分", ret, minute)
	}

	ret = fmt.Sprintf("%s %d秒", ret, second)

	return ret
}

func CreateMessage(last db.LastOfuro) string {
	if last.InOut == "In" {
		return createMessageIn(last.ClickType)
	} else {
		return createMessageOut(last.ClickType, last.DurationSec)
	}
}

func createMessageIn(clickType string) string {

	message := ""
	if clickType == CLICK_DOUBLE {
		message = "シャワる 🛀"
	} else if clickType == CLICK_LONG {
		message = "おふろる 📲🛀"
	} else {
		message = "おふろる 🛀"
	}

	return message
}

func createMessageOut(clickType string, duration_sec int64) string {
	nagasa := secondToText(duration_sec)
	message := "ほかぱい！ ✨ (" + nagasa + ")"

	return message
}
