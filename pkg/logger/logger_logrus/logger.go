package logger_logrus

import "is-theatre/pkg/logger"

type logrusLogger struct {
}

func NewLogrusLogger() (*logger.Logger, error) {
	return nil, nil
}

func (l logrusLogger) Info(aMessage string, aValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logrusLogger) Error(aMessage string, aValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logrusLogger) Warn(aMessage string, aValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logrusLogger) Debug(aMessage string, aValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l logrusLogger) Trace(aMessage string, aValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}
