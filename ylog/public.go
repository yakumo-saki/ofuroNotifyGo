package ylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Logging struct {
	logLevel   int8
	logOutput  int8
	outputType int8
}

const OUTPUT_PLAIN = 0
const OUTPUT_JSON = 1

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
	_, filepath, _, _ := runtime.Caller(1)
	// fn := runtime.FuncForPC(pointer)
	// if !ok || fn == nil {
	// 	return GetLoggerByName("UNKNOWN")
	// }

	// if false {
	// 	return GetLoggerByName(fn.Name())
	// }

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
		panic("SetLogLevel: BAD loglevel " + level)
	}
}

// ログ出力形式を設定します。
// logtype = [PLAIN | JSON]
func SetLogType(logtype string) {
	switch strings.ToUpper(logtype) {
	case "PLAIN":
		logging.outputType = OUTPUT_PLAIN
	case "JSON":
		logging.outputType = OUTPUT_JSON
	default:
		panic("SetLogType: BAD logtype " + logtype)
	}
}

func (l *YLogger) log(level int8, args ...interface{}) {
	switch logging.outputType {
	case OUTPUT_PLAIN:
		l.logOutputPlain(level, args...)
	case OUTPUT_JSON:
		l.logOutputJson(level, args...)
	default:
		panic("UNKNOWN OUTPUT : " + fmt.Sprint(logging.outputType))
	}
}
