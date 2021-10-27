package ylog

type Logger interface {
	T(args ...interface{})
	D(args ...interface{})
	I(args ...interface{})
	W(args ...interface{})
	E(args ...interface{})
	F(args ...interface{})
}

type YLogger struct {
	name      string
	logOutput *int8
	logLevel  *int8
}

const LOG_OUTPUT_STDERR = "STDERR"
const LOG_OUTPUT_STDOUT = "STDOUT"

const LOG_LEVEL_TRACE = "TRACE"
const LOG_LEVEL_DEBUG = "DEBUG"
const LOG_LEVEL_INFO = "INFO"
const LOG_LEVEL_WARN = "WARN"
const LOG_LEVEL_ERROR = "ERROR"
const LOG_LEVEL_FATAL = "FATAL"

const LOG_TYPE_JSON = "JSON"
const LOG_TYPE_PLAIN = "PLAIN"

func (l *YLogger) T(args ...interface{}) {
	l.log(0, args...)
}

func (l *YLogger) D(args ...interface{}) {
	l.log(1, args...)
}

func (l *YLogger) I(args ...interface{}) {
	l.log(2, args...)
}

func (l *YLogger) W(args ...interface{}) {
	l.log(3, args...)
}

func (l *YLogger) E(args ...interface{}) {
	l.log(4, args...)
}

func (l *YLogger) F(args ...interface{}) {
	l.log(5, args...)
}
