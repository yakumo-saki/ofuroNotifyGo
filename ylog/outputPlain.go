package ylog

import (
	"fmt"
	"os"
	"time"
)

// internal log output implementation
func (l *YLogger) logOutputPlain(level int8, args ...interface{}) {
	if *l.logLevel <= level {
		t := time.Now().Format("2006/1/2 15:04:05")
		lv := l.levelToString(level)

		msg := t + " " + l.name + " " + lv + " " + fmt.Sprint(args...)
		switch *l.logOutput {
		case 0:
			fmt.Fprintln(os.Stderr, msg)
		case 1:
			fmt.Fprintln(os.Stdout, msg)
		}

	}
}
