package types

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Autorun struct {
	mu sync.RWMutex
	cron *cron.Cron
	jobs map[string]*Process
}

func NewAutorun() *Autorun {
	c := cron.New()
	c.Start()

	return &Autorun {
		cron: c,
		jobs: make(map[string]*Process),
	}
}