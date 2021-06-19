package team

import "time"

const (
	AssignmentStatusTodo     = "todo"
	AssignmentStatusProgress = "progress"
	AssignmentStatusDone     = "done"
)

type Sprint struct {
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (sprint *Sprint) DaysFromStart() int {
	today := time.Now()
	return int(today.Sub(sprint.Start).Hours()) / 24
}

func (sprint *Sprint) DaysToEnd() int {
	today := time.Now()
	return int(sprint.End.Sub(today).Hours()) / 24
}

type Assignment struct {
	Reference string `json:"reference"`
	Summary   string `json:"summary"`
	Status    string `json:"status"`
}

type Member struct {
	Name        string       `json:"name"`
	Assignments []Assignment `json:"assignments"`
}

type Team struct {
	Cycle   string   `json:"cycle"`
	Sprint  Sprint   `json:"sprint"`
	Name    string   `json:"name"`
	Members []Member `json:"members"`
}
