package team

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	CeremonyStandUp = "StandUp"
)

func RunCeremony(ceremony string, team Team) {
	fmt.Printf("Team    : %s\n", team.Name)
	fmt.Printf("Cycle   : %s\n", team.Cycle)
	fmt.Printf("Sprint  : %s\n", team.Sprint)
	fmt.Printf("Ceremony: %v\n", ceremony)
	fmt.Printf("Date    : %v\n", time.Now().Format("January 2, 2006"))
	fmt.Println("--------------------------------------------------------------------------------")

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
		fmt.Printf("| %s is working on:", member.Name)
		for _, assignment := range member.Assignments {
			fmt.Printf("\n| %s: %s", assignment.Reference, assignment.Summary)
		}
		fmt.Print("\n................................................................................")
		pause()
	}
}

func pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
