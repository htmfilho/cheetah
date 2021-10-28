package team

import (
	"fmt"
	"time"
)

const (
	AssignmentStatusTodo     = "todo"
	AssignmentStatusProgress = "progress"
	AssignmentStatusDone     = "done"

	AssignmentTypeWeekly = "weekly"
)

type Sprint struct {
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (sprint *Sprint) DaysFromStart() int {
	today := time.Now()
	return 1 + int(today.Sub(sprint.Start).Hours())/24
}

func (sprint *Sprint) DaysToEnd() int {
	today := time.Now()
	return int(sprint.End.Sub(today).Hours()) / 24
}

func (sprint *Sprint) DaysToStart() int {
	today := time.Now()
	return int(sprint.Start.Sub(today).Hours()) / 24
}

func (sprint *Sprint) Passed() bool {
	today := time.Now()
	return today.After(sprint.End)
}

func (sprint *Sprint) Started() bool {
	today := time.Now()
	return !today.Before(sprint.Start)
}

type Assignment struct {
	Reference string `json:"reference"`
	Summary   string `json:"summary"`
	Status    string `json:"status"`
	Type      string `json:"type"`
}

func (assignment *Assignment) Print(member Member, members []Member) {
	statusColors := make(map[string]string)
	statusColors[AssignmentStatusTodo] = colorReset
	statusColors[AssignmentStatusProgress] = colorYellow
	statusColors[AssignmentStatusDone] = colorGreen

	if assignment.Type == AssignmentTypeWeekly {
		_, week := time.Now().ISOWeek()
		weeklyAssignedMembers := membersWeeklyAssigned(members)
		if member.Name == weeklyAssignedMembers[week%len(weeklyAssignedMembers)].Name {
			fmt.Printf("\n ! %s%s%s", statusColors[AssignmentStatusProgress], assignment.Summary, colorReset)
		}
	} else {
		fmt.Printf("\n . %s%s: %s%s", statusColors[assignment.Status], assignment.Reference, assignment.Summary, colorReset)
	}
}

func membersWeeklyAssigned(members []Member) []Member {
	assignedOnes := make([]Member, 0)
	for _, member := range members {
		for _, assignment := range member.Assignments {
			if assignment.Type == AssignmentTypeWeekly {
				assignedOnes = append(assignedOnes, member)
				break
			}
		}
	}
	return assignedOnes
}

type Member struct {
	Name        string       `json:"name"`
	Assignments []Assignment `json:"assignments"`
}

type Team struct {
	Cycle   string   `json:"cycle"`
	Sprint  Sprint   `json:"sprint"`
	Name    string   `json:"name"`
	Manager string   `json:"manager"`
	Members []Member `json:"members"`
}
