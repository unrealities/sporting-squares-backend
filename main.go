package main

import (
	"fmt"

	"github.com/unrealities/sporting-squares-backend/cmd/nflStats"
)

func main() {
	games, _ := nflStats.GetGames()
	for i, g := range games {
		fmt.Println("eggs")
		fmt.Printf("Year %d Week %d Type %d Game %v[%v]: %s (%v [%v, %v, %v, %v]) @ %s (%v [%v, %v, %v, %v]) [%v][%v] {%v (%v) [%v]} [%v]\n",
			g.Year, g.Week, g.SeasonType, i, g.ID, g.AwayTeam, g.AwayScore, g.Digits.Away1, g.Digits.Away2, g.Digits.Away3, g.Digits.Away4,
			g.HomeTeam, g.HomeScore, g.Digits.Home1, g.Digits.Home2, g.Digits.Home3, g.Digits.Home4, g.Quarter, g.Time, g.Odds.Details,
			g.Odds.OverUnder, g.GameOver)
	}
}
