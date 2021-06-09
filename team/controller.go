package team

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	// Ceremonies
	CeremonyStandUp = "StandUp"
)

func RunCeremony(ceremony string, team Team) {
	fmt.Printf("Team  : %s\n", team.Name)
	fmt.Printf("Cycle : %s\n", team.Cycle)
	fmt.Printf("Sprint: %s\n", team.Sprint)

	switch ceremony {
	case CeremonyStandUp:
		runStandUp(team)
	default:
		runStandUp(team)
	}
}

func runStandUp(team Team) {
	members := team.Members
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(members), func(i, j int) { members[i], members[j] = members[j], members[i] })
	for _, member := range members {
		fmt.Printf("\n%s is working on:\n", member.Name)
		for _, assignment := range member.Assignments {
			fmt.Printf(" %s: %s\n", assignment.Reference, assignment.Summary)
		}
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
