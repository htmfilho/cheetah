package main

import (
	"flag"
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/htmfilho/cheetah/team"
)

var flgTeamFileName = flag.String("f", "team.json", "Relative or absolute path to the data file.")
var flgCeremony = flag.String("c", team.CerimonyStandUp, "Cerimony to run.")

func main() {
	flag.Parse()

	printLogo()
	printFlags()

	var repo team.Repository
	repo = &team.JsonRepository{
		File: *flgTeamFileName,
	}

	currentTeam := repo.GetCurrentTeam()
	team.RunCerimony(*flgCeremony, currentTeam)
}

func printLogo() {
	logo := figure.NewFigure("Cheetah !", "speed", true)
	logo.Print()
}

func printFlags() {
	fmt.Println("\n-----------------------------------------------------------")
	fmt.Printf("Team File: %s\n", *flgTeamFileName)
	fmt.Printf("Ceremony : %s\n", *flgCeremony)
	fmt.Println("-----------------------------------------------------------")
}
