package ylog

import (
	"path"
	"runtime"
	"strings"
)

type Logging struct {
	logLevel  int8
	logOutput int8
}

var logging Logging

var projectBaseDir string

// Must call this from file in project root dir
func Init() {
	_, filepath, _, ok := runtime.Caller(1)
	if !ok {
		projectBaseDir = ""
	}
	fn := path.Base(filepath)
	projectBaseDir = strings.Replace(filepath, fn, "", 1)
}

func GetLogger() *YLogger {
	pointer, filepath, _, ok := runtime.Caller(1)
	fn := runtime.FuncForPC(pointer)
	if !ok || fn == nil {
		return GetLoggerByName("UNKNOWN")
	}

	if false {
		return GetLoggerByName(fn.Name())
	}

	fp := strings.Replace(filepath, projectBaseDir, "", 1)
	return GetLoggerByName(fp)

}

// Loggerを取得します
// name = 名称。 main とか sub1 とか。 ログに出力される。
func GetLoggerByName(name string) *YLogger {
	logger := &YLogger{
		name:      name,
		logOutput: &logging.logOutput,
		logLevel:  &logging.logLevel,
	}

	return logger
}

// ログ出力先を設定します
// output = [STDERR | STDOUT]
func SetLogOutput(output string) {
	switch strings.ToUpper(output) {
	case "STDERR":
		logging.logOutput = 0
	case "STDOUT":
		logging.logOutput = 1
	default:
		panic("BAD logOutput " + output)
	}
}

// ログ出力しきい値を設定します。
// level = [TRACE | DEBUG | INFO | WARN | ERROR | FATAL]
func SetLogLevel(level string) {
	switch strings.ToUpper(level) {
	case "TRACE":
		logging.logLevel = 0
	case "DEBUG":
		logging.logLevel = 1
	case "INFO":
		logging.logLevel = 2
	case "WARN":
		logging.logLevel = 3
	case "ERROR":
		logging.logLevel = 4
	case "FATAL":
		logging.logLevel = 5
	default:
		panic("BAD loglevel " + level)
	}
}

// func (l *Logging) SetLogOutput(output string) {
// 	switch strings.ToUpper(output) {
// 	case "STDERR":
// 		l.logOutput = 0
// 	case "STDOUT":
// 		l.logOutput = 1
// 	default:
// 		panic("BAD logOutput " + output)
// 	}
// }

// func (l *Logging) SetLogLevel(level string) {
// 	switch strings.ToUpper(level) {
// 	case "TRACE":
// 		l.logLevel = 0
// 	case "DEBUG":
// 		l.logLevel = 1
// 	case "INFO":
// 		l.logLevel = 2
// 	case "WARN":
// 		l.logLevel = 3
// 	case "ERROR":
// 		l.logLevel = 4
// 	case "FATAL":
// 		l.logLevel = 5
// 	default:
// 		panic("BAD loglevel " + level)
// 	}
// }
