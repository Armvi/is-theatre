package logger_zerolog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"time"
)

type Logger struct {
	Log zerolog.Logger
}

func (g Logger) Info(aMessage string) {
	lEvent := g.Log.Info()
	lEvent.Msg(aMessage)
}

func (g Logger) Error(aMessage string, aError error) {
	lEvent := g.Log.Error().Err(aError)
	lEvent.Msg(aMessage)
}

func (g Logger) Warning(aMessage string) {
	lEvent := g.Log.Warn()
	lEvent.Msg(aMessage)
}

func (g Logger) Debug(aMessage string) {
	lEvent := g.Log.Debug()
	lEvent.Msg(aMessage)
}

func (g Logger) Trace(aMessage string) {
	lEvent := g.Log.Trace()
	lEvent.Msg(aMessage)
}

func New(aWriters ...io.Writer) Logger {
	lMulti := zerolog.MultiLevelWriter(aWriters...)
	w := zerolog.ConsoleWriter{
		Out:        lMulti,
		TimeFormat: time.RFC822,
		PartsOrder: []string{"level", "time", "message"},
	}
	return Logger{log.Output(w)}
}
