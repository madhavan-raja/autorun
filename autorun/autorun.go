package autorun

import (
	"log/slog"
	"sync"
	"time"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/madhavan-raja/autorun/internal/types"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger
}

type Autorun struct {
	Target    string
	processes []types.Process
}

func New(target string) Autorun {
	return Autorun{target, []types.Process{}}
}

func (a *Autorun) LoadProcesses() {
	a.processes = []types.Process{
		{
			Name:        "Test",
			Description: "Test Process",
			Cmd:         "echo 'Hello World'",
			RunOnStart:  false,
			Repeat:      true,
			Interval:    5,
		},
	}
}

func (a *Autorun) Start() {
	wg := sync.WaitGroup{}

	for _, p := range a.processes {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if p.RunOnStart {
				logger.Info("Running Process On Start", "process", p.Name)
			}

			if p.Repeat {
				ticker := time.NewTicker(time.Duration(p.Interval) * time.Second)
				defer ticker.Stop()

				for range ticker.C {
					logger.Info("Running Process", "process", p.Name)
				}
			}
		}()
	}

	wg.Wait()
}