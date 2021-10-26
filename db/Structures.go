package db

const LAST_TBL = "LastOfuro"
const LAST_TBL_KEY = "LAST"

const HIST_TBL = "OfuroHistories"
const HIST_TBL_KEY = "OfuroHistories"

// { "UnixTime" : { "N" : "1633432212" },
//   "InOut" : { "S" : "In" },
//   "DateTime" : { "S" : "20211005111012" },
//   "LastIn" : { "NULL" : true } }
type LastOfuro struct {
	Key         string `dynamo:"key,hash"`
	UnixTime    int64  // 更新したときのUnixTime。OfuroHistoryと合わせる
	InOut       string // "In" or "Out"
	DateTime    string // yyyyMMddHH24mmss
	ClickType   string // SINGLE DOUBLE LONG
	DurationSec int64  // Out only.
	Lastin      string // "Out" の時のみ。 InしたときのDateTime
}

type OfuroHistory struct {
	UnixTime    int64  `dynamo:",hash"` // 更新したときのUnixTime
	InOut       string // "In" or "Out"
	DateTime    string // yyyyMMddHH24mmss
	ClickType   string // SINGLE DOUBLE LONG
	DurationSec int64  // Out only.  0 when In
	Lastin      string // "Out" の時のみ。 InしたときのDateTime
}

func LastOfuroToHistory(lastOfuro LastOfuro) OfuroHistory {
	return OfuroHistory{
		UnixTime:    lastOfuro.UnixTime,
		InOut:       lastOfuro.InOut,
		DateTime:    lastOfuro.DateTime,
		ClickType:   lastOfuro.ClickType,
		DurationSec: lastOfuro.DurationSec,
		Lastin:      lastOfuro.Lastin,
	}
}
