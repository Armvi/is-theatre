package logger

type Logger interface {
	Info(aMessage string, aValues ...interface{})
	Error(aMessage string, aValues ...interface{})
	Warn(aMessage string, aValues ...interface{})
	Debug(aMessage string, aValues ...interface{})
	Trace(aMessage string, aValues ...interface{})
}
