package autorun

import (
	"log/slog"
	"sync"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/robfig/cron/v3"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger.WithGroup("autorun")
}

type Process struct {
	Id          uint64
	Name        string
	Description string
	Cmd         string
	CronId      cron.EntryID
}

type Autorun struct {
	mu sync.RWMutex
	cron *cron.Cron
	jobs map[uint64]*Process
}

func NewAutorun() *Autorun {
	c := cron.New(cron.WithSeconds())
	c.Start()

	return &Autorun {
		cron: c,
		jobs: make(map[uint64]*Process),
	}
}

func (a *Autorun) Add(name string, description string, cmd string, schedule string) (uint64, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	maxId := uint64(0)
	for key := range a.jobs {
		maxId = max(maxId, key)
	}
	newId := maxId + 1

	cronId, err := a.cron.AddFunc(schedule, func() { logger.Info("Executing Process", "process", name) })
	if err != nil {
		return 0, err
	}

	a.jobs[newId] = &Process{
		Id: newId,
		Name: name,
		Description: description,
		Cmd: cmd,
		CronId: cronId,
	}

	return newId, nil
}