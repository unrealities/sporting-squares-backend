package main

import (
	"fmt"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

func main() {
	games, _ := nflStats.GetGames()
	for i, g := range games {
		fmt.Printf("Game %v[%v]: %s (%v [%v, %v, %v, %v]) @ %s (%v [%v, %v, %v, %v]) [%v][%v] {%v [%v]}\n",
			i, g.ID, g.AwayTeam, g.AwayScore, g.Digits.Away1, g.Digits.Away2, g.Digits.Away3, g.Digits.Away4,
			g.HomeTeam, g.Digits.Home1, g.Digits.Home2, g.Digits.Home3, g.Digits.Home4, g.HomeScore,
			g.Quarter, g.Time, g.Odds.Details, g.Odds.OverUnder)
	}
}
