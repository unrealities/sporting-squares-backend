package transformers

import (
	"github.com/unrealities/sporting-squares-backend/nflStats"
)

func UltraMagnus(games []nflStats.EspnNFLGame) (Energon, error) {
	// TODO
	/*
		gameIDs := extractGameIDs(resp)

	for i, gameID := range gameIDs {
		g := Game{}

		score, err := GetESPNGame(gameIDs[i])
		if err != nil {
			g.Quarter = -1
			continue
		}

		for _, c := range score.Header.Competitions {
			for _, t := range c.Competitors {
				score, err := strconv.Atoi(t.Score)
				if err != nil {
					score = -1
				}

				if t.HomeAway == "home" {
					g.HomeTeam = t.Team.Abbreviation
					g.HomeScore = score

					if len(t.Linescores) == 0 {
						g.Digits.Home1 = -1
						g.Digits.Home2 = -1
						g.Digits.Home3 = -1
						g.Digits.Home4 = -1
						continue
					}
					// TODO: This is verbose and duplicated
					firstQ, err := strconv.Atoi(t.Linescores[0].DisplayValue)
					if err != nil {
						g.Digits.Home1 = -1
					} else {
						g.Digits.Home1 = firstQ % 10
					}
					if len(t.Linescores) == 1 {
						g.Digits.Home2 = -1
						g.Digits.Home3 = -1
						g.Digits.Home4 = -1
						continue
					}
					secondQ, err := strconv.Atoi(t.Linescores[1].DisplayValue)
					if err != nil {
						g.Digits.Home2 = -1
					} else {
						g.Digits.Home2 = (firstQ + secondQ) % 10
					}
					if len(t.Linescores) == 2 {
						g.Digits.Home3 = -1
						g.Digits.Home4 = -1
						continue
					}
					thirdQ, err := strconv.Atoi(t.Linescores[2].DisplayValue)
					if err != nil {
						g.Digits.Home3 = -1
					} else {
						g.Digits.Home3 = (firstQ + secondQ + thirdQ) % 10
					}
					if len(t.Linescores) == 3 {
						g.Digits.Home4 = -1
						continue
					}
					fourthQ, err := strconv.Atoi(t.Linescores[3].DisplayValue)
					if err != nil {
						g.Digits.Home4 = -1
					} else {
						g.Digits.Home4 = (firstQ + secondQ + thirdQ + fourthQ) % 10
						if len(t.Linescores) > 4 {
							overtime, err := strconv.Atoi(t.Linescores[4].DisplayValue)
							if err != nil {
								g.Digits.Home4 = -1
							} else {
								g.Digits.Home4 = (firstQ + secondQ + thirdQ + fourthQ + overtime) % 10
							}
						}
					}
				} else {
					g.AwayTeam = t.Team.Abbreviation
					g.AwayScore = score

					if len(t.Linescores) == 0 {
						g.Digits.Away1 = -1
						g.Digits.Away2 = -1
						g.Digits.Away3 = -1
						g.Digits.Away4 = -1
						continue
					}
					firstQ, err := strconv.Atoi(t.Linescores[0].DisplayValue)
					if err != nil {
						g.Digits.Away1 = -1
					} else {
						g.Digits.Away1 = firstQ % 10
					}
					if len(t.Linescores) == 1 {
						g.Digits.Away2 = -1
						g.Digits.Away3 = -1
						g.Digits.Away4 = -1
						continue
					}
					secondQ, err := strconv.Atoi(t.Linescores[1].DisplayValue)
					if err != nil {
						g.Digits.Away2 = -1
					} else {
						g.Digits.Away2 = (firstQ + secondQ) % 10
					}
					if len(t.Linescores) == 2 {
						g.Digits.Away3 = -1
						g.Digits.Away4 = -1
						continue
					}
					thirdQ, err := strconv.Atoi(t.Linescores[2].DisplayValue)
					if err != nil {
						g.Digits.Away3 = -1
					} else {
						g.Digits.Away3 = (firstQ + secondQ + thirdQ) % 10
					}
					if len(t.Linescores) == 3 {
						g.Digits.Away4 = -1
						continue
					}
					fourthQ, err := strconv.Atoi(t.Linescores[3].DisplayValue)
					if err != nil {
						g.Digits.Away4 = -1
					} else {
						g.Digits.Away4 = (firstQ + secondQ + thirdQ + fourthQ) % 10
						if len(t.Linescores) > 4 {
							overtime, err := strconv.Atoi(t.Linescores[4].DisplayValue)
							if err != nil {
								g.Digits.Away4 = -1
							} else {
								g.Digits.Away4 = (firstQ + secondQ + thirdQ + fourthQ + overtime) % 10
							}
						}
					}
				}
			}
		}

		g.ID = gameID
		g.SeasonType = score.Header.Season.Type
		g.Year = score.Header.Season.Year
		g.Week = score.Header.Week
		if len(score.Drives.Previous) > 0 {
			g.Quarter = score.Drives.Previous[0].End.Period.Number
			g.Time = score.Drives.Previous[0].TimeElapsed.DisplayValue
		}
		g.GameOver = score.Header.Competitions[0].Status.Type.Completed
		g.Odds.Details = score.Pickcenter[0].Details
		g.Odds.OverUnder = score.Pickcenter[0].OverUnder
		games = append(games, g)
	}
	*/
}
