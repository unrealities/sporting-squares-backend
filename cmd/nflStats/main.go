package nflStats

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// TODO: reduce GetESPNScores struct to app readable data
func GetGames() ([]Game, error) {
	games := []Game{}
	resp, err := GetESPNScores()
	if err != nil {
		return games, fmt.Errorf("nflStats@GetScores: GetESPNScores, error: %s", err)
	}

	for _, e := range resp.Events {
		for _, c := range e.Competitions {
			g := Game{}
			g.Quarter = c.Status.Period // 0 = Game has not started. 4 = 4th or Game Over
			g.Time = c.Status.Clock // 0 = Game is over or has not started
			if len(c.Odds) > 0 {
				g.Odds.Details = c.Odds[0].Details // ex. ARI -2.5 (need to separate out team and spread)
				g.Odds.OverUnder = c.Odds[0].OverUnder // ex. 51.5
			}
			if e.ID == c.ID {
				for _, t := range c.Competitors {
					score, err := strconv.Atoi(t.Score)
					if err != nil {
						score = -1
					}
					if t.HomeAway == "home" {
						g.HomeTeam = t.Team.Abbreviation
						g.HomeScore = score
					} else {
						g.AwayTeam = t.Team.Abbreviation
						g.AwayScore = score
					}
				}
			}
			games = append(games, g)
		}
	}

	return games, nil
}

// GetESPNScores fetches current NFL game scores
func GetESPNScores() (EspnNFLScores, error) {
	URL := "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
	resp, err := http.Get(URL)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: Get %s, error: %s", URL, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	espnNFLScoresResp := EspnNFLScores{}
	json.Unmarshal([]byte(body), &espnNFLScoresResp)
	if err != nil {
		return EspnNFLScores{}, fmt.Errorf("nflStats#GetESPNScores: reading Get response body error: %s", err)
	}

	return espnNFLScoresResp, nil
}
