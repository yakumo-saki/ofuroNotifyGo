package util

import (
	"fmt"
)

const CLICK_SINGLE = "SINGLE"
const CLICK_LONG = "LONG"
const CLICK_DOUBLE = "DOUBLE"

func secondToText(second int64) string {
	return fmt.Sprint(second)
	// delta = timedelta(=int(second) )

	// ret = ""
	// if delta.days > 0 {
	// 	ret = "{delta.days}日"
	// }

	// m, s = divmod(delta.seconds, 60)
	// h, m = divmod(m, 60)

	// if h > 0 {
	// 	ret = "{ret} {h}時間"
	// }

	// ret = "{ret} {m}分 {s}秒"

	// return ret
}

func CreateMessageIn(clickType string) string {

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

func CreateMessageOut(clickType string, duration_sec int64) string {
	nagasa := secondToText(duration_sec)
	message := "ほかぱい！ ✨ (" + nagasa + ")"

	return message
}
