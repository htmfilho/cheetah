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

	colorReset  = "\033[0m"
	colorBlue   = "\033[34m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

func RunCeremony(ceremony string, team Team) {
	fmt.Printf("Team    : %s%s%s\n", colorBlue, team.Name, colorReset)
	fmt.Printf("Cycle   : %s%s%s\n", colorBlue, team.Cycle, colorReset)
	fmt.Printf("Sprint  : %s%s%s\n", colorBlue, team.Sprint.Name, colorReset)
	fmt.Printf("Ceremony: %s%s%s\n", colorBlue, ceremony, colorReset)
	fmt.Printf("Date    : %s%v.", colorBlue, time.Now().Format("January 2, 2006"))
	if !team.Sprint.Passed() {
		fmt.Printf(" %s%s day of the sprint. %d days to go!%s\n", colorRed, ToOrdinalNumber(team.Sprint.DaysFromStart()), team.Sprint.DaysToEnd(), colorReset)
	} else {
		fmt.Printf(" Sprint finished on %v.\n", team.Sprint.End.Format("January 2, 2006"))
	}
	fmt.Println()

	switch ceremony {
	case CeremonyStandUp:
		runStandUp(team)
	default:
		runStandUp(team)
	}
}

func runStandUp(team Team) {
	members := make([]Member, len(team.Members))
	copy(members, team.Members)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(members), func(i, j int) { members[i], members[j] = members[j], members[i] })

	for _, member := range members {
		if len(member.Assignments) == 0 {
			continue
		}

		fmt.Printf("%s%s%s is assigned to:", colorBlue, member.Name, colorReset)
		for _, assignment := range member.Assignments {
			assignment.Print(member, team.Members)
		}
		fmt.Println()
		pause()
	}
	fmt.Printf("%s%s%s makes final considerations.\n", colorBlue, team.Manager, colorReset)
	pause()
}

func pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
