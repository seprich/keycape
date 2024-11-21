package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger = loadLogger()

func loadLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, nil))
}
