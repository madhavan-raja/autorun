package internal

import "github.com/robfig/cron/v3"

type Process struct {
	Id          string
	Name        string
	Description string
	Cmd         string
	CronId      cron.EntryID
}
