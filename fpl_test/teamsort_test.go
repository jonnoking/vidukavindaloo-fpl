package fpl_test

import (
	"log"

	"testing"
)

func TestTeamSortByStrength(t *testing.T) {

	FPL := initFPL()
	log.Println(FPL.Bootstrap.Events.GetCurrentEvent().Name)

	FPL.Bootstrap.Teams.SortByStrength()

	for i := 0; i <= 10; i++ {
		log.Printf("[%d] %s - %d \n", FPL.Bootstrap.Teams.Teams[i].ID, FPL.Bootstrap.Teams.Teams[i].Name, FPL.Bootstrap.Teams.Teams[i].Strength)
	}
}

func TestTeamSortByHomeStrength(t *testing.T) {

	FPL := initFPL()
	log.Println(FPL.Bootstrap.Events.GetCurrentEvent().Name)

	FPL.Bootstrap.Teams.SortByHomeStrength()

	for i := 0; i <= 10; i++ {
		log.Printf("[%d] %s - %d \n", FPL.Bootstrap.Teams.Teams[i].ID, FPL.Bootstrap.Teams.Teams[i].Name, FPL.Bootstrap.Teams.Teams[i].StrengthOverallHome)
	}
}

func TestTeamSortByTotalPoints(t *testing.T) {

	FPL := initFPL()
	log.Println(FPL.Bootstrap.Events.GetCurrentEvent().Name)

	FPL.Bootstrap.Teams.SortByTotalPoints(FPL.Bootstrap.Players)

	for i := 0; i <= 10; i++ {
		log.Printf("[%d] %s - %d \n", FPL.Bootstrap.Teams.Teams[i].ID, FPL.Bootstrap.Teams.Teams[i].Name, FPL.Bootstrap.Teams.Teams[i].Points)
	}
}
