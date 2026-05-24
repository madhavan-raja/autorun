package types

type Process struct {
	Name        string
	Description string
	Cmd         string
	RunOnStart  bool
	Repeat      bool
	Interval    uint64
}
