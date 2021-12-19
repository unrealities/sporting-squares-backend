package main

import (
	"fmt"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

func main() {
	games, _ := nflStats.GetGames()
	for i, g := range games {
		fmt.Printf("Game %v[%v]: %s (%v) @ %s (%v) [%v][%v] {%v [%v]}\n",
			i, g.ID, g.AwayTeam, g.AwayScore, g.HomeTeam, g.HomeScore, g.Quarter, g.Time, g.Odds.Details, g.Odds.OverUnder)
	}
}
