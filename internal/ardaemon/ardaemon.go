package ardaemon

import (
	"fmt"
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
	Interval    uint32
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

func (a *ArDaemon) Add(name string, description string, cmd string, interval uint32) (uint64, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	maxId := uint64(0)
	for key := range a.Jobs {
		maxId = max(maxId, key)
	}
	newId := maxId + 1

	cronId, err := a.cron.AddFunc(fmt.Sprintf("@every %ds", interval), func() { logger.Info("Executing Process", "process", name) })
	if err != nil {
		return 0, err
	}

	a.Jobs[newId] = &Process{
		Id: newId,
		Name: name,
		Description: description,
		Cmd: cmd,
		Interval: interval,
		CronId: cronId,
	}

	return newId, nil
}