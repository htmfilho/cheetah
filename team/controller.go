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

	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
	colorRed   = "\033[31m"
)

func RunCeremony(ceremony string, team Team) {
	fmt.Printf("Team    : %s%s%s\n", colorBlue, team.Name, colorReset)
	fmt.Printf("Cycle   : %s%s%s\n", colorBlue, team.Cycle, colorReset)
	fmt.Printf("Sprint  : %s%s%s\n", colorBlue, team.Sprint.Name, colorReset)
	fmt.Printf("Ceremony: %s%s%s\n", colorBlue, ceremony, colorReset)
	fmt.Printf("Date    : %s%v. %s%s day of the sprint. %d days to go!%s\n",
		colorBlue, time.Now().Format("January 2, 2006"), colorRed, ToOrdinalNumber(team.Sprint.DaysFromStart()),
		team.Sprint.DaysToEnd(), colorReset)
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
		fmt.Printf("%s%s%s is working on:", colorBlue, member.Name, colorReset)
		for _, assignment := range member.Assignments {
			fmt.Printf("\n - %s: %s", assignment.Reference, assignment.Summary)
		}
		fmt.Print("\n................................................................................")
		pause()
	}
}

func pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
