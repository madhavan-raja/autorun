package ardaemon

import (
	"log/slog"
	"sync"

	"github.com/madhavan-raja/autorun/internal"
	"github.com/robfig/cron/v3"
)

var logger *slog.Logger

func init() {
	logger = internal.Logger.WithGroup("ardaemon")
}

type Process struct {
	Id          uint64
	Name        string
	Description string
	Cmd         string
	CronId      cron.EntryID
}

type ArDaemon struct {
	mu sync.RWMutex
	cron *cron.Cron
	Jobs map[uint64]*Process
}

func NewArDaemon() *ArDaemon {
	c := cron.New(cron.WithSeconds())
	c.Start()

	return &ArDaemon {
		cron: c,
		Jobs: make(map[uint64]*Process),
	}
}

func (a *ArDaemon) Add(name string, description string, cmd string, schedule string) (uint64, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	maxId := uint64(0)
	for key := range a.Jobs {
		maxId = max(maxId, key)
	}
	newId := maxId + 1

	cronId, err := a.cron.AddFunc(schedule, func() { logger.Info("Executing Process", "process", name) })
	if err != nil {
		return 0, err
	}

	a.Jobs[newId] = &Process{
		Id: newId,
		Name: name,
		Description: description,
		Cmd: cmd,
		CronId: cronId,
	}

	return newId, nil
}