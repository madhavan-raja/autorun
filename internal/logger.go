package internal

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

var Logger *slog.Logger

func init() {
	w := os.Stdout
	Logger = slog.New(tint.NewHandler(w, nil))
}