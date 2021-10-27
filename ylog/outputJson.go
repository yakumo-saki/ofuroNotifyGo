package ylog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type awsJsonLog struct {
	time    string
	level   string
	message string
}

// internal log output implementation
func (l *YLogger) logOutputJson(level int8, args ...interface{}) {
	if *l.logLevel <= level {
		t := time.Now().Format("2006/1/2 15:04:05")
		lv := l.levelToString(level)

		message := fmt.Sprint(args...)

		out := awsJsonLog{
			time:    t,
			level:   lv,
			message: message,
		}

		bytes, _ := json.MarshalIndent(out, "", "\t")
		msg := string(bytes)

		switch *l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}

	}
}
