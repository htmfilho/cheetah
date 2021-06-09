package main

import (
	"flag"
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/htmfilho/cheetah/team"
)

var flgTeamFileName = flag.String("f", "team.json", "Relative or absolute path to the data file.")
var flgCeremony = flag.String("c", team.CeremonyStandUp, "Cerimony to run.")
var flgDebug = flag.Bool("d", false, "Run in debug mode, showing extra information.")

func main() {
	flag.Parse()

	printLogo()
	printFlags()

	var repo team.Repository
	repo = &team.JsonRepository{
		File: *flgTeamFileName,
	}

	currentTeam := repo.GetCurrentTeam()
	team.RunCeremony(*flgCeremony, currentTeam)
}

func printLogo() {
	logo := figure.NewFigure("Cheetah !", "speed", true)
	logo.Print()
}

func printFlags() {
	if *flgDebug {
		fmt.Println("==========================================================")
		fmt.Printf("\nFile     : %s\n", *flgTeamFileName)
		fmt.Printf("Ceremony : %s\n", *flgCeremony)

	}
	fmt.Println("==========================================================")
}
