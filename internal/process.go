package internal

import "github.com/robfig/cron/v3"

type Process struct {
	Id          uint64
	Name        string
	Description string
	Cmd         string
	CronId      cron.EntryID
}
