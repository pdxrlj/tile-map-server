package logger

import (
	"log/slog"

	"github.com/labstack/gommon/log"
)

type Logger struct {
	*log.Logger
	s slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(""),
		s:      slog.Logger{},
	}
}

func (l Logger) Info(s ...interface{}) {
	l.s.Info("info:", s)
}
