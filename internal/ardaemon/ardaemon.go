package ardaemon

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"sync"

	"github.com/madhavan-raja/autorun/db/sqlc"
	"github.com/madhavan-raja/autorun/internal"
	_ "github.com/mattn/go-sqlite3"
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
	mu      sync.RWMutex
	cron    *cron.Cron
	db      *sql.DB
	cronIds map[uint64]cron.EntryID
}

func NewArDaemon(ctx context.Context) *ArDaemon {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		logger.Error("Cannot access DB", "err", err)
		return nil
	}

	c := cron.New(cron.WithSeconds())
	c.Start()

	a := &ArDaemon{
		cron:    c,
		db:      db,
		cronIds: make(map[uint64]cron.EntryID),
	}

	processes, err := a.List(ctx)
	for _, p := range processes {
		err := a.schedule(ctx, sqlc.Process{
			ID: int64(p.Id),
			Name: p.Name,
			Description: sql.NullString{String: p.Description},
			Command: p.Cmd,
			Interval: int64(p.Interval),
		})

		if err != nil {
			logger.Error("Cannot add existing process", "err", err)
		}
	}

	return a
}

func (a *ArDaemon) schedule(ctx context.Context, p sqlc.Process) error {
	cronId, err := a.cron.AddFunc(fmt.Sprintf("@every %ds", p.Interval), func() { logger.Info("Executing Process", "process", p) })
	if err != nil {
		return err
	}

	a.cronIds[uint64(p.ID)] = cronId

	return nil
}

func (a *ArDaemon) Add(ctx context.Context, name string, description string, cmd string, interval uint32) (uint64, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	queries := sqlc.New(a.db)

	p, err := queries.AddProcess(ctx, sqlc.AddProcessParams{
		Name:        name,
		Description: sql.NullString{String: description},
		Command:     cmd,
		Interval:    int64(interval),
	})

	err = a.schedule(ctx, p)
	if err != nil {
		return 0, err
	}

	return uint64(p.ID), nil
}

func (a *ArDaemon) List(ctx context.Context) ([]Process, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	queries := sqlc.New(a.db)

	processes, err := queries.ListProcesses(ctx)
	if err != nil {
		return nil, err
	}

	processesCvt := []Process{}
	for _, p := range processes {
		processesCvt = append(processesCvt, Process{
			Id:          uint64(p.ID),
			Name:        p.Name,
			Description: p.Description.String,
			Cmd:         p.Command,
			Interval:    uint32(p.Interval),
		})
	}

	return processesCvt, nil
}
