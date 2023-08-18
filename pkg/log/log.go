package log

import (
	"io"
	"log/slog"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	DefaultLogLevel = slog.LevelInfo
)

func Init(level ...slog.Level) {
	if len(level) > 0 {
		DefaultLogLevel = level[0]
	}
	record := io.MultiWriter(os.Stdout, FileWriter())

	slog.SetDefault(slog.New(slog.NewTextHandler(record, &slog.HandlerOptions{
		AddSource: true,
		Level:     DefaultLogLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				a.Value = slog.StringValue(time.Now().Format(time.DateTime))
			}
			return a
		},
	})))
}

func FileWriter() *lumberjack.Logger {
	fileWriter := &lumberjack.Logger{
		Filename:         "./logs/tile.log",
		MaxSize:          1,
		MaxBackups:       3,
		MaxAge:           28,
		Compress:         true,
		LocalTime:        true,
		BackupTimeFormat: time.DateTime,
	}

	return fileWriter
}
