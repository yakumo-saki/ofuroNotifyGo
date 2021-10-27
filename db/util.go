package db

import (
	"time"

	"github.com/yakumo-saki/ofuroNotifyGo/CS"
)

func CreateLastOfuro(inOut string, clickType string, lastInDateTime string) *LastOfuro {

	l := LastOfuro{
		Key:       LAST_TBL_KEY,
		UnixTime:  int64(time.Now().Unix()),
		ClickType: clickType,
		InOut:     inOut,
		DateTime:  time.Now().Format("20060102150405"),
	}

	switch inOut {
	case CS.OFURO_IN:
		l.Lastin = ""
		l.DurationSec = 0
	case CS.OFURO_OUT:
		l.Lastin = lastInDateTime

		l.DurationSec = dateTimeDiff(l.Lastin, l.DateTime)
	}

	return &l
}

// return diff (second)
func dateTimeDiff(before, after string) int64 {
	b4, err := time.Parse(CS.DATETIME_FORMAT, before)
	if err != nil {
		return 0
	}

	aft, err := time.Parse(CS.DATETIME_FORMAT, after)
	if err != nil {
		return 0
	}

	diff := aft.Unix() - b4.Unix()
	return diff
}
