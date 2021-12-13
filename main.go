package main

import (
	"fmt"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

func main() {
	games, _ := nflStats.GetGames()
	for i, g := range games {
		fmt.Printf("Game %v: %s (%v) @ %s (%v) [%v][%v] {%v [%v]}\n",
			i, g.AwayTeam, g.AwayScore, g.HomeTeam, g.HomeScore, g.Quarter, g.Time, g.Odds.Details, g.Odds.OverUnder)
	}
}
