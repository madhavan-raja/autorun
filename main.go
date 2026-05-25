package main

import (
	"log/slog"

	"github.com/madhavan-raja/autorun/autorun"
	"github.com/madhavan-raja/autorun/internal"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger
}

func main() {
	a := autorun.New("")
	a.LoadProcesses()

	a.Start()
}
