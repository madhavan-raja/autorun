package main

import (
	"fmt"
	"log/slog"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/types"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger
}

func main() {
	processes := []types.Process{
		{
			Name: "Test",
			Description: "Test Process",
			Cmd: "echo 'Hello World'",
			RunOnStart: false,
			Repeat: true,
			Interval: 60,
		},
	}

	fmt.Printf("%v\n", processes)
}
