package logger

import (
	"errors"
	"log/slog"
	"os"
)

const layout = "02/01/2006 15:04:05"

func NewLogger() (*slog.Logger, error) {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: replacer,
	})

	f, err := os.Open("logs.json")
	if err != nil {
		return nil, errors.New("error open log file")
	}
	defer f.Close()

	jsonHandler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: replacer,
	})

	return slog.New(slog.NewMultiHandler(textHandler, jsonHandler)), nil
}

func replacer(groups []string, a slog.Attr) slog.Attr {
	if groups == nil && a.Key == slog.TimeKey && a.Value.Kind() == slog.KindTime {
		t := a.Value.Time()
		return slog.String(a.Key, t.Format(layout))
	}

	return a
}
