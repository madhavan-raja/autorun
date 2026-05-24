package internal

import "log/slog"

var Logger *slog.Logger

func init() {
	Logger = slog.Default()
}